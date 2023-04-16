package client

import (
	"context"
	"fmt"

	"github.com/fazilnbr/GoCart-grpc-cart-Service/pkg/config"
	"github.com/fazilnbr/GoCart-grpc-cart-Service/pkg/pb"
	"google.golang.org/grpc"
)

type ProductServiceClient struct {
	Client pb.ProductServiceClient
}

func InitProductServiceClient(cfg config.Config) ProductServiceClient {
	cc, err := grpc.Dial(cfg.ProductSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	c := ProductServiceClient{
		Client: pb.NewProductServiceClient(cc),
	}

	return c
}

func (c *ProductServiceClient) GetProduct(productId int64) (*pb.GetProductResponse, error) {
	req := &pb.GetProductRequest{
		Id: productId,
	}

	return c.Client.GetProduct(context.Background(), req)
}
