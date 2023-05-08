package command

import "fmt"

type Command interface {
	execute()
}

type ConcreteCommand struct {
	receiver Receiver
}

func (c *ConcreteCommand) execute() {
	c.receiver.action()
}

type Receiver struct {
}

func (r *Receiver) action() {
	fmt.Println("Receiver action")
}

type Invoker struct {
	command Command
}

func (i *Invoker) setCommand(c Command) {
	i.command = c
}

func (i *Invoker) executeCommand() {
	i.command.execute()
}
