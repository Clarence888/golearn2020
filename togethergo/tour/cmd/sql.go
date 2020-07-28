package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"tour/internal/sql2struct"
)

var username string
var password string

var host string
var charset string
var dbType string
var dbName string
var tableName string

var sqlCmd = &cobra.Command{
	Use: "sql",
	Short: "sql表转换和处理",
	Long: "sql表转换和处理",
	Run: func(cmd *cobra.Command, args []string) {},
}

var sql2structCmd = &cobra.Command{
	Use: "struct",
	Short: "sql转换",
	Long: "sql转换",
	Run: func(cmd *cobra.Command, args []string) {

		dbInfo := &sql2struct.DBInfo{
			DBType:   dbType,  //注意 此处为赋值  小心手残 不能 "dbType" 否则拿不到相关信息
			Host:     host,
			UserName: username,
			Password: password,
			Charset:  charset,
		}

		dbModel := sql2struct.NewDBModel(dbInfo)

		err := dbModel.Connect()
		if err != nil {
			log.Fatalf("dbmodel.connect err : %v",err)
		}

		columns,err := dbModel.GetColumns(dbName,tableName)
		err = dbModel.Connect()
		if err != nil {
			log.Fatalf("dbmodel.GetColumns err : %v",err)
		}

		template := sql2struct.NewStructTemplate()
		templateColumns := template.AssemblyColumns(columns)
		err = template.Generate(tableName,templateColumns)
		if err != nil {
			log.Fatalf("dbmodel.Generate err : %v",err)
		}
	},


}

func init()  {
	sqlCmd.AddCommand(sql2structCmd)
	sql2structCmd.Flags().StringVarP(&username, "username", "", "", "请输入数据库的账号")
	sql2structCmd.Flags().StringVarP(&password, "password", "", "", "请输入数据库的密码")
	sql2structCmd.Flags().StringVarP(&host, "host", "", "127.0.0.1:3306", "请输入数据库的HOST")
	sql2structCmd.Flags().StringVarP(&charset, "charset", "", "utf8mb4", "请输入数据库的编码")
	sql2structCmd.Flags().StringVarP(&dbType, "type", "", "mysql", "请输入数据库实例类型")
	sql2structCmd.Flags().StringVarP(&dbName, "db", "", "", "请输入数据库名称")
	sql2structCmd.Flags().StringVarP(&tableName, "table", "", "", "请输入表名称")
}


/*
❯ go run main.go sql struct --username  xxxxx --password xxxxx  --db superstar --table star_info --host xxxxx:3306
changliang
type StarInfo struct {
         // 主键ID
         Id     int32   `json:id`
         // 中文名
         NameZh string  `json:name_zh`
         // 英文名
         NameEn string  `json:name_en`
         // 头像
         Avatar string  `json:avatar`
         // 出生日期
         Birthday       string  `json:birthday`
         // 身高，单位cm
         Height int32   `json:height`
         // 体重，单位kg
         Weight int32   `json:weight`
         // 俱乐部
         Club   string  `json:club`
         // 球衣号码以及主打位置
         Jersy  string  `json:jersy`
         // 国籍
         Country        string  `json:country`
         // 出生地
         Birthaddress   string  `json:birthaddress`
         // 个人特点
         Feature        string  `json:feature`
         // 更多介绍
         Moreinfo       string  `json:moreinfo`
         // 状态，默认值 0 正常，1 删除
         SysStatus      int8    `json:sys_status`
         // 创建时间
         SysCreated     int32   `json:sys_created`
         // 最后修改时间
         SysUpdated     int32   `json:sys_updated`
}
func (model StarInfo) TableName() string {
        return "star_info"
}
~/go/src/togethergo/tour master*
❯
*/

/*
CREATE TABLE `star_info` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name_zh` varchar(50) NOT NULL DEFAULT '' COMMENT '中文名',
  `name_en` varchar(50) NOT NULL DEFAULT '' COMMENT '英文名',
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
  `birthday` varchar(50) NOT NULL DEFAULT '' COMMENT '出生日期',
  `height` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '身高，单位cm',
  `weight` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '体重，单位kg',
  `club` varchar(50) NOT NULL DEFAULT '' COMMENT '俱乐部',
  `jersy` varchar(50) NOT NULL DEFAULT '' COMMENT '球衣号码以及主打位置',
  `country` varchar(50) NOT NULL DEFAULT '' COMMENT '国籍',
  `birthaddress` varchar(255) NOT NULL DEFAULT '' COMMENT '出生地',
  `feature` varchar(255) NOT NULL DEFAULT '' COMMENT '个人特点',
  `moreinfo` text COMMENT '更多介绍',
  `sys_status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态，默认值 0 正常，1 删除',
  `sys_created` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `sys_updated` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '最后修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
 */