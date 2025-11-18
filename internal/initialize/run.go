package initialize

import (
	"fmt"
	"go-ecommerce-backend-api/global"

	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	m := global.Config.Mysql
	fmt.Println("Loading configuration mysql", m.Username, m.Password)
	InitLogger()
	global.Logger.Info("Config log ok!!", zap.String("ok", "success"))
	InitMySQL()
	InitRedis()

	r := InitRouter()

	r.Run(":8002")
}