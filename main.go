package main

import (
	"github.com/shivaraj-shanthaiah/code_orbit_notification/pkg/config"
	"github.com/shivaraj-shanthaiah/code_orbit_notification/pkg/rabbitmq"
)

func main() {

	cfg := config.LoadConfig()
	rabbitmq.ConsumeNotificationMessages(cfg)

}
