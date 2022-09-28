package todos

import "gorm.io/gorm"

type Repository interface {
	AddTodo(todos Todo) (Todo, error)
	FindAllTodo(ID int) ([]Todo, error)
	DeleteTodo(ID int) (Todo, error)
	FindByIDTodo(ID int) (Todo, error)
	UpdateTodo(update Todo) (Todo, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) AddTodo(Todo Todo) (Todo, error) {
	err := r.db.Create(&Todo).Error
	if err != nil {
		return Todo, err
	}
	return Todo, nil
}

func (r *repository) FindAllTodo(ID int) ([]Todo, error) {
	var Todo []Todo
	err := r.db.Order("id desc").Where("user_id", ID).Find(&Todo).Error
	if err != nil {
		return Todo, err
	}
	return Todo, nil
}

func (r *repository) DeleteTodo(ID int) (Todo, error) {
	var Todo Todo
	err := r.db.Where("id =?", ID).Delete(&Todo).Error
	if err != nil {
		return Todo, err
	}
	return Todo, nil
}

func (r *repository) FindByIDTodo(ID int) (Todo, error) {
	var Todo Todo
	err := r.db.Where("id =?", ID).Find(&Todo).Error
	if err != nil {
		return Todo, err
	}
	return Todo, nil
}

func (r *repository) UpdateTodo(update Todo) (Todo, error) {
	var Todo Todo
	err := r.db.Save(&update).Error
	if err != nil {
		return Todo, err
	}
	return Todo, nil
}
