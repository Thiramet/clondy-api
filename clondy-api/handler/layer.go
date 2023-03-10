package handler

import (
	"clondy/model"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateLayer ..
func (h *Handler) CreateLayer(c echo.Context) (err error) {

	// Bind
	cly := &model.Layer{ID: primitive.NewObjectID()}
	if err = c.Bind(cly); err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "Unable to connect to the database "}
	}

	// Validate ProjectName
	if cly.ProjectID == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "No project name found "}
	}

	// / Validate Name
	if cly.Name == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Data layer name not found. "}
	}

	// Check LayerName
	Layer, _ := h.Collector.GetLayer(bson.M{"name": cly.Name})
	if Layer != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "The data layer name already exists."}
	}

	cly.Type = "FeatureCollection"

	// Save Project
	iocly, err := h.Collector.InsertLayer(cly)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusExpectationFailed, Message: "Import data has a problem"}
	}

	return c.JSON(http.StatusOK, iocly)
}

// GetLayers ..
func (h *Handler) GetLayers(c echo.Context) (err error) {

	// // Bind
	// cly := new(model.Layer)
	// if err = c.Bind(cly); err != nil {
	// 	return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "Unable to connect to the database "}
	// }

	ProjectID := c.Param("id")

	// / Validate UserID
	if ProjectID == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "No ObjectIDProject found "}
	}

	// Check ProjectID
	_, err = h.Collector.GetLayer(bson.M{"project_id": ProjectID})
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Cannot find ProjectID."}
	}

	// GetLayers
	gly, err := h.Collector.GetLayers(bson.M{"project_id": ProjectID})
	if err != nil {
		return &echo.HTTPError{Code: http.StatusExpectationFailed, Message: "Unable to retrieve information"}
	}

	return c.JSON(http.StatusOK, gly)
}

// UpdateLayer

// DeleteLayer ..
func (h *Handler) DeleteLayer(c echo.Context) (err error) {

	IDLayer := c.Param("id")
	fmt.Println(IDLayer)
	ObjectIDLayer, _ := primitive.ObjectIDFromHex(IDLayer)

	fmt.Println(ObjectIDLayer)

	// Check ProjectName
	_, err = h.Collector.GetLayer(bson.M{"_id": ObjectIDLayer})
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "The name you want to delete cannot be found "}
	}

	// DeleteProject
	_, err = h.Collector.DeleteLayer(bson.M{"_id": ObjectIDLayer})
	if err != nil {
		return &echo.HTTPError{Code: http.StatusExpectationFailed, Message: "There is an error, cannot be deleted "}
	}
	return &echo.HTTPError{Code: http.StatusOK, Message: "Successful deletion"}

}

// UpdateLayer ..
func (h *Handler) UpdateLayer(c echo.Context) (err error) {

	// Bind
	cly := new(model.Layer)
	if err = c.Bind(cly); err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "Unable to connect to the database "}
	}

	filter := bson.M{"_id": cly.ID}
	update := bson.M{"$set": cly}

	// Check ProjectName
	_, err = h.Collector.GetLayer(filter)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "The name you want to delete cannot be found "}
	}

	// UpdateProject
	_, err = h.Collector.UpdateLayer(filter, update)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Can't be updated"}
	}

	return &echo.HTTPError{Code: http.StatusOK, Message: "Successfully updated"}

}
