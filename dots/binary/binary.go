package binary

import "github.com/scryInfo/dot/dot"

type Binary struct {
}

func NewBinary() *Binary {
	return &Binary{}
}

func (c *Binary) Create(l dot.Line) error {
	return nil
}

func (c *Binary) Start(ignore bool) error {
	return nil
}

func (c *Binary) Stop(ignore bool) error {
	return nil
}

func (c *Binary) Destroy(ignore bool) error {
	return nil
}
