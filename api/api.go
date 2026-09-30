package api

type API struct {
	Client *Client
	Users  *UsersClient
}

func NewAPI(baseURL string) *API {
	client := NewClient(baseURL)

	return &API{
		Client: client,
		Users:  client.Users(),
	}
}