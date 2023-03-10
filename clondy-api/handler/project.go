package handler

import (
	"clondy/model"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetProjects ..
func (h *Handler) GetProjects(c echo.Context) (err error) {

	userID := c.Param("id")

	// / Validate UserID
	if userID == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "No ObjectIDUser found "}
	}

	gpj, err := h.Collector.GetProjects(bson.M{"user_id": userID})
	if err != nil {
		return &echo.HTTPError{Code: http.StatusExpectationFailed, Message: "Unable to retrieve information"}
	}

	// c.Set("ProjectID", nil)

	return c.JSON(http.StatusOK, gpj)
}

// CreateProject ..
func (h *Handler) CreateProject(c echo.Context) (err error) {

	// Bind
	cpj := &model.Project{ID: primitive.NewObjectID()}
	if err = c.Bind(cpj); err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "Unable to connect to the database "}
	}

	// Validate ProjectName
	if cpj.Name == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "No project name found "}
	}

	// userID := c.Param("id")
	// cpj.UserID = userID

	// / Validate UserID
	if cpj.UserID == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "No ObjectIDUser found "}
	}

	// // Check ProjectName
	// project, _ := h.Collector.GetProjects(bson.M{"user_id": cpj.UserID})
	// if project != nil {
	// 	return &echo.HTTPError{Code: http.StatusBadRequest, Message: "This project name has already been used "}
	// }

	CreatedAt := time.Now()
	cpj.CreatedAt = CreatedAt
	UpdateAt := time.Now()
	cpj.UpdateAt = UpdateAt

	// Save Project
	iocpj, _ := h.Collector.InsertProjec(cpj)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusExpectationFailed, Message: "Import data has a problem"}
	}

	return c.JSON(http.StatusOK, iocpj)
}

// UpdateProject ..
func (h *Handler) UpdateProject(c echo.Context) (err error) {

	// Bind
	cpj := new(model.Project)
	if err = c.Bind(cpj); err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "Unable to connect to the database "}
	}

	UpdateAt := time.Now()
	cpj.UpdateAt = UpdateAt

	// set filters and updates
	filter := bson.M{"_id": cpj.ID}
	update := bson.M{"$set": cpj}

	fmt.Println(filter)
	fmt.Println(update)

	// Check ProjectName
	_, err = h.Collector.GetProject(filter)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "The name you want to delete cannot be found "}
	}

	// UpdateProject
	_, err = h.Collector.UpdateProject(filter, update)
	if err != nil {
		fmt.Println(err)
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Can't be updated"}
	}

	return &echo.HTTPError{Code: http.StatusOK, Message: "Successfully updated"}

}

// DeleteProject ..
func (h *Handler) DeleteProject(c echo.Context) (err error) {

	IDProject := c.Param("id")
	fmt.Println(IDProject)
	ObjectIDProject, _ := primitive.ObjectIDFromHex(IDProject)

	fmt.Println(ObjectIDProject)

	// // / Validate UserID
	// if ObjectIDProject == primitive.NilObjectID {
	// 	return &echo.HTTPError{Code: http.StatusBadRequest, Message: "No ObjectIDUser found "}
	// }

	// Check LayerName
	_, err = h.Collector.GetLayers(bson.M{"project_id": IDProject})
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "The name you want to delete cannot be found "}
	}

	// DeleteLayer
	_, err = h.Collector.DeleteLayers(bson.M{"project_id": IDProject})
	if err != nil {
		return &echo.HTTPError{Code: http.StatusExpectationFailed, Message: "There is an error, cannot be deleted "}
	}

	// Check ProjectName
	_, err = h.Collector.GetProject(bson.M{"_id": ObjectIDProject})
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "The name you want to delete cannot be found "}
	}

	// DeleteProject
	_, err = h.Collector.DeleteProject(bson.M{"_id": ObjectIDProject})
	if err != nil {
		return &echo.HTTPError{Code: http.StatusExpectationFailed, Message: "There is an error, cannot be deleted "}
	}

	return &echo.HTTPError{Code: http.StatusOK, Message: "Successful deletion"}

}
