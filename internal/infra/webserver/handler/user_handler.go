package handler

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/raulsilva-tech/UserEventsAPI/internal/dto"
	"github.com/raulsilva-tech/UserEventsAPI/internal/entity"
	"github.com/raulsilva-tech/UserEventsAPI/internal/infra/database"
)

type UserHandler struct {
	DAO *database.UserDAO
}

func NewUserHandler(dao *database.UserDAO) *UserHandler {
	return &UserHandler{
		DAO: dao,
	}
}

// Create User godoc
// @Summary			Create User
// @Description		Creates a User in the database
// @Tags 			users
// @Accept			json
// @Produce			json
// @Param			request	body	dto.CreateUserInput	true	"user request"
// @Success 		201
// @Failure 		500
// @Failure 		400
// @Router 			/users	[post]
func (h *UserHandler) CreateUser(c *gin.Context) {

	var data dto.CreateUserInput

	err := c.BindJSON(&data)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	record, err := entity.NewUser(data.Name, data.Email, data.Password, time.Now(), time.Now())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	err = h.DAO.Create(record)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "created successfully"})
}

// Update User godoc
// @Summary Update an user
// @Description Updates a user in the database
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"  Format(uuid)
// @Param user body dto.CreateUserInput true "User data"
// @Success 200 {object} entity.User
// @Failure 400
// @Failure 404
// @Router /users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {

	var data entity.User
	id := c.Param("id")
	if id == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	err := c.BindJSON(&data)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	found, err := h.DAO.GetById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	data.Id = found.Id
	data.CreatedAt = found.CreatedAt
	data.UpdatedAt = time.Now()

	err = h.DAO.Update(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "updated successfully"})
}

// GetUser godoc
// @Summary Get an user
// @Description Get an user by its id
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID" Format(uuid)
// @Success 200 {object} entity.User
// @Failure 400
// @Failure 404
// @Router /users/{id} [get]
func (h *UserHandler) GetUser(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	et, err := h.DAO.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"message": "record not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, et)
}

// Delete User godoc
// @Summary Delete an User
// @Description Deletes an User from the database
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"  Format(uuid)
// @Success 200
// @Failure 400
// @Failure 404
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	et, err := h.DAO.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"message": "record not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}

	err = h.DAO.Delete(et)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted successfully"})
}

// Get All godoc
// @Summary			Gets all users
// @Description		Gets all users in the database
// @Tags 			users
// @Accept			json
// @Produce			json
// @Success 		200	{array}	entity.User
// @Failure 		500
// @Failure			404
// @Router 			/users	[get]
func (h *UserHandler) GetAllUser(c *gin.Context) {

	etList, err := h.DAO.GetAll()
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"message": "record not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, etList)
}
