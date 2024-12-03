package observer

import "fmt"

type Observer interface {
	Update(interface{})
}

type Subject interface {
	Register(observer Observer)
	Unregister(Observer Observer)
	Notify()
}

type Stock struct {
	Name      string
	Price     float64
	Observers []Observer
}

func (s *Stock) Register(observer Observer) {
	s.Observers = append(s.Observers, observer)
}

func (s *Stock) Unregister(observer Observer) {
	for index, currObserver := range s.Observers {
		if currObserver == observer {
			s.Observers = append(s.Observers[:index], s.Observers[index+1:]...)
			return
		}
	}
}

func (s *Stock) SetPrice(newPrice float64) {
	s.Price = newPrice
}

func (s *Stock) Notify() {
	for _, Observer := range s.Observers {
		Observer.Update(s.Price)
	}
}

type Broker struct{}

func (b *Broker) Update(price interface{}) {
	fmt.Println(price)
}
