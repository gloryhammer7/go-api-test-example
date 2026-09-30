package api

import (
	"fmt"
	"net/http"
	"net/url"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UsersClient struct {
	client *Client
}

func (c *Client) Users() *UsersClient {
	return &UsersClient{
		client: c,
	}
}

func (u *UsersClient) Create(req CreateUserRequest) (*http.Response, User, *APIError, error) {

	resp, err := u.client.Do(
		http.MethodPost,
		"/user",
		req,
	)
	if err != nil {
		return nil, User{}, nil, err
	}

	user, apiErr, err := decodeResponse[User](resp)
	if err != nil {
		return nil, User{}, apiErr, err
	}
	
	return resp, user, nil, nil
}

func (u *UsersClient) Get(id string) (*http.Response, User, *APIError, error) {

	path := fmt.Sprintf(
		"/user?id=%s",
		url.PathEscape(id),
	)

	resp, err := u.client.Do(
		http.MethodGet,
		path,
		nil,
	)
	if err != nil {
		return nil, User{}, nil, err
	}

	user, apiErr, err := decodeResponse[User](resp)
	if err != nil {
		return nil, User{}, apiErr, err
	}

	return resp, user, nil, nil
}
