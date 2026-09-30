package users

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"example.com/company/qa-api/api"
	"example.com/company/qa-api/tests"

	"github.com/stretchr/testify/suite"
)

type UsersTestSuite struct {
	tests.APISuite
}

func (s *UsersTestSuite) TestCreateUser() {
	email := fmt.Sprintf(
		"john-%d@example.com",
		time.Now().UnixNano(),
	)

	req := api.CreateUserRequest{
		Name:  "John Doe",
		Email: email,
	}

	resp, user, apiErr, err := s.API.Users.Create(req)

	s.Require().NoError(err)
	s.Require().Nil(apiErr)

	s.Require().NotNil(resp)
	s.Require().Equal(
		http.StatusCreated,
		resp.StatusCode,
	)

	s.Require().NotEmpty(user.ID)
	s.Require().Equal(req.Name, user.Name)
	s.Require().Equal(req.Email, user.Email)
}

func TestUsersTestSuite(t *testing.T) {
	suite.Run(t, new(UsersTestSuite))
}
