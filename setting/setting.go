package setting

import (
	"github.com/fatih/color"
	"github.com/go-ini/ini"
	"log"
)

type App struct {
	Url        string
	ApiKey     string
	SecreteKey string
}

var AppSetting = &App{}

type DataBase struct {
	Type     string // 数据库类型
	User     string // 用户
	PassWord string // 密码
	Host     string // 数据库地址+端口号
	Name     string // 数据库名字
}

var DataBaseSetting = &DataBase{}

func SetUp() {
	Cfg, err := ini.Load("app.ini") //加载配置文件ini
	if err != nil {
		color.Red("ini load err:", err)
	}
	//映射配置
	err = Cfg.Section("app").MapTo(AppSetting)
	if err != nil {
		log.Fatalf("Cfg配置文件映射 AppSetting 错误: %v", err)
	}
	err = Cfg.Section("database").MapTo(DataBaseSetting)
	if err != nil {
		log.Fatalf("Cfg配置文件映射 DatabaseSetting 错误: %v", err)
	}
}
