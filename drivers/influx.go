package drivers

import (
	"log"
	"time"

	"monitorX/config"
	client "github.com/influxdata/influxdb1-client/v2"
)

func SendToInflux(service, name, address string, status int, desc string) {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://" + config.Get().InfluxHost + ":" + config.Get().InfluxPort,
	})
	if err != nil {
		log.Println("InfluxDB client error:", err)
		return
	}

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  config.Get().InfluxDB,
		Precision: "s",
	})

	tags := map[string]string{"address": address}
	fields := map[string]interface{}{
		"name":    name,
		"service": service,
		"status":  status,
		"desc":    desc,
	}

	point, err := client.NewPoint(config.Get().InfluxMeasurement, tags, fields, time.Now())
	if err == nil {
		bp.AddPoint(point)
		c.Write(bp)
	}
}
