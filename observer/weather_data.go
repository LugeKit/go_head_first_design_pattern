package observer

import "github.com/lugekit/design_pattern/tools"

type WeatherData struct {
	observers []Observer
	Data      string
}

func NewWeatherData() *WeatherData {
	return &WeatherData{
		observers: make([]Observer, 0),
		Data:      "This is primary data",
	}
}

func (w *WeatherData) Register(o Observer) {
	w.observers = append(w.observers, o)
}

func (w *WeatherData) Unregister(o Observer) {
	index := len(w.observers)
	for i, registeredO := range w.observers {
		if registeredO == o {
			index = i
			break
		}
	}
	w.observers, _ = tools.Remove(w.observers, index)
}

func (w *WeatherData) SetNewData(d string) {
	w.Data = d
	for _, o := range w.observers {
		o.Update(w)
	}
}
