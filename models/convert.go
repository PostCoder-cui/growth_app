package models

import (
	"time"
	"user_growth/comm"
	"user_growth/pb"
)

var emptyTime time.Time

func CoinTaskToMessage(in *TbCoinTask) *pb.TbCoinTask {
	d := &pb.TbCoinTask{
		Id:         int32(in.Id),
		Task:       in.Task,
		Coin:       int32(in.Coin),
		Limit:      int32(in.Limit),
		Start:      comm.TimeFormat(in.Start),
		SysCreated: comm.TimeFormat(&in.SysCreated),
		SysUpdated: comm.TimeFormat(in.SysUpdated),
		SysStatus:  int32(in.SysStatus),
	}
	return d
}

// CoinTaskToObject message转model
func CoinTaskToObject(data *pb.TbCoinTask) *TbCoinTask {
	d := &TbCoinTask{
		Id:         int(data.Id),
		Task:       data.Task,
		Coin:       int(data.Coin),
		Limit:      int(data.Limit),
		Start:      nil,
		SysCreated: emptyTime,
		SysUpdated: nil,
		SysStatus:  int(data.SysStatus),
	}
	*d.Start = comm.TimeParse(data.Start)
	d.SysCreated = comm.TimeParse(data.SysCreated)
	*d.SysUpdated = comm.TimeParse(data.SysUpdated)
	return d
}

// CoinUserToMessage model转message
func CoinUserToMessage(data *TbCoinUser) *pb.TbCoinUser {
	d := &pb.TbCoinUser{
		Id:         int32(data.Id),
		Uid:        int32(data.Uid),
		Coins:      int32(data.Coins),
		SysCreated: "",
		SysUpdated: "",
	}
	d.SysCreated = comm.TimeFormat(&data.SysCreated)
	if data.SysUpdated != nil {
		d.SysUpdated = comm.TimeFormat(data.SysUpdated)
	}
	return d
}

// CoinUserToObject message转model
func CoinUserToObject(data *pb.TbCoinUser) *TbCoinUser {
	d := &TbCoinUser{
		Id:         int(data.Id),
		Uid:        int(data.Uid),
		Coins:      int(data.Coins),
		SysCreated: emptyTime,
		SysUpdated: nil,
	}
	d.SysCreated = comm.TimeParse(data.SysCreated)
	*d.SysUpdated = comm.TimeParse(data.SysUpdated)
	return d
}

// CoinDetailToMessage model转message
func CoinDetailToMessage(data *TbCoinDetail) *pb.TbCoinDetail {
	d := &pb.TbCoinDetail{
		Id:         int32(data.Id),
		Uid:        int32(data.Uid),
		TaskId:     int32(data.TaskId),
		Coin:       int32(data.Coin),
		SysCreated: "",
		SysUpdate:  "",
	}
	d.SysCreated = comm.TimeFormat(&data.SysCreated)
	d.SysUpdate = comm.TimeFormat(data.SysUpdated)
	return d
}

// CoinDetailToObject message转model
func CoinDetailToObject(data *pb.TbCoinDetail) *TbCoinDetail {
	d := &TbCoinDetail{
		Id:         int(data.Id),
		Uid:        int(data.Uid),
		TaskId:     int(data.TaskId),
		Coin:       int(data.Coin),
		SysCreated: emptyTime,
		SysUpdated: nil,
	}
	d.SysCreated = comm.TimeParse(data.SysCreated)
	*d.SysUpdated = comm.TimeParse(data.SysUpdate)
	return d
}

// GradeInfoToMessage model转message
func GradeInfoToMessage(data *TbGradeInfo) *pb.TbGradeInfo {
	d := &pb.TbGradeInfo{
		Id:          int32(data.Id),
		Title:       data.Title,
		Description: data.Description,
		Score:       int32(data.Score),
		Expired:     int32(data.Expired),
		SysCreated:  "",
		SysUpdated:  "",
	}
	d.SysCreated = comm.TimeFormat(&data.SysCreated)
	d.SysUpdated = comm.TimeFormat(data.SysUpdated)
	return d
}

// GradeInfoToObject message转model
func GradeInfoToObject(data *pb.TbGradeInfo) *TbGradeInfo {
	d := &TbGradeInfo{
		Id:          int(data.Id),
		Title:       data.Title,
		Description: data.Description,
		Score:       int(data.Score),
		Expired:     int(data.Expired),
		SysCreated:  emptyTime,
		SysUpdated:  nil,
	}
	d.SysCreated = comm.TimeParse(data.SysCreated)
	*d.SysUpdated = comm.TimeParse(data.SysUpdated)
	return d
}

// GradePrivilegeToMessage model转message
func GradePrivilegeToMessage(data *TbGradePrivilege) *pb.TbGradePrivilege {
	d := &pb.TbGradePrivilege{
		Id:          int32(data.Id),
		GradeId:     int32(data.GradeId),
		Product:     data.Product,
		Function:    data.Function,
		Description: data.Description,
		Expired:     int32(data.Expired),
		SysCreated:  "",
		SysUpdated:  "",
		SysStatus:   int32(data.SysStatus),
	}
	d.SysCreated = comm.TimeFormat(&data.SysCreated)
	d.SysUpdated = comm.TimeFormat(data.SysUpdated)
	return d
}

// GradePrivilegeToObject message转model
func GradePrivilegeToObject(data *pb.TbGradePrivilege) *TbGradePrivilege {
	d := &TbGradePrivilege{
		Id:          int(data.Id),
		GradeId:     int(data.GradeId),
		Product:     data.Product,
		Function:    data.Function,
		Description: data.Description,
		Expired:     int(data.Expired),
		SysCreated:  emptyTime,
		SysUpdated:  nil,
		SysStatus:   int(data.SysStatus),
	}
	d.SysCreated = comm.TimeParse(data.SysCreated)
	*d.SysUpdated = comm.TimeParse(data.SysUpdated)
	return d
}

// GradeUserToMessage model转message
func GradeUserToMessage(data *TbGradeUser) *pb.TbGradeUser {
	d := &pb.TbGradeUser{
		Id:         int32(data.Id),
		Uid:        int32(data.Uid),
		GradeId:    int32(data.GradeId),
		Expired:    "",
		Score:      int32(data.Score),
		SysCreated: "",
		SysUpdated: "",
	}
	d.Expired = comm.TimeFormat(&data.Expired)
	d.SysCreated = comm.TimeFormat(&data.SysCreated)
	d.SysUpdated = comm.TimeFormat(data.SysUpdated)
	return d
}

// GradeUserToObject message转model
func GradeUserToObject(data *pb.TbGradeUser) *TbGradeUser {
	d := &TbGradeUser{
		Id:         int(data.Id),
		Uid:        int(data.Uid),
		GradeId:    int(data.GradeId),
		Expired:    emptyTime,
		Score:      int(data.Score),
		SysCreated: emptyTime,
		SysUpdated: nil,
	}
	d.Expired = comm.TimeParse(data.Expired)
	d.SysCreated = comm.TimeParse(data.SysCreated)
	*d.SysUpdated = comm.TimeParse(data.SysUpdated)
	return d
}
