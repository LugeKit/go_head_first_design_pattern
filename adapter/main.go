package adapter

func Main() {
	plugin := NormalChinaPlugin{}
	testHongKongPlugin(&ChinaToHongKongAdapter{plugin})
}

func testHongKongPlugin(p HongKongPlugin) {
	p.HongKongPlugIn()
}
