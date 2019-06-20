package teamwork

import ()

type TaskApi struct {
	TaskEndPoint
}

func (t TaskApi) Fetch(id string) (*TodoItem, error) {

	out := TaskResponse{}
	err := t.TaskEndPoint.Fetch(id, &out)

	if err != nil {
		return nil, err
	}

	return out.Task, nil
}

type TaskEndPoint struct {
	Client ApiClient
}

func (t TaskEndPoint) Fetch(id string, out interface{}) error {

	// to simple no need for string interpolation here
	endpoint := "tasks/" + id + ".json"
	err := t.Client.DoRequest("GET", endpoint, nil, out)

	if err != nil {
		return err
	}

	return nil
}

type BoardColumn struct {
	Id    int
	Name  string
	color string
}

type TodoItem struct {
	Id          int
	BoardCol    *BoardColumn `json:"boardColumn"`
	CanComplete bool
}

type TaskResponse struct {
	Status string    `json:"STATUS"`
	Task   *TodoItem `json:"todo-item"`
}
