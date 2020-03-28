package main

import (
	"context"
	"log"
	"os"

	micro "github.com/micro/go-micro/v2"
	pb "github.com/whuangz/shippy/consignment-service/proto/consignment"
	vesselPb "github.com/whuangz/shippy/vessel-service/proto/vessel"
)

const (
	defaultHost = "datastore:27017"
)

func main() {
	srv := micro.NewService(
		micro.Name("shippy.consignment.service"),
	)

	srv.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}

	client, err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())

	consignmentCollection := client.Database("shippy").Collection("consignments")

	repository := &MongoRepository(consignmentCollection)
	vesselClient := vesselPb.NewVesselService("shippy.vessel.service", srv.Client())

	h := &handler{repository, vesselClient}

	pb.RegisterShippingServiceHandler(srv.Server(), h)

	if err := srv.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
