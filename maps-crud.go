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

	"github.com/iancoleman/strcase"
)

// Has checks for the presence of the (deep or exact) key given
func Has(m map[string]interface{}, key string) (present bool) {
	if key == "" || key[0] != '.' {
		// not a deep key prefix, simple check
		_, present = m[key]
		return
	}
	// split after the prefixing dot
	if keys := strings.Split(key[1:], "."); len(keys) == 1 {
		present = Has(m, keys[0])
	} else if v, found := m[keys[0]]; found {
		if ms, strOk := v.(map[string]string); strOk {
			_, present = ms[keys[1]]
		} else if mm, ok := v.(map[string]interface{}); ok {
			present = Has(mm, "."+strings.Join(keys[1:], "."))
		}
	}
	return
}

// Set checks if the key being set is a deep key and if so, sets the value
// deeply within the map structure. Set is a simple map value assignment for
// non-deep key cases
func Set(m map[string]interface{}, key string, value interface{}) (err error) {
	if key == "" || key[0] != '.' {
		m[key] = value
		return
	}
	keys := strings.Split(key[1:], ".")
	keysLen := len(keys)
	switch keysLen {
	case 0: // nop

	case 1:
		// direct
		if name, idx, ok := ParseDeepKeySlice(keys[0]); ok {
			if vs := Get(m, name); vs != nil {
				if slice, ok := vs.([]interface{}); ok {
					count := len(slice)
					if idx == -1 || idx == count {
						// append
						slice = append(slice, value)
					} else if idx > count {
						for i := count; i < idx; i++ {
							slice = append(slice, nil)
						}
						slice = append(slice, value)
					} else {
						// overwrite
						slice[idx] = value
					}
					err = Set(m, name, slice)
				} else {
					// hmm
				}
			} else {
				var slice []interface{}
				if idx <= 0 {
					slice = append(slice, value)
				} else {
					for i := 0; i < idx; i++ {
						slice = append(slice, nil)
					}
					slice = append(slice, value)
				}
				err = Set(m, name, slice)
			}
		} else {
			err = Set(m, keys[0], value)
		}
		return

	default:
		// recursive

		if name, idx, ok := ParseDeepKeySlice(keys[0]); ok {
			var list []interface{}
			if v, ok := m[name]; ok {
				if vl, ok := v.([]interface{}); ok {
					list = vl
				} else {
					err = fmt.Errorf("unexpected sub-context list type: %T", v)
					return
				}
			} else {
				list = make([]interface{}, 0)
				if err = Set(m, name, list); err != nil {
					return
				}
			}

			var mm map[string]interface{}
			count := len(list)
			if idx == -1 || idx >= count {
				// append
				if idx >= count {
					// append more
					for i := count - 1; i < idx; i++ {
						list = append(list, make(map[string]interface{}))
					}
				}
				mm, _ = list[idx].(map[string]interface{})
				if err = Set(m, name, list); err != nil {
					return
				}
			} else {
				// existing
				if vm, ok := list[idx].(map[string]interface{}); ok {
					mm = vm
				} else {
					err = fmt.Errorf("unexpected sub-context value type: %T", list[idx])
					return
				}
			}

			err = Set(mm, "."+strings.Join(keys[1:], "."), value)
			return
		}

		var mm map[string]interface{}
		if v, ok := m[keys[0]]; ok {
			if t, ok := v.(map[string]interface{}); ok {
				mm = t
			} else {
				err = fmt.Errorf("unexpected sub-context value type: %T", v)
				return
			}
		} else {
			mm = make(map[string]interface{})
			if err = Set(m, keys[0], mm); err != nil {
				return
			}
		}
		err = Set(mm, "."+strings.Join(keys[1:], "."), value)
	}
	return
}

func Get(m map[string]interface{}, key string) (value interface{}) {
	if key == "" {
		return
	} else if key != "" && key[0] != '.' {
		if v, ok := m[key]; ok {
			value = v
		}
		return
	}
	keys := strings.Split(key[1:], ".")
	switch len(keys) {
	case 0: // nop
	case 1:
		if name, idx, ok := ParseDeepKeySlice(keys[0]); ok {
			if vs := Get(m, name); vs != nil {
				if slice, ok := vs.([]interface{}); ok {
					if len(slice) < idx {
						return
					}
					value = slice[idx]
				}
			}
		} else {
			value = Get(m, keys[0])
		}
	default:
		if v, ok := m[keys[0]]; ok {
			if ms, ok := v.(map[string]string); ok {
				value, _ = ms[keys[1]]
			} else if mm, ok := v.(map[string]interface{}); ok {
				value = Get(mm, "."+strings.Join(keys[1:], "."))
			}
		}
	}
	return
}

func Delete(m map[string]interface{}, key string) {
	if key == "" {
		return
	} else if key != "" && key[0] != '.' {
		if _, ok := m[key]; ok {
			delete(m, key)
		}
		return
	}
	keys := strings.Split(key[1:], ".")
	switch len(keys) {
	case 0: // nop
	case 1:
		if name, idx, ok := ParseDeepKeySlice(keys[0]); ok {
			if vs := Get(m, name); vs != nil {
				if slice, ok := vs.([]interface{}); ok {
					count := len(slice)
					if idx > -1 && idx < count {
						newSlice := make([]interface{}, 0)
						newSlice = append(newSlice, slice[:idx]...)
						if idx < count-1 {
							newSlice = append(newSlice, slice[idx+1:]...)
						}
						_ = Set(m, name, newSlice)
					}
				}
			}
		} else {
			Delete(m, keys[0])
		}
	default:
		if v, ok := m[keys[0]]; ok {
			if ms, ok := v.(map[string]string); ok {
				delete(ms, keys[1])
			} else if mm, ok := v.(map[string]interface{}); ok {
				Delete(mm, "."+strings.Join(keys[1:], "."))
			}
		}
	}
	return
}

// GetKV finds the relative key given, returning both the actual key and the
// associated interface{} value, using the following procedure:
//
// 1. uses Get if the first character of the key is a period (".")
// 2. checks if the key is already an exact match
// 3. checks for a CamelCased version of the key
// 4. checks for a kebab-cased version of the key
// 5. checks for a snake_cased version of the key
func GetKV(c map[string]interface{}, key string) (k string, v interface{}) {
	if key != "" && key[0] == '.' {
		k = key
		v = Get(c, key)
		return
	}

	k = key
	var ok bool
	if v, ok = c[k]; ok {
		return
	}
	k = strcase.ToCamel(key)
	if v, ok = c[k]; ok {
		return
	}
	k = strcase.ToKebab(key)
	if v, ok = c[k]; ok {
		return
	}
	k = strcase.ToSnake(key)
	if v, ok = c[k]; ok {
		return
	}
	return
}

// SetKV uses Set if the first character of the key is a period ("."), otherwise
// SetKV sets the map value to a CamelCased version of the key
func SetKV(c map[string]interface{}, key string, value interface{}) (err error) {
	if key != "" && key[0] == '.' {
		err = Set(c, key, value)
		return
	}
	k := strcase.ToCamel(key)
	c[k] = value
	return
}

// DeleteKV uses Delete if the first character of the key is a period ("."),
// otherwise DeleteKV uses GetKV to find the correct key and then deletes it
// from the map
func DeleteKV(c map[string]interface{}, key string) (deleted bool) {
	if key != "" && key[0] == '.' {
		Delete(c, key)
		return
	}
	if k, v := GetKV(c, key); v != nil {
		delete(c, k)
		return true
	}
	return false
}
