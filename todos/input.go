package todos

type AddTodo struct {
	Title    string `json:"title" binding:"required"`
	LongText string `json:"long_text" binding:"required"`
	UserID   int    `json:"userid" binding:"required"`
}

type UpdateTodo struct {
	ID       int    `json:"todoid" binding:"required"`
	Title    string `json:"title" binding:"required"`
	LongText string `json:"long_text" binding:"required"`
}

type AllTodo struct {
	ID int `json:"userid" binding:"required"`
}
