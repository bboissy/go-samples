package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
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

func send(ctx *sendContext) {
	var err error

	fmt.Println("send routine started")
	for {
		fmt.Printf("sending datagram to %v\n", ctx.raddr)
		ctx.udpc.WriteToUDP(ctx.segment, ctx.raddr)
		if err != nil {
			fmt.Println("WriteToUDP error")
		}
		time.Sleep(time.Second)
	}

	ctx.wg.Done()

}
func main() {

	destIp := flag.String("i", "127.0.0.1", "destination ip")
	destPort := flag.String("p", "5000", "destination port")
	bindip := flag.String("b", "127.0.0.1", "binding interface")
	bindPortTx := 0
	var ctx sendContext
	var err error

	flag.Parse()

	ctx.laddr, err = net.ResolveUDPAddr("udp", *bindip+":"+strconv.Itoa(bindPortTx))
	if err != nil {
		fmt.Printf("ResolveUDPAddr %v error ; %v", *bindip, err)
		return
	}

	ctx.raddr, err = net.ResolveUDPAddr("udp", *destIp+":"+*destPort)
	if err != nil {
		fmt.Printf("ResolveUDPAddr %v error ; %v", *destIp+":"+*destPort, err)
		return
	}

	fmt.Println("binding an arbitrary port to send data")
	ctx.udpc, err = net.ListenUDP("udp", nil)
	if err != nil {
		fmt.Printf("ListenUDP error : %v", err)
		return
	}

	ctx.segment = make([]byte, 2048)

	ctx.wg.Add(1)
	go send(&ctx)

	ctx.wg.Wait()

	fmt.Println("bye bye !")

}
