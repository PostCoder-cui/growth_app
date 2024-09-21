package ugserver

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
	"user_growth/comm"
	"user_growth/models"
	"user_growth/pb"
	"user_growth/service"
)

type CoinServer struct {
	pb.UnimplementedUserCoinServer
}

// 获取所有积分任务列表
func (c *CoinServer) ListTasks(ctx context.Context, request *pb.ListTasksRequest) (*pb.ListTasksReply, error) {
	//return nil, status.Errorf(codes.Unimplemented, "method ListTasks not implemented")

	coinTaskService := service.NewCoinTaskService(ctx)
	tasks, err := coinTaskService.FindAll()
	if err != nil {
		log.Printf("ListTasks has an error:%v", err)
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	resList := make([]*pb.TbCoinTask, len(tasks))
	for i, task := range tasks {
		resList[i] = models.CoinTaskToMessage(&task)
	}

	out := &pb.ListTasksReply{
		DataList: resList,
	}

	return out, nil
}

// 获取用户积分
func (c *CoinServer) UserCoinInfo(ctx context.Context, request *pb.UserCoinInfoRequest) (*pb.UserCoinInfoReply, error) {
	userService := service.NewCoinUserService(ctx)
	id := int(request.Uid)
	coinInfo, err := userService.GetByUid(id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}
	d := models.CoinUserToMessage(coinInfo)
	out := &pb.UserCoinInfoReply{
		Data: d,
	}

	return out, nil
}

// 获取用户积分明细
func (c *CoinServer) UserCoinDetails(ctx context.Context, request *pb.UserCoinDetailsRequest) (*pb.UserCoinDetailsReply, error) {
	uid := int(request.Uid)
	page := int(request.Page)
	size := int(request.Size)
	detailService := service.NewCoinDetailService(ctx)
	dataList, total, err := detailService.FindByUid(uid, page, size)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", comm.MarkLineErr(err))
	}
	resList := make([]*pb.TbCoinDetail, len(dataList))
	for i, data := range dataList {
		resList[i] = models.CoinDetailToMessage(&data)
	}

	return &pb.UserCoinDetailsReply{DataList: resList, Total: int32(total)}, nil
}

// 调整积分
func (c *CoinServer) UserCoinChange(ctx context.Context, request *pb.UserCoinChangeRequest) (*pb.UserCoinChangeReply, error) {
	//TODO implement me
	fmt.Printf("request: %v\n", request)
	uid := int(request.Uid)
	task := request.TaskName
	coin := int(request.Coin)
	coinTask, err := service.NewCoinTaskService(ctx).GetByTask(task)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", comm.MarkLineErr(err))
	}
	if coinTask == nil {
		log.Printf("CoinTask is nil!%v", request)
		return nil, status.Errorf(codes.NotFound, "%v====请求是：%v", comm.MarkLineErr("任务不存在"), *request)
	}
	log.Printf("CoinTask is:%v", coinTask)
	//新增一组积分详情数据
	if coin == 0 {
		coin = coinTask.Coin
	}
	coinDetail := models.TbCoinDetail{
		Uid:        uid,
		TaskId:     coinTask.Id,
		Coin:       coin,
		SysCreated: time.Time{},
		SysUpdated: nil,
	}
	err = service.NewCoinDetailService(ctx).Save(&coinDetail)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", comm.MarkLineErr(err))
	}

	//更新用户信息
	coinUserSrv := service.NewCoinUserService(ctx)
	userCoin, err := coinUserSrv.GetByUid(uid)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", comm.MarkLineErr(err))
	}
	if userCoin == nil {
		userCoin = &models.TbCoinUser{
			Uid:   uid,
			Coins: coin,
		}
	} else {
		userCoin.Coins = coin
		//userCoin.SysUpdated = time.Now()
	}

	return &pb.UserCoinChangeReply{
		User: models.CoinUserToMessage(userCoin),
	}, nil
}
