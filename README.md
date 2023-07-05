# Percent

[![Go Reference](https://pkg.go.dev/badge/github.com/sentenz/percent.svg)](https://pkg.go.dev/github.com/sentenz/percent)
![Go Version](https://img.shields.io/github/go-mod/go-version/sentenz/percent)
![Go SemVer](https://img.shields.io/github/v/release/sentenz/percent)
[![Go Report](https://goreportcard.com/badge/github.com/sentenz/percent)](https://goreportcard.com/report/github.com/sentenz/percent)

The `percent` package provides a simple and easy-to-use API for working with percentages in Go.

## Installation

Use `go get` to install this package:

```shell
go get github.com/sentenz/percent
```

## Usage

To use the `percent` package in your Go program, import it like this:

```go
import "github.com/sentenz/percent/pkg/percent"
```

The percent package provides functions for working with percentages, including:

- percent.Percent()
  > Calculates the percentage of a given value.

- percent.Change()
  > Calculates the percentage change between two values.

Example of Percent:

```go
package main

import (
  "fmt"

  "github.com/sentenz/percent/pkg/percent"
)

const (
  percentage = -25
  value      = 100
)

func main() {
  result, err := percent.Percent(percentage, value)
  if err != nil {
    log.Printf("calculating percentage %v%% of %v: %v", percentage, value, err)

    return
  }

  log.Printf("%v%% of %v is %.0f\n", percentage, value, result)
}
```

This will output:

```shell
25% of 100 is 25
```

## Contributing

Contributions to the percent package are welcome! If you find a bug or have a feature request, please open an issue on the GitHub repository.

## License

The percent package is available under the MIT License. See the LICENSE file for more information.
