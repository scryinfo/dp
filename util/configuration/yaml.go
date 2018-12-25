package yamlconfig

import (
	"errors"
	"fmt"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"log"
	"reflect"
	"strconv"
	"strings"
)

type ConfigEngine struct {
	data map[interface{}]interface{}
}

func (c *ConfigEngine) Load(path string) error {
	ext := c.guessFileType(path)
	if ext == "" {
		return errors.New("Can not load : " + path + " address.")
	}
	return c.loadFromYaml(path)
}

func (c *ConfigEngine) guessFileType(path string) string {
	s := strings.Split(path, ".")
	ext := s[len(s)-1]
	switch ext {
	case "yaml", "yml":
		return "yaml"
	}
	return ""
}

func (c *ConfigEngine) loadFromYaml(path string) error {
	yamlS, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlS, &c.data)
	if err != nil {
		return errors.New("Can not parse : " + path + " address.")
	}
	return nil
}

func (c *ConfigEngine) Get(configName string) interface{} {
	path := strings.Split(configName, ".")
	data := c.data
	for index, value := range path {
		v, ok := data[value]
		if !ok {
			break
		}
		if (index + 1) == len(path) {
			return v
		}
		if reflect.TypeOf(v).String() == "map[interface {}]interface {}" {
			data = v.(map[interface{}]interface{})
		}
	}
	return nil
}

func (c *ConfigEngine) GetString(configName string) string {
	value := c.Get(configName)
	switch value := value.(type) {
	case string:
		return value
	case bool, float64, int:
		return fmt.Sprint(value)
	default:
		return ""
	}
}

func (c *ConfigEngine) GetInt(configName string) int {
	value := c.Get(configName)
	switch value := value.(type) {
	case string:
		i, _ := strconv.Atoi(value)
		return i
	case int:
		return value
	case bool:
		if value {
			return 1
		}
		return 0
	case float64:
		return int(value)
	default:
		return 0
	}
}

func (c *ConfigEngine) GetBool(configName string) bool {
	value := c.Get(configName)
	switch value := value.(type) {
	case string:
		str, _ := strconv.ParseBool(value)
		return str
	case int:
		if value != 0 {
			return true
		}
		return false
	case bool:
		return value
	case float64:
		if value != 0.0 {
			return true
		}
		return false
	default:
		return false
	}
}

func (c *ConfigEngine) GetFloat64(configName string) float64 {
	value := c.Get(configName)
	switch value := value.(type) {
	case string:
		str, _ := strconv.ParseFloat(value, 64)
		return str
	case int:
		return float64(value)
	case bool:
		if value {
			return float64(1)
		}
		return float64(0)
	case float64:
		return value
	default:
		return float64(0)
	}
}

func (c *ConfigEngine) GetStruct(configName string, s interface{}) interface{} {
	d := c.Get(configName)
	switch d.(type) {
	case string:
		err := c.setField(s, configName, d)
		if err != nil {
			log.Println(err)
			return nil
		}
	case map[interface{}]interface{}:
		c.mapToStruct(d.(map[interface{}]interface{}), s)
	}
	return s
}

func (c *ConfigEngine) mapToStruct(m map[interface{}]interface{}, s interface{}) interface{} {
	for key, value := range m {
		switch key.(type) {
		case string:
			err := c.setField(s, key.(string), value)
			if err != nil {
				log.Println(err)
				return nil
			}
		}
	}
	return s
}

func (c *ConfigEngine) setField(obj interface{}, configName string, value interface{}) error {
	structValue := reflect.Indirect(reflect.ValueOf(obj))
	structFieldValue := structValue.FieldByName(configName)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", configName)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", configName)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)

	if structFieldType.Kind() == reflect.Struct && val.Kind() == reflect.Map {
		vint := val.Interface()

		switch vint.(type) {
		case map[interface{}]interface{}:
			for key, value := range vint.(map[interface{}]interface{}) {
				err := c.setField(structFieldValue.Addr().Interface(), key.(string), value)
				if err != nil {
					log.Println(err)
					return nil
				}
			}
		case map[string]interface{}:
			for key, value := range vint.(map[string]interface{}) {
				err := c.setField(structFieldValue.Addr().Interface(), key, value)
				if err != nil {
					log.Println(err)
					return nil
				}
			}
		}

	} else {
		if structFieldType != val.Type() {
			return errors.New("Provided value type didn't match obj field type")
		}

		structFieldValue.Set(val)
	}

	return nil
}
