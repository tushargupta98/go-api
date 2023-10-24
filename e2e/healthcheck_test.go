package e2e

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type EndToEndSuite struct {
	suite.Suite
	server *httptest.Server
}

func (suite *EndToEndSuite) SetupTest() {
	// Create a new Chi router and add the /health route

	// Create a test server using httptest
	//suite.server = httptest.NewServer(router)
}

func (suite *EndToEndSuite) TearDownTest() {
	suite.server.Close()
}

func (suite *EndToEndSuite) TestHealthHandler() {
	// Send an HTTP GET request to the /health endpoint
	response, err := http.Get(suite.server.URL + "/health")
	assert.NoError(suite.T(), err)
	defer response.Body.Close()

	// Check the response status code
	assert.Equal(suite.T(), http.StatusOK, response.StatusCode)

	// Read and validate the response body
	expectedBody := "OK"
	actualBody := readResponseBody(suite.T(), response)
	assert.Equal(suite.T(), expectedBody, actualBody)
}

func readResponseBody(t assert.TestingT, response *http.Response) string {
	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	return string(body)
}

func TestMySuite(t *testing.T) {
	//suite.Run(t, new(EndToEndSuite))
}
