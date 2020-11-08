package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

type sendContext struct {
	laddr   *net.UDPAddr
	raddr   *net.UDPAddr
	udpc    *net.UDPConn
	segment []byte
	wg      sync.WaitGroup
}

func wait(ctx *sendContext) {
	// var err error

	fmt.Println("wait  routine started")
	for {
		nbytes, remoteAddr, err := ctx.udpc.ReadFromUDP(ctx.segment)
		if err != nil {
			fmt.Println("WriteToUDP error")
		}
		fmt.Printf("received a datagram with size %d from %v\n", nbytes, remoteAddr)
		time.Sleep(time.Second)
	}

	// ctx.wg.Done()

}
func main() {

	ip := flag.String("i", "127.0.0.1", "local ip")
	port := flag.String("p", "5000", "local port")
	var ctx sendContext
	var err error

	flag.Parse()

	ctx.laddr, err = net.ResolveUDPAddr("udp", *ip+":"+*port)

	if err != nil {
		fmt.Printf("error : %v", err)
		return
	}

	ctx.udpc, err = net.ListenUDP("udp", ctx.laddr)
	if err != nil {
		fmt.Printf("error : %v", err)
		return
	}
	fmt.Printf("linstening on %v\n", ctx.laddr)

	ctx.segment = make([]byte, 2048)

	ctx.wg.Add(1)
	go wait(&ctx)

	ctx.wg.Wait()

	fmt.Println("bye bye !")

}
