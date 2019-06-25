package teamwork

import (
	"strconv"
)

type CardApi struct {
	CardEndPoint
}

func (cardApi CardApi) Create(columnId int, taskId int, afterId int) (string, error) {

	payload := struct {
		Card struct {
			TaskId int `json:"taskId"`
		} `json:"card"`
		PositionAfterId int `json:"positionAfterId"`
	}{
		Card: struct {
			TaskId int `json:"taskId"`
		}{
			TaskId: taskId,
		},
		PositionAfterId: afterId,
	}

	return cardApi.CardEndPoint.Create(&payload, columnId)
}

func (cardApi CardApi) Move(cardId, columnId, afterId int) error {

	payload := moveCardRequest{
		CardId:          cardId,
		PositionAfterId: afterId,
		ColumnId:        columnId,
	}

	return cardApi.CardEndPoint.Move(&payload, cardId)
}

type CardEndPoint struct {
	Client ApiClient
}

func (card CardEndPoint) Create(payload interface{}, columnId int) (string, error) {

	endpoint := "boards/columns/" + strconv.Itoa(columnId) + "/cards.json"
	out := newCardResponse{}

	err := card.Client.DoRequest("POST", endpoint, payload, &out)
	if err != nil {
		return "", err
	}

	return out.CardId, nil

}

func (card CardEndPoint) Move(payload interface{}, cardId int) error {

	endpoint := "boards/columns/cards/" + strconv.Itoa(cardId) + "/move.json"
	var out interface{}
	return card.Client.DoRequest("PUT", endpoint, payload, &out)

}

type newCardResponse struct {
	CardId string
	TaskId string
	STATUS string
}

type moveCardRequest struct {
	CardId          int `json:"cardId"`
	PositionAfterId int `json:"positionAfterId"`
	ColumnId        int `json:"columnId"`
}
