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

type TypeMapper interface {
	// TypeMapper represents an interface for mapping interface{} values based
	// on type.
	Map(interface{}) TypeMapper
	// Maps the interface{} value based on its immediate type form reflect.TypeOf
	MapTo(interface{}, interface{}) TypeMapper
	// Maps the interface{} value  base on the pointer of the Interface provided
	// This is really only useful for mapping a value as an interface, as interfaces
	// cannot at this time be referenced directly without a pointer.
	Set(reflect.Type, reflect.Value) TypeMapper
	// Provides a possibility to directly insert a mapping based on type and value.
	// This makes it possible to directly map type arguments not possible to instantiate
	// with reflect like unidirectional channels.
	Set(reflect.Type, reflect.Value) TypeMapper
	Get(reflect.Type) reflect.Value
	// Returns the Value that is mapped to the current type. Returns a zeroed Value if
	// the Type has not been mapped
	Get(reflect.Type) reflect.Value
}
