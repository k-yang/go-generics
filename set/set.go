package set

import (
	"fmt"
)

type Set[T comparable] map[T]struct{}

func New[T comparable](strs ...T) Set[T] {
	set := Set[T]{}
	for _, s := range strs {
		set.Add(s)
	}
	return set
}

func (set Set[T]) Add(s T) {
	set[s] = struct{}{}
}

func (set Set[T]) Remove(s T) {
	delete(set, s)
}

func (set Set[T]) Has(s T) bool {
	_, ok := set[s]
	return ok
}

func (set Set[T]) Len() int {
	return len(set)
}

func (set Set[T]) List() []T {
	var slice []T
	for s := range set {
		slice = append(slice, s)
	}
	return slice
}

func (set Set[T]) ToMap() map[T]struct{} {
	return map[T]struct{}(set)
}

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

func (set Set[T]) Intersection(other Set[T]) Set[T] {
	intersection := Set[T]{}
	for s := range set {
		if other.Has(s) {
			intersection.Add(s)
		}
	}
	return intersection
}

func (set Set[T]) Difference(other Set[T]) Set[T] {
	difference := Set[T]{}
	for s := range set {
		if !other.Has(s) {
			difference.Add(s)
		}
	}
	return difference
}

func (set Set[T]) Iterate(f func(T) bool) {
	for s := range set {
		if !f(s) {
			break
		}
	}
}

func (set Set[T]) IterateAll(f func(T)) {
	for s := range set {
		f(s)
	}
}

func (set Set[T]) Filter(f func(T) bool) Set[T] {
	filtered := Set[T]{}
	for s := range set {
		if f(s) {
			filtered.Add(s)
		}
	}
	return filtered
}

func (set Set[T]) Map(f func(T) T) Set[T] {
	mapped := Set[T]{}
	for s := range set {
		mapped.Add(f(s))
	}
	return mapped
}

func (set Set[T]) Reduce(f func(T, T) T) T {
	var reduced T
	for s := range set {
		reduced = f(reduced, s)
	}
	return reduced
}

func (set Set[T]) ReduceAll(f func(T, T) T, initial T) T {
	reduced := initial
	for s := range set {
		reduced = f(reduced, s)
	}
	return reduced
}

func (set Set[T]) Any(f func(T) bool) bool {
	for s := range set {
		if f(s) {
			return true
		}
	}
	return false
}

func (set Set[T]) All(f func(T) bool) bool {
	for s := range set {
		if !f(s) {
			return false
		}
	}
	return true
}

func (set Set[T]) None(f func(T) bool) bool {
	for s := range set {
		if f(s) {
			return false
		}
	}
	return true
}

func (set Set[T]) IsEmpty() bool {
	return len(set) == 0
}

func (set Set[T]) Clear() {
	for s := range set {
		delete(set, s)
	}
}

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

func (set Set[T]) Clone() Set[T] {
	clone := Set[T]{}
	for s := range set {
		clone.Add(s)
	}
	return clone
}

func (set Set[T]) String() string {
	return fmt.Sprint(set.List())
}

func (set Set[T]) Subset(other Set[T]) bool {
	for s := range set {
		if !other.Has(s) {
			return false
		}
	}
	return true
}

func (set Set[T]) Superset(other Set[T]) bool {
	for s := range other {
		if !set.Has(s) {
			return false
		}
	}
	return true
}

func (set Set[T]) Find(f func(T) bool) (T, bool) {
	for s := range set {
		if f(s) {
			return s, true
		}
	}

	var ret T
	return ret, false
}

func (set Set[T]) FindAll(f func(T) bool) []T {
	var ret []T
	for s := range set {
		if f(s) {
			ret = append(ret, s)
		}
	}
	return ret
}
