package teamwork

import (
	"testing"
)

func TestBuildAuthenticationClient(t *testing.T) {

	token := "1234"
	expUri := "https://api.teamwork.com"
	client := BuildAuthenticationClient(token)

	if client.token != token {
		t.Errorf("token missmatch got %s expecting %s", client.token, token)
	}

	if client.uri != expUri {
		t.Errorf("Uri missmatch got %s expecting %s", client.uri, expUri)
	}

}

func TestFirstRequest(t *testing.T) {

	client := BuildAuthenticationClient(
		"keeyyy",
	)

	auth := Authentication{
		client,
	}

	err := auth.Authenticate()
	if err != nil {
		t.Errorf("error: %v", err)
	}

}
