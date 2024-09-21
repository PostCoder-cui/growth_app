package service

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"time"
	"user_growth/conf"
	"user_growth/dbhelper"
	"user_growth/models"
)

func initDb() {
	time.Local = time.UTC
	conf.LoadConfig()
	dbhelper.InitDb()
}

func TestCoinTaskService_Save2(t *testing.T) {
	initDb()
	s := NewCoinTaskService(context.Background())
	data := models.TbCoinTask{Id: 0, Task: "drink", Coin: 12, Limit: 100} //  <0更新，>0新增
	if err := s.Save(&data); err != nil {
		t.Errorf("Save(%v) err:%v)", data, err)
	} else {
		t.Logf("Save(%v) success", data)
	}
}

func TestCoinUserService_FindAllPager(t *testing.T) {
	initDb()
	s := NewCoinTaskService(context.Background())
	if all, err := s.FindAll(); err != nil {
		t.Errorf("FindAll() err:%v", err)
	} else {
		t.Logf("FindAll(%v) success", all)
	}
}

func TestGradeInfoService_Save(t *testing.T) {
	initDb()
	s := NewGradeInfoService(context.Background())
	data := models.TbGradeInfo{
		Id:          0,
		Title:       "中级",
		Description: "中级用户",
		Score:       1,
		Expired:     0,
	}
	if err := s.Save(&data); err != nil {
		t.Errorf("Save(%v) err:%v)", data, err)
	} else {
		t.Logf("Save(%v) success", data)
	}
}

func TestGradeInfoService_Get(t *testing.T) {
	initDb()
	s := NewGradeInfoService(context.Background())
	if data, err := s.Get(1); err != nil {
		t.Errorf("Get(1) err:%v", err)
	} else {
		t.Logf("Get(%v) success", data)
	}
}

func TestGradeInfoService_FindAll(t *testing.T) {
	initDb()
	s := NewGradeInfoService(context.Background())
	if all, err := s.FindAll(); err != nil {
		t.Errorf("FindAll() err:%v", err)
	} else {
		t.Logf("FindAll(%v) success", all)
	}
}

func TestGradeInfoService_NowGrade(t *testing.T) {
	initDb()
	s := NewGradeInfoService(context.Background())
	if data, err := s.NowGrade(0); err != nil {
		t.Errorf("Get(1) err:%v", err)
	} else {
		t.Logf("Get(%v) success", data)
	}
}
