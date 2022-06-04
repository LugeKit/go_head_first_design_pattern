package command

func Main() {
	controller := NewRemoteController()
	controller.SetSlot(0, &LightOnCommand{"k1"}, &LightOffCommand{"k1"})
	controller.SetSlot(1, &FanOnCommand{"k1"}, &FanOffCommand{"k1"})
	controller.OnButtonPressed(0)
	controller.OnButtonPressed(1)
	controller.OffButtonPressed(0)
	controller.OffButtonPressed(1)
	controller.Undo()
	controller.Undo()
	controller.Undo()
	controller.Undo()
}
