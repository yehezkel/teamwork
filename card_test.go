package teamwork

import (
	"testing"
)

func TestCreateCard(t *testing.T) {

	columnId := 2531
	taskId := 15553648

	cardApi := CardApi{
		CardEndPoint{
			Client: getTestingClient(),
		},
	}

	id, err := cardApi.Create(columnId, taskId, -1)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	t.Log("New CardId: " + id)
}

func TestMoveCard(t *testing.T) {

	columnId := 2531
	cardId := 362383

	cardApi := CardApi{
		CardEndPoint{
			Client: getTestingClient(),
		},
	}

	err := cardApi.Move(cardId, columnId, -1)
	if err != nil {
		t.Errorf("error: %v", err)
	}
}
