package command

import (
	"fmt"
)

type Command interface {
	Excute()
	Undo()
}

type LightOnCommand struct {
	light string
}

func (c *LightOnCommand) Excute() {
	fmt.Printf("light %s is on\n", c.light)
}

func (c *LightOnCommand) Undo() {
	fmt.Printf("light %s is off\n", c.light)
}

type LightOffCommand struct {
	light string
}

func (c *LightOffCommand) Excute() {
	fmt.Printf("light %s is off\n", c.light)
}

func (c *LightOffCommand) Undo() {
	fmt.Printf("light %s is on\n", c.light)
}

type FanOnCommand struct {
	fan string
}

func (c *FanOnCommand) Excute() {
	fmt.Printf("fan %s is on\n", c.fan)
}

func (c *FanOnCommand) Undo() {
	fmt.Printf("fan %s is off\n", c.fan)
}

type FanOffCommand struct {
	fan string
}

func (c *FanOffCommand) Excute() {
	fmt.Printf("fan %s is off\n", c.fan)
}

func (c *FanOffCommand) Undo() {
	fmt.Printf("fan %s if on\n", c.fan)
}

type NoCommand struct{}

func (c *NoCommand) Excute() {
	fmt.Println("no command")
}

func (c *NoCommand) Undo() {
	fmt.Println("no command")
}
