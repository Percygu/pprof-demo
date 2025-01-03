package felidae

import "github.com/pprof-demo/animal"

type Felidae interface {
	animal.Animal
	Climb()
	Sneak()
}
