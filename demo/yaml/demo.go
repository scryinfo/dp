package main

import (
	"github.com/scryInfo/dp/demo"
	configuration2 "github.com/scryInfo/dp/dots/binary/sdk/util/config_yaml"
	rlog "github.com/sirupsen/logrus"
)

//modify the directory before you run this demo.
func main() {
	rv, err := configuration2.GetYAMLStructure("./test/test.yaml", &demo.Conf{})
	if err != nil {
		rlog.Println(err)
		return
	}

	conf, ok := rv.(*demo.Conf)
	if !ok {
		rlog.Error("failed to get yaml structure")
		return
	}

	rlog.Debug("conf: ", conf)

	conf.Contact.EMail = []string{"982200000@qq.com", "mat00000000@foxmail.com", "new e-mail."}
	conf.ForeignLanguage = []string{"CET-4", "no others."}
	err = configuration2.SaveChanges("./test/test.yaml", conf)
	if err != nil {
		rlog.Println(err)
		return
	}

	rlog.Debug("E-mail`: ", conf.Contact.EMail)
}
