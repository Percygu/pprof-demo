package canidae

import "github.com/pprof-demo/animal"

type Canidae interface {
	animal.Animal
	Run()
	Howl()
}
