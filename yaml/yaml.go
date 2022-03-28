package yaml

import (
	"io/ioutil"

	"github.com/jsmzr/bootstrap-config/config"
	"github.com/jsmzr/bootstrap-config/util"

	"gopkg.in/yaml.v2"
)

type YamlConfig struct{}

type YamlContainer struct {
	dict map[string]interface{}
}

func (c *YamlConfig) Load(filename string) (config.Configer, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var dict map[interface{}]interface{}
	err = yaml.Unmarshal(data, &dict)
	if err != nil {
		return nil, err
	}
	return &YamlContainer{
		dict: util.FlatMap(&dict),
	}, nil
}

func (c *YamlContainer) Get(key string) (interface{}, bool) {
	value, ok := c.dict[key]
	return value, ok
}

func (c *YamlContainer) Resolve(prefix string, p interface{}) error {
	return util.ResolveStruct(&c.dict, prefix, p)
}

func init() {
	config.Register("yaml", &YamlConfig{})
}
