package drivers

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"monitorX/config"
)

func SendToTelegram(service, name, address string, status int, desc string) {
	statusText := strconv.Itoa(status)
	url := "https://api.telegram.org/bot" + config.Get().BotToken +
		"/sendMessage?chat_id=-" + config.Get().ChatID +
		"&text=MonitorX Alert: " + service + " on " + name + " " + address + " status " + statusText + " - " + desc

	resp, err := http.Get(url)
	if err != nil {
		log.Println("Telegram send error:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	log.Println("Telegram response:", string(body))
}
