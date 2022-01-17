package srv

import (
	"context"

	"github.com/spf13/viper"
)

type App struct {
	Bootloader *Bootloader
	Config     *viper.Viper
}

func (p *App) Run() error {
	// Run("里面不指定端口号默认为8080")
	p.Bootloader.HttpInitializer.Router.Run(":" + p.Config.GetString("app.port"))
	p.afterStart()
	return nil
}

func (p *App) Init() *App {
	p.Bootloader.Boot(context.Background())
	return p
}

func (p *App) afterStart() error {
	return nil
}
