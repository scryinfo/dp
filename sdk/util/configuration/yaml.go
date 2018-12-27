package my_yaml

type Conf struct {
	Person    `yaml:"Person"`
	Contact   `yaml:"Contact"`
	Education `yaml:"Education"`
	Inline    `yaml:"Inline,inline"`//use ",inline" read simple configration's value.
}

//simple configration item.
type Inline struct {
	NativePlace     string   `yaml:"Native_place"`
	PliticalOutlook string   `yaml:"Political_outlook"`
	MaritalStatus   string   `yaml:"Marital_status"`
	Hobby           []string `yaml:"Hobby,flow"`
	ForeignLanguage []string `yaml:"Foreign_language,flow"`
	AwardsReceived  string   `yaml:"Awards_received"`
}

//complex configration item which has lower level tags.
type Person struct {
	Name `yaml:"Name"`
	Sex  string `yaml:"Sex"`
	Age  int    `yaml:"Age"`
	IDCN int64  `yaml:"ID_card_number"`
}
type Name struct {
	CName string `yaml:"Chinese_name"`
	EName string `yaml:"English_name"`
}

type Contact struct {
	Tel   int64    `yaml:"Telephone"`
	Phone int32    `yaml:"Phone"`
	QQ    int64    `yaml:"QQ"`
	WX    string   `yaml:"Wei_xin"`
	EMail []string `yaml:"E-mail"`
}

type Education struct {
	Primary School `yaml:"Primary"`
	Junior  School `yaml:"Junior"`
	Senior  School `yaml:"Senior"`
	Collage School `yaml:"Collage"`
}
type School struct {
	SName string `yaml:"School_name"`
	SDate string `yaml:"Start_date"`
	EDate string `yaml:"End_date"`
}
