package handler

import (
	m "bam/internal/app/model"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	// "strings"
	// "github.com/gofiber/fiber/v2"
)

type ApproveActions interface {
	RegisterOrdination(reg *m.RegisOrdinary) error
	FindOrdinationByID(id uint) (*m.RegisOrdinary, error)
	UpdateOrdination(user *m.RegisOrdinary) error
	DeleteOrdination(id uint) error
	FindOrdinationByName(name string) ([]*m.RegisOrdinary, error)
	FindOrdinations() ([]*m.RegisOrdinary, error)
	FindOrdinationByStatus(status string) (*m.RegisOrdinary, error)
	UpdateOrdinationStatus(id uint, status, comment string) error
}

type ApproveHandler struct {
	service ApproveActions
}

func NewApproveHandler(service ApproveActions) *ApproveHandler {
	return &ApproveHandler{service: service}
}

func (h *ApproveHandler) RegisterOrdination(c *fiber.Ctx) error {
	reg := new(m.RegisOrdinary)
	if err := c.BodyParser(reg); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Validate gender
	switch reg.Gender {
	case m.Man, m.Woman, m.PreferNotToSay, m.Alternative:
		// Valid gender, proceed
	default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid gender"})
	}

	// Validate and format birthday
	birthday, err := time.Parse("02/01/2006", reg.Birthday)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid birthday format. Use DD/MM/YYYY"})
	}
	reg.Birthday = birthday.Format("02/01/2006")

	err = h.service.RegisterOrdination(reg)
	if err != nil {
		// Check if the error is due to the user already existing
		if strings.Contains(err.Error(), "already exists") {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(reg)
}

func (h *ApproveHandler) FindOrdinationByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	ord, err := h.service.FindOrdinationByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(ord)
}

func (h *ApproveHandler) UpdateOrdination(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	new := &m.RegisOrdinary{}
	if err := c.BodyParser(new); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	new.ID = uint(id)

	if err := h.service.UpdateOrdination(new); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(new)
}

func (h *ApproveHandler) DeleteOrdination(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := h.service.DeleteOrdination(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// return c.SendStatus(fiber.StatusNoContent)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"complete": "Deleted"})
}

func (h *ApproveHandler) FindOrdinations(c *fiber.Ctx) error {
	page := c.Query("page", "1")
	view := c.Query("view", "10")

	pageInt, _ := strconv.Atoi(page)
	viewInt, _ := strconv.Atoi(view)

	ords, err := h.service.FindOrdinations()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	total := len(ords)
	start := (pageInt - 1) * viewInt
	end := start + viewInt
	if end > total {
		end = total
	}

	paginatedOrds := ords[start:end]

	return c.JSON(fiber.Map{"users": paginatedOrds, "pagination": m.Pagination{Page: pageInt, View: viewInt, Total: total}, "message": "success", "status": 200, "success": true})
}

func (h *ApproveHandler) FindOrdinationByStatus(c *fiber.Ctx) error {
	status := c.Query("status")
	ords, err := h.service.FindOrdinationByStatus(status)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"users": ords, "status": 200, "success": true})
}

func (h *ApproveHandler) FindOrdinationByName(c *fiber.Ctx) error {
	search := c.Query("search")

	ord, err := h.service.FindOrdinationByName(search)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(ord)
}

func (h *ApproveHandler) UpdateOrdinationStatus(c *fiber.Ctx) error {
	// Parse request body
	var req struct {
		ID      uint   `json:"id"`
		Status  string `json:"status"`
		Comment string `json:"comment"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Update ordination status and comment
	err := h.service.UpdateOrdinationStatus(req.ID, req.Status, req.Comment)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Ordination status and comment updated successfully", "status": 200, "success": true})
}
