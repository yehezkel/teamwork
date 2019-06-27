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

func TestUpdateTask(t *testing.T) {

	taskApi := TaskApi{
		TaskEndPoint{
			Client: getTestingClient(),
		},
	}

	taskId := "15553648"
	payload := &EditTodoItem{
		ColumnId: 2531,
	}

	err := taskApi.Update(taskId, payload)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	payload.ColumnId = 2532
	err = taskApi.Update(taskId, payload)
	if err != nil {
		t.Errorf("error: %v", err)
	}
}
