package url

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/muhammad-faisal-ashshidiq/go-chaweket/controller"
)

func Web(page *fiber.App) {
	page.Get("/", controller.GetIP)
	page.Get("/ws", websocket.New(controller.Websocket))

}
