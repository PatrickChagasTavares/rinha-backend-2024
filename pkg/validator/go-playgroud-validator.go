package validator

import (
	v10 "github.com/go-playground/validator/v10"
	"github.com/patrickchagastavares/rinha-backend-2024/internal/entities"
)

type validator struct {
	validate *v10.Validate
}

func New() Validator {
	validate := v10.New()

	// Registre uma função de validação personalizada para o formato de data.
	validate.RegisterValidation("typeTransaction", func(fl v10.FieldLevel) bool {
		dataStr := fl.Field().String()

		switch entities.Tipo(dataStr) {
		case entities.TipoCredito, entities.TipoDebito:
			return true
		default:
			return false
		}
	})

	return &validator{
		validate: validate,
	}
}

// Validate is used to validate a struct using rules defined in the 'validate' tag.
// This method return the struct *ValidationError that contains details of the rules violation.
// ValidationError is compatible with 'error' interface and can be returned as error.
func (v *validator) Validate(val any) error {
	return v.validate.Struct(val)
}
