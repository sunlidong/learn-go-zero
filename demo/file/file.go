package main

import (
	"flag"
	"net/http"

	"file/internal/config"
	"file/internal/handler"
	"file/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/rest/router"
)

var configFile = flag.String("f", "etc/file-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	rt := router.NewPatRouter()
	rt.SetNotFoundHandler(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("服务器开小差了,这里可定制"))
	}))

	server := rest.MustNewServer(c.RestConf, rest.WithRouter(rt))
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)
	server.Start()
}
