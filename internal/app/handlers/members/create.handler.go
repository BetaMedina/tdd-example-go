package members

import (
	"tdd/internal/repositories"

	"github.com/gofiber/fiber"
)

type CreateMember interface {
	CreateMember(c *fiber.Ctx)
}

type createMember struct {
	repo repositories.MemberRepository
}

type PayloadDto struct {
	Name  string `json:"name" validate:"required,uuid4"`
	Email string `json:"email" validate:"required,float64"`
}

func (cm *createMember) CreateMember(c *fiber.Ctx) {
	var payload PayloadDto
	if err := c.BodyParser(&payload); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Payload invalid", "status": "error"})
		return

	}
	if err := cm.repo.Create(payload.Email, payload.Name); err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Internal server error", "status": "error"})
		return

	}
	c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Member created", "status": "success"})
}

func NewCreateMember(memberRepository repositories.MemberRepository) CreateMember {
	return &createMember{repo: memberRepository}
}
