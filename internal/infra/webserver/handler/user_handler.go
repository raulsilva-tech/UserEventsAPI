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
