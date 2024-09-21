package dbhelper

import (
	"fmt"
	"log"
	"time"
	"user_growth/conf"
	"xorm.io/xorm"
)

var dbEngine *xorm.Engine

// 初始化数据库
func InitDb() {
	if dbEngine != nil {
		return
	}
	sourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		conf.GlobalConfig.Db.User,
		conf.GlobalConfig.Db.Password,
		conf.GlobalConfig.Db.Host,
		conf.GlobalConfig.Db.Port,
		conf.GlobalConfig.Db.Database,
		conf.GlobalConfig.Db.Charset,
	)
	log.Printf("mysql sourceName: %s", sourceName)
	if engine, err := xorm.NewEngine(conf.GlobalConfig.Db.Engine, sourceName); err != nil {
		log.Printf("%s调用引擎连接失败", sourceName)
		log.Fatalf("dbhelper.InitEngine(%s) error:%v", sourceName, err)
		return
	} else {
		dbEngine = engine
	}

	//其他配置
	if conf.GlobalConfig.Db.MaxIdle > 0 {
		dbEngine.SetMaxIdleConns(conf.GlobalConfig.Db.MaxIdle)
	}
	if conf.GlobalConfig.Db.MaxConns > 0 {
		dbEngine.SetMaxOpenConns(conf.GlobalConfig.Db.MaxConns)
	}
	if conf.GlobalConfig.Db.ConnnMaxLifetime > 0 {
		dbEngine.SetConnMaxLifetime(time.Minute * time.Duration(conf.GlobalConfig.Db.ConnnMaxLifetime))
	}
}

func GetDbEngine() *xorm.Engine {
	return dbEngine
}
