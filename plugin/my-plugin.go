package main

import (
	"fmt"

	"github.com/Kong/go-pdk"
	"github.com/Kong/go-pdk/server"
)

func main() {
	server.StartServer(New, Version, Priority)
}

var Version = "0.2"
var Priority = 1

type Config struct {
	HeaderKey string
}

func New() interface{} {
	return &Config{}
}

func (conf Config) Access(kong *pdk.PDK) {

	header, err := kong.Request.GetHeader(conf.HeaderKey)

	if err != nil {
		kong.Log.Err(fmt.Sprintf("Error reading '%s' header: %s", conf.HeaderKey, err.Error()))
	}

	if header == "" {
		kong.Log.Err(fmt.Sprintf("Header %s is empty", conf.HeaderKey))
		kong.Response.Exit(401, "", nil)
		return
	}

	kong.Log.Info(fmt.Sprintf("Header %s valid :)", conf.HeaderKey))
	kong.ServiceRequest.SetHeader("header_valid", "true")
}
