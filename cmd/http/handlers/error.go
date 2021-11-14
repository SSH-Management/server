package handlers

import (
	"errors"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"

	"github.com/SSH-Management/server/pkg/db"
)

type message struct {
	Message string `json:"message"`
}

func Error(logger zerolog.Logger, translator ut.Translator) fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
		code := fiber.StatusInternalServerError

		redirect := ctx.Query("redirect", "")

		logger.Error().
			Err(err).
			Msg("An error has occurred in application")

		if e, ok := err.(*fiber.Error); ok {
			if redirect != "" {
				redirect, _ = Redirect(redirect, e.Message)
				return ctx.Redirect(redirect, fiber.StatusTemporaryRedirect)
			}

			return ctx.Status(e.Code).JSON(message{
				Message: e.Message,
			})
		}

		if _, ok := err.(*validator.InvalidValidationError); ok {
			if redirect != "" {
				redirect, _ = Redirect(redirect, "Data is invalid")
				return ctx.Redirect(redirect, fiber.StatusTemporaryRedirect)
			}

			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(message{Message: "Data is invalid"})
		}

		if err, ok := err.(validator.ValidationErrors); ok {
			if redirect != "" {
				redirect, _ = Redirect(redirect, "validation error")
				return ctx.Redirect(redirect, fiber.StatusTemporaryRedirect)
			}

			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(err.Translate(translator))
		}

		if errors.Is(err, db.ErrNotFound) {
			if redirect != "" {
				redirect, _ = Redirect(redirect, "Data not found!")
				return ctx.Redirect(redirect, fiber.StatusTemporaryRedirect)
			}

			return ctx.Status(fiber.StatusNotFound).
				JSON(message{Message: "Data not found!"})
		}

		if redirect != "" {
			redirect, _ = Redirect(redirect, "An error has occurred!")
			return ctx.Redirect(redirect, fiber.StatusTemporaryRedirect)
		}

		return ctx.Status(code).
			JSON(message{Message: "An error has occurred!"})
	}
}
