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

	res, err := models.AllTopics(limit, offset)
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
	return nil
}

func (tc *TopicsController) Create(c *echo.Context) error {
	resp := response.New(c)
	defer resp.Render()
	return nil
}

func (tc *TopicsController) Update(c *echo.Context) error {
	resp := response.New(c)
	defer resp.Render()
	return nil
}

func (tc *TopicsController) Destroy(c *echo.Context) error {
	resp := response.New(c)
	defer resp.Render()
	return nil
}
