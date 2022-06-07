package adapter

import "fmt"

type HongKongPlugin interface {
	HongKongPlugIn()
}

type ChinaPlugin interface {
	ChinaPlugIn()
}

type NormalChinaPlugin struct{}

func (p *NormalChinaPlugin) ChinaPlugIn() {
	fmt.Println("China plug in")
}

type ChinaToHongKongAdapter struct {
	NormalChinaPlugin
}

func (a *ChinaToHongKongAdapter) HongKongPlugIn() {
	a.NormalChinaPlugin.ChinaPlugIn()
}
