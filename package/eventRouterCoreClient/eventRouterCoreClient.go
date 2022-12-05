package eventRouterCoreClient

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/bloodyrafo75/PoC-event-router-core-lib/models"
)

const CONTENT_TYPE = "application/json"

var (
	EVENT_ROUTER_API_HOST string
	EVENT_ROUTER_API_PORT string
	EVENT_ROUTER_CLIENTID string
)

func NotifyEvent(msg *models.MessageModel) (*[]byte, error) {

	//To identify who create the event.
	msg.Attributes.Src = EVENT_ROUTER_CLIENTID

	//convert to json
	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	//send to event core-api
	resp, err := http.Post(EVENT_ROUTER_API_HOST+":"+EVENT_ROUTER_API_PORT, CONTENT_TYPE, bytes.NewBuffer(jsonMsg))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &body, nil
}

func SetConfiguration(host string, port string, clientID string) {
	EVENT_ROUTER_API_HOST = host
	EVENT_ROUTER_API_PORT = port
	EVENT_ROUTER_CLIENTID = clientID
}
