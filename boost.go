/*

Booster made to works with GoXp
We do not inject, we boost.

inject -> boost
Injector -> Booster
...

*/

package boost

import (
	"fmt"
	"reflect"
)
// reflection in Go http://blog.golang.org/laws-of-reflection

type Booster interface {
	Applicator
	Invoker
	TypeMapper
	SetParent(Booster)
	// SetParent sets the parent of the booster, If the booster cannot find a
	// dependancy in its Type map it will check its parent before returning an
	// error.
}

type Application interface {
	// Applicator represents an interface for mapping dependencies to a struct.
	Apply(interface{}) error
	// Maps dependencies in the Type map to each field in the struct
	// that is tagged with 'boost', Returns an error if the booster
	// fails.
}

type Invoker interface {
	// Invoker represents an interface for calling functions via reflection.
	Invoker(interface{}) ([]reflect Value, error)
	// Invoke attemps to call the interface() provided as a function,
	// providing dependencies for function arguments based on Type.
	// Returns a slice of reflect.Value representing the returned value of the
	// function
	// Returns an error if the injection fails.
}
