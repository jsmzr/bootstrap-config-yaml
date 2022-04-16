package yaml

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/jsmzr/bootstrap-config/config"

	"gopkg.in/yaml.v2"
)

type YamlConfig struct{}

type YamlContainer struct {
	content string
}

func (c *YamlConfig) Load(filename string) (config.Configer, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var instance map[interface{}]interface{}
	if err := yaml.Unmarshal(data, &instance); err != nil {
		return nil, err
	}
	dict := convertDict(instance)
	if b, err := json.Marshal(dict); err != nil {
		return nil, err
	} else {
		return &YamlContainer{content: string(b)}, nil
	}
}

func convertDict(data map[interface{}]interface{}) map[string]interface{} {
	dict := make(map[string]interface{})
	for k, v := range data {
		ks := fmt.Sprintf("%v", k)
		cv, ok := v.(map[interface{}]interface{})
		if ok {
			dict[ks] = convertDict(cv)
		} else {
			dict[ks] = v
		}
	}
	return dict
}

func (c *YamlContainer) GetJson() string {
	return c.content
}

func init() {
	config.Register("yaml", &YamlConfig{})
}
