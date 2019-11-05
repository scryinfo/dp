// Scry Info.  All rights reserved.
// license that can be found in the license file.

package define

// Conf
type Conf struct {
	Person    `yaml:"Person"`
	Contact   `yaml:"Contact"`
	Education `yaml:"Education"`
	Inline    `yaml:"Inline,inline"` //use ",inline" read simple configuration's value.
	InlineAdd `yaml:"InlineAdd,inline"`
}

// InlineAdd
type InlineAdd struct {
	Height int `yaml:"Height_cm"`
	Heavy  int `yaml:"Heavy_kg"`
}

// Inline simple configuration item.
type Inline struct {
	NativePlace     string   `yaml:"Native_place"`
	PliticalOutlook string   `yaml:"Political_outlook"`
	MaritalStatus   string   `yaml:"Marital_status"`
	Hobby           []string `yaml:"Hobby,flow"`       //"flow" shows in array .
	ForeignLanguage []string `yaml:"Foreign_language"` //without flow shows in several lines begin with "-".
	AwardsReceived  string   `yaml:"Awards_received"`
}

// Person complex configuration item which has lower level tags.
type Person struct {
	Name `yaml:"Name"`
	Sex  string `yaml:"Sex"`
	Age  int    `yaml:"Age"`
	IDCN int64  `yaml:"ID_card_number"`
}
// Name
type Name struct {
	CName string `yaml:"Chinese_name"`
	EName string `yaml:"English_name"`
	//NameUB string `yaml:"Name_used_before"`
}

// Contact
type Contact struct {
	Tel   int64    `yaml:"Telephone"`
	Phone int32    `yaml:"Phone"`
	QQ    int64    `yaml:"QQ"`
	WX    string   `yaml:"Wei_xin"`
	EMail []string `yaml:"E-mail"`
}

// Education
type Education struct {
	Primary School `yaml:"Primary"`
	Junior  School `yaml:"Junior"`
	Senior  School `yaml:"Senior"`
	Collage School `yaml:"Collage"`
}
// School
type School struct {
	SName string `yaml:"School_name"`
	SDate string `yaml:"Start_date"`
	EDate string `yaml:"End_date"`
}
