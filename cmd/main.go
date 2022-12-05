package main

import (
	"fmt"
	"os"

	"github.com/bloodyrafo75/PoC-event-router-core-lib/models"
	"github.com/bloodyrafo75/PoC-event-router-core-lib/package/eventRouterCoreClient"
	"github.com/joho/godotenv"
)

var (
	EVENT_ROUTER_API_HOST string
	EVENT_ROUTER_API_PORT string
	EVENT_ROUTER_CLIENTID string
)

func main() {

	err := getEnvConfiguration()
	if err != nil {
		panic(err)
	}

	eventRouterCoreClient.SetConfiguration(EVENT_ROUTER_API_HOST, EVENT_ROUTER_API_PORT, EVENT_ROUTER_CLIENTID)
	msg := createExampleMsg()
	resp, err := eventRouterCoreClient.NotifyEvent(&msg)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

func createExampleMsg() models.MessageModel {
	attr := models.MessageAttributes{
		Src:   "IAM",
		Prod:  "fake_prod",
		Type:  "fake_type",
		Stype: "fake_stype",
		Op:    "fake_op",
	}
	return models.MessageModel{
		Payload:         "fake_payload",
		SpecificPayload: "fake_specific_payload",
		Attributes:      attr,
	}
}

// get configuration from .env file.
func getEnvConfiguration() error {
	err := godotenv.Load("configs/.env")

	if err != nil {
		return err
	}

	EVENT_ROUTER_API_HOST = os.Getenv("EVENT_ROUTER_API_HOST")
	EVENT_ROUTER_API_PORT = os.Getenv("EVENT_ROUTER_API_PORT")
	EVENT_ROUTER_CLIENTID = os.Getenv("EVENT_ROUTER_CLIENTID")

	return nil
}
