package main

import (
    "encoding/json"

    "github.com/gofiber/fiber/v2"
    zmq "github.com/pebbe/zmq4"
	"os"
)

type Input struct {
    Title   string `json:"title"`
    Content string `json:"content"`
}

func main() {

	port :=  os.Getenv("PUB_PORT")
    app := fiber.New()
    zctx, _ := zmq.NewContext()

    s, _ := zctx.NewSocket(zmq.PUB)

	if port == "" {
		port = "5553"
	}

    s.Bind("tcp://*:" + port)

    app.Post("/", func(c *fiber.Ctx) error {
        input := new(Input)
        if err := c.BodyParser(input); err != nil {
            panic(err)
        }

        if buffer, err := json.Marshal(input); err != nil {
            panic(err)
        } else {
            s.Send("example", zmq.SNDMORE)
            s.Send(string(buffer), 0)
        }

		c.SendString("Message created!")
        return c.SendStatus(201)
    })

    app.Listen(":3000")
}
