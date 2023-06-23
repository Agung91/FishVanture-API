package main

import (
	authhttp "github.com/e-fish/api/auth_http"
	budidayahttp "github.com/e-fish/api/budidaya_http"
	mainconfig "github.com/e-fish/api/main_config"
	"github.com/e-fish/api/migrations"
	"github.com/e-fish/api/pkg/common/helper/logger"
	"github.com/e-fish/api/pkg/common/helper/ptime"
	"github.com/e-fish/api/pkg/common/helper/restsvr"
	transactionhttp "github.com/e-fish/api/transaction_http"
)

func main() {
	ptime.SetDefaultTimeToUTC()

	//get main config
	conf := mainconfig.GetConfig()
	logger.SetupLogger(conf.Debug)

	migrations.Migrations()

	//create new route
	restsvr.NewRoute(conf.AppConfig)

	//register auth http in main
	authhttp.NewAuthHttp()
	//register auth http in main
	budidayahttp.NewBudidayaHttp()

	transactionhttp.NewTransactionHttp()

	restsvr.Run()

}
