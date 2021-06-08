package main

import (
	"database/sql"
	"strconv"

	"github.com/bangarangler/go-fiber-todos/postgres"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	_ "github.com/lib/pq"
)

type Todo struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

var todos = []*Todo{
	{Id: 1, Name: "walk the dog", Completed: false},
	{Id: 2, Name: "walk the cat", Completed: false},
}

type Handlers struct {
	Repo *postgres.Repo
}

func NewHandlers(repo *postgres.Repo) *Handlers {
	return &Handlers{Repo: repo}
}

func main() {
	db, err := sql.Open("postgres", postgres.PgConnStr)
	if err != nil {
		panic(err)
	}
	repo := postgres.NewRepo(db)
	app := fiber.New()
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	app.Get("/", func(ctx *fiber.Ctx) {
		ctx.Send("hello world")
	})

	handlers := NewHandlers(repo)

	// SetupTodosRoutes(app)
	SetupAPIV1(app, handlers)

	err = app.Listen(3000)
	if err != nil {
		panic(err)
	}

}

func SetupAPIV1(app *fiber.App, handlers *Handlers) {
	v1 := app.Group("/v1")
	SetupTodosRoutes(v1, handlers)
}

// func SetupTodosRoutes(app *fiber.App) {
func SetupTodosRoutes(grp fiber.Router, handlers *Handlers) {
	todosRoutes := grp.Group("/todos")

	todosRoutes.Get("/", handlers.GetTodos)
	todosRoutes.Get("/:id", GetTodo)
	todosRoutes.Post("/", CreateTodo)
	todosRoutes.Delete("/:id", DeleteTodo)
	todosRoutes.Patch("/:id", UpdateTodo)
}

func (h *Handlers) GetTodos(ctx *fiber.Ctx) {
	tts, err := h.Repo.GetAllTodos(ctx.Context())
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).Send(err.Error())
		return
	}
	if err := ctx.Status(fiber.StatusOK).JSON(tts); err != nil {
		return
	}
}

func CreateTodo(ctx *fiber.Ctx) {
	type request struct {
		Name string `json:"name"`
	}
	var body request
	err := ctx.BodyParser(&body)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
		return
	}

	todo := &Todo{
		Id:        len(todos) + 1,
		Name:      body.Name,
		Completed: false,
	}

	todos = append(todos, todo)
	ctx.Status(fiber.StatusCreated).JSON(todo)
}

func GetTodo(ctx *fiber.Ctx) {
	paramsId := ctx.Params("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})
		return
	}
	for _, todo := range todos {
		if todo.Id == id {
			ctx.Status(fiber.StatusOK).JSON(todo)
			return
		}
	}

	ctx.Status(fiber.StatusNotFound)
}

func DeleteTodo(ctx *fiber.Ctx) {
	paramsId := ctx.Params("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})
		return
	}

	for i, todo := range todos {
		if todo.Id == id {
			todos = append(todos[0:i], todos[i+1:]...)
			ctx.Status(fiber.StatusNoContent)
			return
		}
	}

	ctx.Status(fiber.StatusNotFound)
}

func UpdateTodo(ctx *fiber.Ctx) {
	type request struct {
		Name      *string `json:"name"`
		Completed *bool   `json:"completed"`
	}

	paramsId := ctx.Params("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})
		return
	}
	var body request
	err = ctx.BodyParser(&body)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse body",
		})
		return
	}

	var todo *Todo

	for _, t := range todos {
		if t.Id == id {
			todo = t
			break
		}
	}

	if todo == nil {
		ctx.Status(fiber.StatusNotFound)
		return
	}

	if body.Name != nil {
		todo.Name = *body.Name
	}

	if body.Completed != nil {
		todo.Completed = *body.Completed
	}

	// think todo is updated but
	ctx.Status(fiber.StatusOK).JSON(todo)
}
