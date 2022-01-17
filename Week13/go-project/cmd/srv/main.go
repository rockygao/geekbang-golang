package main

import (
	"go-project/pkg/app/srv"

	"github.com/quexer/utee"
	log "github.com/sirupsen/logrus"
)

func main() {
	app, cleanup, err := srv.RunSrv()
	utee.Chk(err)
	defer cleanup()

	err = app.Init().Run()
	if err != nil {
		log.Fatalln(err)
	}
}
