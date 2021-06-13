package handlers

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/gofiber/fiber/v2"
	"github.com/jason-shen/clubhouse-clone-biz/ent/user"
	"github.com/jason-shen/clubhouse-clone-biz/utils"
	"net/http"
)

func (r registerRequest) validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Firstname, validation.Required, validation.Length(3, 20)),
		validation.Field(&r.Lastname, validation.Required, validation.Length(3, 20)),
		validation.Field(&r.Email, validation.Required, is.Email),
		validation.Field(&r.Password, validation.Required, validation.Length(6, 12)),
	)
}

func (h *Handler) UserRegister(ctx *fiber.Ctx) error {
	var request registerRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		err = ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid Json",
		})
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
		}
	}

	if err = request.validate(); err != nil {
		ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err,
		})
		return nil
	}

	exist, _ := h.Client.User.Query().Where(user.Email(request.Email)).Only(ctx.Context())
	if exist != nil {
		ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "The Email is Already Taken",
		})
	}

	hashpassword, err := utils.HashPassword(request.Password)
	if err != nil {
		fmt.Errorf("failed hash user password:", err)
		return nil
	}

	_, err = h.Client.User.Create().
		SetEmail(request.Email).
		SetFirstName(request.Firstname).
		SetLastName(request.Lastname).
		SetEmail(request.Email).
		SetAvatar(request.Avatar).
		SetPassword(hashpassword).
		Save(ctx.Context())
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "registered error.",
		})
		return nil
	}

	_ = ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"error":   false,
		"message": "registered successfully",
	})

	return nil
}
