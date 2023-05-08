package factory

type Product interface {
	GetName() string
}

type ConcreteProductA struct{}

func (p *ConcreteProductA) GetName() string {
	return "Product A"
}

type ConcreteProductB struct{}

func (p *ConcreteProductB) GetName() string {
	return "Product B"
}

type Creator interface {
	CreateProduct() Product
}

type ConcreteCreatorA struct{}

func (c *ConcreteCreatorA) CreateProduct() Product {
	return &ConcreteProductA{}
}

type ConcreteCreatorB struct{}

func (c *ConcreteCreatorB) CreateProduct() Product {
	return &ConcreteProductB{}
}
