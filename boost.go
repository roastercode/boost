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
