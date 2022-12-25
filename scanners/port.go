package main

import (
	"fmt"
	"net"
	"sort"
)

func main() {
    pc := make(chan int, 100)
    oc := make(chan int)

    var open []int

    for i := 0; i <= cap(pc); i++ {
        go worker(pc, oc)
    }

    go func() {
        for i := 1; i <= 1024; i++ {
            pc <- i
        }
    }()

    for i := 0; i < 1024; i++ {
        port := <-oc
        if port != 0 {
            open = append(open, port)
        }
    }

    close(pc)
    close(oc)
    sort.Ints(open)

    for _, port := range open {
        fmt.Printf("%d open\n", port)
    }
}

func worker(ports chan int, results chan int) {
    for p := range ports {
        addr := fmt.Sprintf("scanme.nmap.org:%d", p)
        conn, err := net.Dial("tcp", addr)
        defer conn.Close()

        if err != nil {
            results <- 0
            continue
        }
        
        results <- p
    }
}
