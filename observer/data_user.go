package observer

import "fmt"

type DataUser struct {
	data *WeatherData
}

func NewDataUser(d *WeatherData) *DataUser {
	user := &DataUser{
		data: d,
	}
	d.Register(user)
	return user
}

func (u *DataUser) Update(d *WeatherData) {
	fmt.Printf("DataUser get updated data: %v\n", d.Data)
}
