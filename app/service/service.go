package service

import (
	"context"
	"github.com/atrybulkevychglobalgames/grpc-go-kubernetes-load-balancing-example/app/responder"
	"os"
)

type ResponderService struct {
	Node string;
}

func (rs *ResponderService) GetName(ctx context.Context,
	req *responder.GetNameRequest) (*responder.GetNameResponse, error) {
		hostName, _ := os.Hostname()
		resp := responder.GetNameResponse{
			Name: hostName,
			Value: req.Value,
		}
	return &resp, nil
}
