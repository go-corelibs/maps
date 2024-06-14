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
	"reflect"
)

// MakeTypedKey is used to simplify the adding of more map values to a parent
// map without having to check if the value map exists already or needs to be
// created first
//
// Example:
//
//	// Standard way
//	m := make(map[string]map[string]struct{})
//	if _, present := m["top"]; !present {
//	  m["top"] = make(map[string]struct{})
//	}
//	m["top"]["thing"] = struct{}{}
//
//	// Using MakeTypedKey
//	m := make(map[string]map[string]struct{})
//	_ = maps.MakeTypedKey(m, "top")
//	m["top"]["thing"] = struct{}{}
func MakeTypedKey[K comparable, L comparable, V interface{}, M map[L]V](m map[K]M, key K) (made bool) {
	if _, present := m[key]; !present {
		var l L
		var v V
		kt, vt := reflect.TypeOf(l), reflect.TypeOf(v)
		mt := reflect.MapOf(kt, vt)
		mv := reflect.MakeMapWithSize(mt, 0)
		mi := mv.Interface()
		m[key], _ = mi.(M)
		return true
	}
	return
}
