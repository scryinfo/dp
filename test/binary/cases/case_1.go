package cases

import (
    "fmt"
    "github.com/scryinfo/dp/test/binary/utils"
    "io"
    "os/exec"
    "time"
)

func CaseOne() {
    const userNum = 2

    var (
        processes = make([]exec.Cmd, userNum, userNum)
        pipes     = make([]io.ReadCloser, userNum, userNum)

        ch = make(chan int, userNum)
    )
    
    processes[0], pipes[0] = utils.Start("seller")
    processes[1], pipes[1] = utils.Start("buyer")
    
    fmt.Println("Node: show params. ", processes, pipes)

    utils.ReadPipes(pipes, ch)

    time.Sleep(time.Second * 3)

    for i := 0; i < userNum; i++ {
        t := <-ch
        fmt.Println("Node: receive exit flag. ", t)
        utils.Stop(processes[t])
    }
    
    utils.RecordResult([]byte("Test case 1 passed. "))
}
