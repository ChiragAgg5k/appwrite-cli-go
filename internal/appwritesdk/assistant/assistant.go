package assistant

import (
	"encoding/json"
	"errors"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
	"strings"
)

// Assistant service
type Assistant struct {
	client client.Client
}

func New(clt client.Client) *Assistant {
	return &Assistant{
		client: clt,
	}
}

	
// Chat send a prompt to the AI assistant and receive a response. This
// endpoint allows you to interact with Appwrite's AI assistant by sending
// questions or prompts and receiving helpful responses in real-time through a
// server-sent events stream.
func (srv *Assistant) Chat(Prompt string)(*interface{}, error) {
	path := "/console/assistant"
	params := map[string]interface{}{}
	params["prompt"] = Prompt
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type": "application/json",
		"accept": "text/plain",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		var parsed interface{}

		err = json.Unmarshal(bytes, &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	var parsed interface{}
	parsed, ok := resp.Result.(interface{})
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
