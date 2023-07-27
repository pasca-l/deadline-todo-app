package controllers

import (
	"net/http"

	"app/config"
)

func StartServer() error {
	files := http.FileServer(http.Dir(config.Config.Srv.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/", home)
	return http.ListenAndServe(":" + config.Config.Srv.Port, nil)
}
