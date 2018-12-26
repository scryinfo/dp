package main

import (
	"fmt"
	"github.com/go-yaml/yaml"
	y "github.com/mats9693/YAML"
	"io/ioutil"
	"log"
)

//modify the directory before you run this demo.
const basedir = "github.com/mats9693/YAML/test"

func main() {
	yf, err := ioutil.ReadFile(basedir + "/test.yaml")
	if err != nil {
		log.Println("Read file Failed : ", err)
		return
	}

	c := new(y.Conf)
	err = yaml.Unmarshal(yf, c)
	if err != nil {
		log.Println("Unmarshal failed : ", err)
		return
	}

	//c is prepared,you can use any configration item as below:
	//fmt.Println(c)
	fmt.Println("Author : ", c.Person.CName)
	fmt.Println("E-mail : ", c.Contact.EMail)
	fmt.Println("From   : ", c.NativePlace)
}
