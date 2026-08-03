package notifications

import (
	"encoding/json"
	"errors"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/models"
	"strings"
)

// Notifications service
type Notifications struct {
	client client.Client
}

func New(clt client.Client) *Notifications {
	return &Notifications{
		client: clt,
	}
}

type ListOptions struct {
	Queries []string
	enabledSetters map[string]bool
}
func (options ListOptions) New() *ListOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
	}
	return &options
}
type ListOption func(*ListOptions)
func (srv *Notifications) WithListQueries(v []string) ListOption {
	return func(o *ListOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
	
// List get the list of notifications for the currently logged in console
// user. Use queries to filter the results by attributes such as read status,
// view timestamps, or creation date.
func (srv *Notifications) List(optionalSetters ...ListOption)(*models.NotificationList, error) {
	path := "/notifications"
	options := ListOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"accept": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.NotificationList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.NotificationList
	parsed, ok := resp.Result.(models.NotificationList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// Update update a notification by its unique ID. Use the `read` parameter to
// mark the notification as read or unread.
func (srv *Notifications) Update(NotificationId string, Read bool)(*models.Notification, error) {
	r := strings.NewReplacer("{notificationId}", NotificationId)
	path := r.Replace("/notifications/{notificationId}")
	params := map[string]interface{}{}
	params["notificationId"] = NotificationId
	params["read"] = Read
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type": "application/json",
		"accept": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Notification{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Notification
	parsed, ok := resp.Result.(models.Notification)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
