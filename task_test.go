package teamwork

import (
	"testing"
)

func TestFetchTask(t *testing.T) {

	taskApi := TaskApi{
		TaskEndPoint{
			Client: getTestingClient(),
		},
	}

	out, err := taskApi.Fetch("15553648")
	if err != nil {
		t.Errorf("error: %v", err)
	}

	t.Logf("log: %#v", out)
}
