package main

import (
	"fmt"
	"github.com/mats9693/YAML"
	"log"
)

//modify the directory before you run this demo.
const basedir = "github.com/mats9693/YAML/test"

func main() {
	conf, err := my.GetYAMLStructure(basedir + "/test.yaml")
	if err != nil {
		log.Println(err)
		return
	}

	/*
		conf is prepared,you can use any configuration item as below:
	*/

	//fmt.Println(conf)
	fmt.Println("Author : ", conf.Person.EName)
	fmt.Println("E-mail : ", conf.Contact.EMail)
	fmt.Println("From   : ", conf.NativePlace)

	/*
		e.g. add items : (Deprecated)
			 if you add items only , this method will be a better choice.
	*/
	//conf.InlineAdd.Heavy = 95
	//conf.InlineAdd.Height = 183
	//err = my.Add(basedir+"/test.yaml", &conf.InlineAdd)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}

	/*
		e.g. modify items :
			 you can also add or delete items , remember modify the structure file first .
	*/

	//conf.Person.Name.NameUB = "abc"
	conf.Contact.EMail = []string{"982200000@qq.com", "mat00000000@foxmail.com", "new e-mail."}
	conf.ForeignLanguage = []string{"CET-4", "no others."}
	err = my.SaveChanges(basedir+"/test.yaml", conf)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("E-mail`: ", conf.Contact.EMail)
}
