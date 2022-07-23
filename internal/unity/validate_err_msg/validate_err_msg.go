package validate_err_msg

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

type ErrorMessage struct {
	Param   string
	Message string
}

type ValidateErrorMessages []ErrorMessage

func (e ValidateErrorMessages) Error() string {
	switch len(e) {
	case 0:
		return ""
	case 1:
		return "Key: '" + e[0].Param + "' Error: " + e[0].Message
	}

	var result string
	result = "Key: '" + e[0].Param + "' Error: " + e[0].Message
	for _, item := range e[1:] {
		result = result + "\n" + "Key: '" + item.Param + "' Error: " + item.Message
	}
	return result
}

func Transfer(err error) error {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make(ValidateErrorMessages, 0)
		for _, item := range ve {
			out = append(out, ErrorMessage{
				Param:   item.Field(),
				Message: transferMsgTag(item),
			})
		}
		return out
	}
	return err
}

func transferMsgTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "uuid":
		return "Invalid UUID"
	case "numeric":
		return "Field must be numeric"
	case "email":
		return "Invalid email"
	}
	return fe.Error() // default error
}
