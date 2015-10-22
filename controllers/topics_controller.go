package controllers

import (
	"github.com/carrot/go-base-api/models"
	"github.com/carrot/go-base-api/response"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type TopicsController struct{}

func (tc *TopicsController) Index(c *echo.Context) error {
	resp := response.New(c)
	defer resp.Render()

	res, err := models.AllTopics()

	if err != nil {
		// Do some checks
		resp.AddError(response.ErrorNoContent)
		resp.SetResponse(http.StatusInternalServerError, nil)
		return nil
	}

	resp.SetResponse(http.StatusOK, res)
	return nil
}

func (tc *TopicsController) Show(c *echo.Context) error {
	resp := response.New(c)
	defer resp.Render()

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		resp.AddError(response.ErrorInvalidParameters)
		resp.SetResponse(http.StatusBadRequest, nil)
		return nil
	}

	res, err := new(models.Topic).Find(id)

	if err != nil {
		resp.AddError(response.ErrorRecordNotFound)
		resp.SetResponse(http.StatusInternalServerError, nil)
		return nil
	}

	resp.SetResponse(http.StatusOK, res)
	return nil
}

func (tc *TopicsController) Create(c *echo.Context) error {
	resp := response.New(c)
	defer resp.Render()

	topic := models.Topic{
		Copy:  c.Param("copy"),
		Asset: c.Param("asset"),
	}

	err := topic.Create()

	if err != nil {
		resp.AddError(response.ErrorRecordNotCreated)
		resp.SetResponse(http.StatusInternalServerError, nil)
		return nil
	}

	resp.SetResponse(http.StatusCreated, topic)
	return nil
}

func (tc *TopicsController) Update(c *echo.Context) error {
	resp := response.New(c)
	defer resp.Render()

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		resp.AddError(response.ErrorInvalidParameters)
		resp.SetResponse(http.StatusBadRequest, nil)
		return nil
	}

	topic := models.Topic{
		Id:    id,
		Copy:  c.Param("copy"),
		Asset: c.Param("asset"),
	}

	err = topic.Update()

	if err != nil {
		return err
		resp.AddError(response.ErrorRecordNotUpdated)
		resp.SetResponse(http.StatusInternalServerError, nil)
	}

	resp.SetResponse(http.StatusNoContent, nil)

	return nil
}

func (tc *TopicsController) Destroy(c *echo.Context) error {
	resp := response.New(c)
	defer resp.Render()

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		resp.AddError(response.ErrorInvalidParameters)
		resp.SetResponse(http.StatusBadRequest, nil)
		return nil
	}

	topic := models.Topic{Id: id}

	err = topic.Destroy()

	if err != nil {
		resp.AddError(response.ErrorInvalidParameters)
		resp.SetResponse(http.StatusBadRequest, nil)
		return nil

	}

	resp.SetResponse(http.StatusOK, nil)

	return nil
}
