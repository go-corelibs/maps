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
	"bytes"

	"github.com/gookit/goutil/dump"
)

var (
	gDumper *dump.Dumper
)

// Dump is a convenience wrapper around [github.com/gookit/goutil/dump]
// to return a human-readable representation of the given map
//
// Only really useful during ad-hoc development cycles
func Dump[T comparable, V interface{}](m map[T]V) (pretty string) {
	if gDumper == nil {
		gDumper = dump.NewWithOptions(
			dump.WithoutPosition(),
			dump.WithoutColor(),
			dump.WithoutOutput(&bytes.Buffer{}),
		)
	}
	var buf bytes.Buffer
	gDumper.Fprint(&buf, m)
	return buf.String()
}
