package main

import "fmt"

type Car struct {
	gas_pedal uint16 // min 0, max 65535
	brake_pedal uint16
	steering_wheel int16
	top_speed_kmh float64
}

func main() {
	a_car := Car {
		gas_pedal: 22341,
		brake_pedal: 0,
		steering_wheel: 12561,
		top_speed_kmh: 225.0};
	// a_car := Car { 22341, 0, 12561, 225.0 };

	fmt.Println(a_car.gas_pedal);
}