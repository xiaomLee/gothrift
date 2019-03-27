package main

import (
	"github.com/xiaomLee/gothrift/example"
	"fmt"
	"strings"
	"git.apache.org/thrift.git/lib/go/thrift"
	"context"
	"time"
)

type FormatDataImpl struct {
}

const (
	HOST = "localhost"
	PORT = "8080"
)


func (fdi *FormatDataImpl) Ping(ctx context.Context) (r string, err error) {
	fmt.Printf("ctx:%+v \n", ctx)
	fmt.Println("ping", time.Now())
	res := time.Now().String() + "hello world"
	return res, nil
}

func (fdi *FormatDataImpl) DoFormat(ctx context.Context, data *example.Data) (r *example.Data, err error) {
	fmt.Println("context: ", ctx)
	fmt.Println("input data: ", data.Text)
	var res example.Data
	res.Text = strings.ToUpper(data.Text)
	return &res, nil
}

func main() {
	handler := &FormatDataImpl{}
	processor := example.NewFormatDataProcessor(handler)
	serverTransport, err := thrift.NewTServerSocket(HOST + ":" + PORT)
	if err != nil {
		fmt.Println("Error:", err)
	}
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("Running at:", HOST + ":" + PORT)
	server.Serve()
}