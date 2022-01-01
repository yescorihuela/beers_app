package api

type DescriptionResponse struct {
	Description string      `json:"description"`
	Content     interface{} `json:"content,omitempty"`
}

func (d *DescriptionResponse) NewDescriptionResponse(description string, content interface{}) {
	d.Description = description
	d.Content = content
}
