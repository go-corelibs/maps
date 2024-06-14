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
	"sort"
	"testing"

	"github.com/maruel/natural"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSort(t *testing.T) {

	Convey("Keys", t, func() {
		m := map[string]struct{}{
			"one":  {},
			"two":  {},
			"many": {},
			"v1":   {},
			"v10":  {},
			"v2":   {},
		}
		keys := Keys(m)
		sort.Sort(natural.StringSlice(keys))
		So(keys, ShouldEqual, []string{
			"many", "one", "two", "v1", "v2", "v10",
		})
	})

	Convey("SortedKeys", t, func() {
		m := map[string]struct{}{
			"one":  {},
			"two":  {},
			"many": {},
			"v1":   {},
			"v10":  {},
			"v2":   {},
		}
		So(SortedKeys(m), ShouldEqual, []string{
			"many", "one", "two", "v1", "v2", "v10",
		})
	})

	Convey("SortedNumbers", t, func() {
		mInt := map[int]struct{}{
			2: {},
			1: {},
			0: {},
		}
		So(SortedNumbers(mInt), ShouldEqual, []int{
			0, 1, 2,
		})
		mFloat := map[float32]struct{}{
			2.5: {},
			1.0: {},
			3.4: {},
		}
		So(SortedNumbers(mFloat), ShouldEqual, []float32{
			1.0, 2.5, 3.4,
		})
	})

	Convey("ReverseSortedNumbers", t, func() {
		mInt := map[int]struct{}{
			2: {},
			1: {},
			0: {},
		}
		So(ReverseSortedNumbers(mInt), ShouldEqual, []int{
			2, 1, 0,
		})
		mFloat := map[float32]struct{}{
			2.5: {},
			1.0: {},
			3.4: {},
		}
		So(ReverseSortedNumbers(mFloat), ShouldEqual, []float32{
			3.4, 2.5, 1.0,
		})
	})

	Convey("ValuesSortedByKeys", t, func() {
		m := map[string]string{
			"v0.0.1":  "first",
			"v0.0.2":  "second",
			"v0.0.10": "last",
		}
		So(ValuesSortedByKeys(m), ShouldEqual, []string{
			"first", "second", "last",
		})
	})

	Convey("ValuesSortedByNumbers", t, func() {
		m := map[int]string{
			1:  "first",
			2:  "second",
			10: "last",
		}
		So(ValuesSortedByNumbers(m), ShouldEqual, []string{
			"first", "second", "last",
		})
	})

	Convey("SortedKeyLengths", t, func() {
		m := map[string]struct{}{
			"one":  {},
			"more": {},
			"go":   {},
			"two":  {},
		}
		So(SortedKeyLengths(m), ShouldEqual, []string{
			"more", "one", "two", "go",
		})
	})

	Convey("ReverseSortedKeys", t, func() {
		m := map[string]struct{}{
			"one":  {},
			"two":  {},
			"many": {},
			"v1":   {},
			"v10":  {},
			"v2":   {},
		}
		So(ReverseSortedKeys(m), ShouldEqual, []string{
			"v10", "v2", "v1", "two", "one", "many",
		})
	})

	Convey("OrderedKeys", t, func() {
		m := map[int]struct{}{
			20: {},
			1:  {},
			10: {},
		}
		So(OrderedKeys(m), ShouldEqual, []int{
			1, 10, 20,
		})
	})

	Convey("ReverseOrderedKeys", t, func() {
		m := map[int]struct{}{
			20: {},
			1:  {},
			10: {},
		}
		So(ReverseOrderedKeys(m), ShouldEqual, []int{
			20, 10, 1,
		})
	})

	Convey("SortedKeysByLastName", t, func() {
		m := map[string]struct{}{
			"Alan Watts":                {},
			"John Ronald Reuel Tolkien": {},
			"Frankly Orwell":            {},
			"George Orwell":             {},
			"Gene Roddenberry":          {},
		}
		So(SortedKeysByLastName(m), ShouldEqual, []string{
			"Frankly Orwell", "George Orwell", "Gene Roddenberry", "John Ronald Reuel Tolkien", "Alan Watts",
		})
	})

}
