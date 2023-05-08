package adapter

type Target interface {
	Request() string
}

type Adaptee interface {
	SpecificRequest() string
}

type AdapteeImpl struct{}

func (a *AdapteeImpl) SpecificRequest() string {
	return "specific request"
}

type Adapter struct {
	adaptee Adaptee
}

func (a *Adapter) Request() string {
	return a.adaptee.SpecificRequest()
}
