package utils

import (
    "fmt"
    "io"
    "io/ioutil"
    "os/exec"
    "strconv"
    "time"
)

const (
    // todo: config file?
    dir        = "D:/GoPath/src/github.com/scryinfo/dp/test/binary/users/"
    resultFile = "D:/GoPath/src/github.com/scryinfo/dp/test/binary/result.txt"

    deployerAddr = "0xd280b60c38bc8db9d309fa5a540ffec499f0a3e8"
    deployerPwd  = "111111"
)

func Start(identify string, caseNo string) (user exec.Cmd, pipe io.ReadCloser) {
    user = exec.Cmd{
        Path: "main.exe",
        Dir:  dir + identify + "/" + caseNo,
    }

    var err error

    if pipe, err = user.StdoutPipe(); err != nil {
        fmt.Println("make ", identify, " pipe failed", err)
    }

    if err = user.Start(); err != nil {
        fmt.Println("sub process start failed", err)
    }

    fmt.Println("> Start sub process: ", identify, " pid:", user.Process.Pid)
    
    time.Sleep(time.Second * 3)
    
    return
}

func ReadPipes(ch chan int, pipes ...io.ReadCloser) {
    for i := range pipes {
        go func(pipe io.ReadCloser, i int, ch chan int) {
            var buffer = make([]byte, 4096)
            for {
                n, err := pipe.Read(buffer)
                if err != nil {
                    if err == io.EOF {
                        fmt.Println("pipe ", i, " has closed. ")
                        break
                    } else {
                        fmt.Println("Read pipe", i, " failed. ")
                    }
                }
                if string(buffer[:n]) == "[exit]" {
                    fmt.Println("Node: pipe ", i, " exit. ")
                    ch <- i
                }
                fmt.Println(string(buffer[:n]))
            }
        }(pipes[i], i, ch)
    }
}

func Stop(ch chan int, chStr chan string, userNum int, caseNo int, processes ...exec.Cmd) {
    for i := 0; i < userNum; i++ {
        t := <-ch
        fmt.Println("Node: receive exit flag. ", t)
        stop(processes[t])
    }

    chStr <- "Test case " + strconv.Itoa(caseNo) + " passed. "
}

func stop(process exec.Cmd) {
    process.Process.Kill()

    if err := process.Wait(); err != nil {
        fmt.Printf("> Release sub process: %d, with error: %v\n", process.Process.Pid, err)
    } else {
        fmt.Println("> Release sub process: ", process.Process.Pid)
    }
}

func RegisterVerifiers(chStr chan string) {
    process, pipe := Start("verifiers", "register")

    var ch = make(chan int)
    ReadPipes(ch, pipe)

    time.Sleep(time.Second * 3)

    Stop(ch, chStr, 1, 0, process)
}

func RecordResult(ch chan string, caseNum int) {
    var str string
    for i := 0; i < caseNum; i++ {
        t := <- ch
        str += t + "\r\n"
    }

    _ = ioutil.WriteFile(resultFile, []byte(str), 0666)
}
