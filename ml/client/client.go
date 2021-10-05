package client

import (
	"google.golang.org/grpc"
	pb "seim/ml/client/pb"
)

type Client struct {
	conn    *grpc.ClientConn
	Service pb.ImageServiceClient
}

func NewClient(url string) (*Client, error) {
	// client connectionを生成
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	// articleサービスのclientを生成
	c := pb.NewImageServiceClient(conn)

	// articleサービスのclientを返す
	return &Client{conn, c}, nil
}

func (c *Client) Close() {
	c.conn.Close()
}
