package global

import (
	"blog-service/pkg/logger"
	"blog-service/pkg/setting"
)

/*
配置全局变量
*/
var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger //定义一个Logger对象 便于使用
)
