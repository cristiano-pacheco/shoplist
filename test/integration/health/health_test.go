package health

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/cristiano-pacheco/shoplist/test"
	"github.com/stretchr/testify/suite"
)

type HealthCheckTestSuite struct {
	suite.Suite
	itest *test.IntegrationTest
}

func (s *HealthCheckTestSuite) SetupSuite() {
	s.itest = test.Setup()
}

func (s *HealthCheckTestSuite) TearDownSuite() {
	s.itest.Close()
}

func TestHealthCheckTestSuite(t *testing.T) {
	suite.Run(t, new(HealthCheckTestSuite))
}

func (s *HealthCheckTestSuite) TestGetTestReadinessEndpoint() {
	baseURL := s.itest.Config.App.BaseURL
	healthCheckURL := fmt.Sprintf("%s/readyz", baseURL)

	resp, err := http.Get(healthCheckURL)
	if err != nil {
		s.T().Fatalf("Failed to get health check: %v", err)
	}

	defer resp.Body.Close()

	s.Equal(http.StatusOK, resp.StatusCode)
}

func (s *HealthCheckTestSuite) TestGetLivenessEndpoint() {
	baseURL := s.itest.Config.App.BaseURL
	healthCheckURL := fmt.Sprintf("%s/livez", baseURL)

	resp, err := http.Get(healthCheckURL)
	if err != nil {
		s.T().Fatalf("Failed to get health check: %v", err)
	}

	defer resp.Body.Close()

	s.Equal(http.StatusOK, resp.StatusCode)
}
