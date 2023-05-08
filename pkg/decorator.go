package decorator

import "fmt"

type Component interface {
	Operation() string
}

type ConcreteComponent struct{}

func (c *ConcreteComponent) Operation() string {
	return "ConcreteComponent"
}

type Decorator interface {
	Component
}

type ConcreteDecoratorA struct {
	component Component
}

func NewConcreteDecoratorA(component Component) *ConcreteDecoratorA {
	return &ConcreteDecoratorA{
		component: component,
	}
}

func (d *ConcreteDecoratorA) Operation() string {
	return fmt.Sprintf("ConcreteDecoratorA(%s)", d.component.Operation())
}

type ConcreteDecoratorB struct {
	component Component
}

func NewConcreteDecoratorB(component Component) *ConcreteDecoratorB {
	return &ConcreteDecoratorB{
		component: component,
	}
}

func (d *ConcreteDecoratorB) Operation() string {
	return fmt.Sprintf("ConcreteDecoratorB(%s)", d.component.Operation())
}
