package main

import (
	"context"

	pb "github.com/whuangz/shippy/vessel-service/proto/vessel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository interface {
	FindAvailable(ctx context.Context, spec *pb.Specification) (*pb.Vessel, error)
	Create(ctx context.Context, vessel *Vessel) error
}

type VesselRepository struct {
	collection *mongo.Collection
}

func (repository *VesselRepository) FindAvailable(ctx context.Context, spec *pb.Specification) (*pb.Vessel, error) {
	filter := bson.D{{
		"capacity",
		bson.D{{
			"$lte",
			spec.Capacity,
		}, {
			"$lte",
			spec.MaxWeight,
		}},
	}}
	vessel := &pb.Vessel{}
	if err := repository.collection.FindOne(ctx, filter).Decode(vessel); err != nil {
		return nil, err
	}
	return vessel, nil
}
