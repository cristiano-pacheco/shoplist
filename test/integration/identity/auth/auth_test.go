package auth

import (
	"github.com/cristiano-pacheco/shoplist/test"
	"github.com/stretchr/testify/suite"
)

type AuthTestSuite struct {
	suite.Suite
	itest *test.IntegrationTest
}

func (s *AuthTestSuite) TestGenerateTokenSuccess() {

}

func (s *AuthTestSuite) TestGenerateTokenFailed() {

}
