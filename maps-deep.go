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
	"strconv"

	"github.com/go-corelibs/rxp"
)

var (
	rxKeySlice = rxp.Pattern{}.
		Caret().
		Add(rxp.IsFieldKey("c")).
		Text("[").
		D("*", "c").
		Text("]").
		Dollar()
)

func ParseDeepKeySlice(input string) (key string, idx int, ok bool) {
	km := rxKeySlice.FindStringSubmatch(input)
	if idx, ok = -1, len(km) == 3; ok {
		if km[2] != "" {
			idx, _ = strconv.Atoi(km[2])
		}
		key = km[1]
	}
	return
}
