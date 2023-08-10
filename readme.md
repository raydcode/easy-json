# easy-json - Simplified JSON Parsing for Go

The `easy-json` package provides a collection of utility functions designed to simplify the process of JSON parsing within Go applications. It facilitates effortless conversion of JSON strings into Go data structures, eases the extraction of values from JSON objects, and simplifies the management of intricate JSON hierarchies.

## Features

- Effortlessly parse JSON strings and transform them into Go-native data structures.
- Seamlessly retrieve string, integer, array, and object values from JSON data.
- Simplify the handling of deeply nested and complex JSON structures.
- Built-in error handling for gracefully managing missing keys and incorrect data types.

## Installation

To seamlessly integrate the `easy-json` package into your Go project, import it using the following import path:

```go
import "github.com/raydcode/easy-json"
```

## Usage

- Consider the following basic example of utilizing the easy-json package:

```go
package main

import (
	"fmt"
	"github.com/raydcode/easy-json"
)

func main() {
	jsonStr := `{"name": "Ray", "age": 25}`

	data, err := easyjson.ParseString(jsonStr)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	name, _ := easyjson.GetString(data, "name")
	age, _ := easyjson.GetInt(data, "age")

	fmt.Printf("Name: %s, Age: %d\n", name, age)
}


```

## Contributions and Issues

- Contributions, bug reports, and feature requests are heartily encouraged! In case you encounter any issues or have ideas for improvements, please open an issue on GitHub.

## License

-This package is released under the MIT License.


[![PkgGoDev](https://pkg.go.dev/badge/github.com/raydcode/easy-json)](https://pkg.go.dev/github.com/raydcode/easy-json)
