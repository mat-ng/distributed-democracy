package core

import (
	"context"
	"fmt"
	"os/exec"
	"strings"

	"google.golang.org/grpc"
)

type WorkerNode struct {
	conn *grpc.ClientConn
	c    NodeServiceClient
}

func (n *WorkerNode) Init() (err error) {
	n.conn, err = grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		return err
	}

	n.c = NewNodeServiceClient(n.conn)

	return nil
}

func (n *WorkerNode) Start() {
	fmt.Println("worker node running...")

	_, _ = n.c.ReportStatus(context.Background(), &Request{})

	stream, _ := n.c.AssignIssue(context.Background(), &Request{})
	for {
		res, err := stream.Recv()
		if err != nil {
			return
		}

		if res.Data != "" {
			issue := strings.TrimSpace(res.Data)

			fmt.Println("Received Issue: ", issue)
			if err := exec.Command("cmd", "/C", fmt.Sprintf("echo %s >> ./issues.txt", issue)).Run(); err != nil {
				fmt.Println(err)
			}
		}
	}
}

var workerNode *WorkerNode

func GetWorkerNode() *WorkerNode {
	if workerNode == nil {
		workerNode = &WorkerNode{}

		if err := workerNode.Init(); err != nil {
			panic(err)
		}
	}

	return workerNode
}
