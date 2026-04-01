package main

import (
	"context"
	"fmt"
	"log"

	"github.com/chiaf1/solar-api/internal/config"
	"github.com/chiaf1/solar-api/internal/influx"
)

const CONFIG_PATH = `C:\Generale\Progetti\Lettura produzione e consumo arduino\solar-api\config.yaml`

func main() {
	// Load configs from file
	var conf config.Config
	err := conf.Load(CONFIG_PATH)
	if err != nil {
		log.Fatal(err)
	}
	err = conf.Validate()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Config Loaded")

	// Creating an influx db client and testing some querys
	infClient := influx.New(conf.InfluxDB.Url, conf.InfluxDB.Token, conf.InfluxDB.Org, conf.InfluxDB.Bucket)

	//test query
	flux := fmt.Sprintf(`from(bucket: "%s")
	|> range(start: today(), stop: now())
	|> filter(fn: (r) => r._measurement == "energy")
	|> aggregateWindow(every: 1m, fn: mean)
	|> yield()`, conf.InfluxDB.Bucket)
	test, err := infClient.Query(context.Background(), flux)
	if err != nil {
		panic(err)
	}
	fmt.Println(test)
}
