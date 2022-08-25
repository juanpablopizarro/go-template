package main

import (
	"fmt"
	"go-template/service"
	"net/http"
	"os"
	"strconv"

	"github.com/go-kit/log"
	"github.com/spf13/viper"
)

func main() {
	viper.SetEnvPrefix("APP")
	viper.BindEnv("PORT")

	port := viper.GetInt("PORT")
	fmt.Println(port)

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "listen", strconv.Itoa(port), "caller", log.DefaultCaller)

	srv := service.NewService()
	srv = service.NewLoggingMiddleware(logger, srv)
	r := service.NewHTTPServer(srv, logger)
	logger.Log("msg", "HTTP", "addr", strconv.Itoa(port))
	logger.Log("err", http.ListenAndServe(fmt.Sprintf(":%v", port), r))
}
