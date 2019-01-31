package configuration

import (
    "github.com/go-yaml/yaml"
    "io/ioutil"
    "os"
    "reflect"
)

func GetYAMLStructure(fileAddr string, v interface{}) (interface{}, error) {
	yf, err := ioutil.ReadFile(fileAddr)
	if err != nil {
		return nil, err
	}

	t := reflect.TypeOf(v).Elem()
	conf := reflect.New(t).Interface()
	err = yaml.Unmarshal(yf, conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}


/*
    This function will delete items in .yaml file which not in structure.go,
    so make sure items in structure.go is not less than in .yaml file.
*/
func SaveChanges(fileAddr string, conf interface{}) error {
	err := writeFile(fileAddr, conf, os.O_TRUNC) //O_TRUNC param will rewrite the file
	if err != nil {
		return err
	}

	return nil
}

func writeFile(fileAddr string, ymlConf interface{}, opt int) error {
	wf, err := os.OpenFile(fileAddr, os.O_WRONLY|os.O_EXCL|opt, 0777)
	if err != nil {
		return err
	}

	ycb, err := yaml.Marshal(ymlConf)
	if err != nil {
		return err
	}

	_, err = wf.Write(ycb)
	if err != nil {
		return err
	}

	return nil
}
/*
func Add(fileAddr string, inlineAdd *InlineAdd) error {
	addNow, err := getInlineAdd(fileAddr)
	if err != nil {
		return err
	}
	if *addNow != *new(InlineAdd) {
		return errors.New("please ensure items in \"inlineAdd\" all not in .yaml file first")
	}

	err = writeFile(fileAddr, inlineAdd, os.O_APPEND)
	if err != nil {
		return err
	}

	return nil
}

func getInlineAdd(fileAddr string) (*InlineAdd, error) {
	yf, err := ioutil.ReadFile(fileAddr)
	if err != nil {
		return nil, err
	}

	inlineAdd := new(InlineAdd)
	err = yaml.Unmarshal(yf, inlineAdd)
	if err != nil {
		return nil, err
	}

	return inlineAdd, nil
}*/
