package todos

type Service interface {
	AddTodo(Todo AddTodo) (Todo, error)
	FindAllTodo(ID int) ([]Todo, error)
	DeleteTodo(ID int) (Todo, error)
	UpdateTodo(update UpdateTodo) (Todo, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) AddTodo(Todos AddTodo) (Todo, error) {
	Todo := Todo{}
	Todo.Title = Todos.Title
	Todo.LongText = Todos.LongText
	Todo.UserID = Todos.UserID

	notifikasi, err := s.repository.AddTodo(Todo)
	if err != nil {
		return notifikasi, err
	}
	return notifikasi, nil
}

func (s *service) FindAllTodo(ID int) ([]Todo, error) {
	Todo, err := s.repository.FindAllTodo(ID)

	if err != nil {
		return Todo, nil
	}
	return Todo, nil
}

func (s *service) DeleteTodo(ID int) (Todo, error) {
	Todo, err := s.repository.DeleteTodo(ID)

	if err != nil {
		return Todo, nil
	}
	return Todo, nil
}

func (s *service) UpdateTodo(update UpdateTodo) (Todo, error) {
	TodoData, err := s.repository.FindByIDTodo(update.ID)

	if err != nil {
		return TodoData, nil
	}
	if TodoData.Id != update.ID {
		return TodoData, nil
	}
	TodoData.Title = update.Title
	TodoData.LongText = update.LongText
	TodoData.Id = update.ID
	Todo, err := s.repository.UpdateTodo(TodoData)

	if err != nil {
		return Todo, nil
	}
	return TodoData, nil
}
