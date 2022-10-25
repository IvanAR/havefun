package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Person struct {
	name string
	age  int
}

func TestFilter(t *testing.T) {
	slice := FunList[Person]{{
		name: "Raisa",
		age:  30,
	},
		{
			name: "Ivan",
			age:  32,
		},
	}
	result := slice.Filter(func(p *Person) bool {
		if p.age == 30 {
			return true
		}
		return false
	})

	assert.Len(t, result, 1)
}

func TestMap(t *testing.T) {
	slice := FunList[Person]{{
		name: "Raisa",
		age:  30,
	},
		{
			name: "Ivan",
			age:  32,
		},
	}
	result := slice.Map(func(p *Person) *Person {
		if p.age == 30 {
			return &Person{
				name: "New person",
				age:  20,
			}
		}
		return p
	})
	assert.Len(t, result, 2)
	assert.Equal(t, 20, result[0].age)
}

func TestFilterChainMap(t *testing.T) {
	slice := FunList[Person]{{
		name: "Raisa",
		age:  30,
	},
		{
			name: "Ivan",
			age:  32,
		},
	}
	result := slice.Filter(func(p *Person) bool {
		if p.age == 30 {
			return true
		}
		return false
	}).Map(func(p *Person) *Person {
		return &Person{
			name: "Equals 30 person",
			age:  p.age,
		}
	})

	assert.Len(t, result, 1)
	assert.Equal(t, 30, result[0].age)
}

func TestMapNonNilValues(t *testing.T) {
	slice := FunList[Person]{{
		name: "Raisa",
		age:  30,
	},
		nil,
	}
	result := slice.Map(func(p *Person) *Person {
		if p.age == 30 {
			return &Person{
				name: "New person",
				age:  20,
			}
		}
		return p
	})
	assert.Len(t, result, 2)
	assert.Equal(t, 20, result[0].age)
}
