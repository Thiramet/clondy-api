package handler

import (
	"clondy/database"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

// Handler ..
type Handler struct {
	Collector database.Collector
}

// Index ..
func (h *Handler) Index(c echo.Context) error {

	users, _ := h.Collector.GetLayers(bson.M{})

	return c.JSON(http.StatusOK, users)

}
