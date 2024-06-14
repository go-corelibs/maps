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
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTransform(t *testing.T) {

	Convey("RenameKeys", t, func() {

		for idx, test := range []struct {
			input  map[string]interface{}
			output map[string]interface{}
			rename func(original string) (renamed string)
		}{

			{
				input: map[string]interface{}{
					"one-two":    2,
					"many_other": 10,
				},
				output: map[string]interface{}{
					"one-two":    2,
					"many_other": 10,
				},
				rename: func(original string) (renamed string) {
					return original
				},
			},

			{
				input: map[string]interface{}{
					"one-two":    2,
					"many_other": 10,
				},
				output: map[string]interface{}{
					"ONE-TWO":    2,
					"MANY_OTHER": 10,
				},
				rename: func(original string) (renamed string) {
					return strings.ToUpper(original)
				},
			},
		} {
			RenameKeys(test.input, test.rename)
			SoMsg(fmt.Sprintf("test #%d", idx), test.input, ShouldEqual, test.output)
		}

		Convey("CamelizeKeys", func() {
			m := map[string]interface{}{
				"one-two":    2,
				"many_other": 10,
			}
			CamelizeKeys(m)
			So(m, ShouldEqual, map[string]interface{}{
				"OneTwo":    2,
				"ManyOther": 10,
			})
		})

		Convey("KebabKeys", func() {
			m := map[string]interface{}{
				"one-two":    2,
				"many_other": 10,
			}
			KebabKeys(m)
			So(m, ShouldEqual, map[string]interface{}{
				"one-two":    2,
				"many-other": 10,
			})
		})

		Convey("SnakeKeys", func() {
			m := map[string]interface{}{
				"one-two":    2,
				"many_other": 10,
			}
			SnakeKeys(m)
			So(m, ShouldEqual, map[string]interface{}{
				"one_two":    2,
				"many_other": 10,
			})
		})

		Convey("ScreamingSnakeKeys", func() {
			m := map[string]interface{}{
				"one-two":    2,
				"many_other": 10,
			}
			ScreamingSnakeKeys(m)
			So(m, ShouldEqual, map[string]interface{}{
				"ONE_TWO":    2,
				"MANY_OTHER": 10,
			})
		})

	})

	Convey("ToEnviron", t, func() {

		So(ToEnviron(map[string]string{"one": "two", "many": "other"}), ShouldEqual, []string{
			"MANY=other",
			"ONE=two",
		})

		So(ToEnviron(map[string]int{"one": 2, "many": 10}), ShouldEqual, []string{
			"MANY=10",
			"ONE=2",
		})

	})

}
