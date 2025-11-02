package items

import (
	"context"

	pb "github.com/marketplace-go/contracts/items"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	client pb.ItemServiceClient
}

func New(addr string) *Client {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("error while getting conn")
	}

	c := pb.NewItemServiceClient(conn)

	return &Client{client: c}
}

func (c *Client) GetItem(ctx context.Context, id int32) (*pb.GetItemResponse, error) {
	item, err := c.client.GetItem(ctx, &pb.GetItemRequest{Id: id})
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (c *Client) CreateItem(ctx context.Context, item *pb.CreateItemRequest) (*pb.CreateItemResponse, error) {
	itemResponse, err := c.client.CreateItem(ctx, item)
	if err != nil {
		return nil, err
	}

	return itemResponse, nil
}
