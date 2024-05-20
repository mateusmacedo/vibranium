package contract

import "github.com/mateusmacedo/vibranium/validation/errors"

type Presenter interface {
    Present(errors *errors.Errors) string
}
