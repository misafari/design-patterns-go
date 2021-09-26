package restaurant_example

type Cook struct {
	Commands []Command
}

func (c *Cook) ExecuteCommands() {
	for _, c := range c.Commands {
		c.Execute()
	}
}
