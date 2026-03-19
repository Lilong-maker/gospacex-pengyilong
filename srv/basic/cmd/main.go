package main

import (
	"flag"
	"fmt"
	_ "gospacex-pengyilong/srv/basic/inits"
	"gospacex-pengyilong/srv/handler/service/payment"
	"gospacex-pengyilong/srv/handler/service/product"
	__ "gospacex-pengyilong/srv/proto/goods"
	order2 "gospacex-pengyilong/srv/proto/order"
	"log"
	"net"

	"github.com/Lilong-maker/consul"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	err := consul.ConsulInit()
	if err != nil {
		fmt.Println("consul初始化失败")
	}
	fmt.Println("consul初始化成功")
	balancer, err := consul.GetServiceWithLoadBalancer("service")
	if err != nil {

		log.Printf("获取用户服务失败:%s", err)
	} else {
		log.Printf("获取用户服务成功:%s ,地址:%s:%d", balancer.Service, balancer.Address, balancer.Port)
	}
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	__.RegisterGoodsServer(s, &product.Server{})
	order2.RegisterOrderServer(s, &payment.Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	err = consul.ConsulShutdown()
	if err != nil {
		return
	}
	fmt.Println("服务已退出")
}
