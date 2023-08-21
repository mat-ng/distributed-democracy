package core

import (
	"context"
)

type NodeServiceGrpcServer struct {
	UnimplementedNodeServiceServer

	Issue chan string
}

func (n NodeServiceGrpcServer) ReportStatus(ctx context.Context, request *Request) (*Response, error) {
	return &Response{Data: "ok"}, nil
}

func (n NodeServiceGrpcServer) AssignIssue(request *Request, server NodeService_AssignIssueServer) error {
	for {
		select {
		case issue := <-n.Issue:
			if err := server.Send(&Response{Data: issue}); err != nil {
				return err
			}
		}
	}
}

var server *NodeServiceGrpcServer

func GetNodeServiceGrpcServer() *NodeServiceGrpcServer {
	if server == nil {
		server = &NodeServiceGrpcServer{
			Issue: make(chan string),
		}
	}
	return server
}
