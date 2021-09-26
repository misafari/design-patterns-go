package tv_example

type OnCommand struct {
	Device
}

func (o *OnCommand) Execute() {
	o.On()
}

type OffCommand struct {
	Device
}

func (o *OffCommand) Execute() {
	o.Off()
}
