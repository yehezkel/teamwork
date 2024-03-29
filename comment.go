package teamwork

type CommentApi struct {
	CommentEndPoint
}

func (c CommentApi) Fetch(id string) (*Comment, error) {

	out := FetchCommentResponse{}
	err := c.CommentEndPoint.Fetch(id, &out)

	if err != nil {
		return nil, err
	}

	return out.Comment, nil
}

func (c CommentApi) Create(n *NewComment, resource string, resourceId string) (string, error) {

	payload := &newCommentRequest{
		Comment: n,
	}

	id, err := c.CommentEndPoint.Create(payload, resource, resourceId)
	if err != nil {
		return "", err
	}

	return id, nil
}

type CommentEndPoint struct {
	Client ApiClient
}

func (c CommentEndPoint) Fetch(id string, out interface{}) error {

	endpoint := "comments/" + id + ".json"
	err := c.Client.DoRequest("GET", endpoint, nil, out)

	if err != nil {
		return err
	}

	return nil
}

func (c CommentEndPoint) Create(payload interface{}, resource string, resourceId string) (string, error) {

	endpoint := resource + "/" + resourceId + "/comments.json"
	out := newCommentResponse{}

	err := c.Client.DoRequest("POST", endpoint, payload, &out)
	if err != nil {
		return "", err
	}

	return out.CommentId, nil
}

type FetchCommentResponse struct {
	Status  string `json:"STATUS"`
	Comment *Comment
}

type Comment struct {
	ProjectId       string `json:"project-id"`
	Id              string
	ItemName        string `json:"item-name"`
	ItemType        string `json:"type"`
	Body            string
	HtmlBody        string `json:"html-body"`
	AuthorId        string `json:"author-id"`
	AuthorFirstName string `json:"author-firstname"`
	ContentType     string `json:"content-type"`
}

type NewComment struct {
	Body        string `json:"body"`
	Notify      string `json:"notify"`
	Private     bool   `json:"isprivate"`
	ContentType string `json:"content-type"`
}

type newCommentRequest struct {
	Comment *NewComment `json:"comment"`
}

type newCommentResponse struct {
	CommentId string
	STATUS    string
}
