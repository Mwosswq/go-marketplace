package items

import (
	"context"

	pb "github.com/marketplace-go/contracts/items"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	client pb.ItemServiceClient
	conn   *grpc.ClientConn
}

func New(addr string) *Client {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("error while getting conn")
	}

	c := pb.NewItemServiceClient(conn)

	return &Client{client: c, conn: conn}
}

func (c *Client) Close() {
	_ = c.conn.Close()
}

func (c *Client) GetItem(ctx context.Context, id int32) (*pb.GetItemResponse, error) {
	item, err := c.client.GetItem(ctx, &pb.GetItemRequest{Id: id})

	if err != nil {
		return &pb.GetItemResponse{}, err
	}

	return item, nil
}

func (c *Client) CreateItem(ctx context.Context, item *pb.CreateItemRequest) (*pb.CreateItemResponse, error) {
	itemResponse, err := c.client.CreateItem(ctx, item)

	if err != nil {
		return &pb.CreateItemResponse{}, err
	}

	return itemResponse, nil
}
