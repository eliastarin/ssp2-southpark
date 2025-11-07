package domain

type Message struct {
	Author string `json:"author"`
	Body   string `json:"body"`
}

func (m Message) Validate() error {
	if m.Author == "" {
		return ErrInvalidAuthor
	}
	if m.Body == "" {
		return ErrInvalidBody
	}
	return nil
}
