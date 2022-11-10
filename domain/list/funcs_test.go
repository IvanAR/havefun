package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Person struct {
	name string
	age  int
}

func TestFilterWithFoundValue(t *testing.T) {
	slice := FunList[Person]{{
		name: "Galadriel",
		age:  30,
	},
		{
			name: "Frodo",
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

func TestFilterWithNOTFoundValue(t *testing.T) {
	slice := FunList[Person]{{
		name: "Galadriel",
		age:  30,
	},
		{
			name: "Frodo",
			age:  32,
		},
	}
	result := slice.Filter(func(p *Person) bool {
		if p.age == 40 {
			return true
		}
		return false
	})

	assert.Len(t, result, 0)
}

func TestFilterWithWithMoreThanOneFoundValue(t *testing.T) {
	slice := FunList[Person]{{
		name: "Galadriel",
		age:  30,
	},
		{
			name: "Gandalf",
			age:  32,
		},
		{
			name: "Frodo",
			age:  32,
		},
	}
	result := slice.Filter(func(p *Person) bool {
		if p.age == 32 {
			return true
		}
		return false
	})

	assert.Len(t, result, 2)
	for _, p := range result {
		assert.Equal(t, 32, p.age)
	}
}

func TestMap(t *testing.T) {
	slice := FunList[Person]{{
		name: "Galadriel",
		age:  30,
	},
		{
			name: "Frodo",
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
		name: "Galadriel",
		age:  30,
	},
		{
			name: "Frodo",
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
		name: "Galadriel",
		age:  30,
	},
		nil,
	}
	result := slice.MapNonNil(func(p *Person) Person {
		if p.age == 30 {
			return Person{
				name: "New person",
				age:  20,
			}
		}
		return *p
	})
	assert.Len(t, result, 1)
	assert.Equal(t, 20, result[0].age)
}

func TestAnyFunMustReturnFirstAssertion(t *testing.T) {
	slice := FunList[Person]{{
		name: "Galadriel",
		age:  30,
	},
		nil,
		{
			name: "Frodo",
			age:  32,
		},
	}

	result := slice.Find(func(p *Person) bool {
		if p != nil && p.age == 32 {
			return true
		}
		return false
	}).GetOrDefault(func() Person { return Person{} })

	assert.Equal(t, result.age, 32)
}
