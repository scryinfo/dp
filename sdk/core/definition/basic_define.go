package definition

const (
	Document int = iota
	Picture
	Video
	Audio
	Other
)

type DescriptionData struct {
	id uint
	title string
	publishTime uint64
}

type ProofData struct {
	id uint
	hint string
	stream []byte
	dataType int
	publishTime uint64
}

const (
	Yes int = iota
	No
)

type Vote struct {
	voteResult int
	comment    string
}