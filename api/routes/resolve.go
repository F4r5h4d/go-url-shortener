package routes

import (
	"go-url-shortener/api/database"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func ResolveURL(c *fiber.Ctx) error {
	
	url := c.Params("url")

	r := database.CreateClient(0)
	defer r.Close()

	value, err: = r.Get(database.Ctx, url).Result()

	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{ "error": 
		"short not found in database" 
		})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			 "error": "Can't connect to DB" 
		})
	}

	rInr := database.CreateClient(1)
	defer rInr.Close()

	_ = rInr.Incr(database.Ctx, "counter")

	return c.Redirect(value, user)


}