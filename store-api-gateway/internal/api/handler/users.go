package handler

import (
	"api-gateway-service/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	userUrl string
}

func NewUserHandler(userUrl string) *UserHandler {
	return &UserHandler{userUrl}
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body user.Request true "User data"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /users [post]
func (u *UserHandler) CreateUser(c *gin.Context) {
	req, err := http.NewRequest(http.MethodPost, u.userUrl, c.Request.Body)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()
	res, err := response.ParseResponse(resp)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(resp.StatusCode, res)
}

// ListUsers godoc
// @Summary List all users
// @Description List all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /users [get]
func (u *UserHandler) ListUsers(c *gin.Context) {
	req, err := http.NewRequest("GET", u.userUrl, nil)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()
	res, err := response.ParseResponse(resp)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(resp.StatusCode, res)
}

// GetUser godoc
// @Summary Get user by id
// @Description Get user by id
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /users/{id} [get]
func (u *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	req, err := http.NewRequest("GET", u.userUrl+"/"+id, nil)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()
	res, err := response.ParseResponse(resp)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(resp.StatusCode, res)

}

// UpdateUser godoc
// @Summary Update user by id
// @Description Update user by id
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Param user body user.Request true "User data"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /users/{id} [put]
func (u *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	req, err := http.NewRequest("PUT", u.userUrl+"/"+id, c.Request.Body)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()
	res, err := response.ParseResponse(resp)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(resp.StatusCode, res)
}

// DeleteUser godoc
// @Summary Delete user by id
// @Description Delete user by id
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /users/{id} [delete]
func (u *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	req, err := http.NewRequest("DELETE", u.userUrl+"/"+id, nil)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()
	res, err := response.ParseResponse(resp)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(resp.StatusCode, res)
}

// SearchUser godoc
// @Summary Search user
// @Description Search user
// @Tags users
// @Accept  json
// @Produce  json
// @Param filter query string true "Search filter"
// @Param value query string true "Search value"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /users/search [get]
func (u *UserHandler) SearchUser(c *gin.Context) {
	filter := c.Query("filter")
	val := c.Query("value")
	req, err := http.NewRequest("GET", u.userUrl+"/search?filter="+filter+"&value="+val, nil)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()
	res, err := response.ParseResponse(resp)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(resp.StatusCode, res)
}
