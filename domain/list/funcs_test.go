package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Person struct {
	name string
	age  int
}

//func TestMap(t *testing.T) {
//	Map[Person, bool](func(p Person) bool {
//		return true
//	})
//}

func TestFilter(t *testing.T) {
	slice := []*Person{{
		name: "Raisa",
		age:  30,
	},
		{
			name: "Ivan",
			age:  32,
		},
	}
	result := Filter[Person](slice, func(p *Person) bool {
		if p.age == 30 {
			return true
		}
		return false
	})
	assert.Len(t, result, 1)
}
