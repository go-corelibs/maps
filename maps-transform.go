// Copyright (c) 2023  The Go-Enjin Authors
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

	"github.com/iancoleman/strcase"
)

// RenameKeys ranges over the given data and calls the rename function for
// each key, if the renamed string is different from the original string,
// the renamed key is set on the data and the original key is deleted
func RenameKeys[K ~string, V interface{}](data map[K]V, rename func(original K) (renamed K)) {
	for original, v := range data {
		if renamed := rename(original); renamed != original {
			if _, present := data[renamed]; !present {
				data[renamed] = v
			}
			delete(data, original)
		}
	}
}

// CamelizeKeys is a CamelCase wrapper around RenameKeys
func CamelizeKeys[K ~string, V interface{}](data map[K]V) {
	RenameKeys(data, func(original K) (renamed K) {
		return K(strcase.ToCamel(string(original)))
	})
	return
}

// KebabKeys is a kebab-case wrapper around RenameKeys
func KebabKeys[K ~string, V interface{}](data map[K]V) {
	RenameKeys(data, func(original K) (renamed K) {
		return K(strcase.ToKebab(string(original)))
	})
	return
}

// SnakeKeys is a snake_case wrapper around RenameKeys
func SnakeKeys[K ~string, V interface{}](data map[K]V) {
	RenameKeys(data, func(original K) (renamed K) {
		return K(strcase.ToSnake(string(original)))
	})
	return
}

// ScreamingSnakeKeys converts all keys to kebab-case
//
// For each key/value pair:
// - if the key is already kebab-cased, does nothing
// - if the kebab-cased key does not exist, sets it
// - deletes the original, not-kebab-case key
func ScreamingSnakeKeys[K ~string, V interface{}](data map[K]V) {
	RenameKeys(data, func(original K) (renamed K) {
		return K(strcase.ToScreamingSnake(string(original)))
	})
	return
}

// ToEnviron transforms the map into a SortedKeys os.Environ slice of KEY=value
// pairs. ToEnviron uses fmt.Sprintf("%v") to transform the values
//
// Example:
//
//	environ := ToEnviron(map[string]interface{}{"one":"two", "many": 10})
//	// environ == []string{"ONE=two", "MANY=10"}
func ToEnviron[K ~string, V interface{}](data map[K]V) (environ []string) {
	for _, k := range SortedKeys(data) {
		key := strcase.ToScreamingSnake(string(k))
		value := fmt.Sprintf("%v", data[k])
		environ = append(environ, key+"="+value)
	}
	return
}
