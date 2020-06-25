package main

import (
	"testing"

	"github.com/dsrhub/dsrhub/idl_dsrhub"
	"github.com/stretchr/testify/assert"
)

func TestValidConfig(t *testing.T) {
	var cfg DsrHubGRPCConfig

	cfg = DsrHubGRPCConfig{}
	assert.Error(t, validConfig(&cfg))

	cfg = DsrHubGRPCConfig{
		URL: "localhost:50051",
		Request: idl_dsrhub.CreateDSRRequest{
			ApiVersion:         "2.0",
			Regulation:         "gdpr",
			StatusCallbackUrl:  "http://localhost:9999/dsrhub/callback",
			SubjectRequestId:   "123456",
			SubjectRequestType: "erasure",
			IdentityType:       "email",
			IdentityFormat:     "raw",
			IdentityValue:      "test@example.com",
		},
		Timeout: 10,
	}
	assert.NoError(t, validConfig(&cfg))
}
