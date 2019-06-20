package teamwork

const (
	AUTHPREFIX = "api"
)

// BuildAuthenticationClient build a default api client base on the particular uri
// pattern for the authentication api endpoint
func BuildAuthenticationClient(token string, options ...ClientOption) *DefaultClient {

	return NewClient(AUTHPREFIX, token, options...)
}

type Authentication struct {
	ApiClient *DefaultClient
}

func (auth Authentication) Authenticate() error {

	endpoint := "authenticate.json"
	//var out interface{}
	out := AccountResponse{}
	err := auth.ApiClient.DoRequest("GET", endpoint, nil, &out)

	if err != nil {
		return err
	}
	auth.ApiClient.log.Printf("Response: %#v", out.Account)
	return nil

}

type Account struct {
	Firstname string `json:"firstname"`
	Id        string `json:"id"`
}

type AccountResponse struct {
	Status  string   `json:"STATUS"`
	Account *Account `json:"account"`
}
