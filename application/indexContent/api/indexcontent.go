package main

import (
	"MyRainFarm/pkg/xrrcode"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"

	"MyRainFarm/application/indexContent/api/internal/config"
	"MyRainFarm/application/indexContent/api/internal/handler"
	"MyRainFarm/application/indexContent/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/indexcontent.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	httpx.SetErrorHandler(xrrcode.ErrHandler)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
