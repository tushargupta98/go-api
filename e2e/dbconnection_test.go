package e2e

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type DatabaseConnectionSuite struct {
	suite.Suite
	dbURL string
}

func (suite *DatabaseConnectionSuite) SetupTest() {
	// Provide the URL to your PostgreSQL database.
	suite.dbURL = "host=postgres user=myuser dbname=todo sslmode=disable password=mypassword"
}

func (suite *DatabaseConnectionSuite) TestDatabaseConnection() {
	pool, err := pgxpool.Connect(context.Background(), suite.dbURL)
	assert.NoError(suite.T(), err, "Failed to open the database connection")
	defer pool.Close()
}

func TestDatabaseConnectionSuite(t *testing.T) {
	//suite.Run(t, new(DatabaseConnectionSuite))
}
