package dao

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type (
	// todoModel 包括了 todoModel 的字段类型
	TodoModel struct {
		gorm.Model
		Title     string `json:"title"`
		Completed int    `json:"completed"`
	}

	// transformedTodo 代表格式化的 todo 结构体
	TransformedTodo struct {
		ID        uint   `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
)

func DoCreateData(title string, completed int) error {
	todo := TodoModel{Title: title, Completed: completed}
	if err := orm.Save(&todo).Error; err != nil {
		return err
	}
	return nil
}

func DoFetchAllTodo() ([]TransformedTodo, error) {
	var _todos []TransformedTodo

	var todos []TodoModel
	if err := orm.Find(&todos).Error; err != nil {
		return nil, err
	}
	//转化 todos 数据，用来格式化
	for _, item := range todos {
		completed := false
		if item.Completed == 1 {
			completed = true
		} else {
			completed = false
		}
		_todos = append(_todos, TransformedTodo{ID: item.ID, Title: item.Title, Completed: completed})
	}
	return _todos, nil
}

func DoFetchSingleTodo(todoID string) (*TransformedTodo, error) {
	var todo TodoModel
	if err := orm.First(&todo, todoID).Error; err != nil {
		return nil, err
	}
	completed := false
	if todo.Completed == 1 {
		completed = true
	} else {
		completed = false
	}
	_todo := TransformedTodo{ID: todo.ID, Title: todo.Title, Completed: completed}
	return &_todo, nil
}

func DoUpdateTodo(todoID, title string, completed int) error {
	var todo TodoModel
	orm.First(&todo, todoID)
	if todo.ID == 0 {
		return errors.New("there is a error")
	}
	orm.Model(&todo).Update("title", title)
	orm.Model(&todo).Update("completed", completed)
	return nil
}

func DoDeleteTodo(todoID string) error {
	var todo TodoModel
	orm.First(&todo, todoID)
	orm.Delete(&todo)
	return nil
}
