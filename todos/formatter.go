package todos

type UserFormatter struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	LongText string `json:"long_text"`
	UserID   int    `json:"userid"`
}

func FormatTodo(todos Todo) UserFormatter {
	formatter := UserFormatter{
		Id:       todos.Id,
		Title:    todos.Title,
		LongText: todos.LongText,
		UserID:   todos.UserID,
	}

	return formatter
}

func FormatterAllTodo(Todo []Todo) []UserFormatter {
	TodosFormatter := []UserFormatter{}
	for _, teac := range Todo {
		userFormatter := FormatTodo(teac)
		TodosFormatter = append(TodosFormatter, userFormatter)

	}

	return TodosFormatter

}
