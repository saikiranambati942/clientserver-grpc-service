package main

import "fmt"

//In GO we have two types of methods
// 1) Value Receivers- They do calculations on values
// 2) Pointer Receivers - If you want to change some value in struct then need Pointer Receiver
//There is no concept of classes in GO. In GO we have methods & we associate them with Structs

const usixteenbitmax float64 = 65536
const kmh_multiple float64 = 1.609

type car struct {
	gas_pedal      uint16 //0 to 65535
	brake_pedal    uint16
	steering_wheel int16 //-32k to +32k
	top_speed_kmh  float64
}

func (c car) kmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbitmax)
}
func (c car) mph() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbitmax / kmh_multiple)
}
func (c *car) newTopSpeed(new_top_speed float64) {
	c.top_speed_kmh = new_top_speed
}
func main() {
	a_car := car{gas_pedal: 65000, brake_pedal: 0, steering_wheel: 12561, top_speed_kmh: 225.0}
	//a_car := car{2236, 0, 1256, 225.0}  //we can use this format also but this is not legible
	fmt.Println(a_car)
	fmt.Println(a_car.gas_pedal)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())
	a_car.newTopSpeed(500)

	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())
}
