package handler

import (
	"todolist-go/internal/middleware"
	"todolist-go/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	router     fiber.Router
	usecase    usecase.Usecase
	middleware middleware.Middleware
}

func NewHandler(router fiber.Router, usecase usecase.Usecase, middleware middleware.Middleware) *Handler {
	return &Handler{
		router:     router,
		usecase:    usecase,
		middleware: middleware,
	}
}

func (h *Handler) MountEndpoint() {
	h.router.Get("/", h.Test)
	userGroup := h.router.Group("/users")
	userGroup.Post("/register", h.Register)
	userGroup.Post("/login", h.Login)
}

//func (h *Handler) Run() {
//	address := os.Getenv("ADDRESS")
//	port := os.Getenv("PORT")
//	err := h.router.(*fiber.App).Listen(address + ":" + port)
//	if err != nil {
//		log.Fatalf(err.Error())
//	}
//	fmt.Printf("Listening on %s:%s\n", address, port)
//}

func (h *Handler) Test(ctx *fiber.Ctx) error {
	return ctx.Status(200).SendString("test")
}
