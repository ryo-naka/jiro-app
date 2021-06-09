package v1

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"artics-api/src/internal/application"
	"artics-api/src/internal/application/command"
	"artics-api/src/internal/application/query"
	"artics-api/src/internal/domain"
	"artics-api/src/internal/usecase"
	"artics-api/src/internal/usecase/request"
	"artics-api/src/internal/usecase/response"
	"artics-api/src/pkg"
)

// V1UserHandler - v1 user handler
type V1UserHandler interface {
	Create(c *fiber.Ctx) error
	Show(c *fiber.Ctx) error
	Followings(c *fiber.Ctx) error
	Followers(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Suspend(c *fiber.Ctx) error
}

type v1UserHandler struct {
	u   usecase.UserUsecase
	app application.Application
}

// NewV1UserHandler - setups v1 user handler
func NewV1UserHandler(u usecase.UserUsecase, app application.Application) V1UserHandler {
	return &v1UserHandler{u, app}
}

func (h *v1UserHandler) Create(c *fiber.Ctx) error {
	req := &request.CreateUser{}
	if err := c.BodyParser(req); err != nil {
		return domain.UnableParseJSON.New(err)
	}

	if err := h.app.Commands.CreateUser.Handle(pkg.Context{Ctx: c}, command.CreateUser{
		Email:                req.Email,
		Password:             req.Password,
		PasswordConfirmation: req.PasswordConfirmation,
	}); err != nil {
		return err
	}

	return c.JSON(nil)
}

func (h *v1UserHandler) Show(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return domain.UnableParseJSON.New(err)
	}

	ctx := pkg.Context{Ctx: c}
	u, err := h.app.Queries.GetUser.Handle(ctx, id)
	if err != nil {
		return err
	}

	favoriteContents, err := h.app.Queries.GetFavoriteContents.Handle(ctx, query.GetFavoriteContents{
		UserID: id,
		Limit:  3,
	})

	resFavoriteContents := make([]*response.Content, len(favoriteContents))
	for i, c := range favoriteContents {
		resFavoriteContents[i] = &response.Content{
			ID:    c.ID,
			Title: c.Title,
		}
	}

	res := &response.ShowUser{
		ID:               u.ID,
		Nickname:         u.Nickname,
		Followingcount:   u.FollowingCount,
		Followercount:    u.FollowerCount,
		FavoriteContents: resFavoriteContents,
	}

	return c.JSON(res)
}

func (h *v1UserHandler) Followings(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return domain.UnableParseJSON.New(err)
	}

	res, err := h.u.Followings(pkg.Context{Ctx: c}, id)
	if err != nil {
		return err
	}

	return c.JSON(res)
}

func (h *v1UserHandler) Followers(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return domain.UnableParseJSON.New(err)
	}

	res, err := h.u.Followers(pkg.Context{Ctx: c}, id)
	if err != nil {
		return err
	}

	return c.JSON(res)
}

func (h *v1UserHandler) Update(c *fiber.Ctx) error {
	req := &request.UpdateUser{}
	if err := c.BodyParser(req); err != nil {
		return domain.UnableParseFormData.New(err)
	}

	res, err := h.u.Update(pkg.Context{Ctx: c}, req)
	if err != nil {
		return err
	}

	return c.JSON(res)
}

func (h *v1UserHandler) Suspend(c *fiber.Ctx) error {
	if err := h.u.Suspend(pkg.Context{Ctx: c}); err != nil {
		return err
	}

	return c.JSON(nil)
}
