package configs

import (
	"github.com/spf13/viper"
)

// type ConfigStruct struct {
// 	name  string   `yaml:"name"`
// 	mysql []string `yaml:"mysql"`
// 	redis []string `yaml:"redis"`
// }

// func main() {

// 	var configViperConfig = viper.New()
// 	configViperConfig.AddConfigPath("./")
// 	configViperConfig.SetConfigName("configs")
// 	configViperConfig.SetConfigType("yaml")
// 	//读取配置文件内容
// 	if err := configViperConfig.ReadInConfig(); err != nil {
// 		panic(err)
// 	}
// 	var c ConfigStruct
// 	if err := configViperConfig.Unmarshal(&c); err != nil {
// 		panic(err)
// 	}
// 	err := configViperConfig.ReadInConfig() // Find and read the config file
// 	if err != nil {                         // Handle errors reading the config file
// 		panic(fmt.Errorf("Fatal error config file: %w \n", err))
// 	}
// }

type Config struct {
	HttpAddress string
	GrpcAddress string
}

func InitConfig() (*viper.Viper, error) {
	viper.AddConfigPath("./../week4/config/")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {

		return nil, err
	}
	return viper.GetViper(), nil
}

func NewConfig(v *viper.Viper) *Config {
	return &Config{
		HttpAddress: ":" + v.Sub("http").GetString("port"),
		GrpcAddress: ":" + v.Sub("grpc").GetString("port"),
	}
}
