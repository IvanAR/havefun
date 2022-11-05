package obj

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Mage struct {
	kind string
	age  int
	name string
}

func TestGetOrDefaultReturningDefault(t *testing.T) {
	var mage Optional[Mage]

	result := mage.GetOrDefault(func() Mage {
		return Mage{
			kind: "",
			age:  0,
			name: "",
		}
	})
	assert.Equal(t, 0, result.age)
}

func TestGetOrDefaultReturningOriginalValue(t *testing.T) {
	mage := Optional[Mage]{
		value: &Mage{
			kind: "Original Mage",
			age:  1500,
			name: "The One",
		},
	}

	result := mage.GetOrDefault(func() Mage {
		return Mage{
			kind: "",
			age:  0,
			name: "",
		}
	})
	assert.Equal(t, "Original Mage", result.kind)
}

func TestGetOrNil(t *testing.T) {
	mage := NewOptional[any](nil)

	result := mage.GetOrNil()
	assert.Nil(t, result)
}
