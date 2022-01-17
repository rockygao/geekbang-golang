package configs

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// type Config struct {
// 	ServerHttpAddr string
// 	DatabaseDriver string
// 	DatabaseDsn    string
// }

// Init 初始化viper
func NewConfig() (*viper.Viper, error) {
	var (
		err error
		v   = viper.New()
	)
	path, _ := os.Getwd()
	//fmt.Println(path)
	v.AddConfigPath(path + "/configs")
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	//v.SetConfigFile("config.json")
	if err := v.ReadInConfig(); err == nil {
		fmt.Printf("use config file -> %s\n", v.ConfigFileUsed())
	} else {
		return nil, err
	}
	// config := Config{
	// 	ServerHttpAddr: v.GetString("server.http.addr"),
	// 	DatabaseDriver: v.GetString("data.database.driver"),
	// 	DatabaseDsn:    v.GetString("data.database.source"),
	// }
	// fmt.Println(config)

	return v, err
}

//var ProviderSet = wire.NewSet(NewConfig)
