package delivery

import (
	"context"
	"github.com/itss-academy/imago/core/domain/profile"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ProfileHttpDelivery struct {
	api     *echo.Group
	interop profile.ProfileInterop
}

func (p ProfileHttpDelivery) GetById(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	//if token is empty return error
	if token == "" {
		return profile.ErrTokenEmpty
	}
	//using query param to get id
	profileData, err := p.interop.GetById(c.Request().Context(), token, c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, profileData)
}

func (p ProfileHttpDelivery) GetMine(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	//if token is empty return error
	if token == "" {
		return profile.ErrTokenEmpty
	}
	profileData, err := p.interop.GetMine(c.Request().Context(), token)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())

	}
	return c.JSON(http.StatusOK, profileData)
}

func (p ProfileHttpDelivery) GetAll(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	//if token is empty return error
	if token == "" {
		return profile.ErrTokenEmpty
	}
	profileData, err := p.interop.GetAll(c.Request().Context(), token)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, profileData)
}

func (p ProfileHttpDelivery) Create(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return c.JSON(http.StatusBadRequest, "token is empty")
	}

	profileData := &profile.Profile{}
	if err := c.Bind(profileData); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	err := p.interop.Create(c.Request().Context(), token, profileData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, profileData)
}

func (p ProfileHttpDelivery) Update(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return c.JSON(http.StatusBadRequest, "token is empty")
	}

	profileData := &profile.Profile{}
	if err := c.Bind(profileData); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	err := p.interop.Update(c.Request().Context(), token, profileData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, profileData)
}

func (p ProfileHttpDelivery) Follow(ctx context.Context, token string, profileId string, profileOther string) error {
	//TODO implement me
	panic("implement me")
}

func (p ProfileHttpDelivery) Unfollow(ctx context.Context, token string, profileId string, profileOther string) error {
	//TODO implement me
	panic("implement me")
}

func NewProfileHttpDelivery(api *echo.Group, interop profile.ProfileInterop) *ProfileHttpDelivery {
	handler := &ProfileHttpDelivery{
		api:     api,
		interop: interop,
	}
	api.GET("/all", handler.GetAll)
	api.GET("", handler.GetById)
	api.GET("/mine", handler.GetMine)
	api.POST("/mine", handler.Create)
	api.PUT("/mine", handler.Update)
	return handler
}
