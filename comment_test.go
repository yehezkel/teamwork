package teamwork

import (
	"testing"
	"time"
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

func TestAddComment(t *testing.T) {

	commentApi := CommentApi{
		CommentEndPoint{
			Client: getTestingClient(),
		},
	}

	comment := &NewComment{
		Body:        "Testing body: " + time.Now().String(),
		Notify:      "",
		Private:     true,
		ContentType: "TEXT",
	}

	id, err := commentApi.Create(comment, "tasks", "15553648")
	if err != nil {
		t.Errorf("Error: %v", err)
		return
	}

	t.Logf("New Comment Id: %s", id)

}
