package main

import "fmt"

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

type Car struct {
	gas_pedal uint16 // min 0, max 65535
	brake_pedal uint16
	steering_wheel int16
	top_speed_kmh float64
}

func (c Car) kmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbitmax);
}

func (c Car) mph() float64 {
	return c.kmh() / kmh_multiple;
}

func (c *Car) new_top_speed(newspeed float64) {
	c.top_speed_kmh = newspeed;
}

func newer_top_speed(c Car, newspeed float64) Car {
	c.top_speed_kmh = newspeed;
	return c;
}

func main() {
	a_car := Car {
		gas_pedal: 65000,
		brake_pedal: 0,
		steering_wheel: 12561,
		top_speed_kmh: 225.0};

	fmt.Printf("Gas pedal: %d, kmh: %v, mph: %v\n", a_car.gas_pedal, a_car.kmh(), a_car.mph());
	a_car.new_top_speed(500);
	// a_car = newer_top_speed(a_car, 500);
	fmt.Printf("Gas pedal: %d, kmh: %v, mph: %v\n", a_car.gas_pedal, a_car.kmh(), a_car.mph());
}