// Copyright (c) 2024  The Go-CoreLibs Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package maps

import (
	"cmp"
	"sort"

	"github.com/maruel/natural"

	"github.com/go-corelibs/maths"
	"github.com/go-corelibs/strings"
)

// Keys returns the list of map keys, in whatever order Go produces from a
// simple range over the data
func Keys[K comparable, V interface{}](data map[K]V) (keys []K) {
	for key := range data {
		keys = append(keys, key)
	}
	return
}

// SortedKeys returns a slice of natural-sorted keys from the given map
func SortedKeys[K ~string, V interface{}](data map[K]V) (keys []K) {
	for key := range data {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) (less bool) {
		less = natural.Less(string(keys[i]), string(keys[j]))
		return
	})
	return
}

// SortedNumbers returns a slice of ascending sorted keys from the given map
func SortedNumbers[K maths.Number, V interface{}](data map[K]V) (keys []K) {
	for key := range data {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	return
}

// ReverseSortedNumbers returns a slice of descending sorted keys from the given map
func ReverseSortedNumbers[K maths.Number, V interface{}](data map[K]V) (keys []K) {
	for key := range data {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})
	return
}

// ValuesSortedByKeys returns a slice of values, ordered by SortedKeys
func ValuesSortedByKeys[K ~string, V interface{}](data map[K]V) (values []V) {
	for _, k := range SortedKeys(data) {
		values = append(values, data[k])
	}
	return
}

// ValuesSortedByNumbers returns a slice of values, ordered by SortedNumbers
func ValuesSortedByNumbers[K maths.Number, V interface{}](data map[K]V) (values []V) {
	for _, k := range SortedNumbers(data) {
		values = append(values, data[k])
	}
	return
}

// SortedKeyLengths returns the list of keys, from longest to shortest and
// natural sorted for same length keys
func SortedKeyLengths[K ~string, V interface{}](data map[K]V) (keys []K) {
	for key := range data {
		keys = append(keys, key)
	}
	// longest -> shortest, natural sorted when same lengths
	sort.Slice(keys, func(i, j int) (less bool) {
		if il, jl := len(keys[i]), len(keys[j]); il == jl {
			less = natural.Less(string(keys[i]), string(keys[j]))
		} else {
			less = il > jl // long > short
		}
		return
	})
	return
}

// ReverseSortedKeys returns the list of keys in reverse natural sorted order
func ReverseSortedKeys[K ~string, V interface{}](data map[K]V) (keys []K) {
	for key := range data {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) (less bool) {
		less = natural.Less(string(keys[j]), string(keys[i]))
		return
	})
	return
}

// OrderedKeys returns the list of cmp.Ordered keys in ascending order
func OrderedKeys[K cmp.Ordered, V interface{}](data map[K]V) (keys []K) {
	for key := range data {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) (less bool) {
		less = cmp.Less(keys[i], keys[j]) // i vs j
		return
	})
	return
}

// ReverseOrderedKeys returns the list of cmp.Ordered keys in descending order
func ReverseOrderedKeys[K cmp.Ordered, V interface{}](data map[K]V) (keys []K) {
	for key := range data {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) (less bool) {
		less = cmp.Less(keys[j], keys[i]) // j vs i
		return
	})
	return
}

// SortedKeysByLastName assumes that the keys in the map are human names and
// uses [strings.LastName] produce a list of keys which are sorted by last name
// and uses a natural sorting for keys where the last names are the same
func SortedKeysByLastName[K ~string, V interface{}](data map[K]V) (keys []K) {
	lookup := make(map[K]string)
	for key := range data {
		lookup[key] = strings.LastName(string(key))
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) (less bool) {
		a, b := keys[i], keys[j]
		la, lb := lookup[a], lookup[b]
		if la == lb {
			less = natural.Less(string(a), string(b))
			return
		}
		less = natural.Less(la, lb)
		return less
	})
	return
}
