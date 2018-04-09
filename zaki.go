package main

import (
    "fmt"
    "log"
    "syscall"
)

func main() {
    h, e := syscall.LoadLibrary("square.dll")
    if e != nil {
        log.Fatal(e)
    }
    defer syscall.FreeLibrary(h)
    proc, e := syscall.GetProcAddress(h, "dll_math_example")
    if e != nil {
        log.Fatal(e)
    }
    
    n, _, _ := syscall.Syscall9(uintptr(proc), 0, 2, 2, 2, 2, 0, 0, 0, 0, 0)
    fmt.Printf("Hello dll function returns %d\n", n)
}