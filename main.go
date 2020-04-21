package main

import (
	"config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"model"
	"os"
	"route"
	"strconv"
	"time"
)

func init() {
	db, err := gorm.Open(config.DBConfig.Dialect, config.DBConfig.URL)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	if config.DBConfig.SQLLog {
		db.LogMode(true)
	}

	db.DB().SetMaxIdleConns(config.DBConfig.MaxIdleConns)
	db.DB().SetMaxOpenConns(config.DBConfig.MaxOpenConns)

	model.DB = db
}

func main() {
	app := iris.New()
	logger := app.Logger()
	logger.SetLevel("debug")

	sess := sessions.New(sessions.Config{
		// Cookie string, the session's client cookie name, for example: "mysessionid"
		Cookie: config.ServerConfig.SessionID,
		// it's time.Duration, from the time cookie is created, how long it can be alive?
		// 0 means no expire.
		// -1 means expire when browser closes
		// or set a value, like 2 hours:
		Expires: time.Hour * 2,
		// if you want to invalid cookies on different subdomains
		// of the same host, then enable it.
		// Defaults to false.
		DisableSubdomainPersistence: false,
	})

	app.Use(sess.Handler()) // session is always non-nil inside handlers now.
	logger.Infof("==============")
	route.Route(app)
	logger.Infof("==============")

	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		_, _ = ctx.JSON(iris.Map{
			"errNo": model.ErrorCode.NotFound,
			"msg":   "Not Found",
			"data":  iris.Map{},
		})

	})

	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		_, _ = ctx.JSON(iris.Map{
			"errNo": model.ErrorCode.ERROR,
			"msg":   "error",
			"data":  iris.Map{},
		})
	})

	logger.Infof("------------------")
	logger.Infof("Listening %v port....", config.ServerConfig.Port)
	logger.Infof("------------------")
	app.Listen(":" + strconv.Itoa(config.ServerConfig.Port))
}
