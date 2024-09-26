package observer

import "fmt"

// Observer interface
type Observer interface {
	Update(temp float64) // Update method to notify observers of the change
}

// Subject interface
type Subject interface {
	Register(observer Observer)   // Register an observer
	Deregister(observer Observer) // Deregister an observer
	Notify()                      // Notify all observers of a state change
}

// TemperatureSensor is the concrete subject
type TemperatureSensor struct {
	observers []Observer // List of observers
	temp      float64    // Current temperature
}

func NewTemperatureSensor() *TemperatureSensor {
	return &TemperatureSensor{}
}

func (s *TemperatureSensor) Register(observer Observer) {
	s.observers = append(s.observers, observer) // Add an observer to the list
}

func (s *TemperatureSensor) Deregister(observer Observer) {
	for i, obs := range s.observers {
		if obs == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...) // Remove the observer
			break
		}
	}
}

func (s *TemperatureSensor) Notify() {
	for _, observer := range s.observers {
		observer.Update(s.temp) // Notify each observer with the updated temperature
	}
}

func (s *TemperatureSensor) SetTemperature(temp float64) {
	fmt.Printf("TemperatureSensor: setting temperature to %.2f\n", temp)
	s.temp = temp
	s.Notify() // Notify observers about the temperature change
}

// CurrentConditionsDisplay is a concrete observer that displays the current temperature
type CurrentConditionsDisplay struct {
	Sensor *TemperatureSensor
}

func (d *CurrentConditionsDisplay) Update(temp float64) {
	fmt.Printf("CurrentConditionsDisplay: current temperature is %.2f°C\n", temp)
}

// ForecastDisplay is another concrete observer that displays a temperature forecast
type ForecastDisplay struct {
	Sensor *TemperatureSensor
}

func (d *ForecastDisplay) Update(temp float64) {
	forecast := temp + 2.0 // Simple forecast logic (add 2 degrees)
	fmt.Printf("ForecastDisplay: predicted temperature is %.2f°C\n", forecast)
}
