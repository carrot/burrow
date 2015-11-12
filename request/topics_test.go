package request

import (
	"github.com/BrandonRomano/wrecker"
	db "github.com/carrot/go-base-api/db/postgres"
	"github.com/carrot/go-base-api/models"
	"github.com/carrot/go-base-api/response"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/tylerb/graceful"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestApiTestSuite(t *testing.T) {
	suite.Run(t, new(ApiTestSuite))
}

type ApiTestSuite struct {
	suite.Suite
	wreckerClient *wrecker.Wrecker
}

func (suite *ApiTestSuite) SetupTest() {
	// Loading .env file.  We're not failing here in the
	// event that the .env file can't be loaded. The
	// project might be running through a CI system
	// and the environment variables.  Will fail later
	// rather than early
	godotenv.Load("../.env")

	// Starting database
	db.Open()

	// Getting our Echo instance
	e := BuildEcho()

	// Running the server
	port := os.Getenv("PORT")
	suite.wreckerClient = wrecker.New("http://localhost:" + port)
	log.Println("Server started on :" + port)
	go graceful.ListenAndServe(e.Server(":"+port), 5*time.Second) // Graceful shutdown
}

func (suite *ApiTestSuite) TearDownTest() {
	db.Close()
}

func (suite *ApiTestSuite) TestPost() {
	// Testing [POST] /topics
	postResp := new(response.Response)
	postResp.Content = new(models.Topic)

	suite.wreckerClient.Post("/topics").
		FormParam("name", randomString(10)).
		Into(&postResp).
		Execute()

	assert.NotNil(suite.T(), postResp.Content)
	if postResp.Content == nil {
		return
	}
	postTopic := postResp.Content.(*models.Topic)

	// Testing [GET] /topics/{id}
	getResp := new(response.Response)
	getResp.Content = new(models.Topic)

	suite.wreckerClient.Get("/topics/" + strconv.FormatInt(postTopic.Id, 10)).
		Into(&getResp).
		Execute()

	assert.NotNil(suite.T(), getResp.Content)
	if getResp.Content == nil {
		return
	}
	getTopic := getResp.Content.(*models.Topic)
	assert.Equal(suite.T(), getTopic.Name, postTopic.Name)

	// Testing [PUT] /topics/{id}
	putResp := new(response.Response)
	putResp.Content = new(models.Topic)

	newName := randomString(10)
	suite.wreckerClient.Put("/topics/"+strconv.FormatInt(postTopic.Id, 10)).
		FormParam("name", newName).
		Into(&putResp).
		Execute()

	assert.NotNil(suite.T(), putResp.Content)
	if putResp.Content == nil {
		return
	}
	putTopic := putResp.Content.(*models.Topic)
	assert.Equal(suite.T(), putTopic.Name, newName)

	// Testing [GET] /topics/{id} (verifying that it was updated)
	getResp = new(response.Response)
	getResp.Content = new(models.Topic)

	suite.wreckerClient.Get("/topics/" + strconv.FormatInt(postTopic.Id, 10)).
		Into(&getResp).
		Execute()

	assert.NotNil(suite.T(), getResp.Content)
	if getResp.Content == nil {
		return
	}
	getTopic = getResp.Content.(*models.Topic)
	assert.Equal(suite.T(), getTopic.Name, newName)

	// Testing [DELETE] /topics/{id}
	deleteResp := new(response.Response)
	deleteResp.Content = new(models.Topic)

	httpResp, _ := suite.wreckerClient.Delete("/topics/" + strconv.FormatInt(postTopic.Id, 10)).
		Execute()
	assert.Equal(suite.T(), httpResp.StatusCode, http.StatusOK)

	// Testing [GET] /topics/{id} (verifying that it was actually deleted)
	httpResp, _ = suite.wreckerClient.Get("/topics/" + strconv.FormatInt(postTopic.Id, 10)).
		Execute()
	assert.Equal(suite.T(), httpResp.StatusCode, http.StatusNotFound)
}

// ============================
// ========== Helper ==========
// ============================

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func randomString(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
