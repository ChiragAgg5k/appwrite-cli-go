package projects

import (
	"encoding/json"
	"errors"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/models"
	"net/url"
	"strings"
)

// Projects service
type Projects struct {
	client client.Client
}

func New(clt client.Client) *Projects {
	return &Projects{
		client: clt,
	}
}

	
// ListAddons list all billing addons for a project.
func (srv *Projects) ListAddons(ProjectId string)(*models.AddonList, error) {
	r := strings.NewReplacer("{projectId}", url.PathEscape(ProjectId))
	path := r.Replace("/projects/{projectId}/addons")
	params := map[string]interface{}{}
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

		parsed := models.AddonList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AddonList
	parsed, ok := resp.Result.(models.AddonList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// CreatePremiumGeoDBAddon create a Premium Geo DB addon for a project.
func (srv *Projects) CreatePremiumGeoDBAddon(ProjectId string)(*models.Addon, error) {
	r := strings.NewReplacer("{projectId}", url.PathEscape(ProjectId))
	path := r.Replace("/projects/{projectId}/addons/premium-geo-db")
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type": "application/json",
		"accept": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Addon{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Addon
	parsed, ok := resp.Result.(models.Addon)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// GetAddon get the details of a billing addon for a project.
func (srv *Projects) GetAddon(ProjectId string, AddonId string)(*models.Addon, error) {
	r := strings.NewReplacer("{projectId}", url.PathEscape(ProjectId), "{addonId}", url.PathEscape(AddonId))
	path := r.Replace("/projects/{projectId}/addons/{addonId}")
	params := map[string]interface{}{}
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

		parsed := models.Addon{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Addon
	parsed, ok := resp.Result.(models.Addon)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// DeleteAddon delete a billing addon for a project.
func (srv *Projects) DeleteAddon(ProjectId string, AddonId string)(*interface{}, error) {
	r := strings.NewReplacer("{projectId}", url.PathEscape(ProjectId), "{addonId}", url.PathEscape(AddonId))
	path := r.Replace("/projects/{projectId}/addons/{addonId}")
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
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
			
// ConfirmAddonPayment confirm payment for a billing addon for a project.
func (srv *Projects) ConfirmAddonPayment(ProjectId string, AddonId string)(*models.Addon, error) {
	r := strings.NewReplacer("{projectId}", url.PathEscape(ProjectId), "{addonId}", url.PathEscape(AddonId))
	path := r.Replace("/projects/{projectId}/addons/{addonId}/confirmations")
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type": "application/json",
		"accept": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Addon{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Addon
	parsed, ok := resp.Result.(models.Addon)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// GetAddonPrice get the price details for a billing addon for a project,
// including the prorated amount for the remaining days in the current billing
// cycle.
func (srv *Projects) GetAddonPrice(ProjectId string, Addon string)(*models.AddonPrice, error) {
	r := strings.NewReplacer("{projectId}", url.PathEscape(ProjectId), "{addon}", url.PathEscape(Addon))
	path := r.Replace("/projects/{projectId}/addons/{addon}/price")
	params := map[string]interface{}{}
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

		parsed := models.AddonPrice{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AddonPrice
	parsed, ok := resp.Result.(models.AddonPrice)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// UpdateConsoleAccess record console access to a project. This endpoint
// updates the last accessed timestamp for the project to track console
// activity.
func (srv *Projects) UpdateConsoleAccess(ProjectId string)(*interface{}, error) {
	r := strings.NewReplacer("{projectId}", url.PathEscape(ProjectId))
	path := r.Replace("/projects/{projectId}/console-access")
	params := map[string]interface{}{}
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
type ListDevKeysOptions struct {
	Queries []string
	enabledSetters map[string]bool
}
func (options ListDevKeysOptions) New() *ListDevKeysOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
	}
	return &options
}
type ListDevKeysOption func(*ListDevKeysOptions)
func (srv *Projects) WithListDevKeysQueries(v []string) ListDevKeysOption {
	return func(o *ListDevKeysOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
			
// ListDevKeys list all the project\'s dev keys. Dev keys are project specific
// and allow you to bypass rate limits and get better error logging during
// development.'
func (srv *Projects) ListDevKeys(ProjectId string, optionalSetters ...ListDevKeysOption)(*models.DevKeyList, error) {
	r := strings.NewReplacer("{projectId}", url.PathEscape(ProjectId))
	path := r.Replace("/projects/{projectId}/dev-keys")
	options := ListDevKeysOptions{}.New()
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

		parsed := models.DevKeyList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DevKeyList
	parsed, ok := resp.Result.(models.DevKeyList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// GetDevKey get a project\'s dev key by its unique ID. Dev keys are project
// specific and allow you to bypass rate limits and get better error logging
// during development.
func (srv *Projects) GetDevKey(ProjectId string, KeyId string)(*models.DevKey, error) {
	r := strings.NewReplacer("{projectId}", url.PathEscape(ProjectId), "{keyId}", url.PathEscape(KeyId))
	path := r.Replace("/projects/{projectId}/dev-keys/{keyId}")
	params := map[string]interface{}{}
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

		parsed := models.DevKey{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DevKey
	parsed, ok := resp.Result.(models.DevKey)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
							
// UpdateDevKey update a project\'s dev key by its unique ID. Use this
// endpoint to update a project\'s dev key name or expiration time.'
func (srv *Projects) UpdateDevKey(ProjectId string, KeyId string, Name string, Expire string)(*models.DevKey, error) {
	r := strings.NewReplacer("{projectId}", url.PathEscape(ProjectId), "{keyId}", url.PathEscape(KeyId))
	path := r.Replace("/projects/{projectId}/dev-keys/{keyId}")
	params := map[string]interface{}{}
	params["name"] = Name
	params["expire"] = Expire
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type": "application/json",
		"accept": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.DevKey{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DevKey
	parsed, ok := resp.Result.(models.DevKey)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// DeleteDevKey delete a project\'s dev key by its unique ID. Once deleted,
// the key will no longer allow bypassing of rate limits and better logging of
// errors.
func (srv *Projects) DeleteDevKey(ProjectId string, KeyId string)(*interface{}, error) {
	r := strings.NewReplacer("{projectId}", url.PathEscape(ProjectId), "{keyId}", url.PathEscape(KeyId))
	path := r.Replace("/projects/{projectId}/dev-keys/{keyId}")
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type": "application/json",
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
type ListSchedulesOptions struct {
	Queries []string
	Total bool
	enabledSetters map[string]bool
}
func (options ListSchedulesOptions) New() *ListSchedulesOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Total": false,
	}
	return &options
}
type ListSchedulesOption func(*ListSchedulesOptions)
func (srv *Projects) WithListSchedulesQueries(v []string) ListSchedulesOption {
	return func(o *ListSchedulesOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Projects) WithListSchedulesTotal(v bool) ListSchedulesOption {
	return func(o *ListSchedulesOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}
			
// ListSchedules get a list of all the project's schedules. You can use the
// query params to filter your results.
func (srv *Projects) ListSchedules(ProjectId string, optionalSetters ...ListSchedulesOption)(*models.ScheduleList, error) {
	r := strings.NewReplacer("{projectId}", url.PathEscape(ProjectId))
	path := r.Replace("/projects/{projectId}/schedules")
	options := ListSchedulesOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["Total"] {
		params["total"] = options.Total
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

		parsed := models.ScheduleList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ScheduleList
	parsed, ok := resp.Result.(models.ScheduleList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateScheduleOptions struct {
	Active bool
	Data interface{}
	enabledSetters map[string]bool
}
func (options CreateScheduleOptions) New() *CreateScheduleOptions {
	options.enabledSetters = map[string]bool{
		"Active": false,
		"Data": false,
	}
	return &options
}
type CreateScheduleOption func(*CreateScheduleOptions)
func (srv *Projects) WithCreateScheduleActive(v bool) CreateScheduleOption {
	return func(o *CreateScheduleOptions) {
		o.Active = v
		o.enabledSetters["Active"] = true
	}
}
func (srv *Projects) WithCreateScheduleData(v interface{}) CreateScheduleOption {
	return func(o *CreateScheduleOptions) {
		o.Data = v
		o.enabledSetters["Data"] = true
	}
}
									
// CreateSchedule create a new schedule for a resource.
func (srv *Projects) CreateSchedule(ProjectId string, ResourceType string, ResourceId string, Schedule string, optionalSetters ...CreateScheduleOption)(*models.Schedule, error) {
	r := strings.NewReplacer("{projectId}", url.PathEscape(ProjectId))
	path := r.Replace("/projects/{projectId}/schedules")
	options := CreateScheduleOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["resourceType"] = ResourceType
	params["resourceId"] = ResourceId
	params["schedule"] = Schedule
	if options.enabledSetters["Active"] {
		params["active"] = options.Active
	}
	if options.enabledSetters["Data"] {
		params["data"] = options.Data
	}
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type": "application/json",
		"accept": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Schedule{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Schedule
	parsed, ok := resp.Result.(models.Schedule)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// GetSchedule get a schedule by its unique ID.
func (srv *Projects) GetSchedule(ProjectId string, ScheduleId string)(*models.Schedule, error) {
	r := strings.NewReplacer("{projectId}", url.PathEscape(ProjectId), "{scheduleId}", url.PathEscape(ScheduleId))
	path := r.Replace("/projects/{projectId}/schedules/{scheduleId}")
	params := map[string]interface{}{}
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

		parsed := models.Schedule{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Schedule
	parsed, ok := resp.Result.(models.Schedule)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// ListStages get the onboarding stages for the current project, including
// each stage’s SDK method key and status (for example pending, completed,
// or skipped).
func (srv *Projects) ListStages(ProjectId string)(*models.StageList, error) {
	r := strings.NewReplacer("{projectId}", url.PathEscape(ProjectId))
	path := r.Replace("/projects/{projectId}/stages")
	params := map[string]interface{}{}
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

		parsed := models.StageList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.StageList
	parsed, ok := resp.Result.(models.StageList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateStageOptions struct {
	Skip bool
	enabledSetters map[string]bool
}
func (options UpdateStageOptions) New() *UpdateStageOptions {
	options.enabledSetters = map[string]bool{
		"Skip": false,
	}
	return &options
}
type UpdateStageOption func(*UpdateStageOptions)
func (srv *Projects) WithUpdateStageSkip(v bool) UpdateStageOption {
	return func(o *UpdateStageOptions) {
		o.Skip = v
		o.enabledSetters["Skip"] = true
	}
}
					
// UpdateStage update an onboarding stage for the current project. Use this
// endpoint to skip a stage or leave it unchanged without performing the
// related API action.
func (srv *Projects) UpdateStage(ProjectId string, StageId string, optionalSetters ...UpdateStageOption)(*models.Stage, error) {
	r := strings.NewReplacer("{projectId}", url.PathEscape(ProjectId), "{stageId}", url.PathEscape(StageId))
	path := r.Replace("/projects/{projectId}/stages/{stageId}")
	options := UpdateStageOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Skip"] {
		params["skip"] = options.Skip
	}
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

		parsed := models.Stage{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Stage
	parsed, ok := resp.Result.(models.Stage)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// UpdateStatus update the status of a project. Can be used to archive/restore
// projects, and to restore paused projects. When restoring a paused project,
// the console fingerprint header must be provided and the project must not be
// blocked for any reason other than inactivity.
func (srv *Projects) UpdateStatus(ProjectId string, Status string)(*interface{}, error) {
	r := strings.NewReplacer("{projectId}", url.PathEscape(ProjectId))
	path := r.Replace("/projects/{projectId}/status")
	params := map[string]interface{}{}
	params["status"] = Status
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
			
// UpdateTeam update the team ID of a project allowing for it to be
// transferred to another team.
func (srv *Projects) UpdateTeam(ProjectId string, TeamId string)(*models.Project, error) {
	r := strings.NewReplacer("{projectId}", url.PathEscape(ProjectId))
	path := r.Replace("/projects/{projectId}/team")
	params := map[string]interface{}{}
	params["teamId"] = TeamId
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

		parsed := models.Project{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Project
	parsed, ok := resp.Result.(models.Project)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
