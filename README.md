# Vibranium Specification Pattern

## Overview

The Vibranium Specification Pattern project provides a set of tools and utilities for building and combining specifications in Go. It implements the Specification design pattern, allowing you to create complex business rules and validation logic in a modular and reusable way.

## Table of Contents

- [Vibranium Specification Pattern](#vibranium-specification-pattern)
  - [Overview](#overview)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
    - [Steps](#steps)
  - [Usage](#usage)
    - [Running Tests](#running-tests)
    - [Example](#example)
  - [Features](#features)
  - [Contributing](#contributing)
    - [Code of Conduct](#code-of-conduct)
  - [License](#license)

## Installation

To install the project, you need to have [Go](https://golang.org/doc/install) installed on your system.

### Steps

1. Clone the repository:
    ```sh
    git clone https://github.com/mateusmacedo/vibranium.git
    ```
2. Navigate to the project directory:
    ```sh
    cd vibranium/specification
    ```
3. Install dependencies:
    ```sh
    go mod tidy
    ```

## Usage

### Running Tests

You can run the tests using the provided Makefile:
```sh
make test
```

### Example

Here is an example of how to use the specification pattern:

```go
package main

import (
    "fmt"
    "github.com/mateusmacedo/vibranium/specification"
    "github.com/mateusmacedo/vibranium/specification/contract"
)

// Define a candidate type
type Product struct {
    Price float64
}

// Define a specification
type PriceSpecification struct {
    MinPrice float64
}

func (s *PriceSpecification) IsSatisfiedBy(candidate Product) bool {
    return candidate.Price >= s.MinPrice
}

func main() {
    product := Product{Price: 100.0}
    spec := &PriceSpecification{MinPrice: 50.0}

    fmt.Println(spec.IsSatisfiedBy(product)) // Output: true
}
```

## Features

- **Specification Interface**: Defines a contract for specifications.
- **Builder Pattern**: Provides a builder interface for constructing complex specifications.
- **Logical Operators**: Combine specifications using logical AND, OR, and NOT operators.

## Contributing

Contributions are welcome! Please follow these steps to contribute:

1. Fork the repository.
2. Create a new branch:
    ```sh
    git checkout -b feature/your-feature
    ```
3. Make your changes.
4. Commit your changes:
    ```sh
    git commit -m 'Add some feature'
    ```
5. Push to the branch:
    ```sh
    git push origin feature/your-feature
    ```
6. Open a pull request.

### Code of Conduct

Please adhere to the project's [Code of Conduct](./CODE_OF_CONDUCT.md).

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.