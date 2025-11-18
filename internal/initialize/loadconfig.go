package initialize

import (
	"fmt"
	"go-ecommerce-backend-api/global"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./config/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	//read configuration
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read configuration %w", err))
	}

	fmt.Println("Server Port::", viper.GetInt("server.port"))
	fmt.Println("Server Port::", viper.GetString("security.jwt.key"))

	// configure structure
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}
	// fmt.Println("Server port::", config.Server.Port)

	// for _, db := range config.Databases {
	// 	fmt.Printf("databases User: %s, password: %s, host: %s \n", db.User, db.Password, db.Host)
	// }
}