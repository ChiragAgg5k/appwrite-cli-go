package console

import (
	"encoding/json"
	"errors"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/models"
	"net/url"
	"strings"
)

// Console service
type Console struct {
	client client.Client
}

func New(clt client.Client) *Console {
	return &Console{
		client: clt,
	}
}

	
// GetCampaign receive the details of a campaign using its ID.
func (srv *Console) GetCampaign(CampaignId string)(*models.Campaign, error) {
	r := strings.NewReplacer("{campaignId}", url.PathEscape(CampaignId))
	path := r.Replace("/console/campaigns/{campaignId}")
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

		parsed := models.Campaign{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Campaign
	parsed, ok := resp.Result.(models.Campaign)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// GetCoupon get the details of a coupon using it's coupon ID.
func (srv *Console) GetCoupon(CouponId string)(*models.Coupon, error) {
	r := strings.NewReplacer("{couponId}", url.PathEscape(CouponId))
	path := r.Replace("/console/coupons/{couponId}")
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

		parsed := models.Coupon{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Coupon
	parsed, ok := resp.Result.(models.Coupon)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type ListDatabasesOptions struct {
	Queries []string
	Search string
	Total bool
	enabledSetters map[string]bool
}
func (options ListDatabasesOptions) New() *ListDatabasesOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Search": false,
		"Total": false,
	}
	return &options
}
type ListDatabasesOption func(*ListDatabasesOptions)
func (srv *Console) WithListDatabasesQueries(v []string) ListDatabasesOption {
	return func(o *ListDatabasesOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Console) WithListDatabasesSearch(v string) ListDatabasesOption {
	return func(o *ListDatabasesOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}
func (srv *Console) WithListDatabasesTotal(v bool) ListDatabasesOption {
	return func(o *ListDatabasesOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}
	
// ListDatabases get a list of all the project's databases. You can use the
// query params to filter your results. This returns every database across all
// types and product APIs in a single call.
func (srv *Console) ListDatabases(optionalSetters ...ListDatabasesOption)(*models.DatabaseList, error) {
	path := "/console/databases"
	options := ListDatabasesOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["Search"] {
		params["search"] = options.Search
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

		parsed := models.DatabaseList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DatabaseList
	parsed, ok := resp.Result.(models.DatabaseList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// ListOAuth2Providers list all OAuth2 providers supported by the Appwrite
// server, along with the parameters required to configure each provider. The
// response excludes mock providers but includes sandbox providers.
func (srv *Console) ListOAuth2Providers()(*models.ConsoleOAuth2ProviderList, error) {
	path := "/console/oauth2-providers"
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

		parsed := models.ConsoleOAuth2ProviderList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ConsoleOAuth2ProviderList
	parsed, ok := resp.Result.(models.ConsoleOAuth2ProviderList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type GetPlansOptions struct {
	Platform string
	enabledSetters map[string]bool
}
func (options GetPlansOptions) New() *GetPlansOptions {
	options.enabledSetters = map[string]bool{
		"Platform": false,
	}
	return &options
}
type GetPlansOption func(*GetPlansOptions)
func (srv *Console) WithGetPlansPlatform(v string) GetPlansOption {
	return func(o *GetPlansOptions) {
		o.Platform = v
		o.enabledSetters["Platform"] = true
	}
}
	
// GetPlans return a list of all available plans.
func (srv *Console) GetPlans(optionalSetters ...GetPlansOption)(*models.BillingPlanList, error) {
	path := "/console/plans"
	options := GetPlansOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Platform"] {
		params["platform"] = options.Platform
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

		parsed := models.BillingPlanList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.BillingPlanList
	parsed, ok := resp.Result.(models.BillingPlanList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// GetPlan get the details of a plan using its plan ID.
func (srv *Console) GetPlan(PlanId string)(*models.BillingPlan, error) {
	r := strings.NewReplacer("{planId}", url.PathEscape(PlanId))
	path := r.Replace("/console/plans/{planId}")
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

		parsed := models.BillingPlan{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.BillingPlan
	parsed, ok := resp.Result.(models.BillingPlan)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// ListPostgresExtensions get the catalog of Postgres extensions that can be
// installed on a dedicated Postgres database.
func (srv *Console) ListPostgresExtensions()(*models.PostgresExtensionList, error) {
	path := "/console/postgres-extensions"
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

		parsed := models.PostgresExtensionList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.PostgresExtensionList
	parsed, ok := resp.Result.(models.PostgresExtensionList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// GetProgram receive the details of a program using its ID.
func (srv *Console) GetProgram(ProgramId string)(*models.Program, error) {
	r := strings.NewReplacer("{programId}", url.PathEscape(ProgramId))
	path := r.Replace("/console/programs/{programId}")
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

		parsed := models.Program{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Program
	parsed, ok := resp.Result.(models.Program)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// CreateProgramMembership create a new membership for an account to a
// program.
func (srv *Console) CreateProgramMembership(ProgramId string)(*models.Organization, error) {
	r := strings.NewReplacer("{programId}", url.PathEscape(ProgramId))
	path := r.Replace("/console/programs/{programId}/memberships")
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

		parsed := models.Organization{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Organization
	parsed, ok := resp.Result.(models.Organization)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// ListRegions get all available regions for the console.
func (srv *Console) ListRegions()(*models.ConsoleRegionList, error) {
	path := "/console/regions"
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

		parsed := models.ConsoleRegionList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ConsoleRegionList
	parsed, ok := resp.Result.(models.ConsoleRegionList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// GetResource check if a resource ID is available.
func (srv *Console) GetResource(Value string, Type string)(*interface{}, error) {
	path := "/console/resources"
	params := map[string]interface{}{}
	params["value"] = Value
	params["type"] = Type
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
	}

	resp, err := srv.client.Call("GET", path, headers, params)
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

// ListOrganizationScopes list all scopes available for organization API keys,
// along with a description for each scope.
func (srv *Console) ListOrganizationScopes()(*models.ConsoleKeyScopeList, error) {
	path := "/console/scopes/organization"
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

		parsed := models.ConsoleKeyScopeList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ConsoleKeyScopeList
	parsed, ok := resp.Result.(models.ConsoleKeyScopeList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// ListProjectScopes list all scopes available for project API keys, along
// with a description for each scope.
func (srv *Console) ListProjectScopes()(*models.ConsoleKeyScopeList, error) {
	path := "/console/scopes/project"
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

		parsed := models.ConsoleKeyScopeList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ConsoleKeyScopeList
	parsed, ok := resp.Result.(models.ConsoleKeyScopeList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateSourceOptions struct {
	Ref string
	Referrer string
	UtmSource string
	UtmCampaign string
	UtmMedium string
	enabledSetters map[string]bool
}
func (options CreateSourceOptions) New() *CreateSourceOptions {
	options.enabledSetters = map[string]bool{
		"Ref": false,
		"Referrer": false,
		"UtmSource": false,
		"UtmCampaign": false,
		"UtmMedium": false,
	}
	return &options
}
type CreateSourceOption func(*CreateSourceOptions)
func (srv *Console) WithCreateSourceRef(v string) CreateSourceOption {
	return func(o *CreateSourceOptions) {
		o.Ref = v
		o.enabledSetters["Ref"] = true
	}
}
func (srv *Console) WithCreateSourceReferrer(v string) CreateSourceOption {
	return func(o *CreateSourceOptions) {
		o.Referrer = v
		o.enabledSetters["Referrer"] = true
	}
}
func (srv *Console) WithCreateSourceUtmSource(v string) CreateSourceOption {
	return func(o *CreateSourceOptions) {
		o.UtmSource = v
		o.enabledSetters["UtmSource"] = true
	}
}
func (srv *Console) WithCreateSourceUtmCampaign(v string) CreateSourceOption {
	return func(o *CreateSourceOptions) {
		o.UtmCampaign = v
		o.enabledSetters["UtmCampaign"] = true
	}
}
func (srv *Console) WithCreateSourceUtmMedium(v string) CreateSourceOption {
	return func(o *CreateSourceOptions) {
		o.UtmMedium = v
		o.enabledSetters["UtmMedium"] = true
	}
}
	
// CreateSource create a new source.
func (srv *Console) CreateSource(optionalSetters ...CreateSourceOption)(*interface{}, error) {
	path := "/console/sources"
	options := CreateSourceOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Ref"] {
		params["ref"] = options.Ref
	}
	if options.enabledSetters["Referrer"] {
		params["referrer"] = options.Referrer
	}
	if options.enabledSetters["UtmSource"] {
		params["utmSource"] = options.UtmSource
	}
	if options.enabledSetters["UtmCampaign"] {
		params["utmCampaign"] = options.UtmCampaign
	}
	if options.enabledSetters["UtmMedium"] {
		params["utmMedium"] = options.UtmMedium
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
type SuggestColumnsOptions struct {
	Context string
	Min int
	Max int
	enabledSetters map[string]bool
}
func (options SuggestColumnsOptions) New() *SuggestColumnsOptions {
	options.enabledSetters = map[string]bool{
		"Context": false,
		"Min": false,
		"Max": false,
	}
	return &options
}
type SuggestColumnsOption func(*SuggestColumnsOptions)
func (srv *Console) WithSuggestColumnsContext(v string) SuggestColumnsOption {
	return func(o *SuggestColumnsOptions) {
		o.Context = v
		o.enabledSetters["Context"] = true
	}
}
func (srv *Console) WithSuggestColumnsMin(v int) SuggestColumnsOption {
	return func(o *SuggestColumnsOptions) {
		o.Min = v
		o.enabledSetters["Min"] = true
	}
}
func (srv *Console) WithSuggestColumnsMax(v int) SuggestColumnsOption {
	return func(o *SuggestColumnsOptions) {
		o.Max = v
		o.enabledSetters["Max"] = true
	}
}
					
// SuggestColumns suggests column names and their size limits based on the
// provided table name. The API will also analyze other tables in the same
// database to provide context-aware suggestions, ensuring consistency across
// schema design. Users may optionally provide custom context to further
// refine the suggestions.
func (srv *Console) SuggestColumns(DatabaseId string, TableId string, optionalSetters ...SuggestColumnsOption)(*models.ColumnList, error) {
	path := "/console/suggestions/columns"
	options := SuggestColumnsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["tableId"] = TableId
	if options.enabledSetters["Context"] {
		params["context"] = options.Context
	}
	if options.enabledSetters["Min"] {
		params["min"] = options.Min
	}
	if options.enabledSetters["Max"] {
		params["max"] = options.Max
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

		parsed := models.ColumnList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ColumnList
	parsed, ok := resp.Result.(models.ColumnList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type SuggestIndexesOptions struct {
	Min int
	Max int
	enabledSetters map[string]bool
}
func (options SuggestIndexesOptions) New() *SuggestIndexesOptions {
	options.enabledSetters = map[string]bool{
		"Min": false,
		"Max": false,
	}
	return &options
}
type SuggestIndexesOption func(*SuggestIndexesOptions)
func (srv *Console) WithSuggestIndexesMin(v int) SuggestIndexesOption {
	return func(o *SuggestIndexesOptions) {
		o.Min = v
		o.enabledSetters["Min"] = true
	}
}
func (srv *Console) WithSuggestIndexesMax(v int) SuggestIndexesOption {
	return func(o *SuggestIndexesOptions) {
		o.Max = v
		o.enabledSetters["Max"] = true
	}
}
					
// SuggestIndexes suggests database indexes for table columns based on the
// provided table structure and existing columns. The API will also analyze
// the table's column types, names, and patterns to recommend optimal indexes
// that improve query performance for common database operations like
// filtering, sorting, and searching.
func (srv *Console) SuggestIndexes(DatabaseId string, TableId string, optionalSetters ...SuggestIndexesOption)(*models.ColumnIndexList, error) {
	path := "/console/suggestions/indexes"
	options := SuggestIndexesOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["tableId"] = TableId
	if options.enabledSetters["Min"] {
		params["min"] = options.Min
	}
	if options.enabledSetters["Max"] {
		params["max"] = options.Max
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

		parsed := models.ColumnIndexList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ColumnIndexList
	parsed, ok := resp.Result.(models.ColumnIndexList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type SuggestQueriesOptions struct {
	DatabaseId string
	TableId string
	enabledSetters map[string]bool
}
func (options SuggestQueriesOptions) New() *SuggestQueriesOptions {
	options.enabledSetters = map[string]bool{
		"DatabaseId": false,
		"TableId": false,
	}
	return &options
}
type SuggestQueriesOption func(*SuggestQueriesOptions)
func (srv *Console) WithSuggestQueriesDatabaseId(v string) SuggestQueriesOption {
	return func(o *SuggestQueriesOptions) {
		o.DatabaseId = v
		o.enabledSetters["DatabaseId"] = true
	}
}
func (srv *Console) WithSuggestQueriesTableId(v string) SuggestQueriesOption {
	return func(o *SuggestQueriesOptions) {
		o.TableId = v
		o.enabledSetters["TableId"] = true
	}
}
					
// SuggestQueries suggest valid Appwrite query JSON objects for a supported
// list resource from free-text user intent. The endpoint picks a validator
// based on `resource` — for system resources it uses the static validator
// and its allowed attributes, and for user-owned table rows it loads the
// table schema and validates against those attributes at request time. The
// returned queries are guaranteed to parse and pass the relevant queries
// validator.
func (srv *Console) SuggestQueries(Resource string, Input string, optionalSetters ...SuggestQueriesOption)(*interface{}, error) {
	path := "/console/suggestions/queries"
	options := SuggestQueriesOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["resource"] = Resource
	params["input"] = Input
	if options.enabledSetters["DatabaseId"] {
		params["databaseId"] = options.DatabaseId
	}
	if options.enabledSetters["TableId"] {
		params["tableId"] = options.TableId
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
type GetEmailTemplateOptions struct {
	Locale string
	enabledSetters map[string]bool
}
func (options GetEmailTemplateOptions) New() *GetEmailTemplateOptions {
	options.enabledSetters = map[string]bool{
		"Locale": false,
	}
	return &options
}
type GetEmailTemplateOption func(*GetEmailTemplateOptions)
func (srv *Console) WithGetEmailTemplateLocale(v string) GetEmailTemplateOption {
	return func(o *GetEmailTemplateOptions) {
		o.Locale = v
		o.enabledSetters["Locale"] = true
	}
}
			
// GetEmailTemplate get the Appwrite built-in default email template for the
// specified type and locale. Always returns the unmodified default, ignoring
// any custom project overrides.
func (srv *Console) GetEmailTemplate(TemplateId string, optionalSetters ...GetEmailTemplateOption)(*models.EmailTemplate, error) {
	r := strings.NewReplacer("{templateId}", url.PathEscape(TemplateId))
	path := r.Replace("/console/templates/email/{templateId}")
	options := GetEmailTemplateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Locale"] {
		params["locale"] = options.Locale
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

		parsed := models.EmailTemplate{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.EmailTemplate
	parsed, ok := resp.Result.(models.EmailTemplate)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// Variables get all Environment Variables that are relevant for the console.
func (srv *Console) Variables()(*models.ConsoleVariables, error) {
	path := "/console/variables"
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

		parsed := models.ConsoleVariables{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ConsoleVariables
	parsed, ok := resp.Result.(models.ConsoleVariables)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
