package strategy

type Strategy interface {
	Execute(a, b int) int
}

type Add struct{}

func (Add) Execute(a, b int) int {
	return a + b
}

type Subtract struct{}

func (Subtract) Execute(a, b int) int {
	return a - b
}

type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(s Strategy) {
	c.strategy = s
}

func (c *Context) ExecuteStrategy(a, b int) int {
	return c.strategy.Execute(a, b)
}
