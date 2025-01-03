package animal

import (
	"github.com/pprof-demo/animal/canidae/dog"
	"github.com/pprof-demo/animal/canidae/wolf"
	"github.com/pprof-demo/animal/felidae/cat"
	"github.com/pprof-demo/animal/felidae/tiger"
	"github.com/pprof-demo/animal/muridae/mouse"
)

var (
	AllAnimals = []Animal{
		&dog.Dog{},
		&wolf.Wolf{},

		&cat.Cat{},
		&tiger.Tiger{},

		&mouse.Mouse{},
	}
)

type Animal interface {
	Name() string
	Live()

	Eat()
	Drink()
	Shit()
	Pee()
}
