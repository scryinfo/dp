package main

import (
    "github.com/scryinfo/dp/test/binary/cases"
    "github.com/scryinfo/dp/test/binary/utils"
)

const (
    caseNum = 1
)

func main() {
    // todo: format print to stdout -> write log file

    var ch = make(chan string, caseNum)
    
    // todo: register a group of verifiers

    // run cases
    // idea: simulate concurrent with goroutine: PT
    //       without goroutine: FT

    // todo: add the other cases
    go cases.CaseOne(ch)
    
    // record result
    utils.RecordResult(ch, caseNum)
}
