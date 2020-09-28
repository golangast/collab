package event

// Event struct
type Event struct {
	ID          int    `JSON:"id"`
	Title       string `JSON:"title"`
	Description string `JSON:"description"`
	Date        string `JSON:"date"`
}

// NewEvent func
func NewEvent() *Event {
	e := &Event{}
	return e
}
