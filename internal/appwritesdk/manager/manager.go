package manager

import (
	"encoding/json"
	"errors"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/models"
	"strings"
)

// Manager service
type Manager struct {
	client client.Client
}

func New(clt client.Client) *Manager {
	return &Manager{
		client: clt,
	}
}

type CreateBlockOptions struct {
	ResourceId string
	Mode string
	Reason string
	ExpiredAt string
	enabledSetters map[string]bool
}
func (options CreateBlockOptions) New() *CreateBlockOptions {
	options.enabledSetters = map[string]bool{
		"ResourceId": false,
		"Mode": false,
		"Reason": false,
		"ExpiredAt": false,
	}
	return &options
}
type CreateBlockOption func(*CreateBlockOptions)
func (srv *Manager) WithCreateBlockResourceId(v string) CreateBlockOption {
	return func(o *CreateBlockOptions) {
		o.ResourceId = v
		o.enabledSetters["ResourceId"] = true
	}
}
func (srv *Manager) WithCreateBlockMode(v string) CreateBlockOption {
	return func(o *CreateBlockOptions) {
		o.Mode = v
		o.enabledSetters["Mode"] = true
	}
}
func (srv *Manager) WithCreateBlockReason(v string) CreateBlockOption {
	return func(o *CreateBlockOptions) {
		o.Reason = v
		o.enabledSetters["Reason"] = true
	}
}
func (srv *Manager) WithCreateBlockExpiredAt(v string) CreateBlockOption {
	return func(o *CreateBlockOptions) {
		o.ExpiredAt = v
		o.enabledSetters["ExpiredAt"] = true
	}
}
					
// CreateBlock creates a new resource block.
func (srv *Manager) CreateBlock(ProjectId string, ResourceType string, optionalSetters ...CreateBlockOption)(*models.Block, error) {
	path := "/manager/blocks"
	options := CreateBlockOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["projectId"] = ProjectId
	params["resourceType"] = ResourceType
	if options.enabledSetters["ResourceId"] {
		params["resourceId"] = options.ResourceId
	}
	if options.enabledSetters["Mode"] {
		params["mode"] = options.Mode
	}
	if options.enabledSetters["Reason"] {
		params["reason"] = options.Reason
	}
	if options.enabledSetters["ExpiredAt"] {
		params["expiredAt"] = options.ExpiredAt
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
		"accept": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Block{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Block
	parsed, ok := resp.Result.(models.Block)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type DeleteBlockOptions struct {
	ResourceId string
	enabledSetters map[string]bool
}
func (options DeleteBlockOptions) New() *DeleteBlockOptions {
	options.enabledSetters = map[string]bool{
		"ResourceId": false,
	}
	return &options
}
type DeleteBlockOption func(*DeleteBlockOptions)
func (srv *Manager) WithDeleteBlockResourceId(v string) DeleteBlockOption {
	return func(o *DeleteBlockOptions) {
		o.ResourceId = v
		o.enabledSetters["ResourceId"] = true
	}
}
					
// DeleteBlock deletes resource blocks for a project.
func (srv *Manager) DeleteBlock(ProjectId string, ResourceType string, optionalSetters ...DeleteBlockOption)(*models.BlockDelete, error) {
	path := "/manager/blocks"
	options := DeleteBlockOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["projectId"] = ProjectId
	params["resourceType"] = ResourceType
	if options.enabledSetters["ResourceId"] {
		params["resourceId"] = options.ResourceId
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
		"accept": "application/json",
	}

	resp, err := srv.client.Call("DELETE", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.BlockDelete{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.BlockDelete
	parsed, ok := resp.Result.(models.BlockDelete)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// ListBlocks lists all resource blocks for a project.
func (srv *Manager) ListBlocks(ProjectId string)(*models.BlockList, error) {
	r := strings.NewReplacer("{projectId}", ProjectId)
	path := r.Replace("/manager/blocks/{projectId}")
	params := map[string]interface{}{}
	params["projectId"] = ProjectId
	headers := map[string]interface{}{
		"accept": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.BlockList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.BlockList
	parsed, ok := resp.Result.(models.BlockList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type DeleteCacheOptions struct {
	Region string
	Cache string
	All bool
	Database string
	ProjectId string
	CollectionId string
	DocumentId string
	enabledSetters map[string]bool
}
func (options DeleteCacheOptions) New() *DeleteCacheOptions {
	options.enabledSetters = map[string]bool{
		"Region": false,
		"Cache": false,
		"All": false,
		"Database": false,
		"ProjectId": false,
		"CollectionId": false,
		"DocumentId": false,
	}
	return &options
}
type DeleteCacheOption func(*DeleteCacheOptions)
func (srv *Manager) WithDeleteCacheRegion(v string) DeleteCacheOption {
	return func(o *DeleteCacheOptions) {
		o.Region = v
		o.enabledSetters["Region"] = true
	}
}
func (srv *Manager) WithDeleteCacheCache(v string) DeleteCacheOption {
	return func(o *DeleteCacheOptions) {
		o.Cache = v
		o.enabledSetters["Cache"] = true
	}
}
func (srv *Manager) WithDeleteCacheAll(v bool) DeleteCacheOption {
	return func(o *DeleteCacheOptions) {
		o.All = v
		o.enabledSetters["All"] = true
	}
}
func (srv *Manager) WithDeleteCacheDatabase(v string) DeleteCacheOption {
	return func(o *DeleteCacheOptions) {
		o.Database = v
		o.enabledSetters["Database"] = true
	}
}
func (srv *Manager) WithDeleteCacheProjectId(v string) DeleteCacheOption {
	return func(o *DeleteCacheOptions) {
		o.ProjectId = v
		o.enabledSetters["ProjectId"] = true
	}
}
func (srv *Manager) WithDeleteCacheCollectionId(v string) DeleteCacheOption {
	return func(o *DeleteCacheOptions) {
		o.CollectionId = v
		o.enabledSetters["CollectionId"] = true
	}
}
func (srv *Manager) WithDeleteCacheDocumentId(v string) DeleteCacheOption {
	return func(o *DeleteCacheOptions) {
		o.DocumentId = v
		o.enabledSetters["DocumentId"] = true
	}
}
	
// DeleteCache clears internal cache.
func (srv *Manager) DeleteCache(optionalSetters ...DeleteCacheOption)(*interface{}, error) {
	path := "/manager/cache"
	options := DeleteCacheOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Region"] {
		params["region"] = options.Region
	}
	if options.enabledSetters["Cache"] {
		params["cache"] = options.Cache
	}
	if options.enabledSetters["All"] {
		params["all"] = options.All
	}
	if options.enabledSetters["Database"] {
		params["database"] = options.Database
	}
	if options.enabledSetters["ProjectId"] {
		params["projectId"] = options.ProjectId
	}
	if options.enabledSetters["CollectionId"] {
		params["collectionId"] = options.CollectionId
	}
	if options.enabledSetters["DocumentId"] {
		params["documentId"] = options.DocumentId
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
		"accept": "application/json",
	}

	resp, err := srv.client.Call("DELETE", path, headers, params)
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
type UpdateUserStatusOptions struct {
	UserId string
	Email string
	Reason string
	enabledSetters map[string]bool
}
func (options UpdateUserStatusOptions) New() *UpdateUserStatusOptions {
	options.enabledSetters = map[string]bool{
		"UserId": false,
		"Email": false,
		"Reason": false,
	}
	return &options
}
type UpdateUserStatusOption func(*UpdateUserStatusOptions)
func (srv *Manager) WithUpdateUserStatusUserId(v string) UpdateUserStatusOption {
	return func(o *UpdateUserStatusOptions) {
		o.UserId = v
		o.enabledSetters["UserId"] = true
	}
}
func (srv *Manager) WithUpdateUserStatusEmail(v string) UpdateUserStatusOption {
	return func(o *UpdateUserStatusOptions) {
		o.Email = v
		o.enabledSetters["Email"] = true
	}
}
func (srv *Manager) WithUpdateUserStatusReason(v string) UpdateUserStatusOption {
	return func(o *UpdateUserStatusOptions) {
		o.Reason = v
		o.enabledSetters["Reason"] = true
	}
}
			
// UpdateUserStatus updates a console user status using a user ID or email
// address.
func (srv *Manager) UpdateUserStatus(Status bool, optionalSetters ...UpdateUserStatusOption)(*models.User, error) {
	path := "/manager/users/status"
	options := UpdateUserStatusOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["status"] = Status
	if options.enabledSetters["UserId"] {
		params["userId"] = options.UserId
	}
	if options.enabledSetters["Email"] {
		params["email"] = options.Email
	}
	if options.enabledSetters["Reason"] {
		params["reason"] = options.Reason
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
		"accept": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.User{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.User
	parsed, ok := resp.Result.(models.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
