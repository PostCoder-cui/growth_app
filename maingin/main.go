package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"user_growth/pb"
)

// 设置域名白名单
//var AllowOrigin = map[string]bool{
//	"http://www.google.com": true,
//}

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//新建grpcClient客户端
	addr := flag.String("addr", "0.0.0.0:8085", "server address")
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("new grpc client error: %v", err)
	}
	defer conn.Close()

	//建立相应的客户端pb
	coinClient := pb.NewUserCoinClient(conn)
	gradeClient := pb.NewUserGradeClient(conn)

	//设置跨域CORS
	originRouter := router.Group("/v1", func(context *gin.Context) {
		//origin := context.GetHeader("origin")
		//if AllowOrigin[origin] == true {
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("access-control-allow-methods", "GET,POST,PUT,DELETE,OPTIONS")
		context.Header("access-control-allow-headers", "*")
		context.Header("Access-Control-Allow-Credentials", "true")
		//}
		context.Next()
	})

	//用户积分服务的方法
	gUserCoin := originRouter.Group("/UserGrowth.UserCoin")
	gUserCoin.GET("/ListTasks", func(c *gin.Context) {
		//grpc pb客户端调用处理
		listTasks, err2 := coinClient.ListTasks(c, &pb.ListTasksRequest{})
		if err2 != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err2.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"listTasks": listTasks,
			})
		}

	})
	gUserCoin.POST("/UserCoinChange", func(c *gin.Context) {
		body := &pb.UserCoinChangeRequest{}
		if err := c.ShouldBind(body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}
		reply, err2 := coinClient.UserCoinChange(c, body)
		if err2 != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err2.Error(),
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"reply": reply,
		})

	})

	//用户等级服务的方法
	gUserGrade := router.Group("/v1/UserGrowth.UserGrade")
	gUserGrade.GET("/ListGrades", func(c *gin.Context) {
		gradeclient, err2 := gradeClient.ListGrades(c, &pb.ListGradesRequest{})
		if err2 != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err2.Error(),
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"grades": gradeclient,
		})
	})

	//封装成http2处理器
	h2Handler := h2c.NewHandler(router, &http2.Server{})

	//构建server服务
	server := &http.Server{
		Addr:    ":8080",
		Handler: h2Handler,
	}

	//启动服务
	server.ListenAndServe()

}
