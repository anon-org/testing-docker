package main

import (
	"net/http"
	"github.com/okhome/gohome/pkg/application"
	"github.com/okhome/gohome/pkg/request"
	"github.com/okhome/gohome/pkg/gohome"
)

func main() {
	app := application.Default()
	app.Get("/", func(req request.Payload, api *gohome.API) http.ResponseWriter {
		return api.ResponseSuccess(200, "ok")
	})

	app.Start()
}