package main

import (
	"fmt"

	"github.com/mateusmacedo/vibranium/validation"
	"github.com/mateusmacedo/vibranium/validation/contract"
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
    validator *validation.Composite[User]
}

func NewUserValidator() *UserValidator {
    nameValidator := &validators.StringNonEmpty{}
    ageValidator := &validators.PositiveNumber{}
    streetValidator := &validators.StringNonEmpty{}
    cityValidator := &validators.StringNonEmpty{}
    zipValidator := &validators.PositiveNumber{}

    // Compondo validações para PhoneNumber
    phoneNumberValidator := validation.NewComposite[PhoneNumber]()
    phoneNumberValidator.Add(contract.ValidatorFunc[PhoneNumber](func(phone PhoneNumber) error {
        return validators.StringNonEmpty{}.Validate(phone.Number)
    }))
    phoneNumberValidator.Add(contract.ValidatorFunc[PhoneNumber](func(phone PhoneNumber) error {
        return validators.DigitsOnly{}.Validate(phone.Number)
    }))
    phoneNumberValidator.Add(contract.ValidatorFunc[PhoneNumber](func(phone PhoneNumber) error {
        return validators.ExactLength{Length: 10}.Validate(phone.Number)
    }))

    collectionPhoneValidator := validation.NewCollection(phoneNumberValidator)

    builder := validation.NewBuilder[User]().
        Add(contract.ValidatorFunc[User](func(user User) error {
            return nameValidator.Validate(user.Name)
        })).
        Add(contract.ValidatorFunc[User](func(user User) error {
            return ageValidator.Validate(user.Age)
        })).
        Add(contract.ValidatorFunc[User](func(user User) error {
            return streetValidator.Validate(user.Address.Street)
        })).
        Add(contract.ValidatorFunc[User](func(user User) error {
            return cityValidator.Validate(user.Address.City)
        })).
        Add(contract.ValidatorFunc[User](func(user User) error {
            return zipValidator.Validate(user.Address.Zip)
        })).
        Add(contract.ValidatorFunc[User](func(user User) error {
            return collectionPhoneValidator.Validate(user.PhoneNumbers)
        }))

    return &UserValidator{validator: builder.Build()}
}

func (uv *UserValidator) Validate(user User) error {
    return uv.validator.Validate(user)
}

func main() {
    userValidator := NewUserValidator()

    // Exemplo de usuário
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

    // Valida o usuário
    if err := userValidator.Validate(user); err != nil {
        fmt.Println("Validation errors:", err)
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
        fmt.Println("Validation errors:", err)
    } else {
        fmt.Println("User is valid")
    }
}
