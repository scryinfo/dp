package main

import (
	"fmt"
	"log"
	"../../configuration"
	"./define"
)

//modify the directory before you run this demo.
func main() {
	rv, err := configuration.GetYAMLStructure("./test/test.yaml", &define.Conf{})
	if err != nil {
		log.Println(err)
		return
	}

	conf, ok := rv.(*define.Conf)
    if !ok {
        fmt.Println("failed to get yaml structure")
        return
    }

	fmt.Println("conf: ", conf)

	conf.Contact.EMail = []string{"982200000@qq.com", "mat00000000@foxmail.com", "new e-mail."}
	conf.ForeignLanguage = []string{"CET-4", "no others."}
	err = configuration.SaveChanges("./test/test.yaml", conf)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("E-mail`: ", conf.Contact.EMail)
}
