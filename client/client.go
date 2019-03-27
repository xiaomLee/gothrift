package main

import (
	"net"
	"git.apache.org/thrift.git/lib/go/thrift"
	"fmt"
	"github.com/xiaomLee/gothrift/example"
	"context"
)


const (
	HOST = "localhost"
	PORT = "8080"
)

func main() {
	tSocket, err := thrift.NewTSocket(net.JoinHostPort(HOST, PORT))
	if err != nil {
		fmt.Println("tSocket error:", err)
	}
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	transport, _ := transportFactory.GetTransport(tSocket)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	client := example.NewFormatDataClientFactory(transport, protocolFactory)

	if err := transport.Open(); err != nil {
		fmt.Println("Error opening:", HOST + ":" + PORT)
	}
	defer transport.Close()



	res, err := client.Ping(context.Background())
	fmt.Println("ping result: ", res)
	data := example.Data{Text:"hello,world!"}
	d, err := client.DoFormat(context.Background(), &data)
	fmt.Println("doformat result: ", d.Text)

}
