package main

import (
	"fmt"
	"gateway/unpack"
	"net"
)

func main()  {
	//建立tcp连接
	conn,err := net.Dial("tcp","localhost:9091")
	if err != nil {
		fmt.Printf("connect server err:%v\n",err)
		return
	}
	fmt.Printf("connect server:%v\n",conn.RemoteAddr())
	defer conn.Close()
	//发送消息
	content := "hello lk!"
	if err := unpack.Encode(conn,content);err != nil {
		fmt.Printf("send msg err:%v\n",err)
		return
	}
	fmt.Printf("send msg:%v\n",content)

	fmt.Printf("close connect")
}
