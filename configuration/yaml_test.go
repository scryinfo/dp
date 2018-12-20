package my_yaml

import (
	"testing"
)

func TestConfigEngine_Load(t *testing.T) {
	ce := ConfigEngine{}
	err := ce.Load("testdata/test1.yaml")
	if err != nil {
		t.Error("Load failed : ", err)
	}
}

func TestConfigEngine_Load1(t *testing.T) {
	ce := ConfigEngine{}
	err := ce.Load("testdata/test2.yml")
	if err != nil {
		t.Error("Load failed : ", err)
	}
}

func TestConfigEngine_Load2(t *testing.T) {
	ce := ConfigEngine{}
	err := ce.Load("testdata/abc.yaml")
	if err != nil {
		t.Error("Load failed : ", err)
	}
}

func TestConfigEngine_Load_Negative_Path(t *testing.T) {
	ce := ConfigEngine{}
	//the absolute path is wrong. remove "\\wrong" can make it right.
	err := ce.Load("C:\\Users\\马同帅\\go\\src\\codeTest\\testdata\\wrong\\abc.yaml")
	if err == nil {
		t.Error("Load wrong path succeed.")
	}
}

func TestConfigEngine_Load_Negative_Extension(t *testing.T) {
	ce := ConfigEngine{}
	//.yaml or .yml extension is needed.
	err := ce.Load("testdata/abc.txt")
	if err == nil {
		t.Error("Load wrong extension succeed.")
	}
}

var c = ConfigEngine{}
var err = c.Load("testdata/abc.yaml")

func TestConfigEngine_Get(t *testing.T) {
	v := c.Get("Nginx")
	if v == nil {
		t.Error("Get failed : ", v, " (expected nil).")
	}
}

func TestConfigEngine_GetBool(t *testing.T) {
	v := c.GetBool("TypeBoolean")
	if v != true {
		t.Error("GetBool failed : ", v, " (expected true).")
	}
}

func TestConfigEngine_GetInt(t *testing.T) {
	v := c.GetInt("TypeInt")
	if v != 1996 {
		t.Error("GetInt failed : ", v, " (expected 1996).")
	}
}

func TestConfigEngine_GetFloat64(t *testing.T) {
	v := c.GetFloat64("TypeFloat64")
	if v != 3.19 {
		t.Error("GetFloat64 failed : ", v, " (expected 3.19).")
	}
}

func TestConfigEngine_GetString(t *testing.T) {
	v := c.GetString("TypeString")
	if v != "I've been pretending to work hard, but you're really growing up." {
		t.Error("GetString failed : ", v,
			" (expected I've been pretending to work hard, but you're really growing up.).")
	}
}

func TestConfigEngine_Get_Negative_WrongName(t *testing.T) {
	//WrongConfigName is not exist in config file.
	v := c.Get("WrongConfigName")
	if v != nil {
		t.Error("Get wrong config name succeed.")
	}
}

type SiteConfig struct {
	HttpPort  int
	HttpsOn   bool
	Domain    string
	HttpsPort int
}

func TestConfigEngine_GetStruct(t *testing.T) {
	sc := SiteConfig{}
	v := c.GetStruct("Site", &sc)
	switch v.(type) {
	case SiteConfig:
		sc = v.(SiteConfig)
	}
	scExp := SiteConfig{HttpsPort: 443, Domain: "github.com", HttpsOn: false, HttpPort: 8080}
	if sc != scExp {
		t.Error("GetStuct failed : ", v, " (expected ", scExp, ").")
	}
}

func TestConfigEngine_GetStruct_Negative_WrongName(t *testing.T) {
	sc := SiteConfig{}
	v := c.GetStruct("Site111", &sc) // Site111 is not exist in config file.
	switch v.(type) {
	case *SiteConfig:
		sc = *(v.(*SiteConfig))
	}
	if sc != *new(SiteConfig) {
		t.Error("Get wrong struct name succeed.")
	}
}
