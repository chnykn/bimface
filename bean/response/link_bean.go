package response

type Link struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Guid      string `json:"guid"`
	Transform string `json:"transform"`
}

type LinkBean Link
