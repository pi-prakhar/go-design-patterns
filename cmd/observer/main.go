package main

import "go-design-patterns/pkg/observer"

func main() {
	sensor := observer.NewTemperatureSensor()

	// Create observers
	currentDisplay := &observer.CurrentConditionsDisplay{Sensor: sensor}
	forecastDisplay := &observer.ForecastDisplay{Sensor: sensor}

	// Register observers with the sensor
	sensor.Register(currentDisplay)
	sensor.Register(forecastDisplay)

	// Change temperature (this will notify all observers)
	sensor.SetTemperature(25.5)
	sensor.SetTemperature(30.0)

	// Deregister one observer
	sensor.Deregister(currentDisplay)

	// Change temperature again (only forecast display will be notified)
	sensor.SetTemperature(35.0)
}
