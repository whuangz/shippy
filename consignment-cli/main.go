package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	micro "github.com/micro/go-micro/v2"
	pb "github.com/whuangz/shippy/consignment-service/proto/consignment"
)

const (
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {

	srv := micro.NewService(
		micro.Name("shippy.consignment.cli"),
	)

	srv.Init()
	client := pb.NewShippingService("shippy.consignment.service", srv.Client())

	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consigment, err := parseFile(file)
	if err != nil {
		log.Fatal("Could not parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consigment)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}

	log.Printf("Created: %t", r.Created)

	getAll, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatal("could not list consignments: %v", err)
	}

	for _, v := range getAll.Consignments {
		log.Println(v)
	}

}
