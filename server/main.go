package main

import (
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
	"user_growth/conf"
	"user_growth/dbhelper"
	"user_growth/pb"
	"user_growth/ugserver"
)

func initDb() {
	time.Local = time.UTC
	conf.LoadConfig()
	dbhelper.InitDb()
}

func main() {
	initDb()
	//服务端监听（占用）一个服务端口
	lsr, err := net.Listen("tcp", "0.0.0.0:8085")
	if err != nil {
		log.Fatalf("failed to listen: %v", err.Error())
	}
	defer lsr.Close()
	//创建grpc服务
	my_server := grpc.NewServer()

	//注册服务
	pb.RegisterUserCoinServer(my_server, &ugserver.CoinServer{})
	pb.RegisterUserGradeServer(my_server, &ugserver.GradeServer{})

	//反射grpc注册
	reflection.Register(my_server)

	log.Printf("server listening at %v", lsr.Addr())
	if err := my_server.Serve(lsr); err != nil { //写到if里面err:=也能用不报错了
		log.Fatalf("failed to serve: %v", err)
	}
}
