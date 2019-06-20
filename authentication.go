package teamwork

const (
	AUTHPREFIX = "api"
)

func NewAuthentication(client ApiClient) Authentication {

	return Authentication{
		AuthEndPoint{
			Client: client,
		},
	}
}

// BuildAuthenticationClient build a default ApiClient instance base on the particular uri
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
	Client ApiClient
}

func (ap AuthEndPoint) Authenticate(out interface{}) error {

	endpoint := "authenticate.json"
	err := ap.Client.DoRequest("GET", endpoint, nil, out)

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
