package main

import "fmt"

type TemperatureMeasurer interface {
	MeasureTemperature() float64
}

type FahrenheitMeasurer struct {
	temperature float64
}

func (f *FahrenheitMeasurer) MeasureTemperature() float64 {
	return f.temperature
}

type FahrenheitToCelsiusAdapter struct {
	fahrenheitMeasurer *FahrenheitMeasurer
}

func (a *FahrenheitToCelsiusAdapter) MeasureTemperature() float64 {
	fahrenheitTemperature := a.fahrenheitMeasurer.MeasureTemperature()
	celsiusTemperature := (fahrenheitTemperature - 32) * 5 / 9
	return celsiusTemperature
}

func main() {
	fahrenheitMeasurer := &FahrenheitMeasurer{temperature: 100}
	adapter := &FahrenheitToCelsiusAdapter{fahrenheitMeasurer: fahrenheitMeasurer}

	celsiusTemperature := adapter.MeasureTemperature()
	fmt.Printf("Temperature in Celsius: %.2f\n", celsiusTemperature)
}
