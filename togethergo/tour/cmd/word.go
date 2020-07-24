package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"strings"
	"tour/internal/word"
)



const (
	MODE_UPPER = iota + 1 //全部转为大写
	MODE_LOWER
	MODE_UNDERSCORE_TO_UPPER_CAMELCASE
	MODE_UNDERSCORE_TO_LOWER_CAMELCASE
	MODE_CAMELCASE_TO_UNDERSCORE
)

var desc = strings.Join([]string{
	"该子命令支持各种单词格式转换，模式如下",
	"1:全转大写",
	"2:全转小写",
	"3:下划线转大驼峰",
	"4:下划线转小驼峰",
	"5:驼峰转下划线",
},"\n")

var wordCmd = &cobra.Command{
	Use: "word",
	Short: "单词格式转换",
	Long: "支持多种单词格式转换",
	Run: func(cmd *cobra.Command, args []string) {

		var content string
		switch mode {
		case MODE_UPPER:
			content = word.ToUpper(str)
		case MODE_LOWER:
			content = word.ToLower(str)
		case MODE_UNDERSCORE_TO_UPPER_CAMELCASE:
			content = word.UnderscoreToUpperCamelCase(str)
		case MODE_UNDERSCORE_TO_LOWER_CAMELCASE:
			content = word.UnderscoreToLowerCamelCase(str)
		case MODE_CAMELCASE_TO_UNDERSCORE:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatalf("暂时不支持此格式,请执行helo word")
		}
		log.Printf("输出结果：%s",content)
	},
}

var str string
var mode int8

func init()  {
	wordCmd.Flags().StringVarP(&str,"str","s","","请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode,"mode","m",0,"请输入要转换的格式")
}