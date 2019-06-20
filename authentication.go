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
	//ApiClient *DefaultClient
	AuthEndPoint
}

func (auth Authentication) Authenticate() (*Account, error) {

	out := AccountResponse{}
	err := auth.AuthEndPoint.Authenticate(&out)

	if err != nil {
		return nil, err
	}

	return out.Account, nil
}

type AuthEndPoint struct {
	ApiClient *DefaultClient
}

func (ap AuthEndPoint) Authenticate(out interface{}) error {

	endpoint := "authenticate.json"
	err := ap.ApiClient.DoRequest("GET", endpoint, nil, out)

	if err != nil {
		return err
	}

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
