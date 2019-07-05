package main

import (
    "github.com/scryinfo/dp/demo/src/sdk/util/configuration"
    "github.com/scryinfo/dp/demo/yaml/define"
    rlog "github.com/sirupsen/logrus"
)

//modify the directory before you run this demo.
func main() {
    rv, err := configuration.GetYAMLStructure("./test/test.yaml", &define.Conf{})
    if err != nil {
        rlog.Println(err)
        return
    }

    conf, ok := rv.(*define.Conf)
    if !ok {
        rlog.Error("failed to get yaml structure")
        return
    }

    rlog.Debug("conf: ", conf)

    conf.Contact.EMail = []string{"982200000@qq.com", "mat00000000@foxmail.com", "new e-mail."}
    conf.ForeignLanguage = []string{"CET-4", "no others."}
    err = configuration.SaveChanges("./test/test.yaml", conf)
    if err != nil {
        rlog.Println(err)
        return
    }

    rlog.Debug("E-mail`: ", conf.Contact.EMail)
}
