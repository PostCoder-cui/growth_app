package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
	"user_growth/pb"
)

func main() {
	//链接到服务器
	addr := flag.String("addr", "192.168.137.151:8085", "server address")
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to connect server: %v", err)
	}
	defer conn.Close()

	//创建客户端
	coinClient := pb.NewUserCoinClient(conn)
	gradeClient := pb.NewUserGradeClient(conn)

	//测试
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*3)
	defer cancelFunc()

	if tasks, err := coinClient.ListTasks(ctx, &pb.ListTasksRequest{}); err != nil {
		log.Printf("fail to list tasks: %v", err)
	} else {
		log.Printf("list tasks: %v", tasks.GetDataList())
	}

	if grades, err := gradeClient.ListGrades(ctx, &pb.ListGradesRequest{}); err != nil {
		log.Printf("fail to list grades: %v", err)
	} else {
		log.Printf("grades: %v", grades.GetDataList())
	}

}
