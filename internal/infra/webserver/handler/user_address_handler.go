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

// Create UserAddress godoc
// @Summary			Create User Address
// @Description		Creates a User Address  in the database
// @Tags 			user address
// @Accept			json
// @Produce			json
// @Param			request	body	dto.CreateUserAddressInput	true	"user address request"
// @Success 		201
// @Failure 		500
// @Failure 		400
// @Router 			/user_address	[post]
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

// Update user address godoc
// @Summary Update an user address
// @Description Updates an user address in the database
// @Tags user address
// @Accept json
// @Produce json
// @Param id path string true "user address ID"  Format(uuid)
// @Param user_address body entity.UserAddress true "user address data"
// @Success 200 {object} entity.UserAddress
// @Failure 400
// @Failure 404
// @Router /user_address/{id} [put]
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

// GetUserAddress godoc
// @Summary Get an user Address
// @Description Get an user Address by its id
// @Tags user address
// @Accept json
// @Produce json
// @Param id path string true "User Address ID" Format(uuid)
// @Success 200 {object} entity.UserAddress
// @Failure 400
// @Failure 404
// @Router /user_address/{id} [get]
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

// Delete User Address godoc
// @Summary Delete an User Address
// @Description Deletes an User Address from the database
// @Tags user address
// @Accept json
// @Produce json
// @Param id path string true "User Address ID"  Format(uuid)
// @Success 200
// @Failure 400
// @Failure 404
// @Router /user_address/{id} [delete]
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

// Get All godoc
// @Summary			Gets all users address
// @Description		Gets all users address in the database
// @Tags 			user address
// @Accept			json
// @Produce			json
// @Success 		200	{array}	entity.UserAddress
// @Failure 		500
// @Failure 		404
// @Router 			/user_address	[get]
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
