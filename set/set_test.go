package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	set := New("New York", "San Francisco")
	assert.Contains(t, set, "New York")
	assert.NotContains(t, set, "Tokyo")

	set.Add("Tokyo")
	assert.Contains(t, set, "Tokyo")
}

func TestRemove(t *testing.T) {
	set := New("New York", "San Francisco", "Tokyo")
	set.Remove("San Francisco")
	assert.NotContains(t, set, "San Francisco")
}

func TestHas(t *testing.T) {
	set := New("New York", "San Francisco", "Tokyo")
	assert.True(t, set.Has("New York"))
	assert.True(t, set.Has("San Francisco"))
	assert.True(t, set.Has("Tokyo"))
	assert.False(t, set.Has("London"))
}

func TestLen(t *testing.T) {
	set := New("New York", "San Francisco", "Tokyo")
	assert.Equal(t, set.Len(), 3)

	set.Add("London")
	assert.Equal(t, set.Len(), 4)
}

func TestList(t *testing.T) {
	set := New("New York", "San Francisco", "Tokyo")
	list := set.List()

	assert.Contains(t, list, "New York")
	assert.Contains(t, list, "San Francisco")
	assert.Contains(t, list, "Tokyo")
}

func TestToMap(t *testing.T) {
	set := New("New York", "San Francisco", "Tokyo")
	m := set.ToMap()

	assert.Contains(t, m, "New York")
	assert.Contains(t, m, "San Francisco")
	assert.Contains(t, m, "Tokyo")
	assert.NotContains(t, m, "London")
}

func TestUnion(t *testing.T) {
	set1 := New("New York", "San Francisco", "Tokyo")
	set2 := New("New York", "San Francisco", "London")
	set3 := set1.Union(set2)

	assert.Contains(t, set3, "New York")
	assert.Contains(t, set3, "San Francisco")
	assert.Contains(t, set3, "Tokyo")
	assert.Contains(t, set3, "London")
}

func TestIntersection(t *testing.T) {
	set1 := New("New York", "San Francisco", "Tokyo")
	set2 := New("New York", "San Francisco", "London")
	set3 := set1.Intersection(set2)

	assert.Contains(t, set3, "New York")
	assert.Contains(t, set3, "San Francisco")
	assert.NotContains(t, set3, "Tokyo")
	assert.NotContains(t, set3, "London")
}

func TestDifference(t *testing.T) {
	set1 := New("New York", "San Francisco", "Tokyo")
	set2 := New("New York", "San Francisco", "London")
	set3 := set1.Difference(set2)

	assert.NotContains(t, set3, "New York")
	assert.NotContains(t, set3, "San Francisco")
	assert.Contains(t, set3, "Tokyo")
	assert.NotContains(t, set3, "London")
}

func TestSubset(t *testing.T) {
	set1 := New("New York", "San Francisco", "Tokyo")
	set2 := New("New York", "San Francisco")

	assert.False(t, set1.Subset(set2))
	assert.True(t, set2.Subset(set1))
}

func TestSuperset(t *testing.T) {
	set1 := New("New York", "San Francisco", "Tokyo")
	set2 := New("New York", "San Francisco")

	assert.True(t, set1.Superset(set2))
	assert.False(t, set2.Superset(set1))
}

func TestEqual(t *testing.T) {
	set1 := New("New York", "San Francisco", "Tokyo")
	set2 := New("New York", "San Francisco", "Tokyo")
	set3 := New("London", "Berlin", "Paris")
	set4 := New[string]()

	assert.True(t, set1.Equal(set2))
	assert.False(t, set1.Equal(set3))
	assert.False(t, set1.Equal(set4))
}

func TestClear(t *testing.T) {
	set := New("New York", "San Francisco", "Tokyo")
	set.Clear()
	assert.Equal(t, set.Len(), 0)
}

func TestClone(t *testing.T) {
	set := New("New York", "San Francisco", "Tokyo")
	clone := set.Clone()

	assert.True(t, set.Equal(clone))
}

func TestString(t *testing.T) {
	set := New("New York", "San Francisco", "Tokyo")
	str := set.String()
	assert.Contains(t, str, "New York")
	assert.Contains(t, str, "San Francisco")
	assert.Contains(t, str, "Tokyo")
}

func TestIterate(t *testing.T) {
	set := New("New York", "San Francisco", "Tokyo")
	set.Iterate(func(elem string) bool {
		assert.True(t, set.Has(elem))
		return false
	})
}

func TestIterateAll(t *testing.T) {
	set := New("New York", "San Francisco", "Tokyo")

	set.IterateAll(func(elem string) {
		assert.True(t, set.Has(elem))
	})
}

func TestMap(t *testing.T) {
	set := New("New York", "San Francisco", "Tokyo")
	mapped := set.Map(func(elem string) string {
		return elem + "!"
	})

	assert.Contains(t, mapped, "New York!")
	assert.Contains(t, mapped, "San Francisco!")
	assert.Contains(t, mapped, "Tokyo!")
}

func TestFilter(t *testing.T) {
	set := New("New York", "San Francisco", "Tokyo")
	filtered := set.Filter(func(elem string) bool {
		return elem == "New York"
	})

	assert.Contains(t, filtered, "New York")
	assert.NotContains(t, filtered, "San Francisco")
	assert.NotContains(t, filtered, "Tokyo")
}

func TestReduce(t *testing.T) {
	set := New("New York", "San Francisco", "Tokyo")
	reduced := set.Reduce(func(acc string, elem string) string {
		return acc + elem
	})

	assert.Contains(t, reduced, "New York")
	assert.Contains(t, reduced, "San Francisco")
	assert.Contains(t, reduced, "Tokyo")
}

func TestReduceAll(t *testing.T) {
	set := New("New York", "San Francisco", "Tokyo")
	reduced := set.ReduceAll(func(acc string, elem string) string {
		return acc + elem
	}, "City: ")

	assert.Contains(t, reduced, "City: ")
	assert.Contains(t, reduced, "New York")
	assert.Contains(t, reduced, "San Francisco")
	assert.Contains(t, reduced, "Tokyo")
}

func TestAny(t *testing.T) {
	set := New("New York", "San Francisco", "Tokyo")
	any := set.Any(func(elem string) bool {
		return elem == "New York"
	})

	assert.True(t, any)

	any = set.Any(func(elem string) bool {
		return elem == "London"
	})

	assert.False(t, any)
}

func TestAll(t *testing.T) {
	set := New("New York", "San Francisco", "Tokyo")
	all := set.All(func(elem string) bool {
		return elem == "New York"
	})

	assert.False(t, all)

	all = set.All(func(elem string) bool {
		return elem == "New York" || elem == "San Francisco" || elem == "Tokyo"
	})

	assert.True(t, all)
}

func TestNone(t *testing.T) {
	set := New("New York", "San Francisco", "Tokyo")
	none := set.None(func(elem string) bool {
		return elem == "New York"
	})

	assert.False(t, none)

	none = set.None(func(elem string) bool {
		return elem == "London"
	})

	assert.True(t, none)
}

func TestFind(t *testing.T) {
	set := New("New York", "San Francisco", "Tokyo")
	elem, found := set.Find(func(elem string) bool {
		return elem == "New York"
	})

	assert.Equal(t, elem, "New York")
	assert.True(t, found)

	elem, found = set.Find(func(elem string) bool {
		return elem == "London"
	})

	assert.Equal(t, elem, "")
	assert.False(t, found)
}

func TestFindAll(t *testing.T) {
	set := New("New York", "San Francisco", "Tokyo")
	found := set.FindAll(func(elem string) bool {
		return elem == "New York" || elem == "Tokyo"
	})

	assert.Contains(t, found, "New York")
	assert.Contains(t, found, "Tokyo")
}

func TestIsEmpty(t *testing.T) {
	set := New[string]()
	assert.True(t, set.IsEmpty())

	set.Add("New York")
	assert.False(t, set.IsEmpty())
}
