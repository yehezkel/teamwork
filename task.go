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

func (t TaskApi) Update(id string, task *EditTaskPayload) (*TodoItem, error) {

	var out interface{}
	payload := struct {
		task *EditTodoItem `json:"todo-item"`
	}{
		task,
	}

	err := t.TaskEndPoint.Update(id, payload, &out)

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

func (t TaskEndPoint) Update(id string, in, out interface{}) error {

	// to simple no need for string interpolation here
	endpoint := "tasks/" + id + ".json"
	return t.Client.DoRequest("PUT", endpoint, in, out)
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

type EditTaskPayload struct {
	Status string        `json:"STATUS"`
	Task   *EditTodoItem `json:"todo-item"`
}

type EditTodoItem struct {
	ColumnId int `json:"columnId"`
}
