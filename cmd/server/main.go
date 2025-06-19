package main

import (
	"log"
	"net/http"

	"monitorX/config"
	"monitorX/drivers"

	"github.com/labstack/echo/v4"
)

func healthHandler(c echo.Context) error {
	c.Response().Header().Set("X-MIT-Response", "OK")
	return c.String(http.StatusOK, "You're set!")
}

func monitorHandler(c echo.Context) error {
	service := c.QueryParam("service")
	name := c.QueryParam("name")
	address := c.QueryParam("address")
	statusStr := c.QueryParam("status")
	desc := c.QueryParam("desc")

	status := 0
	if statusStr == "up" {
		status = 1
	}

	log.Printf("Received: %s %s %s %d %s\n", service, name, address, status, desc)
	drivers.SendToTelegram(service, name, address, status, desc)
	drivers.SendToInflux(service, name, address, status, desc)

	return c.String(http.StatusOK, "OK")
}

func main() {
	e := echo.New()

	e.GET("/", healthHandler)
	e.GET("/monitor", monitorHandler)

	log.Println("Listening on port " + config.Get().ListenPort)
	err := e.Start(":" + config.Get().ListenPort)
	if err != nil {
		log.Fatal("Server failed:", err)
	}
}
