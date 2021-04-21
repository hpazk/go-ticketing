package event

type Services interface {
	SaveEvent(req *request) (Event, error)
	FetchEvents() ([]Event, error)
	FetchEvent(id uint) (Event, error)
	EditEvent(id uint) (Event, error)
	RemoveEvent(id uint) error
}

type services struct {
	repo repository
}

func eventService(repo repository) *services {
	return &services{repo}
}

func (s *services) SaveEvent(req *request) (Event, error) {
	var event Event
	return event, nil
}

func (s *services) FetchEvents() ([]Event, error) {
	var event []Event
	return event, nil
}

func (s *services) FetchEvent(id uint) (Event, error) {
	var event Event
	return event, nil
}

func (s *services) EditEvent(id uint) (Event, error) {
	var event Event
	return event, nil
}

func (s *services) RemoveEvent(id uint) error {
	return nil
}
