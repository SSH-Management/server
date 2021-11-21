package container

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
)

func (c *Container) GetValidator() *validator.Validate {
	if c.validator == nil && c.translator == nil {
		english := en.New()
		uni := ut.New(english, english)

		c.translator, _ = uni.GetTranslator("en")
		c.validator = validator.New()

		if err := entranslations.RegisterDefaultTranslations(c.validator, c.translator); err != nil {
			c.GetDefaultLogger().Fatal().Err(err).Msg("Error while registering english translations")
		}
	}

	return c.validator
}
