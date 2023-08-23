package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type Albums struct {
	ID     string
	Title  string
	Artist string
	Price  float64
}

type Handler struct{}

// albums slice
var albums = []Albums{
	{ID: "1", Title: "Greatest Hits 2Pac", Artist: "Tupac Shakur", Price: 12.99},
	{ID: "2", Title: "Hans Zimmer Masterpieces", Artist: "Hans Zimmer", Price: 13.99},
	{ID: "3", Title: "Chill Out Classics", Artist: "Various Artists", Price: 15.99},
}

func (c *Handler) Controller(f *fiber.Ctx) error {
	return f.SendString("here is a string")
}

func NewHandler() *Handler {
	return &Handler{}
}

func GetAlbums(c *fiber.Ctx) error {
	return c.JSON(albums)
}

func GetAlbumsById(c *fiber.Ctx) error {
	id := c.Params("id")
	for _, album := range albums {
		if album.ID == id {
			return c.JSON(album)
		}

	}
	return c.SendString("id does not exist")
}
