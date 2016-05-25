package controllers

import (
	"github.com/carrot/burrow/constants"
	"github.com/carrot/burrow/controllers/helper"
	"github.com/carrot/burrow/models"
	"github.com/carrot/burrow/response"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type TopicsController struct{}

func (tc *TopicsController) Index(c echo.Context) error {
	resp := response.New(c)
	defer resp.Render()

	// Getting limit
	limit, helperError := helper.GetLimit(c)
	if helperError != nil {
		return helper.PrepareResponse(resp, helperError)
	}

	// Getting offset
	offset, helperError := helper.GetOffset(c)
	if helperError != nil {
		return helper.PrepareResponse(resp, helperError)
	}

	// Fetching models
	res, err := models.AllTopics(limit, offset)
	if err != nil {
		resp.SetResponse(http.StatusInternalServerError, nil)
		return nil
	}

	resp.SetResponse(http.StatusOK, res)
	return nil
}

func (tc *TopicsController) Show(c echo.Context) error {
	resp := response.New(c)
	defer resp.Render()

	// Getting id
	id, err := strconv.ParseInt(c.Param(constants.ID), 10, 64)
	if err != nil {
		resp.AddErrorDetail(response.ErrorInvalidIdParameter)
		resp.SetResponse(http.StatusBadRequest, nil)
		return nil
	}

	// Fetching model
	topic := models.NewTopic()
	topic.Id = id
	err = topic.Load()
	if err != nil {
		resp.SetResponse(http.StatusNotFound, nil)
		return nil
	}

	resp.SetResponse(http.StatusOK, topic)
	return nil
}

/**
 * @api {post} /topics Creates a new topic
 * @apiName CreateTopic
 * @apiGroup Topics
 *
 * @apiParam {String} name The name of the topic
 */
func (tc *TopicsController) Create(c echo.Context) error {
	resp := response.New(c)
	defer resp.Render()

	// Getting params
	name := c.FormValue(constants.NAME)
	if name == "" {
		resp.AddErrorDetail(response.ErrorMissingNameParameter)
		resp.SetResponse(http.StatusBadRequest, nil)
		return nil
	}

	// Creating the topic
	topic := models.NewTopic()
	topic.Name = name
	err := topic.Insert()
	if err != nil {
		resp.SetResponse(http.StatusConflict, nil)
		return nil
	}

	resp.SetResponse(http.StatusOK, topic)
	return nil
}

func (tc *TopicsController) Update(c echo.Context) error {
	resp := response.New(c)
	defer resp.Render()

	// Getting params
	id, err := strconv.ParseInt(c.Param(constants.ID), 10, 64)
	if err != nil {
		resp.SetResponse(http.StatusBadRequest, nil)
		resp.AddErrorDetail(response.ErrorInvalidIdParameter)
		return nil
	}

	// Loading the topic
	topic := models.NewTopic()
	topic.Id = id
	err = topic.Load()
	if err != nil {
		resp.SetResponse(http.StatusNotFound, nil)
		return nil
	}

	name := c.FormValue(constants.NAME)
	if name != "" {
		topic.Name = name
	}

	err = topic.Update()
	if err != nil {
		resp.SetResponse(http.StatusConflict, nil)
		return nil
	}

	resp.SetResponse(http.StatusOK, topic)
	return nil
}

func (tc *TopicsController) Delete(c echo.Context) error {
	resp := response.New(c)
	defer resp.Render()

	// Getting Params
	id, err := strconv.ParseInt(c.Param(constants.ID), 10, 64)
	if err != nil {
		resp.SetResponse(http.StatusBadRequest, nil)
		resp.AddErrorDetail(response.ErrorInvalidIdParameter)
		return nil
	}

	// Deleting topic
	topic := models.NewTopic()
	topic.Id = id
	err = topic.Delete()
	if err != nil {
		resp.SetResponse(http.StatusNotFound, nil)
		return nil
	}

	resp.SetResponse(http.StatusOK, nil)
	return nil
}
