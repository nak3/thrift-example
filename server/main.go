package main

import (
	"fmt"
	"os"
	"sync"

	"../gen-go/thrift/example"
	"git.apache.org/thrift.git/lib/go/thrift"
)

const (
	NetworkAddr = "127.0.0.1:9090"
)

type customerService struct {
	customers []*example.Person
	m         sync.Mutex
}

func (cs *customerService) ListPerson() ([]*example.Person, error) {
	cs.m.Lock()
	defer cs.m.Unlock()
	return cs.customers, nil
}

func (cs *customerService) AddPerson(p *example.Person) (err error) {
	cs.m.Lock()
	defer cs.m.Unlock()
	cs.customers = append(cs.customers, p)
	return nil
}

func main() {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	serverTransport, err := thrift.NewTServerSocket(NetworkAddr)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	handler := &customerService{}
	processor := example.NewCustomerServiceProcessor(handler)

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("thrift server in", NetworkAddr)
	err = server.Serve()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
