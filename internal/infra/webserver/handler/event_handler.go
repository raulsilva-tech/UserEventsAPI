package handler

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/raulsilva-tech/UserEventsAPI/internal/dto"
	"github.com/raulsilva-tech/UserEventsAPI/internal/entity"
	"github.com/raulsilva-tech/UserEventsAPI/internal/infra/database"
)

type EventHandler struct {
	DAO *database.EventDAO
}

func NewEventHandler(dao *database.EventDAO) *EventHandler {
	return &EventHandler{
		DAO: dao,
	}
}

// Create Event godoc
// @Summary			Create Event
// @Description		Creates a Event  in the database
// @Tags 			events
// @Accept			json
// @Produce			json
// @Param			request	body	dto.CreateEventInput	true	"Event request"
// @Success 		201
// @Failure 		500
// @Failure 		400
// @Router 			/events	[post]
func (h *EventHandler) CreateEvent(c *gin.Context) {

	var data dto.CreateEventInput

	err := c.BindJSON(&data)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	userId, _ := uuid.Parse(data.UserId)
	etId, _ := uuid.Parse(data.EventTypeId)

	record, err := entity.NewEvent(entity.User{Id: userId}, entity.EventType{Id: etId}, data.PerformedAt)
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

// Update event godoc
// @Summary Update an event
// @Description Updates an event in the database
// @Tags events
// @Accept json
// @Produce json
// @Param id path string true "event ID"  Format(uuid)
// @Param event body entity.Event true "event data"
// @Success 200 {object} entity.Event
// @Failure 400
// @Failure 404
// @Router /events/{id} [put]
func (h *EventHandler) UpdateEvent(c *gin.Context) {

	var data entity.Event
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

	err = h.DAO.Update(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "updated successfully"})
}

// GetEvent godoc
// @Summary Get an event
// @Description Get an event by its id
// @Tags events
// @Accept json
// @Produce json
// @Param id path string true "Event ID" Format(uuid)
// @Success 200 {object} entity.Event
// @Failure 400
// @Failure 404
// @Router /events/{id} [get]
func (h *EventHandler) GetEvent(c *gin.Context) {

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

// Delete Event godoc
// @Summary Delete an event
// @Description Deletes an event from the database
// @Tags events
// @Accept json
// @Produce json
// @Param id path string true "Event ID"  Format(uuid)
// @Success 200
// @Failure 400
// @Failure 404
// @Router /events/{id} [delete]
func (h *EventHandler) DeleteEvent(c *gin.Context) {

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
// @Summary			Gets all events
// @Description		Gets all events in the database
// @Tags 			events
// @Accept			json
// @Produce			json
// @Success 		200	{array}	entity.Event
// @Failure 		500
// @Failure			404
// @Router 			/events	[get]
func (h *EventHandler) GetAllEvent(c *gin.Context) {

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
