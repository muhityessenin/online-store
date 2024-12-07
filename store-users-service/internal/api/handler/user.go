package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"users-service/internal/domain/user"
	interfaces "users-service/internal/service/interface"
	"users-service/pkg/response"
)

type UserHandler struct {
	userService interfaces.UserService
}

func NewUserHandler(service interfaces.UserService) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags users
// @Accept json
// @Produce json
// @Param user body user.Request true "User Request"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /users [post]
func (uh *UserHandler) CreateUser(c *gin.Context) {
	req := user.Request{}
	if err := c.BindJSON(&req); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are wrong", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	if err := req.Validate(); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are wrong", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	res, err := uh.userService.CreateUser(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, user.ErrorNotFound) {
			errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are wrong", nil, err.Error())
			c.JSON(http.StatusBadRequest, errRes)
			return
		}
		errRes := response.ClientResponse(http.StatusInternalServerError, "failed to create user", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusCreated, "the user was successfully created", res, nil)
	c.JSON(http.StatusCreated, successRes)
}

// ListUsers godoc
// @Summary List all users
// @Description Get a list of all users
// @Tags users
// @Produce json
// @Success 200 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /users [get]
func (uh *UserHandler) ListUsers(c *gin.Context) {
	res, err := uh.userService.ListUsers(c.Request.Context())
	if err != nil {
		if errors.Is(err, user.ErrorNotFound) {
			errRes := response.ClientResponse(http.StatusNotFound, "no users found", nil, err.Error())
			c.JSON(http.StatusNotFound, errRes)
			return
		}
		errRes := response.ClientResponse(http.StatusInternalServerError, "failed to list users", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "the users list", res, nil)
	c.JSON(http.StatusOK, successRes)
}

// GetUser godoc
// @Summary Get a user by ID
// @Description Get details of a user by its ID
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /users/{id} [get]
func (uh *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	res, err := uh.userService.GetUser(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, user.ErrorNotFound) {
			errRes := response.ClientResponse(http.StatusNotFound, "user not found", nil, err.Error())
			c.JSON(http.StatusNotFound, errRes)
			return
		}
		errRes := response.ClientResponse(http.StatusInternalServerError, "failed to get user", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "the user details", res, nil)
	c.JSON(http.StatusOK, successRes)
}

// UpdateUser godoc
// @Summary Update a user by ID
// @Description Update details of a user by its ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body user.Request true "User Request"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /users/{id} [put]
func (uh *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	req := user.Request{}
	if err := c.BindJSON(&req); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are wrong", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	if err := req.Validate(); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are wrong", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	err := uh.userService.UpdateUser(c.Request.Context(), id, req)
	if err != nil {
		if errors.Is(err, user.ErrorNotFound) {
			errRes := response.ClientResponse(http.StatusNotFound, "user not found", nil, err.Error())
			c.JSON(http.StatusNotFound, errRes)
			return
		}
		errRes := response.ClientResponse(http.StatusInternalServerError, "failed to update user", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "the user was successfully updated", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// DeleteUser godoc
// @Summary Delete a user by ID
// @Description Delete a user by its ID
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /users/{id} [delete]
func (uh *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	err := uh.userService.DeleteUser(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, user.ErrorNotFound) {
			errRes := response.ClientResponse(http.StatusNotFound, "user not found", nil, err.Error())
			c.JSON(http.StatusNotFound, errRes)
			return
		}
		errRes := response.ClientResponse(http.StatusInternalServerError, "failed to delete user", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "the user was successfully deleted", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// SearchUsers godoc
// @Summary Search users by name or email
// @Description Search users by name or email
// @Tags users
// @Produce json
// @Param filter query string true "Query"
// @Param val query string true "Value"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /users/search [get]
func (uh *UserHandler) SearchUsers(c *gin.Context) {
	filter := c.Query("filter")
	val := c.Query("val")

	if filter == "" || val == "" {
		errRes := response.ClientResponse(http.StatusBadRequest, "filter and val query parameters are required", nil, "missing query parameters")
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	res, err := uh.userService.SearchUser(c.Request.Context(), filter, val)
	if err != nil {
		if errors.Is(err, user.ErrorInvalidSearch) {
			errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are wrong", nil, err.Error())
			c.JSON(http.StatusBadRequest, errRes)
			return
		}
		if errors.Is(err, user.ErrorNotFound) {
			errRes := response.ClientResponse(http.StatusNotFound, "no users found", nil, err.Error())
			c.JSON(http.StatusNotFound, errRes)
			return
		}
		errRes := response.ClientResponse(http.StatusInternalServerError, "failed to search users", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "the users list", res, nil)
	c.JSON(http.StatusOK, successRes)
}
