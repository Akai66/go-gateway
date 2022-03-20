package main

import (
	"fmt"
	"gateway/unpack"
	"net"
)

func main()  {
	//simple tcp server
	//listen port
	listener,err := net.Listen("tcp","0.0.0.0:9091")
	if err != nil {
		fmt.Printf("listen port err:%v\n",err)
		return
	}

	//accept conn
	for {
		conn,err := listener.Accept()
		if err != nil {
			fmt.Printf("accept err:%v\n",err)
			continue
		}
		fmt.Printf("connect client:%v\n",conn.RemoteAddr())
		//goroutine read and write conn
		go process(conn)
	}
}

func process(conn net.Conn)  {
	defer conn.Close()
	//cycle read
	for {
		body,err := unpack.Decode(conn)
		if err != nil {
			fmt.Printf("read err:%v\n",err)
			break
		}
		fmt.Printf("read data:%v\n",string(body))
	}
}
