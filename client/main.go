package main

import (
	"fmt"
	"os"
	"strconv"

	"../gen-go/thrift/example"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/mattn/sc"
)

const (
	NetworkAddr = "127.0.0.1:9090"
)

func add(client *example.CustomerServiceClient, name string, age int) error {
	person := &example.Person{
		Name: name,
		Age:  int32(age),
	}
	return client.AddPerson(person)
}

func list(client *example.CustomerServiceClient) error {
	clist, err := client.ListPerson()
	if err != nil {
		return err
	}
	for _, v := range clist {
		fmt.Printf("name:\"%s\" age:%d\n", v.Name, v.Age)
	}
	return nil
}

func main() {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(NetworkAddr)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	useTransport := transportFactory.GetTransport(transport)
	client := example.NewCustomerServiceClientFactory(useTransport, protocolFactory)
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer transport.Close()

	(&sc.Cmds{
		{
			Name: "list",
			Desc: "list: listing person",
			Run: func(c *sc.C, args []string) error {
				return list(client)
			},
		},
		{
			Name: "add",
			Desc: "add [name] [age]: add person",
			Run: func(c *sc.C, args []string) error {
				if len(args) != 2 {
					return sc.UsageError
				}
				name := args[0]
				age, err := strconv.Atoi(args[1])
				if err != nil {
					return err
				}
				return add(client, name, age)
			},
		},
	}).Run(&sc.C{
		Desc: "thrift example",
	})
}
