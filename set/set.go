package set

import (
	"fmt"
)

type Set[T comparable] map[T]struct{}

// New returns a new set with the given elements.
func New[T comparable](strs ...T) Set[T] {
	set := Set[T]{}
	for _, s := range strs {
		set.Add(s)
	}
	return set
}

// Add adds the given element to the set.
func (set Set[T]) Add(s T) {
	set[s] = struct{}{}
}

// Remove removes the given element from the set.
func (set Set[T]) Remove(s T) {
	delete(set, s)
}

// Has returns true if the given element is in the set.
func (set Set[T]) Has(s T) bool {
	_, ok := set[s]
	return ok
}

// Len returns the number of elements in the set.
func (set Set[T]) Len() int {
	return len(set)
}

// List returns the elements as a slice.
func (set Set[T]) List() []T {
	var slice []T
	for s := range set {
		slice = append(slice, s)
	}
	return slice
}

// ToMap returns the elements as a map.
func (set Set[T]) ToMap() map[T]struct{} {
	return map[T]struct{}(set)
}

// Union returns the union of the set and the given set.
func (set Set[T]) Union(other Set[T]) Set[T] {
	union := Set[T]{}
	for s := range set {
		union.Add(s)
	}
	for s := range other {
		union.Add(s)
	}
	return union
}

// Intersection returns the intersection of the set and the given set.
func (set Set[T]) Intersection(other Set[T]) Set[T] {
	intersection := Set[T]{}
	for s := range set {
		if other.Has(s) {
			intersection.Add(s)
		}
	}
	return intersection
}

// Difference returns the difference of the set and the given set.
func (set Set[T]) Difference(other Set[T]) Set[T] {
	difference := Set[T]{}
	for s := range set {
		if !other.Has(s) {
			difference.Add(s)
		}
	}
	return difference
}

// Iterate calls the given function for each element in the set, stopping if the function returns true.
func (set Set[T]) Iterate(f func(T) bool) {
	for s := range set {
		if f(s) {
			break
		}
	}
}

// IterateAll calls the given function for each element in the set.
func (set Set[T]) IterateAll(f func(T)) {
	for s := range set {
		f(s)
	}
}

// Filter returns a new set with the elements for which the given function returns true.
func (set Set[T]) Filter(f func(T) bool) Set[T] {
	filtered := Set[T]{}
	for s := range set {
		if f(s) {
			filtered.Add(s)
		}
	}
	return filtered
}

// Map returns a new set with the elements transformed by the given function.
func (set Set[T]) Map(f func(T) T) Set[T] {
	mapped := Set[T]{}
	for s := range set {
		mapped.Add(f(s))
	}
	return mapped
}

// Reduce returns the result of applying the given function to each element in the set, starting with the first element.
func (set Set[T]) Reduce(f func(T, T) T) T {
	var reduced T
	for s := range set {
		reduced = f(reduced, s)
	}
	return reduced
}

// ReduceAll returns the result of applying the given function to each element in the set, starting with the given initial value.
func (set Set[T]) ReduceAll(f func(T, T) T, initial T) T {
	reduced := initial
	for s := range set {
		reduced = f(reduced, s)
	}
	return reduced
}

// Any returns true if the given function returns true for any element in the set.
func (set Set[T]) Any(f func(T) bool) bool {
	for s := range set {
		if f(s) {
			return true
		}
	}
	return false
}

// All returns true if the given function returns true for all elements in the set.
func (set Set[T]) All(f func(T) bool) bool {
	for s := range set {
		if !f(s) {
			return false
		}
	}
	return true
}

// None returns true if the given function returns true for no elements in the set.
func (set Set[T]) None(f func(T) bool) bool {
	for s := range set {
		if f(s) {
			return false
		}
	}
	return true
}

// IsEmpty returns true if the set is empty.
func (set Set[T]) IsEmpty() bool {
	return len(set) == 0
}

// Clear removes all elements from the set.
func (set Set[T]) Clear() {
	for s := range set {
		delete(set, s)
	}
}

// Equal returns true if the set is equal to the given set.
func (set Set[T]) Equal(other Set[T]) bool {
	if set.Len() != other.Len() {
		return false
	}
	for s := range set {
		if !other.Has(s) {
			return false
		}
	}
	return true
}

// Clone returns a copy of the set.
func (set Set[T]) Clone() Set[T] {
	clone := Set[T]{}
	for s := range set {
		clone.Add(s)
	}
	return clone
}

// String returns a string representation of the set.
func (set Set[T]) String() string {
	return fmt.Sprint(set.List())
}

// Subset returns true if the set is a subset of the given set.
func (set Set[T]) Subset(other Set[T]) bool {
	for s := range set {
		if !other.Has(s) {
			return false
		}
	}
	return true
}

// Superset returns true if the set is a superset of the given set.
func (set Set[T]) Superset(other Set[T]) bool {
	for s := range other {
		if !set.Has(s) {
			return false
		}
	}
	return true
}

// Find returns the first element in the set for which the given function returns true.
func (set Set[T]) Find(f func(T) bool) (T, bool) {
	for s := range set {
		if f(s) {
			return s, true
		}
	}

	var ret T
	return ret, false
}

// FindAll returns all elements in the set for which the given function returns true.
func (set Set[T]) FindAll(f func(T) bool) []T {
	var ret []T
	for s := range set {
		if f(s) {
			ret = append(ret, s)
		}
	}
	return ret
}
