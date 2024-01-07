[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/go-corelibs/maps)
[![codecov](https://codecov.io/gh/go-corelibs/maps/graph/badge.svg?token=7gcECAViv5)](https://codecov.io/gh/go-corelibs/maps)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-corelibs/maps)](https://goreportcard.com/report/github.com/go-corelibs/maps)

# maps - go map type utilities

maps is a package for general purpose map-type things, such as getting a
list of naturally `SortedKeys`.

# Installation

``` shell
> go get github.com/go-corelibs/maps@latest
```

# Examples

## SortedKeys

``` go
func main() {
    m := map[string]struct{}{
        "one": {},
        "two": {},
        "many": {},
        "v1": {},
        "v10": {},
        "v2": {},
    }
    keys := maps.SortedKeys(m)
    // keys == []string{"many", "one", "two", "v1", "v2", "v10"}
}
```

# Go-CoreLibs

[Go-CoreLibs] is a repository of shared code between the [Go-Curses] and
[Go-Enjin] projects.

# License

```
Copyright 2023 The Go-CoreLibs Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use file except in compliance with the License.
You may obtain a copy of the license at

 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```

[Go-CoreLibs]: https://github.com/go-corelibs
[Go-Curses]: https://github.com/go-curses
[Go-Enjin]: https://github.com/go-enjin
