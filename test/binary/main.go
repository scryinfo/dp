package main

import (
    "github.com/scryinfo/dp/test/binary/cases"
    "github.com/scryinfo/dp/test/binary/utils"
)

const (
    caseNum = 1
)

func main() {
    // think: format print to stdout -> write log file, is it necessary?

    var (
        ch = make(chan string, caseNum)
    )

    // todo: concurrent register
    utils.RegisterVerifiers(ch)

    // run cases
    // idea:
    //   simulate concurrent with goroutine: PT
    //   without goroutine: FT
    // todo: add the other cases
    cases.CaseOne(ch)

    // record result
    utils.RecordResult(ch, caseNum)
}
