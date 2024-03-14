package router

import (
	handlers "tdd/internal/app/handlers/members"
	infra "tdd/internal/infra/database"
	"tdd/internal/repositories"

	"github.com/gofiber/fiber"
)

var memberHandler = handlers.NewCreateMember(repositories.NewMemberRepository(
	infra.GetConnection("member")),
)

func StatementRouter(app *fiber.App) {
	app.Get("/member", memberHandler.CreateMember)
}
