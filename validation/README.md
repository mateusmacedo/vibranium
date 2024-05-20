# Vibranium Validation Package

## Overview

The Vibranium Validation Package provides a set of tools and utilities for building and combining validation logic in Go. It allows you to create complex validation rules in a modular and reusable way.

## Table of Contents

- [Vibranium Validation Package](#vibranium-validation-package)
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
    cd vibranium/validation
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

You can also run the tests with coverage:

```sh
make coverage
```

### Example

Here is an example of how to use the validation framework:

```go
package main

import (
    "fmt"
    "github.com/mateusmacedo/vibranium/validation/contract"
    "github.com/mateusmacedo/vibranium/validation/errors"
    "github.com/mateusmacedo/vibranium/validation/presenter"
    "github.com/mateusmacedo/vibranium/validation/validator"
    "github.com/mateusmacedo/vibranium/validation/validators"
)

type Address struct {
    Street string
    City   string
    Zip    int
}

type PhoneNumber struct {
    Number string
}

type User struct {
    Name         string
    Age          int
    Address      Address
    PhoneNumbers []PhoneNumber
}

type UserValidator struct {
    validator *validator.Composite[User]
}

func NewUserValidator() *UserValidator {
    nameValidator := &validators.StringNonEmpty{}
    ageValidator := &validators.PositiveNumber{}
    streetValidator := &validators.StringNonEmpty{}
    cityValidator := &validators.StringNonEmpty{}
    zipValidator := &validators.PositiveNumber{}

    // Compondo validações para PhoneNumber
    phoneNumberValidator := validator.NewComposite[PhoneNumber]()
    phoneNumberValidator.AddValidator("Number", contract.ValidationFunc[PhoneNumber](func(phone PhoneNumber) error {
        return validators.StringNonEmpty{}.Validate(phone.Number)
    }))
    phoneNumberValidator.AddValidator("Number", contract.ValidationFunc[PhoneNumber](func(phone PhoneNumber) error {
        return validators.DigitsOnly{}.Validate(phone.Number)
    }))
    phoneNumberValidator.AddValidator("Number", contract.ValidationFunc[PhoneNumber](func(phone PhoneNumber) error {
        return validators.ExactLength{Length: 10}.Validate(phone.Number)
    }))

    collectionPhoneValidator := validator.NewCollection(phoneNumberValidator)

    builder := validator.NewBuilder[User]().
        Add("Name", contract.ValidationFunc[User](func(user User) error {
            return nameValidator.Validate(user.Name)
        })).
        Add("Age", contract.ValidationFunc[User](func(user User) error {
            return ageValidator.Validate(user.Age)
        })).
        Add("Address.Street", contract.ValidationFunc[User](func(user User) error {
            return streetValidator.Validate(user.Address.Street)
        })).
        Add("Address.City", contract.ValidationFunc[User](func(user User) error {
            return cityValidator.Validate(user.Address.City)
        })).
        Add("Address.Zip", contract.ValidationFunc[User](func(user User) error {
            return zipValidator.Validate(user.Address.Zip)
        })).
        Add("PhoneNumbers", contract.ValidationFunc[User](func(user User) error {
            return collectionPhoneValidator.Validate(user.PhoneNumbers)
        }))

    return &UserValidator{validator: builder.Build()}
}

func (uv *UserValidator) Validate(user User) error {
    return uv.validator.Validate(user)
}

func main() {
    userValidator := NewUserValidator()

    // Exemplo de usuário válido
    user := User{
        Name: "John",
        Age:  30,
        Address: Address{
            Street: "123 Main St",
            City:   "Springfield",
            Zip:    12345,
        },
        PhoneNumbers: []PhoneNumber{
            {Number: "1234567890"},
            {Number: "0987654321"},
        },
    }

    // Valida o usuário válido
    if err := userValidator.Validate(user); err != nil {
        presenter := &presenter.JSONPresenter{}
        fmt.Println("Validation errors:")
        fmt.Println(presenter.Present(err.(*errors.Errors)))
    } else {
        fmt.Println("User is valid")
    }

    // Exemplo de usuário inválido
    invalidUser := User{
        Name: "",
        Age:  -5,
        Address: Address{
            Street: "",
            City:   "",
            Zip:    -123,
        },
        PhoneNumbers: []PhoneNumber{
            {Number: ""},
            {Number: "98765abcde"},
        },
    }

    // Valida o usuário inválido
    if err := userValidator.Validate(invalidUser); err != nil {
        presenter := &presenter.JSONPresenter{}
        fmt.Println("Validation errors:")
        fmt.Println(presenter.Present(err.(*errors.Errors)))
    } else {
        fmt.Println("User is valid")
    }
}
```

## Features

- **Validation Interface**: Defines a contract for validators.
- **Builder Pattern**: Provides a builder interface for constructing complex validators.
- **Collection Validator**: Validate collections of items.
- **Error Presentation**: Present errors in different formats (JSON, XML).

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
