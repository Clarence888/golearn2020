//该包负责转换 摄氏温度 和 华氏温度
package zhizuoyigebao

import "fmt"

type Celsius float64 //摄氏温度
type Fahrenheit float64 //华氏温度

const (
	AbsoluteZeroC Celsius = -273.5
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)

func (c Celsius) String() string  {
	return fmt.Sprintf("%g 摄氏度",c)
}

func (f Fahrenheit) String() string  {
	return fmt.Sprintf("%g 华氏度",f)
}

