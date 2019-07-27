package cases

import (
    "github.com/scryinfo/dp/test/binary/utils"
    "io"
    "os/exec"
    "time"
)

func CaseOne(chStr chan string) {
    const userNum = 2

    var (
        processes = make([]exec.Cmd, userNum, userNum)
        pipes     = make([]io.ReadCloser, userNum, userNum)

        ch = make(chan int, userNum)
    )
    
    processes[0], pipes[0] = utils.Start("seller", "case_1")
    processes[1], pipes[1] = utils.Start("buyer", "case_1")

    utils.ReadPipes(pipes, ch)

    time.Sleep(time.Second * 3)

    utils.Stop(processes, ch, chStr, userNum)
}
