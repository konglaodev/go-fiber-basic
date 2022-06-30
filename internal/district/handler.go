package district

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

	app.Get("/district", h.GetAll)
	app.Post("/district", h.Create)
	app.Get("/district/:myid", h.GetByID)
	app.Patch("/district/:myid", h.Update)
	app.Delete("/district/:myid", h.Delete)
}

type DistrictRequest struct {
	ID         uint   `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	NameEn     string `json:"name_en,omitempty"`
	ProvinceID uint   `json:"province_id,omitempty"`
}

func (p DistrictRequest) Validate() error {
	if p.NameEn == "" || p.Name == "" || p.ProvinceID == 0 {
		return errors.New("invalid name_en")
	}
	return nil
}

func (h handler) GetAll(c *fiber.Ctx) error {
	i, err := h.us.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("error")
	}
	var response []DistrictRequest
	for _, item := range i {
		p := DistrictRequest{ID: item.ID, Name: item.Name, NameEn: item.NameEn, ProvinceID: item.ProvinceID}
		response = append(response, p)
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (h handler) Create(c *fiber.Ctx) error {
	var body DistrictRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("invalid body")
	}
	// validate
	if err := body.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("invalid body")
	}
	fmt.Printf("body:%v\n", body)
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
	response := DistrictRequest{Name: i.Name, NameEn: i.NameEn, ProvinceID: i.ProvinceID}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (h handler) Update(c *fiber.Ctx) error {
	id := c.Params("myid") // return string
	u64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	provinceID := uint(u64)

	var body DistrictRequest
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
