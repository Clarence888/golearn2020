package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
	"time"
	"tour/internal/timer"
)

var calculateTime string
var duration string

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间处理格式",
	Long:  "时间格式处理",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "获取当前时间",
	Long:  "获取当前的时间",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		log.Printf("输出结果： %s, %d", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix())
	},
}

//使用预定义格式的代码  标注库time
//t := time.Now().Format(time.RFC3339)

var calculateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "计算所需要的时间",
	Long:  "计算所需要的时间",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTimer time.Time
		var layout = "2006-01-02 15:04:05" //注意 这种格式化时间 必须写这个值 其他值无效
		// "2020-01-01 13:04:04" 写成这种会出错

		if calculateTime == "" {
			currentTimer = timer.GetNowTime()
		} else {
			var err error
			if !strings.Contains(calculateTime, " ") {
				layout = "2007-09-08"
			}

			currentTimer, err = time.Parse(layout, calculateTime)
			if err != nil {
				t, _ := strconv.Atoi(calculateTime)
				currentTimer = time.Unix(int64(t), 0)
			}
		}

		calculateTime, err := timer.GetCalculateTime(currentTimer, duration)
		if err != nil {
			log.Fatalf("timer.GetcalculateTime err: %v", err)
		}

		fmt.Println(calculateTime)
		log.Printf("输出结果： %s, %d", calculateTime.Format(layout), calculateTime.Unix())

		//time.Format 格式化  还有 time.Parse 请注意 可能会导致时区问题 因此使用 time.ParseInLocation(layout,inputTime,location)
		//所有解析和格式化的最好都带上时区。 保障万无一失 同时需要注意运维部署的相关时区

	},
}

func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calculateTimeCmd)

	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", "需要计算的时间，有效单位为时间戳或者格式化后的时间")
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", "持续时间，有效单位为'ns','us','ms','s','m','h'")
}

/*
~/go/src/togethergo/tour master*
❯ go run main.go time calc  -c="2020-11-11 00:00:00" -d=2m
2020-11-11 00:02:00 +0000 UTC
2020/07/27 11:16:07 输出结果： 2020-11-11 00:02:00, 1605052920

~/go/src/togethergo/tour master*
❯ go run main.go time calc  -c="2020-11-11 00:00:00" -d=-2m
2020-11-10 23:58:00 +0000 UTC
2020/07/27 11:16:14 输出结果： 2020-11-10 23:58:00, 1605052680

~/go/src/togethergo/tour master*
❯ go run main.go time now
2020/07/27 11:16:24 输出结果： 2020-07-07 77:1111:11, 1595819784

~/go/src/togethergo/tour master*

*/

//时间相关的额需要注意失去  Local  UTC
//	time.LoadLocation()
