package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

// 算术运算请求结构体
type ArithRequest struct {
	A int
	B int
}

// 算术运算响应结构体
type ArithResponse struct {
	Pro int
	Quo int
	Rem int
}

func main() {
	conn, err := jsonrpc.Dial("tcp", "127.0.0.1:8096")
	if err != nil {
		log.Fatalln("dialing error:", err)
	}

	req := ArithRequest{9, 2}
	var res ArithResponse

	err = conn.Call("Arith.Multiply", req, &res)
	if err != nil {
		log.Fatalln("arith error:", err)
	}
	fmt.Printf("%d * %d = %d\n", req.A, req.B, res.Pro)

	err = conn.Call("Arith.Divide", req, &res)
	if err != nil {
		log.Fatalln("arith error:", err)
	}
	fmt.Printf("%d / %d, quo is %d, rem is %d\n", req.A, req.B, res.Quo, res.Rem)
}
