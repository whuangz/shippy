package main

import (
	"context"
	"errors"
	"log"

	micro "github.com/micro/go-micro/v2"
	pb "github.com/whuangz/shippy/vessel-service/proto/vessel"
)

type repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
}

type VesselRepository struct {
	vessels []*pb.Vessel
}

func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	for _, vessel := range repo.vessels {
		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}
	return nil, errors.New("No vessel found by that spec")
}

//grpc service handler
type service struct {
	repo repository
}

//matching the proto
func (s *service) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	vessel, err := s.repo.FindAvailable(req)
	if err != nil {
		return err
	}
	res.Vessel = vessel
	return nil
}

func main() {

	vessels := []*pb.Vessel{
		&pb.Vessel{Id: "vessel001", Name: "John Don Boat", MaxWeight: 100000, Capacity: 500},
	}

	repo := &VesselRepository{vessels}

	srv := micro.NewService(
		micro.Name("shippy.vessel.service"),
	)

	srv.Init()

	pb.RegisterVesselServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
