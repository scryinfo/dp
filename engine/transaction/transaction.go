package transaction

const (
	Start int = iota
	Created
	Voted
	ReadyForDownload
	Arbitrating
	Payed
	Closed
)


type Transaction struct {
	id int
	state int
	publishId string
}