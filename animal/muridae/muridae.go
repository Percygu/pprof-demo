package muridae

import "github.com/pprof-demo/animal"

type Muridae interface {
	animal.Animal
	Hole()
	Steal()
}
