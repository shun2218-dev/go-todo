package router

import (
	"net/http"
	"todoList/model"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)


func GetTaskHandler(c echo.Context) error {
	tasks, err := model.GetTasks()
	
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}
	
	return c.JSON(http.StatusOK, tasks)	
}

type ReqTask struct {
	Name string `json:"name"`
}

func AddTaskHandler(c echo.Context) error {
	var req ReqTask

	err := c.Bind(&req)

	if err !=nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}
	
	var task *model.Task

	task, err = model.AddTask(req.Name)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	return c.JSON(http.StatusOK, task)
}

func ChangeFinishedTaskHandler(c echo.Context) error {
	taskID, err := uuid.Parse(c.Param("taskID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	err = model.ChangeFinishedTask(taskID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "BadRequest")
	}

	return c.NoContent(http.StatusOK)
}

func DeleteTaskHandler(c echo.Context) error {
	
	taskID, err := uuid.Parse(c.Param("taskID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	err = model.DeleteTask(taskID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "BadRequest")
	}

	return c.NoContent(http.StatusOK)
}