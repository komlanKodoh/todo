package controller

import (
	"fmt"
	"github.com/devfeel/mapper"
	"gorm.io/gorm"
	"time"
	"todo.app/repository"
	"todo.app/utils"
)

// CreateTodo / Create TodoItem
func CreateTodo(api *utils.ApiContext) {
	todo, err := utils.BindJSON[AddTodoDto](api)

	if err != nil {
		return
	}

	var newTodo = repository.Todo{
		Title:       todo.Title,
		Description: todo.Description,
	}

	api.Db.Create(&newTodo)
	api.Log.Info().Msg("Todo Successfully created")

	api.Success(
		GetTodoDto{
			Title:       newTodo.Title,
			Description: newTodo.Description,
			CompletedAt: newTodo.CompletedAt,
		},
	)
}

// PatchTodo / Create TodoItem
func PatchTodo(api *utils.ApiContext) {
	todoUpdate, err := utils.BindJSON[PatchTodoDto](api)

	if err != nil {
		return
	}

	var idParam = api.Param("id")

	var todo repository.Todo
	result := api.Db.First(&todo, idParam)

	if result.RowsAffected == 0 {
		api.NotFound("Todo with specified id not found")
		return
	}

	if todoUpdate.Title != nil {
		todo.Title = *todoUpdate.Title
	}

	if todoUpdate.Description != nil {
		todo.Description = *todoUpdate.Description
	}

	// If Mark as completed flag is true ensure that it has not been completed yet and mark it as completed
	if todoUpdate.Completed != nil && *todoUpdate.Completed {
		if todo.CompletedAt.Valid {
			api.NotFound("Todo has already been completed")
			return
		}

		todo.CompletedAt = gorm.DeletedAt{
			Time:  time.Now(),
			Valid: true,
		}
	}

	// If Mark as completed flag is false, mark it as completed
	if todoUpdate.Completed != nil && !*todoUpdate.Completed {
		todo.CompletedAt = gorm.DeletedAt{}
	}

	api.Db.Save(&todo)
	api.Log.Info().Msg("Todo Successfully updated")

	var getTodoDto GetTodoDto

	err = mapper.AutoMapper(&todo, &getTodoDto)

	if err != nil {
		api.InternalServerError("Error occurred while mapping todos to dto")
		return
	}

	api.Success(
		getTodoDto,
	)
}

// GetTodos / Get TodoItems with provided filters
func GetTodos(api *utils.ApiContext) {
	var todoFilter, err = utils.BindJSON[RequestTodoDto](api)

	if err != nil {
		return
	}

	var query = &api.Db

	if todoFilter.Title != nil {
		query = query.Where("title like ?", utils.SqlLike(*todoFilter.Title))
	}

	if todoFilter.Description != nil {
		query = query.Where("title = ?", utils.SqlLike(*todoFilter.Title))
	}

	if todoFilter.Completed != nil && *todoFilter.Completed {
		query = query.Where("completed_at IS NOT NULL")
	}

	if todoFilter.Completed != nil && !*todoFilter.Completed {
		query = query.Where("completed_at IS NULL")
	}

	todosDto := []GetTodoDto{}
	var todos []repository.Todo

	var _ = query.Find(&todos)

	err = mapper.MapperSlice(&todos, &todosDto)

	if err != nil {
		api.InternalServerError("Error occurred while mapping todos to dto")
		return
	}

	api.Success(todosDto)
}

// DeleteTodo / Delete TodoItem with specified id
func DeleteTodo(api *utils.ApiContext) {
	var id = api.Param("id")

	api.Db.Delete(&repository.Todo{}, id)
	api.Log.Info().Msg("Todo Successfully deleted")

	api.ReturnSuccess(
		fmt.Sprintf("Todo with id %s Successfully deleted", id),
	)
}
