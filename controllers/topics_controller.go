package controllers

import (
	"github.com/carrot/go-base-api/models"
	"github.com/carrot/go-base-api/response"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type TopicsController struct{}

/**
 * @api {get} /topics Get a list of topics
 * @apiName GetTopics
 * @apiGroup Topics
 *
 * @apiParam {Number} [limit=10] The maximum number of items to return
 * @apiParam {Number} [offset=0] The offset relative to the number of items (not page number)
 */
func (tc *TopicsController) Index(c *echo.Context) error {
	resp := response.New(c)
	defer resp.Render()

	// Defaults
	var limit int64 = 10
	var offset int64 = 0

	// Getting limit
	limitInt, err := strconv.ParseInt(c.Query("limit"), 10, 64)
	if err == nil {
		limit = limitInt
	}

	// Getting offset
	offsetInt, err := strconv.ParseInt(c.Query("offset"), 10, 64)
	if err == nil {
		offset = offsetInt
	}

	// Fetching models
	res, err := models.AllTopics(limit, offset)
	if err != nil {
		resp.AddError(response.ErrorNoContent)
		resp.SetResponse(http.StatusInternalServerError, nil)
		return nil
	}

	resp.SetResponse(http.StatusOK, res)
	return nil
}

/**
 * @api {get} /topics/{id} Get a topic
 * @apiName GetTopic
 * @apiGroup Topics
 *
 * @apiParam {Number} id The id of the topic
 */
func (tc *TopicsController) Show(c *echo.Context) error {
	resp := response.New(c)
	defer resp.Render()

	// Getting id
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		resp.AddError(response.ErrorInvalidParameters)
		resp.SetResponse(http.StatusBadRequest, nil)
		return nil
	}

	// Fetching model
	topic := new(models.Topic)
	err = topic.Load(id)
	if err != nil {
		resp.AddError(response.ErrorRecordNotFound)
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
func (tc *TopicsController) Create(c *echo.Context) error {
	resp := response.New(c)
	defer resp.Render()

	// Getting params
	name := c.Form("name")
	if name == "" {
		resp.AddError(response.ErrorInvalidParameters)
		resp.SetResponse(http.StatusBadRequest, nil)
		return nil
	}

	// Creating the topic
	topic := new(models.Topic)
	topic.Name = name
	err := topic.Save()
	if err != nil {
		resp.SetResponse(http.StatusInternalServerError, nil)
		return nil
	}

	resp.SetResponse(http.StatusOK, topic)
	return nil
}

/**
 * @api {put} /topics/{id} Updates a topic
 * @apiName UpdateTopic
 * @apiGroup Topics
 *
 * @apiParam {Number} id The id of the topic to update
 * @apiParam {String} [name] The new name of the topic
 */
func (tc *TopicsController) Update(c *echo.Context) error {
	resp := response.New(c)
	defer resp.Render()

	// Getting params
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		resp.AddError(response.ErrorInvalidParameters)
		resp.SetResponse(http.StatusBadRequest, nil)
		return nil
	}

	// Loading the topic
	topic := new(models.Topic)
	err = topic.Load(id)
	if err != nil {
		resp.AddError(response.ErrorRecordNotFound)
		resp.SetResponse(http.StatusNotFound, nil)
		return nil
	}

	name := c.Form("name")
	if name != "" {
		topic.Name = name
	}

	err = topic.Update()
	if err != nil {
		resp.SetResponse(http.StatusInternalServerError, nil)
		return nil
	}

	resp.SetResponse(http.StatusOK, topic)
	return nil
}

/**
 * @api {delete} /topics/{id} Deletes a topic
 * @apiName DeleteTopic
 * @apiGroup Topics
 *
 * @apiParam {Number} id The id of the topic to delete
 */
func (tc *TopicsController) Delete(c *echo.Context) error {
	resp := response.New(c)
	defer resp.Render()

	// Getting Params
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		resp.AddError(response.ErrorInvalidParameters)
		resp.SetResponse(http.StatusBadRequest, nil)
		return nil
	}

	// Deleting topic
	topic := new(models.Topic)
	topic.Id = id
	err = topic.Delete()
	if err != nil {
		resp.SetResponse(http.StatusNotFound, nil)
		return nil
	}

	resp.SetResponse(http.StatusOK, nil)
	return nil
}
