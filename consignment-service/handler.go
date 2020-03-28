package main

import (
	"context"
	"github.com/pkg/errors"

	pb "github.com/whuangz/shippy-consignment-service/proto/consignment"
	vesselPb "github.com/whuangz/shippy-vessel-service/proto/vessel"
)

type handler struct {
	repository
	vesselClient vesselPb.VesselService
}

func (h *handler)CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	vesselRes, err = h.vesselClient.FindAvailable(ctx, &vesselPb.Specification){
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	}

	if vesselRes == nil {
		return errors.New("Error fetching vessel, returned nil")
	}

	if err != nil {
		return err
	}


	req.VesselId = vesselRes.Vessel.Id

	if err := h.repository.Create(ctx, MarshalConsignment(req)); err != nil {
		return err
	}

	res.Created = true
	res.Consignment = req
	return nil

}

func (h * handler)GetAll(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments, err = h.repository.GetAll(ctx)
	if err != nil {
		return nil
	}

	res.Consignments = UnmarshalConsignmentCollection(consignments)
	return nil
}
