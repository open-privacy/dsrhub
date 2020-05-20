package model

import (
	"io/ioutil"

	"github.com/jinzhu/copier"
	"gopkg.in/yaml.v2"
)

type DSL struct {
	Version   string               `yaml:"version"`
	Workflows map[string]*Workflow `yaml:"workflows"`

	yaml string `yaml:"-"`
}

func (dsl *DSL) Load(path string) error {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	dsl.yaml = string(dat)
	return yaml.UnmarshalStrict(dat, dsl)
}

func (dsl *DSL) Clone() *DSL {
	n := &DSL{}
	copier.Copy(n, dsl)
	return n
}

func (dsl *DSL) YAML() string {
	return dsl.yaml
}
