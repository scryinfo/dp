package main

import (
	"fmt"
	"github.com/scryInfo/dot/dot"
	_ "github.com/scryInfo/dot/dot"
	"github.com/scryInfo/dot/dots/line"
)



func main()  {
	l, err := line.BuildAndStart(func(l dot.Line) error {
		//todo
		return nil
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	line.StopAndDestroy(l, true)
}
