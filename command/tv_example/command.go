package tv_example

type Command interface {
	Execute()
}

type Button struct {
	Command
}

func (b *Button) Press() {
	b.Execute()
}
