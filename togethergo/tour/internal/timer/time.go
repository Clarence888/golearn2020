package timer

import "time"

func GetNowTime() time.Time  {
	//return time.Now()

	//加上时区之后
	location,_ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(location)
}

//推算时间
func GetCalculateTime(currentTimer time.Time,d string) (time.Time,error)  {
	//主要原因是 不知道传入的是什么 也没有规定 如果规定好了 那么就可以不用先parseDuration转换了
	duration,err := time.ParseDuration(d)
	if err != nil {
		return time.Time{},err
	}

	return currentTimer.Add(duration),nil
}