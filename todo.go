package main

import (
	"github.com/goadesign/goa"
	"github.com/jarifibrahim/todo/app"
)

// TodoController implements the todo resource.
type TodoController struct {
	*goa.Controller
	repository todoRepository
}

// NewTodoController creates a todo controller.
func NewTodoController(service *goa.Service) *TodoController {
	return &TodoController{Controller: service.NewController("TodoController")}
}

// Create runs the create action.
func (c *TodoController) Create(ctx *app.CreateTodoContext) error {
	todo, err := c.repository.Create(ctx, &Todo{
		Title:       ctx.Payload.Title,
		Description: ctx.Payload.Description,
	})

	if err != nil {
		// Todo - return internal server error
		return nil
	}
	return ctx.Created(&app.Todo{
		ID:          int(todo.ID),
		Title:       todo.Title,
		Description: todo.Description,
	})
}

// Delete runs the delete action.
func (c *TodoController) Delete(ctx *app.DeleteTodoContext) error {
	err := c.repository.Delete(ctx, uint(ctx.ID))
	if err != nil {
		// Todo - Send error message
		return ctx.NotFound()
	}
	// Todo - return success message
	return ctx.OK([]byte{})
}

// List runs the list action.
func (c *TodoController) List(ctx *app.ListTodoContext) error {
	todoList, err := c.repository.List(ctx)
	if err != nil {
		// TODO - send error message
		return ctx.NotFound()
	}
	resp := []*app.Todo{}
	for _, t := range todoList {
		resp = append(resp, &app.Todo{
			ID:          int(t.ID),
			Title:       t.Title,
			Description: t.Description,
		})
	}
	return ctx.OK(resp)
}

// Show runs the show action.
func (c *TodoController) Show(ctx *app.ShowTodoContext) error {
	todoItem, err := c.repository.Show(ctx, uint(ctx.ID))
	if err != nil {
		// Todo - return NotFound with error message
		return ctx.NotFound()
	}

	res := &app.Todo{
		ID:          int(todoItem.ID),
		Title:       todoItem.Title,
		Description: todoItem.Description,
	}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *TodoController) Update(ctx *app.UpdateTodoContext) error {
	newTodo, err := c.repository.Update(ctx, &Todo{
		Title:       ctx.Payload.Title,
		Description: ctx.Payload.Description,
	}, uint(ctx.ID))
	if err != nil {
		// TODO - This should be internal server error
		return ctx.BadRequest()
	}
	res := &app.Todo{
		ID:          int(newTodo.ID),
		Title:       newTodo.Title,
		Description: newTodo.Description,
	}
	return ctx.OK(res)
}
