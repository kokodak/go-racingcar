package racingcar

type Car struct {
	name     string
	position int
}

func NewCar(name string) *Car {
	return &Car{
		name:     name,
		position: 0,
	}
}

func (c *Car) Move() {
	c.position++
}

func (c *Car) Name() string {
	return c.name
}

func (c *Car) Position() int {
	return c.position
}
