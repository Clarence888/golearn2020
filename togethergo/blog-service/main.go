package main

import (
	"blog-service/global"
	"blog-service/internal/model"
	"blog-service/internal/routers"
	"blog-service/pkg/logger"
	"blog-service/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = setDBEngine() //注意这里的不是 :=
	if err != nil {
		log.Fatalf("init.setDBEngine err: %v", err)
	}

	//初始化日志
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setuoLogger err: %v", err)
	}

}

func main() {
	/*
		r := gin.Default()
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(200,gin.H{"message":"pongaaa"})
		})
		r.Run(":9999")
	*/

	/*演化版本
	router := routers.NewRouter()
	s := &http.Server{
		Addr: ":9999",
		Handler: router,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	*/
	//修改版本

	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,  //允许读取的最大时间
		WriteTimeout:   global.ServerSetting.WriteTimeout, //允许写入最大时间
		MaxHeaderBytes: 1 << 20,                           //请求头最大字节数
	}

	//打印日志
	global.Logger.Infof("%s: togergo/%s", "hello", "blog-service")

	s.ListenAndServe()
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("JWT",&global.JWTSetting)
	if err != nil {
		return err
	}
	global.JWTSetting.Expire *= time.Second

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil
}

func setDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDbEngine(global.DatabaseSetting)
	//注意 这里的不可以写成 global.DBEngine,err := xxxx
	//因为 := 会重新声明并创建新局部变量 会导致 其他包在调用global。DBENgine变量时候 还是nil
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(
		&lumberjack.Logger{
			Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
			MaxSize:   600, //最大占用空间
			MaxAge:    10,  //最大生存周期
			LocalTime: true,
		},
		"",
		log.LstdFlags).WithCaller(2)
	return nil
}

/*
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

*/
