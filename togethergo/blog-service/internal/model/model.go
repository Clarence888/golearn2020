package model

import (
	"blog-service/global"
	"blog-service/pkg/setting"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//公共model 一些数据创建和更新删除信息
type Model struct {
	ID uint32 `gorm:"primary_key" json:"id"`
	CreateBy string `json:"create_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeleteOn uint32 `json:"delete_on"`
	IsDel uint8 `json:"is_del"`
}

func NewDbEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB,error)  {

	fmt.Println(databaseSetting)

	db,err := gorm.Open(
		databaseSetting.DBType,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
			databaseSetting.UserName,
			databaseSetting.Password,
			databaseSetting.Host,
			databaseSetting.DBName,
			databaseSetting.Charset,
			databaseSetting.ParseTime,
			))
	if err != nil {
		return nil,err
	}
	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return db,nil
}