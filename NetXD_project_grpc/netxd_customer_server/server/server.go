package main

import (
	"context"
	"fmt"
	"net"

	pro "github.com/kishorens18/NetXD_Project/Netxd_Customer"
	"github.com/kishorens18/NetXD_Project/netxd_customer_config/config"
	"github.com/kishorens18/NetXD_Project/netxd_customer_config/constants"
	"github.com/kishorens18/NetXD_Project/netxd_customer_dal/services"
	"github.com/kishorens18/NetXD_Project/netxd_customer_server/controllers"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func initDatabase(client *mongo.Client) {
	customerCollection := config.GetCollection(client, "profiles")
	controllers.CustomerService = services.InitCustomerService(customerCollection, context.Background())
}

func main() {
	mongoclient, err := config.ConnectDatabase()
	defer mongoclient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	initDatabase(mongoclient)
	lis, err := net.Listen("tcp", constants.Port)
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	pro.RegisterProfileServiceServer(s, &controllers.RPCserver{})

	fmt.Println("Server listening on", constants.Port)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}
}
