package province

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	us Usecase
}

func NewHandler(app *fiber.App, usecase Usecase) {
	h := &handler{
		us: usecase,
	}

	app.Get("/province", h.GetAll)
	app.Post("/province", h.Create)
	app.Get("/province/:myid", h.GetByID)
	app.Patch("/province/:myid", h.Update)
	app.Delete("/province/:myid", h.Delete)
}

type ProvinceRequest struct {
	ID     uint   `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	NameEn string `json:"name_en,omitempty"`
}

func (p ProvinceRequest) Validate() error {
	if p.NameEn == "" || p.Name == "" {
		return errors.New("invalid name_en")
	}
	return nil
}

func (h handler) GetAll(c *fiber.Ctx) error {
	i, err := h.us.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("error")
	}
	var response []ProvinceRequest
	for _, item := range i {
		p := ProvinceRequest{ID: item.ID, Name: item.Name, NameEn: item.NameEn}
		response = append(response, p)
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (h handler) Create(c *fiber.Ctx) error {
	var body ProvinceRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("invalid body")
	}
	// validate
	if err := body.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("invalid body")
	}
	// create
	id, err := h.us.Create(body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("error")
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "successfull.",
		"id":      id,
	})
}

func (h handler) GetByID(c *fiber.Ctx) error {
	id := c.Params("myid")
	u64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	provinceID := uint(u64)
	// get province
	i, err := h.us.GetByID(provinceID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("error")
	}
	response := ProvinceRequest{Name: i.Name, NameEn: i.NameEn}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (h handler) Update(c *fiber.Ctx) error {
	id := c.Params("myid") // return string
	u64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	provinceID := uint(u64)

	var body ProvinceRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("invalid body")
	}
	err = h.us.Update(provinceID, body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("error")
	}
	return c.Status(fiber.StatusOK).JSON("successfully.")
}

func (h handler) Delete(c *fiber.Ctx) error {
	id := c.Params("myid") // return string
	u64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	provinceID := uint(u64)
	if err := h.us.Delete(provinceID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("error")
	}
	return c.Status(fiber.StatusOK).JSON("successfully.")
}
