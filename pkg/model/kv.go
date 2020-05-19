package model

import (
	"crypto/rand"
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/dsrhub/dsrhub/pkg/config"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"golang.org/x/crypto/nacl/secretbox"
)

const (
	secretKeySize   = 32
	secretNonceSize = 24
)

var (
	errEmptyMasterKeys     = errors.New("empty DSRHUB_MASTER_KEYS")
	errEmptyTx             = errors.New("empty tx")
	errScopeSecretNotFound = errors.New("scope secret not found")
	errDecryptionFailed    = errors.New("decryption failed")
)

// EncryptionScope describe a table to store the scope ID => Nonce mapping
// Useful for deleting the nonce for compliance
type EncryptionScope struct {
	BaseModel
	EncryptionType string
	ScopeID        string `gorm:"unique_index"`
	Nonce          string
}

func (esc *EncryptionScope) BeforeCreate(scope *gorm.Scope) error {
	var n [secretNonceSize]byte
	_, err := rand.Reader.Read(n[:])
	if err != nil {
		return err
	}
	return scope.SetColumn("Nonce", string(n[:]))
}

// EncryptionKV is a KV like type that we can use to represent encrypted
// map structure in a table. It leverages the EncryptionScope table to store
// the nonce.
type EncryptionKV struct {
	EncryptionScopeID string `json:"encryption_scope_id"`
	Ciphertext        string `json:"ciphertext"`

	m      map[string]interface{}
	loaded bool
	tx     *gorm.DB
}

func NewEncryptionKV(tx *gorm.DB, scopeID string, m map[string]interface{}) (*EncryptionKV, error) {
	ekv := &EncryptionKV{
		EncryptionScopeID: scopeID,
		m:                 m,
		loaded:            true,
		tx:                tx,
	}

	esc := &EncryptionScope{
		EncryptionType: "secretbox",
	}
	if err := tx.FirstOrCreate(esc, scopeID).Error; err != nil {
		return nil, err
	}

	err := ekv.encrypt(esc)
	if err != nil {
		return nil, err
	}
	return ekv, nil
}

// WithTx overrides the Tx
func (ekv *EncryptionKV) WithTx(tx *gorm.DB) *EncryptionKV {
	ekv.tx = tx
	return ekv
}

// GET loads the EncryptionKV.m and returns the value of a key
func (ekv *EncryptionKV) Get(key string) (interface{}, bool, error) {
	if err := ekv.load(); err != nil {
		return "", false, err
	}
	s, found := ekv.m[key]
	return s, found, nil
}

// Set loads the EncryptionKV.m and set a key/value pair without saving it to DB
func (ekv *EncryptionKV) Set(key string, value interface{}) error {
	if err := ekv.load(); err != nil {
		return err
	}
	ekv.m[key] = value
	return nil
}

// load loads the nonce from EncryptionScope by EncryptionScopeID
// and then decrypt to get the EncryptionKV.m, and then sets loaded to true
func (ekv *EncryptionKV) load() error {
	if ekv.loaded {
		return nil
	}
	if ekv.tx == nil {
		return errEmptyTx
	}

	esc := &EncryptionScope{}
	result := ekv.tx.First(esc, ekv.EncryptionScopeID)
	if result.RecordNotFound() {
		config.Logger.Info(
			errScopeSecretNotFound.Error(),
			zap.String("EncryptionScopeID", ekv.EncryptionScopeID),
		)
		return errScopeSecretNotFound
	}
	if result.Error != nil {
		return result.Error
	}
	err := ekv.decrypt(esc)
	if err != nil {
		return err
	}

	ekv.loaded = true

	return nil
}

func (ekv *EncryptionKV) encrypt(esc *EncryptionScope) error {
	if len(config.ENV.MasterKeys) == 0 {
		return errEmptyMasterKeys
	}
	var key [secretKeySize]byte
	var nonce [secretNonceSize]byte
	var box []byte
	copy(key[:], []byte(config.ENV.MasterKeys[0]))
	copy(nonce[:], []byte(esc.Nonce))
	payload, _ := json.Marshal(ekv.m)
	box = secretbox.Seal(box[:0], payload, &nonce, &key)
	ekv.Ciphertext = string(box[:])
	return nil
}

func (ekv *EncryptionKV) decrypt(esc *EncryptionScope) error {
	for _, k := range config.ENV.MasterKeys {
		var key [secretKeySize]byte
		var nonce [secretNonceSize]byte
		var opened []byte
		copy(key[:], []byte(k))
		copy(nonce[:], []byte(esc.Nonce))
		opened, ok := secretbox.Open(opened[:0], []byte(ekv.Ciphertext), &nonce, &key)
		if ok {
			return json.Unmarshal(opened[:], &ekv.m)
		}
	}
	return errDecryptionFailed
}

func (ekv *EncryptionKV) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	s := cast.ToString(value)
	if err := json.Unmarshal([]byte(s), ekv); err != nil {
		return err
	}
	return nil
}

func (ekv EncryptionKV) Value() (driver.Value, error) {
	bytes, err := json.Marshal(ekv)
	if err != nil {
		return nil, err
	}
	return string(bytes), nil
}

type KV map[string]interface{}

func (kv *KV) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	s := cast.ToString(value)
	if err := json.Unmarshal([]byte(s), kv); err != nil {
		return err
	}
	return nil
}

func (kv KV) Value() (driver.Value, error) {
	bytes, err := json.Marshal(kv)
	if err != nil {
		return nil, err
	}
	return string(bytes), nil
}
