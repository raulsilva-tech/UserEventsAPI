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

type EventTypeHandler struct {
	DAO *database.EventTypeDAO
}

func NewEventTypeHandler(dao *database.EventTypeDAO) *EventTypeHandler {
	return &EventTypeHandler{
		DAO: dao,
	}
}

// Create Event type godoc
// @Summary			Create Event Type
// @Description		Creates a Event type  in the database
// @Tags 			event types
// @Accept			json
// @Produce			json
// @Param			request	body	dto.CreateEventTypeInput	true	"Event type request"
// @Success 		201
// @Failure 		500
// @Failure 		400
// @Router 			/event_types	[post]
func (h *EventTypeHandler) CreateEventType(c *gin.Context) {

	var data dto.CreateEventTypeInput

	err := c.BindJSON(&data)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	et, err := entity.NewEventType(data.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	err = h.DAO.Create(et)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "created successfully"})
}

// Update event type godoc
// @Summary Update an event type 
// @Description Updates an event in the database
// @Tags event types
// @Accept json
// @Produce json
// @Param id path string true "event type ID"  Format(uuid)
// @Param event_type body entity.EventType true "event type data"
// @Success 200 {object} entity.EventType
// @Failure 400
// @Failure 404
// @Router /event_types/{id} [put]
func (h *EventTypeHandler) UpdateEventType(c *gin.Context) {

	var data entity.EventType
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

// GetEventType godoc
// @Summary Get an event type
// @Description Get an event type by its id
// @Tags event types
// @Accept json
// @Produce json
// @Param id path string true "Event type ID" Format(uuid)
// @Success 200 {object} entity.EventType
// @Failure 400
// @Failure 404
// @Router /event_types/{id} [get]
func (h *EventTypeHandler) GetEventType(c *gin.Context) {

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

// Delete Event type godoc
// @Summary Delete an event type
// @Description Deletes an event type from the database
// @Tags event types
// @Accept json
// @Produce json
// @Param id path string true "Event type ID"  Format(uuid)
// @Success 200
// @Failure 400
// @Failure 404
// @Router /event_types/{id} [delete]
func (h *EventTypeHandler) DeleteEventType(c *gin.Context) {

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
// @Summary			Gets all event types
// @Description		Gets all event types in the database
// @Tags 			event types
// @Accept			json
// @Produce			json
// @Success 		200	{array}	entity.EventType
// @Failure 		500
// @Failure			404
// @Router 			/event_types	[get]
func (h *EventTypeHandler) GetAllEventType(c *gin.Context) {

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
