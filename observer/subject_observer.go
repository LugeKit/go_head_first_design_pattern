package observer

type Subject interface {
	Register(o Observer)
	Unregister(o Observer)
}

type Observer interface {
	Update(data *WeatherData)
}
