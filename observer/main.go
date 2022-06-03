package observer

func Main() {
	wData := NewWeatherData()
	user := NewDataUser(wData)
	wData.SetNewData("This is second data.")
	wData.Unregister(user)
	wData.SetNewData("This is third data.")
}
