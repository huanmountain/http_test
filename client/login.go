package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"time"

	"http/common/message"
)

func login(userId  int ,userPwd string) (err error) {
	//连接到服务器
	conn,err := net.Dial("tcp","localhost:8889")
	if err != nil {
		fmt.Println("net.dial err=",err)
		return
	}
	defer conn.Close()
	var mes message.Message
	mes.Type =message.LoginMesType
	var loginMes message.LoginMes
	loginMes.UserId =userId
	loginMes.UserPwd= userPwd


	//mes.Data =loginMes??? 实现需要序列化
	//将loginMes(包括userid和userPwd) 序列化
	data,err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println( "json.Marshal err =",err)
		return err
	}
	//把data赋值给mes.data
	mes.Date= string(data)

	//将mes序列化
	data,err = json.Marshal(mes)
	if err != nil {
		fmt.Println( "json.Marshal err =",err)
		return err
	}

	//data就是要发送的数据,
	//为了防止丢包,先把data长度发送给服务器
	//先把data长度发送给服务器,先获取到data的长度-->转成byte长度的切片(write发送的是byte切片)
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4],pkgLen)
	//发送长度
	n,err :=conn.Write(buf[0:4])
	if n!=4 || err !=nil {
		fmt.Println("conn.weite(bytes) err =" ,err)
		return
	}
	//fmt.Printf("client sent pkgLeb seccuss,len= %d,context = %s",len(data),string(data))


	//发送消息本身
	_,err =conn.Write(data)
	if  err !=nil {
		fmt.Println("conn.weite(data) err =" ,err)
		return
	}
	time.Sleep(20*time.Second)
	fmt.Println("休眠20s")
	return
}