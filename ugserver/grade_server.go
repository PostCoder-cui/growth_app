package ugserver

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
	"user_growth/comm"
	"user_growth/models"
	"user_growth/pb"
	"user_growth/service"
)

type GradeServer struct {
	pb.UnimplementedUserGradeServer
}

// 获取所有等级信息列表
func (g *GradeServer) ListGrades(ctx context.Context, request *pb.ListGradesRequest) (*pb.ListGradesReply, error) {
	//TODO implement me
	gradeSrv := service.NewGradeInfoService(ctx)
	gradeInfos, err := gradeSrv.FindAll()
	if err != nil {
		return nil, status.Errorf(codes.Unimplemented, comm.MarkLineErr(err))
	}
	resList := make([]*pb.TbGradeInfo, len(gradeInfos))
	for i, gradeInfo := range gradeInfos {
		resList[i] = models.GradeInfoToMessage(&gradeInfo)
	}
	return &pb.ListGradesReply{
		DataList: resList,
	}, nil
}

// 获取等级特权列表
func (g *GradeServer) ListGradePrivilege(ctx context.Context, request *pb.ListGradePrivilegeRequest) (*pb.ListGradePrivilegeReply, error) {
	//TODO implement me
	gradeId := int(request.GradeId)
	privilegeService := service.NewGradePrivilegeService(ctx)
	var dataList []models.TbGradePrivilege
	var err error
	if gradeId > 0 {
		dataList, err = privilegeService.FindByGrade(gradeId)
	} else {
		dataList, err = privilegeService.FindAll()
	}
	if err != nil {
		return nil, status.Errorf(codes.Unimplemented, comm.MarkLineErr(err))
	}

	var resList []*pb.TbGradePrivilege
	for _, data := range dataList {
		resList = append(resList, models.GradePrivilegeToMessage(&data))
	}

	return &pb.ListGradePrivilegeReply{
		DataList: resList,
	}, nil
}

// 检查用户是否有某产品的特权
func (g *GradeServer) CheckUserPrivilege(ctx context.Context, request *pb.CheckUserPrivilegeRequest) (*pb.CheckUserPrivilegeReply, error) {
	//TODO implement me
	uid := int(request.Uid)
	product := request.Product
	funct := request.Function
	privilegeService := service.NewGradePrivilegeService(ctx)
	userSrv := service.NewGradeUserService(ctx)
	user, err := userSrv.GetByUid(uid)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, comm.MarkLineErr(err))
	}
	privilegeList, err := privilegeService.FindByGrade(user.GradeId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, comm.MarkLineErr(err))
	}

	var isOk bool
	for _, v := range privilegeList {
		if v.Product == product && v.Function == funct {
			isOk = true
			break
		}
	}

	return &pb.CheckUserPrivilegeReply{
		Data: isOk,
	}, nil
}

// 查询用户等级信息
func (g *GradeServer) UserGradeInfo(ctx context.Context, request *pb.UserGradeInfoRequest) (*pb.UserGradeInfoReply, error) {
	//TODO implement me
	uid := int(request.Uid)
	userService := service.NewGradeUserService(ctx)
	infoService := service.NewGradeInfoService(ctx)
	UserGrade, err := userService.GetByUid(uid)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, comm.MarkLineErr(err))
	}
	gradeInfo, err2 := infoService.Get(UserGrade.GradeId)
	if err2 != nil {
		return nil, status.Errorf(codes.InvalidArgument, comm.MarkLineErr(err2))
	}

	data := models.GradeInfoToMessage(gradeInfo)

	return &pb.UserGradeInfoReply{
		Data: data,
	}, nil
}

func (g *GradeServer) UserGradeChange(ctx context.Context, request *pb.UserGradeChangeRequest) (*pb.UserGradeChangeReply, error) {
	//TODO implement me
	uid := int(request.Uid)
	score := int(request.Score)
	userService := service.NewGradeUserService(ctx)
	infoService := service.NewGradeInfoService(ctx)
	user, err := userService.GetByUid(uid)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, comm.MarkLineErr(err))
	}
	if user == nil {
		user = &models.TbGradeUser{
			Uid: uid,
		}
	}
	user.Score += score
	grade, err := infoService.NowGrade(user.Score)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, comm.MarkLineErr(err))
	}
	newData := models.TbGradeUser{
		Id:         user.Id,
		Uid:        0,
		GradeId:    0,
		Expired:    time.Time{},
		Score:      user.Score,
		SysCreated: time.Time{},
		SysUpdated: nil,
	}

	//等级更新
	if user.GradeId != grade.Id {
		newData.GradeId = grade.Id
		exp := comm.Now().AddDate(10, 0, 0)
		if grade.Expired > 0 {
			newData.Expired = exp
		}
	}
	err = userService.Save(&newData)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, comm.MarkLineErr(err))
	}

	return &pb.UserGradeChangeReply{
		Data: models.GradeUserToMessage(&newData),
	}, nil
}
