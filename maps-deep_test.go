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
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDeep(t *testing.T) {
	Convey("ParseDeepKeySlice", t, func() {

		for idx, test := range []struct {
			input string
			key   string
			index int
			ok    bool
		}{
			{"this[10]", "this", 10, true},
			{"this-thing[10]", "this-thing", 10, true},
			{"this[]", "this", -1, true},
			{"this[nope]", "", -1, false},
		} {
			prefix := fmt.Sprintf("test #%d ", idx)
			key, index, ok := ParseDeepKeySlice(test.input)
			SoMsg(prefix+"(ok)", ok, ShouldEqual, test.ok)
			SoMsg(prefix+"(key)", key, ShouldEqual, test.key)
			SoMsg(prefix+"(idx)", index, ShouldEqual, test.index)
		}

	})
}
