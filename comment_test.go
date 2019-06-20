package teamwork

import (
	"testing"
)

func TestFetchComment(t *testing.T) {

	commentApi := CommentApi{
		CommentEndPoint{
			Client: getTestingClient(),
		},
	}

	out, err := commentApi.Fetch("5309104")
	if err != nil {
		t.Errorf("error: %v", err)
	}

	t.Logf("log: %#v", out)
}
