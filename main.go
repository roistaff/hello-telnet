package main

import (
    "context"
    "fmt"
    "os"
    "os/signal"
    "syscall"
    "time"

    telnet "github.com/reiver/go-telnet"
)

var counter int = 0

func main() {
    var handler telnet.Handler = MainHandler{}
    fmt.Println("start server")
    go func() {
        err := telnet.ListenAndServe(":10000", handler)
        if nil != err {
            panic(err)
        }
    }()
    ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt, os.Kill)
    defer stop()
    <-ctx.Done()
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
}
type MainHandler struct{}

func (handler MainHandler) ServeTELNET(ctx telnet.Context, w telnet.Writer, r telnet.Reader) {
    str := `
Hello World from Internet.I am netrunnner.
`
        w.Write([]byte(str))
    }
