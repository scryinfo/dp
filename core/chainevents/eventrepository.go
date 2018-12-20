package chainevents

type EventRepository struct {
	mapEventExecutor map[string]EventCallback
}


func NewEventRepository() (*EventRepository) {
	return &EventRepository{
		mapEventExecutor: make(map[string]EventCallback),
	}
}