package handler

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/raulsilva-tech/UserEventsAPI/internal/dto"
	"github.com/raulsilva-tech/UserEventsAPI/internal/entity"
	"github.com/raulsilva-tech/UserEventsAPI/internal/infra/database"
)

type UserAddressHandler struct {
	DAO *database.UserAddressDAO
}

func NewUserAddressHandler(dao *database.UserAddressDAO) *UserAddressHandler {
	return &UserAddressHandler{
		DAO: dao,
	}
}

func (h *UserAddressHandler) CreateUserAddress(c *gin.Context) {

	var data dto.CreateUserAddressInput

	err := c.BindJSON(&data)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	userId, _ := uuid.Parse(data.UserId)
	user := entity.User{Id: userId}
	record, err := entity.NewUserAddress(user, data.CEP)
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

func (h *UserAddressHandler) UpdateUserAddress(c *gin.Context) {

	var data entity.UserAddress
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

func (h *UserAddressHandler) GetUserAddress(c *gin.Context) {

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

func (h *UserAddressHandler) DeleteUserAddress(c *gin.Context) {

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

func (h *UserAddressHandler) GetAllUserAddress(c *gin.Context) {

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
