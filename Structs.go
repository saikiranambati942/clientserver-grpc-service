package main

import "fmt"

//There is no concept of classes in GO. In GO we have methods & we associate them with Structs
type car struct {
	gas_pedal      uint16 //0 to 65535
	brake_pedal    uint16
	steering_wheel int16 //-32k to +32k
	top_speed_kmh  float64
}

func main() {
	a_car := car{gas_pedal: 2236, brake_pedal: 0, steering_wheel: 1256, top_speed_kmh: 225.0}
	//a_car := car{2236, 0, 1256, 225.0}  //we can use this format also but this is not legible
	fmt.Println(a_car)
	fmt.Println(a_car.gas_pedal)

}
