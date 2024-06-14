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
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCopy(t *testing.T) {
	Convey("CopyBaseType", t, func() {
		a := map[string]int{"one": 1, "two": 2}
		b := map[string]int{"one": 1, "two": 2}
		So(CopyBaseType(a), ShouldEqual, b)
	})

	Convey("DeepCopy", t, func() {
		a := map[string]interface{}{"one": 1, "two": "2", "many": map[string]int{"more": -1}}
		b := map[string]interface{}{"one": 1, "two": "2", "many": map[string]int{"more": -1}}
		So(DeepCopy(a), ShouldEqual, b)
	})
}
