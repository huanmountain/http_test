package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"

	"http/common/message"
)

func readPkg(conn net.Conn)(mes message.Message ,err error)  {
	buf := make([]byte,8096)
	fmt.Println("等待客户端发送的数据...")
	//connet 只有在没有关闭的时候才会阻塞
	_,err = conn.Read(buf[:4])
	if err != nil{
		//fmt.Println("conn.read err =",err)
		err =errors.New("read pkg head err")
		return
	}
	fmt.Println("读到的buff =",buf[:4])
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])
	//根据pkglen读取消息内容
	n,err :=conn.Read(buf[:pkgLen])
	if uint32(n) != pkgLen || err!= nil {
		// fmt.Println("connet.read err =",err)
		//err = errors.New("read pkg body err")
		return
	}
	//把pkglen反序列化成 message.message (需要使用&)
	err = json.Unmarshal(buf[:pkgLen],&mes)
	if err != nil {
		fmt.Println("json.Unmarshal err =",err)
		return
	}
	return

}


//处理和客户端的通讯
func process(conn net.Conn)  {
	//这里需要延时关闭conn
	defer conn.Close()


	//循环读取客户端发送的信息
	for{
		//奖读取数据包封装成一个函数readPKg(返回err和Message)
		mes,err := readPkg(conn)
		if err != nil {

			if err == io.EOF {
				fmt.Println("客户端退出,服务器也退出")
				return
			}

			fmt.Println("readPkg err =",err)
		}
		fmt.Println("mes=",mes)
	}
}




func main()  {

	//提示信息
	fmt.Println("服务器在8889端口监听...")
	listen,err :=net.Listen("tcp","0.0.0.0:8889")
	defer  listen.Close()
	if err != nil {
		fmt.Println("listen err= ",err)
		return
	}
	for {
		fmt.Println("等待客户端连接服务器...")
		conn,err :=listen.Accept()
		if err != nil {
			fmt.Println("listen accept err =",err)
		}
		//连接成功,则启动一个携程和客户端保持通讯
		go process(conn)

	}
}