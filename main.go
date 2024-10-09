// Code generated by hertz generator.

package main

import (
	"bluebook/config"
	"bluebook/dal"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func init() {
	config.Init()
	dal.Init()
}
func main() {
	conf := config.Config
	h := server.Default(
		server.WithHostPorts(conf.System.Host + ":" + conf.System.Port),
	)

	register(h)
	h.Spin()
}
