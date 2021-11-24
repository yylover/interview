package adapter

import "fmt"

//适配器，进行协议转换

type Volts220 struct {

}

func (v Volts220) OutputPower() {
	fmt.Println("output power 220v")
}

// 适配器接口
type Adaptee interface {
	OutputPower()
}

// 目标接口
type Target interface {
	ConvertTo5V()
}


// Adapter 适配器结构体，嵌套Adaptee接口，实现target接口
type Adapter struct {
	Adaptee
}

func (ad *Adapter) ConvertTo5V() {
	ad.OutputPower()
	fmt.Println("conver to 5v")
}
