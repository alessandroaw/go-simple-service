package users

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Use Case
type UserHandler struct {
	UUseCase UserUseCase
}

// NewArticleHandler will initialize the user resources endpoint
func NewUserHandler(e *echo.Echo, uuc UserUseCase) *UserHandler {
	handler := &UserHandler{
		UUseCase: uuc,
	}

	// Routes\
	e.GET("/users", handler.GetAll)
	e.GET("/users/:id", handler.GetById)
	e.POST("/users", handler.Create)
	e.PUT("/users/:id", handler.Update)
	e.DELETE("/users/:id", handler.Delete)

	return handler
}

func (h *UserHandler) GetAll(c echo.Context) error {
	users, err := h.UUseCase.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, users)
}

func (u *UserHandler) GetById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user, err := u.UUseCase.GetById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, user)
}

func (u *UserHandler) Create(c echo.Context) error {
	usr := &User{}

	if err := c.Bind(usr); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	res, err := u.UUseCase.Create(usr)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, res)
}

func (u *UserHandler) Update(c echo.Context) error {
	usr := &User{}
	if err := c.Bind(usr); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	res, err := u.UUseCase.Update(id, usr)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, res)
}

func (u *UserHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = u.UUseCase.Delete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusNoContent)
}
