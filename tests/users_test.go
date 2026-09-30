package tests

import (
	"fmt"
	"testing"
	"time"

	"example.com/company/qa-api/api"
	"github.com/stretchr/testify/require"
)

func newAPI(t *testing.T) *api.Client {
	t.Helper()

	baseURL := "http://localhost:8080"

	if baseURL == "" {
		t.Fatal("API_URL environment variable is not set")
	}

	return api.NewClient(baseURL)
}

func TestUser_CreateAndGet(t *testing.T) {
	client := newAPI(t)
	users := client.Users()

	email := fmt.Sprintf(
		"john-%d@example.com",
		time.Now().UnixNano(),
	)

	createRequest := api.CreateUserRequest{
		Name:  "John Doe",
		Email: email,
	}

	// ---------------------------------------------------------
	// 1. CREATE USER
	// ---------------------------------------------------------

	httpResp, createdUser,_, err := users.Create(
		createRequest,
	)

	require.NoError(t, err)

	require.Equal(
		t,
		201,
		httpResp.StatusCode,
	)

	require.NotEmpty(
		t,
		createdUser.ID,
	)

	require.Equal(
		t,
		createRequest.Name,
		createdUser.Name,
	)

	require.Equal(
		t,
		createRequest.Email,
		createdUser.Email,
	)

	// ---------------------------------------------------------
	// 2. GET USER
	// ---------------------------------------------------------

	httpResp, getResponse, _, err := users.Get(
		createdUser.ID,
	)

	require.NoError(t, err)

	require.Equal(
		t,
		200,
		httpResp.StatusCode,
	)

	// ---------------------------------------------------------
	// 3. VERIFY
	// ---------------------------------------------------------

	expectedUser := api.User{
		ID:    createdUser.ID,
		Name:  createRequest.Name,
		Email: createRequest.Email,
	}

	require.Equal(
		t,
		expectedUser,
		getResponse,
	)
}
