package initialize

import (
	"fmt"
	"go-ecommerce-backend-api/global"
	"go-ecommerce-backend-api/internal/po"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func checkErrorPanic(err error, errString string){
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
	}
}

func InitMySQL() {
	m := global.Config.Mysql
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Dbname)
  	db, err := gorm.Open(mysql.Open(s), &gorm.Config{})
	checkErrorPanic(err," Init Mysql initializing err")
	global.Logger.Info("Initializing MySQL Successfully")
	global.Mdb = db

	SetPool()
	MigrateTables()
}

func SetPool() {
	m := global.Config.Mysql
	sqlDb, err := global.Mdb.DB()
	if err !=nil {
		fmt.Printf("mysql error: %s::", err)
	}
	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))

}

func MigrateTables() {
	err := global.Mdb.AutoMigrate(
		&po.User{},
		&po.Role{},
	)

	if err != nil {
		fmt.Println("Migrating tables error:", err)
	}

}