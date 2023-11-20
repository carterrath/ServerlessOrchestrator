package api_tests

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type APITestSuite struct {
	suite.Suite
	// fields and setup code for the test suite
}

func (suite *APITestSuite) TestExample() {
	// test code here
}

func TestModelTestSuite(t *testing.T) {
	suite.Run(t, new(APITestSuite))
}
