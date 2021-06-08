package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jason-shen/clubhouse-clone-biz/ent/user"
	"net/http"
)

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

	exist, _ := h.Client.User.Query().Where(user.Email(request.Email)).Only(ctx.Context())
	if exist != nil {
		ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "The Email is Already Taken",
		})
	}

	_, err = h.Client.User.Create().
		SetEmail(request.Email).
		SetFirstName(request.Firstname).
		SetLastName(request.Lastname).
		SetEmail(request.Email).
		SetAvatar(request.Avatar).
		SetPassword(request.Password).
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
