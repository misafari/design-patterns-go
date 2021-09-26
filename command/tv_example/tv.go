package tv_example

import "fmt"

type Device interface {
	On()
	Off()
}

type TV struct {
	isRunning bool
}

func (t *TV) On() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

func (t *TV) Off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}
