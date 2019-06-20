package teamwork

import (
	"testing"
)

func TestBuildAuthenticationClient(t *testing.T) {

	token := "1234"
	expUri := "https://api.teamwork.com"
	client := BuildAuthenticationClient(token)

	if client.token != EncodeApiKey(token) {
		t.Errorf("token missmatch got %s expecting %s", client.token, token)
	}

	if client.uri != expUri {
		t.Errorf("Uri missmatch got %s expecting %s", client.uri, expUri)
	}

}

func TestFirstRequest(t *testing.T) {

	client := BuildAuthenticationClient(
		"123",
	)

	auth := Authentication{
		AuthEndPoint{client},
	}

	out, err := auth.Authenticate()
	if err != nil {
		t.Errorf("error: %v", err)
	}

	t.Logf("log: %v", out)

}

func TestSecondRequest(t *testing.T) {

	client := BuildAuthenticationClient(
		"123",
	)

	out := struct {
		SUCCESS string
		Account struct {
			Firstname string
		}
	}{}

	auth := AuthEndPoint{client}

	err := auth.Authenticate(&out)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	t.Logf("log: %v", out.Account.Firstname)

}
