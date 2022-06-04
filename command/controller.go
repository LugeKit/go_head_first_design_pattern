package command

import "github.com/lugekit/design_pattern/tools"

type RemoteController struct {
	onCommands  [7]Command
	offCommands [7]Command
	undoList    []Command
}

func NewRemoteController() *RemoteController {
	return &RemoteController{
		onCommands:  [7]Command{&NoCommand{}},
		offCommands: [7]Command{&NoCommand{}},
		undoList:    make([]Command, 0),
	}
}

func (c *RemoteController) SetSlot(slot int, onCommand Command, offCommand Command) {
	c.onCommands[slot] = onCommand
	c.offCommands[slot] = offCommand
}

func (c *RemoteController) OnButtonPressed(slot int) {
	command := c.onCommands[slot]
	command.Excute()
	c.undoList = append(c.undoList, command)
}

func (c *RemoteController) OffButtonPressed(slot int) {
	command := c.offCommands[slot]
	command.Excute()
	c.undoList = append(c.undoList, command)
}

func (c *RemoteController) Undo() {
	if len(c.undoList) == 0 {
		return
	}

	lastIndex := len(c.undoList) - 1
	c.undoList[lastIndex].Undo()
	c.undoList, _ = tools.Remove(c.undoList, lastIndex)
}
