package handler

import (
	"context"

	micro "github.com/micro/go-micro/v2"
	pb "github.com/whuangz/shippy/vessel-service/proto/vessel"
)

type handler struct {
	repository
}

func NewHandler(client *mongo.Client) *handler {

	vesselCollection := client.Database("shippy").Collection("vessel")

	return &handler{
		&VesselRepository{
			vesselCollection
		}
	}
}


func (h *handler)FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	vessel, err := h.repository.FindAvailable(ctx, req)
	if err != nil {
		return err
	}

	// Set the vessel as part of the response message type
	res.Vessel = vessel
	return nil
}