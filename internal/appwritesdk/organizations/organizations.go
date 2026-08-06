package organizations

import (
	"encoding/json"
	"errors"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/models"
	"net/url"
	"strings"
)

// Organizations service
type Organizations struct {
	client client.Client
}

func New(clt client.Client) *Organizations {
	return &Organizations{
		client: clt,
	}
}

type ListOptions struct {
	Queries        []string
	Search         string
	enabledSetters map[string]bool
}

func (options ListOptions) New() *ListOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Search":  false,
	}
	return &options
}

type ListOption func(*ListOptions)

func (srv *Organizations) WithListQueries(v []string) ListOption {
	return func(o *ListOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Organizations) WithListSearch(v string) ListOption {
	return func(o *ListOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}

// List get a list of all the teams in which the current user is a member. You
// can use the parameters to filter your results.
func (srv *Organizations) List(optionalSetters ...ListOption) (*models.OrganizationList, error) {
	path := "/organizations"
	options := ListOptions{}.New()
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
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"accept":             "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.OrganizationList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OrganizationList
	parsed, ok := resp.Result.(models.OrganizationList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateOptions struct {
	PaymentMethodId  string
	BillingAddressId string
	Invites          []string
	CouponId         string
	TaxId            string
	Budget           int
	Platform         string
	enabledSetters   map[string]bool
}

func (options CreateOptions) New() *CreateOptions {
	options.enabledSetters = map[string]bool{
		"PaymentMethodId":  false,
		"BillingAddressId": false,
		"Invites":          false,
		"CouponId":         false,
		"TaxId":            false,
		"Budget":           false,
		"Platform":         false,
	}
	return &options
}

type CreateOption func(*CreateOptions)

func (srv *Organizations) WithCreatePaymentMethodId(v string) CreateOption {
	return func(o *CreateOptions) {
		o.PaymentMethodId = v
		o.enabledSetters["PaymentMethodId"] = true
	}
}
func (srv *Organizations) WithCreateBillingAddressId(v string) CreateOption {
	return func(o *CreateOptions) {
		o.BillingAddressId = v
		o.enabledSetters["BillingAddressId"] = true
	}
}
func (srv *Organizations) WithCreateInvites(v []string) CreateOption {
	return func(o *CreateOptions) {
		o.Invites = v
		o.enabledSetters["Invites"] = true
	}
}
func (srv *Organizations) WithCreateCouponId(v string) CreateOption {
	return func(o *CreateOptions) {
		o.CouponId = v
		o.enabledSetters["CouponId"] = true
	}
}
func (srv *Organizations) WithCreateTaxId(v string) CreateOption {
	return func(o *CreateOptions) {
		o.TaxId = v
		o.enabledSetters["TaxId"] = true
	}
}
func (srv *Organizations) WithCreateBudget(v int) CreateOption {
	return func(o *CreateOptions) {
		o.Budget = v
		o.enabledSetters["Budget"] = true
	}
}
func (srv *Organizations) WithCreatePlatform(v string) CreateOption {
	return func(o *CreateOptions) {
		o.Platform = v
		o.enabledSetters["Platform"] = true
	}
}

// Create create a new organization.
func (srv *Organizations) Create(OrganizationId string, Name string, BillingPlan string, optionalSetters ...CreateOption) (models.Model, error) {
	path := "/organizations"
	options := CreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["organizationId"] = OrganizationId
	params["name"] = Name
	params["billingPlan"] = BillingPlan
	if options.enabledSetters["PaymentMethodId"] {
		params["paymentMethodId"] = options.PaymentMethodId
	}
	if options.enabledSetters["BillingAddressId"] {
		params["billingAddressId"] = options.BillingAddressId
	}
	if options.enabledSetters["Invites"] {
		params["invites"] = options.Invites
	}
	if options.enabledSetters["CouponId"] {
		params["couponId"] = options.CouponId
	}
	if options.enabledSetters["TaxId"] {
		params["taxId"] = options.TaxId
	}
	if options.enabledSetters["Budget"] {
		params["budget"] = options.Budget
	}
	if options.enabledSetters["Platform"] {
		params["platform"] = options.Platform
	}
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type":       "application/json",
		"accept":             "application/json",
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
	parsed, ok := resp.Result.(models.Model)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return parsed, nil

}

type EstimationCreateOrganizationOptions struct {
	PaymentMethodId string
	Invites         []string
	CouponId        string
	Platform        string
	enabledSetters  map[string]bool
}

func (options EstimationCreateOrganizationOptions) New() *EstimationCreateOrganizationOptions {
	options.enabledSetters = map[string]bool{
		"PaymentMethodId": false,
		"Invites":         false,
		"CouponId":        false,
		"Platform":        false,
	}
	return &options
}

type EstimationCreateOrganizationOption func(*EstimationCreateOrganizationOptions)

func (srv *Organizations) WithEstimationCreateOrganizationPaymentMethodId(v string) EstimationCreateOrganizationOption {
	return func(o *EstimationCreateOrganizationOptions) {
		o.PaymentMethodId = v
		o.enabledSetters["PaymentMethodId"] = true
	}
}
func (srv *Organizations) WithEstimationCreateOrganizationInvites(v []string) EstimationCreateOrganizationOption {
	return func(o *EstimationCreateOrganizationOptions) {
		o.Invites = v
		o.enabledSetters["Invites"] = true
	}
}
func (srv *Organizations) WithEstimationCreateOrganizationCouponId(v string) EstimationCreateOrganizationOption {
	return func(o *EstimationCreateOrganizationOptions) {
		o.CouponId = v
		o.enabledSetters["CouponId"] = true
	}
}
func (srv *Organizations) WithEstimationCreateOrganizationPlatform(v string) EstimationCreateOrganizationOption {
	return func(o *EstimationCreateOrganizationOptions) {
		o.Platform = v
		o.enabledSetters["Platform"] = true
	}
}

// EstimationCreateOrganization get estimation for creating an organization.
func (srv *Organizations) EstimationCreateOrganization(BillingPlan string, optionalSetters ...EstimationCreateOrganizationOption) (*models.Estimation, error) {
	path := "/organizations/estimations/create-organization"
	options := EstimationCreateOrganizationOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["billingPlan"] = BillingPlan
	if options.enabledSetters["PaymentMethodId"] {
		params["paymentMethodId"] = options.PaymentMethodId
	}
	if options.enabledSetters["Invites"] {
		params["invites"] = options.Invites
	}
	if options.enabledSetters["CouponId"] {
		params["couponId"] = options.CouponId
	}
	if options.enabledSetters["Platform"] {
		params["platform"] = options.Platform
	}
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type":       "application/json",
		"accept":             "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Estimation{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Estimation
	parsed, ok := resp.Result.(models.Estimation)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// Delete delete an organization.
func (srv *Organizations) Delete(OrganizationId string) (*interface{}, error) {
	r := strings.NewReplacer("{organizationId}", OrganizationId)
	path := r.Replace("/organizations/{organizationId}")
	params := map[string]interface{}{}
	params["organizationId"] = OrganizationId
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type":       "application/json",
		"accept":             "application/json",
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

// ListAddons list all billing addons for an organization.
func (srv *Organizations) ListAddons(OrganizationId string) (*models.AddonList, error) {
	r := strings.NewReplacer("{organizationId}", OrganizationId)
	path := r.Replace("/organizations/{organizationId}/addons")
	params := map[string]interface{}{}
	params["organizationId"] = OrganizationId
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"accept":             "application/json",
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

// CreateBaaAddon create the BAA billing addon for an organization.
func (srv *Organizations) CreateBaaAddon(OrganizationId string) (*models.Addon, error) {
	r := strings.NewReplacer("{organizationId}", OrganizationId)
	path := r.Replace("/organizations/{organizationId}/addons/baa")
	params := map[string]interface{}{}
	params["organizationId"] = OrganizationId
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type":       "application/json",
		"accept":             "application/json",
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

// CreatePremiumGeoDBAddon create a Premium Geo DB addon for an organization.
func (srv *Organizations) CreatePremiumGeoDBAddon(OrganizationId string) (*models.Addon, error) {
	r := strings.NewReplacer("{organizationId}", OrganizationId)
	path := r.Replace("/organizations/{organizationId}/addons/premium-geo-db")
	params := map[string]interface{}{}
	params["organizationId"] = OrganizationId
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type":       "application/json",
		"accept":             "application/json",
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

// GetAddon get the details of a billing addon for an organization.
func (srv *Organizations) GetAddon(OrganizationId string, AddonId string) (*models.Addon, error) {
	r := strings.NewReplacer("{organizationId}", OrganizationId, "{addonId}", AddonId)
	path := r.Replace("/organizations/{organizationId}/addons/{addonId}")
	params := map[string]interface{}{}
	params["organizationId"] = OrganizationId
	params["addonId"] = AddonId
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"accept":             "application/json",
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

// DeleteAddon delete a billing addon for an organization.
func (srv *Organizations) DeleteAddon(OrganizationId string, AddonId string) (*interface{}, error) {
	r := strings.NewReplacer("{organizationId}", OrganizationId, "{addonId}", AddonId)
	path := r.Replace("/organizations/{organizationId}/addons/{addonId}")
	params := map[string]interface{}{}
	params["organizationId"] = OrganizationId
	params["addonId"] = AddonId
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type":       "application/json",
		"accept":             "application/json",
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

// ConfirmAddonPayment confirm payment for a billing addon for an
// organization.
func (srv *Organizations) ConfirmAddonPayment(OrganizationId string, AddonId string) (*models.Addon, error) {
	r := strings.NewReplacer("{organizationId}", OrganizationId, "{addonId}", AddonId)
	path := r.Replace("/organizations/{organizationId}/addons/{addonId}/confirmations")
	params := map[string]interface{}{}
	params["organizationId"] = OrganizationId
	params["addonId"] = AddonId
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type":       "application/json",
		"accept":             "application/json",
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

// GetAddonPrice get the price details for a billing addon for an
// organization.
func (srv *Organizations) GetAddonPrice(OrganizationId string, Addon string) (*models.AddonPrice, error) {
	r := strings.NewReplacer("{organizationId}", OrganizationId, "{addon}", Addon)
	path := r.Replace("/organizations/{organizationId}/addons/{addon}/price")
	params := map[string]interface{}{}
	params["organizationId"] = OrganizationId
	params["addon"] = Addon
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"accept":             "application/json",
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

type ListAggregationsOptions struct {
	Queries        []string
	enabledSetters map[string]bool
}

func (options ListAggregationsOptions) New() *ListAggregationsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
	}
	return &options
}

type ListAggregationsOption func(*ListAggregationsOptions)

func (srv *Organizations) WithListAggregationsQueries(v []string) ListAggregationsOption {
	return func(o *ListAggregationsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}

// ListAggregations get a list of all aggregations for an organization.
func (srv *Organizations) ListAggregations(OrganizationId string, optionalSetters ...ListAggregationsOption) (*models.AggregationTeamList, error) {
	r := strings.NewReplacer("{organizationId}", OrganizationId)
	path := r.Replace("/organizations/{organizationId}/aggregations")
	options := ListAggregationsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["organizationId"] = OrganizationId
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"accept":             "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.AggregationTeamList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AggregationTeamList
	parsed, ok := resp.Result.(models.AggregationTeamList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type GetAggregationOptions struct {
	Limit          int
	Offset         int
	enabledSetters map[string]bool
}

func (options GetAggregationOptions) New() *GetAggregationOptions {
	options.enabledSetters = map[string]bool{
		"Limit":  false,
		"Offset": false,
	}
	return &options
}

type GetAggregationOption func(*GetAggregationOptions)

func (srv *Organizations) WithGetAggregationLimit(v int) GetAggregationOption {
	return func(o *GetAggregationOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *Organizations) WithGetAggregationOffset(v int) GetAggregationOption {
	return func(o *GetAggregationOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}

// GetAggregation get a specific aggregation using it's aggregation ID.
func (srv *Organizations) GetAggregation(OrganizationId string, AggregationId string, optionalSetters ...GetAggregationOption) (*models.AggregationTeam, error) {
	r := strings.NewReplacer("{organizationId}", OrganizationId, "{aggregationId}", AggregationId)
	path := r.Replace("/organizations/{organizationId}/aggregations/{aggregationId}")
	options := GetAggregationOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["organizationId"] = OrganizationId
	params["aggregationId"] = AggregationId
	if options.enabledSetters["Limit"] {
		params["limit"] = options.Limit
	}
	if options.enabledSetters["Offset"] {
		params["offset"] = options.Offset
	}
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"accept":             "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.AggregationTeam{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AggregationTeam
	parsed, ok := resp.Result.(models.AggregationTeam)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// SetBillingAddress set a billing address for an organization.
func (srv *Organizations) SetBillingAddress(OrganizationId string, BillingAddressId string) (*models.Organization, error) {
	r := strings.NewReplacer("{organizationId}", OrganizationId)
	path := r.Replace("/organizations/{organizationId}/billing-address")
	params := map[string]interface{}{}
	params["organizationId"] = OrganizationId
	params["billingAddressId"] = BillingAddressId
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type":       "application/json",
		"accept":             "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
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

// DeleteBillingAddress delete a team's billing address.
func (srv *Organizations) DeleteBillingAddress(OrganizationId string) (*interface{}, error) {
	r := strings.NewReplacer("{organizationId}", OrganizationId)
	path := r.Replace("/organizations/{organizationId}/billing-address")
	params := map[string]interface{}{}
	params["organizationId"] = OrganizationId
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type":       "application/json",
		"accept":             "application/json",
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

// GetBillingAddress get a billing address using it's ID.
func (srv *Organizations) GetBillingAddress(OrganizationId string, BillingAddressId string) (*models.BillingAddress, error) {
	r := strings.NewReplacer("{organizationId}", OrganizationId, "{billingAddressId}", BillingAddressId)
	path := r.Replace("/organizations/{organizationId}/billing-addresses/{billingAddressId}")
	params := map[string]interface{}{}
	params["organizationId"] = OrganizationId
	params["billingAddressId"] = BillingAddressId
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"accept":             "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.BillingAddress{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.BillingAddress
	parsed, ok := resp.Result.(models.BillingAddress)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// SetBillingEmail set the current billing email for the organization.
func (srv *Organizations) SetBillingEmail(OrganizationId string, BillingEmail string) (*models.Organization, error) {
	r := strings.NewReplacer("{organizationId}", OrganizationId)
	path := r.Replace("/organizations/{organizationId}/billing-email")
	params := map[string]interface{}{}
	params["organizationId"] = OrganizationId
	params["billingEmail"] = BillingEmail
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type":       "application/json",
		"accept":             "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
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

type UpdateBudgetOptions struct {
	Alerts         []int
	enabledSetters map[string]bool
}

func (options UpdateBudgetOptions) New() *UpdateBudgetOptions {
	options.enabledSetters = map[string]bool{
		"Alerts": false,
	}
	return &options
}

type UpdateBudgetOption func(*UpdateBudgetOptions)

func (srv *Organizations) WithUpdateBudgetAlerts(v []int) UpdateBudgetOption {
	return func(o *UpdateBudgetOptions) {
		o.Alerts = v
		o.enabledSetters["Alerts"] = true
	}
}

// UpdateBudget update the budget limit for an organization.
func (srv *Organizations) UpdateBudget(OrganizationId string, Budget int, optionalSetters ...UpdateBudgetOption) (*models.Organization, error) {
	r := strings.NewReplacer("{organizationId}", OrganizationId)
	path := r.Replace("/organizations/{organizationId}/budget")
	options := UpdateBudgetOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["organizationId"] = OrganizationId
	params["budget"] = Budget
	if options.enabledSetters["Alerts"] {
		params["alerts"] = options.Alerts
	}
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type":       "application/json",
		"accept":             "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
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

type ListCreditsOptions struct {
	Queries        []string
	enabledSetters map[string]bool
}

func (options ListCreditsOptions) New() *ListCreditsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
	}
	return &options
}

type ListCreditsOption func(*ListCreditsOptions)

func (srv *Organizations) WithListCreditsQueries(v []string) ListCreditsOption {
	return func(o *ListCreditsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}

// ListCredits list all credits for an organization.
func (srv *Organizations) ListCredits(OrganizationId string, optionalSetters ...ListCreditsOption) (*models.CreditList, error) {
	r := strings.NewReplacer("{organizationId}", OrganizationId)
	path := r.Replace("/organizations/{organizationId}/credits")
	options := ListCreditsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["organizationId"] = OrganizationId
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"accept":             "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.CreditList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.CreditList
	parsed, ok := resp.Result.(models.CreditList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// AddCredit add credit to an organization using a coupon.
func (srv *Organizations) AddCredit(OrganizationId string, CouponId string) (*models.Credit, error) {
	r := strings.NewReplacer("{organizationId}", OrganizationId)
	path := r.Replace("/organizations/{organizationId}/credits")
	params := map[string]interface{}{}
	params["organizationId"] = OrganizationId
	params["couponId"] = CouponId
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type":       "application/json",
		"accept":             "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Credit{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Credit
	parsed, ok := resp.Result.(models.Credit)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetAvailableCredits get total available valid credits for an organization.
func (srv *Organizations) GetAvailableCredits(OrganizationId string) (*models.CreditAvailable, error) {
	r := strings.NewReplacer("{organizationId}", OrganizationId)
	path := r.Replace("/organizations/{organizationId}/credits/available")
	params := map[string]interface{}{}
	params["organizationId"] = OrganizationId
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"accept":             "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.CreditAvailable{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.CreditAvailable
	parsed, ok := resp.Result.(models.CreditAvailable)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetCredit get credit details.
func (srv *Organizations) GetCredit(OrganizationId string, CreditId string) (*models.Credit, error) {
	r := strings.NewReplacer("{organizationId}", OrganizationId, "{creditId}", CreditId)
	path := r.Replace("/organizations/{organizationId}/credits/{creditId}")
	params := map[string]interface{}{}
	params["organizationId"] = OrganizationId
	params["creditId"] = CreditId
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"accept":             "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Credit{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Credit
	parsed, ok := resp.Result.(models.Credit)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// EstimationDeleteOrganization get estimation for deleting an organization.
func (srv *Organizations) EstimationDeleteOrganization(OrganizationId string) (*models.EstimationDeleteOrganization, error) {
	r := strings.NewReplacer("{organizationId}", OrganizationId)
	path := r.Replace("/organizations/{organizationId}/estimations/delete-organization")
	params := map[string]interface{}{}
	params["organizationId"] = OrganizationId
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type":       "application/json",
		"accept":             "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.EstimationDeleteOrganization{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.EstimationDeleteOrganization
	parsed, ok := resp.Result.(models.EstimationDeleteOrganization)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type EstimationUpdatePlanOptions struct {
	Invites        []string
	CouponId       string
	enabledSetters map[string]bool
}

func (options EstimationUpdatePlanOptions) New() *EstimationUpdatePlanOptions {
	options.enabledSetters = map[string]bool{
		"Invites":  false,
		"CouponId": false,
	}
	return &options
}

type EstimationUpdatePlanOption func(*EstimationUpdatePlanOptions)

func (srv *Organizations) WithEstimationUpdatePlanInvites(v []string) EstimationUpdatePlanOption {
	return func(o *EstimationUpdatePlanOptions) {
		o.Invites = v
		o.enabledSetters["Invites"] = true
	}
}
func (srv *Organizations) WithEstimationUpdatePlanCouponId(v string) EstimationUpdatePlanOption {
	return func(o *EstimationUpdatePlanOptions) {
		o.CouponId = v
		o.enabledSetters["CouponId"] = true
	}
}

// EstimationUpdatePlan get estimation for updating the organization plan.
func (srv *Organizations) EstimationUpdatePlan(OrganizationId string, BillingPlan string, optionalSetters ...EstimationUpdatePlanOption) (*models.EstimationUpdatePlan, error) {
	r := strings.NewReplacer("{organizationId}", OrganizationId)
	path := r.Replace("/organizations/{organizationId}/estimations/update-plan")
	options := EstimationUpdatePlanOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["organizationId"] = OrganizationId
	params["billingPlan"] = BillingPlan
	if options.enabledSetters["Invites"] {
		params["invites"] = options.Invites
	}
	if options.enabledSetters["CouponId"] {
		params["couponId"] = options.CouponId
	}
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type":       "application/json",
		"accept":             "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.EstimationUpdatePlan{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.EstimationUpdatePlan
	parsed, ok := resp.Result.(models.EstimationUpdatePlan)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// CreateDowngradeFeedback submit feedback about downgrading from a paid plan
// to a lower tier. This helps the team understand user experience and improve
// the platform.
func (srv *Organizations) CreateDowngradeFeedback(OrganizationId string, Reason string, Message string, FromPlanId string, ToPlanId string) (*models.DowngradeFeedback, error) {
	r := strings.NewReplacer("{organizationId}", OrganizationId)
	path := r.Replace("/organizations/{organizationId}/feedbacks/downgrade")
	params := map[string]interface{}{}
	params["organizationId"] = OrganizationId
	params["reason"] = Reason
	params["message"] = Message
	params["fromPlanId"] = FromPlanId
	params["toPlanId"] = ToPlanId
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type":       "application/json",
		"accept":             "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.DowngradeFeedback{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DowngradeFeedback
	parsed, ok := resp.Result.(models.DowngradeFeedback)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type ListInvoicesOptions struct {
	Queries        []string
	enabledSetters map[string]bool
}

func (options ListInvoicesOptions) New() *ListInvoicesOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
	}
	return &options
}

type ListInvoicesOption func(*ListInvoicesOptions)

func (srv *Organizations) WithListInvoicesQueries(v []string) ListInvoicesOption {
	return func(o *ListInvoicesOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}

// ListInvoices list all invoices for an organization.
func (srv *Organizations) ListInvoices(OrganizationId string, optionalSetters ...ListInvoicesOption) (*models.InvoiceList, error) {
	r := strings.NewReplacer("{organizationId}", OrganizationId)
	path := r.Replace("/organizations/{organizationId}/invoices")
	options := ListInvoicesOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["organizationId"] = OrganizationId
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"accept":             "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.InvoiceList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.InvoiceList
	parsed, ok := resp.Result.(models.InvoiceList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetInvoice get an invoice by its unique ID.
func (srv *Organizations) GetInvoice(OrganizationId string, InvoiceId string) (*models.Invoice, error) {
	r := strings.NewReplacer("{organizationId}", OrganizationId, "{invoiceId}", InvoiceId)
	path := r.Replace("/organizations/{organizationId}/invoices/{invoiceId}")
	params := map[string]interface{}{}
	params["organizationId"] = OrganizationId
	params["invoiceId"] = InvoiceId
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"accept":             "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Invoice{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Invoice
	parsed, ok := resp.Result.(models.Invoice)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetInvoiceDownload download invoice in PDF
func (srv *Organizations) GetInvoiceDownload(OrganizationId string, InvoiceId string) (*[]byte, error) {
	r := strings.NewReplacer("{organizationId}", OrganizationId, "{invoiceId}", InvoiceId)
	path := r.Replace("/organizations/{organizationId}/invoices/{invoiceId}/download")
	params := map[string]interface{}{}
	params["organizationId"] = OrganizationId
	params["invoiceId"] = InvoiceId
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"accept":             "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		var parsed []byte

		err = json.Unmarshal(bytes, &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	var parsed []byte
	parsed, ok := resp.Result.([]byte)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetInvoiceDownloadURL download invoice in PDF
// Returns the URL for the resource instead of the content.
func (srv *Organizations) GetInvoiceDownloadURL(OrganizationId string, InvoiceId string) (*string, error) {
	r := strings.NewReplacer("{organizationId}", OrganizationId, "{invoiceId}", InvoiceId)
	path := r.Replace("/organizations/{organizationId}/invoices/{invoiceId}/download")
	u, err := url.Parse(srv.client.Endpoint + path)
	if err != nil {
		return nil, err
	}
	q := u.Query()
	u.RawQuery = q.Encode()
	result := u.String()
	return &result, nil
}

// CreateInvoicePayment initiate payment for failed invoice to pay live from
// console
func (srv *Organizations) CreateInvoicePayment(OrganizationId string, InvoiceId string, PaymentMethodId string) (*models.Invoice, error) {
	r := strings.NewReplacer("{organizationId}", OrganizationId, "{invoiceId}", InvoiceId)
	path := r.Replace("/organizations/{organizationId}/invoices/{invoiceId}/payments")
	params := map[string]interface{}{}
	params["organizationId"] = OrganizationId
	params["invoiceId"] = InvoiceId
	params["paymentMethodId"] = PaymentMethodId
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type":       "application/json",
		"accept":             "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Invoice{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Invoice
	parsed, ok := resp.Result.(models.Invoice)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// ValidateInvoice validates the payment linked with the invoice and updates
// the invoice status if the payment status is changed.
func (srv *Organizations) ValidateInvoice(OrganizationId string, InvoiceId string) (*models.Invoice, error) {
	r := strings.NewReplacer("{organizationId}", OrganizationId, "{invoiceId}", InvoiceId)
	path := r.Replace("/organizations/{organizationId}/invoices/{invoiceId}/status")
	params := map[string]interface{}{}
	params["organizationId"] = OrganizationId
	params["invoiceId"] = InvoiceId
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type":       "application/json",
		"accept":             "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Invoice{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Invoice
	parsed, ok := resp.Result.(models.Invoice)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetInvoiceView view invoice in PDF
func (srv *Organizations) GetInvoiceView(OrganizationId string, InvoiceId string) (*[]byte, error) {
	r := strings.NewReplacer("{organizationId}", OrganizationId, "{invoiceId}", InvoiceId)
	path := r.Replace("/organizations/{organizationId}/invoices/{invoiceId}/view")
	params := map[string]interface{}{}
	params["organizationId"] = OrganizationId
	params["invoiceId"] = InvoiceId
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"accept":             "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		var parsed []byte

		err = json.Unmarshal(bytes, &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	var parsed []byte
	parsed, ok := resp.Result.([]byte)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetInvoiceViewURL view invoice in PDF
// Returns the URL for the resource instead of the content.
func (srv *Organizations) GetInvoiceViewURL(OrganizationId string, InvoiceId string) (*string, error) {
	r := strings.NewReplacer("{organizationId}", OrganizationId, "{invoiceId}", InvoiceId)
	path := r.Replace("/organizations/{organizationId}/invoices/{invoiceId}/view")
	u, err := url.Parse(srv.client.Endpoint + path)
	if err != nil {
		return nil, err
	}
	q := u.Query()
	u.RawQuery = q.Encode()
	result := u.String()
	return &result, nil
}

// SetDefaultPaymentMethod set a organization's default payment method.
func (srv *Organizations) SetDefaultPaymentMethod(OrganizationId string, PaymentMethodId string) (*models.Organization, error) {
	r := strings.NewReplacer("{organizationId}", OrganizationId)
	path := r.Replace("/organizations/{organizationId}/payment-method")
	params := map[string]interface{}{}
	params["organizationId"] = OrganizationId
	params["paymentMethodId"] = PaymentMethodId
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type":       "application/json",
		"accept":             "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
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

// DeleteDefaultPaymentMethod delete the default payment method for an
// organization.
func (srv *Organizations) DeleteDefaultPaymentMethod(OrganizationId string) (*models.Organization, error) {
	r := strings.NewReplacer("{organizationId}", OrganizationId)
	path := r.Replace("/organizations/{organizationId}/payment-method")
	params := map[string]interface{}{}
	params["organizationId"] = OrganizationId
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type":       "application/json",
		"accept":             "application/json",
	}

	resp, err := srv.client.Call("DELETE", path, headers, params)
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

// SetBackupPaymentMethod set an organization's backup payment method.
func (srv *Organizations) SetBackupPaymentMethod(OrganizationId string, PaymentMethodId string) (*models.Organization, error) {
	r := strings.NewReplacer("{organizationId}", OrganizationId)
	path := r.Replace("/organizations/{organizationId}/payment-method/backup")
	params := map[string]interface{}{}
	params["organizationId"] = OrganizationId
	params["paymentMethodId"] = PaymentMethodId
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type":       "application/json",
		"accept":             "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
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

// DeleteBackupPaymentMethod delete a backup payment method for an
// organization.
func (srv *Organizations) DeleteBackupPaymentMethod(OrganizationId string) (*models.Organization, error) {
	r := strings.NewReplacer("{organizationId}", OrganizationId)
	path := r.Replace("/organizations/{organizationId}/payment-method/backup")
	params := map[string]interface{}{}
	params["organizationId"] = OrganizationId
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type":       "application/json",
		"accept":             "application/json",
	}

	resp, err := srv.client.Call("DELETE", path, headers, params)
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

// GetPaymentMethod get an organization's payment method using it's payment
// method ID.
func (srv *Organizations) GetPaymentMethod(OrganizationId string, PaymentMethodId string) (*models.PaymentMethod, error) {
	r := strings.NewReplacer("{organizationId}", OrganizationId, "{paymentMethodId}", PaymentMethodId)
	path := r.Replace("/organizations/{organizationId}/payment-methods/{paymentMethodId}")
	params := map[string]interface{}{}
	params["organizationId"] = OrganizationId
	params["paymentMethodId"] = PaymentMethodId
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"accept":             "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.PaymentMethod{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.PaymentMethod
	parsed, ok := resp.Result.(models.PaymentMethod)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetPlan get the details of the current billing plan for an organization.
func (srv *Organizations) GetPlan(OrganizationId string) (*models.BillingPlan, error) {
	r := strings.NewReplacer("{organizationId}", OrganizationId)
	path := r.Replace("/organizations/{organizationId}/plan")
	params := map[string]interface{}{}
	params["organizationId"] = OrganizationId
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"accept":             "application/json",
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

type UpdatePlanOptions struct {
	PaymentMethodId  string
	BillingAddressId string
	Invites          []string
	CouponId         string
	TaxId            string
	Budget           int
	enabledSetters   map[string]bool
}

func (options UpdatePlanOptions) New() *UpdatePlanOptions {
	options.enabledSetters = map[string]bool{
		"PaymentMethodId":  false,
		"BillingAddressId": false,
		"Invites":          false,
		"CouponId":         false,
		"TaxId":            false,
		"Budget":           false,
	}
	return &options
}

type UpdatePlanOption func(*UpdatePlanOptions)

func (srv *Organizations) WithUpdatePlanPaymentMethodId(v string) UpdatePlanOption {
	return func(o *UpdatePlanOptions) {
		o.PaymentMethodId = v
		o.enabledSetters["PaymentMethodId"] = true
	}
}
func (srv *Organizations) WithUpdatePlanBillingAddressId(v string) UpdatePlanOption {
	return func(o *UpdatePlanOptions) {
		o.BillingAddressId = v
		o.enabledSetters["BillingAddressId"] = true
	}
}
func (srv *Organizations) WithUpdatePlanInvites(v []string) UpdatePlanOption {
	return func(o *UpdatePlanOptions) {
		o.Invites = v
		o.enabledSetters["Invites"] = true
	}
}
func (srv *Organizations) WithUpdatePlanCouponId(v string) UpdatePlanOption {
	return func(o *UpdatePlanOptions) {
		o.CouponId = v
		o.enabledSetters["CouponId"] = true
	}
}
func (srv *Organizations) WithUpdatePlanTaxId(v string) UpdatePlanOption {
	return func(o *UpdatePlanOptions) {
		o.TaxId = v
		o.enabledSetters["TaxId"] = true
	}
}
func (srv *Organizations) WithUpdatePlanBudget(v int) UpdatePlanOption {
	return func(o *UpdatePlanOptions) {
		o.Budget = v
		o.enabledSetters["Budget"] = true
	}
}

// UpdatePlan update the billing plan for an organization.
func (srv *Organizations) UpdatePlan(OrganizationId string, BillingPlan string, optionalSetters ...UpdatePlanOption) (*models.Organization, error) {
	r := strings.NewReplacer("{organizationId}", OrganizationId)
	path := r.Replace("/organizations/{organizationId}/plan")
	options := UpdatePlanOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["organizationId"] = OrganizationId
	params["billingPlan"] = BillingPlan
	if options.enabledSetters["PaymentMethodId"] {
		params["paymentMethodId"] = options.PaymentMethodId
	}
	if options.enabledSetters["BillingAddressId"] {
		params["billingAddressId"] = options.BillingAddressId
	}
	if options.enabledSetters["Invites"] {
		params["invites"] = options.Invites
	}
	if options.enabledSetters["CouponId"] {
		params["couponId"] = options.CouponId
	}
	if options.enabledSetters["TaxId"] {
		params["taxId"] = options.TaxId
	}
	if options.enabledSetters["Budget"] {
		params["budget"] = options.Budget
	}
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type":       "application/json",
		"accept":             "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
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

// CancelDowngrade cancel the downgrade initiated for an organization.
func (srv *Organizations) CancelDowngrade(OrganizationId string) (*models.Organization, error) {
	r := strings.NewReplacer("{organizationId}", OrganizationId)
	path := r.Replace("/organizations/{organizationId}/plan/cancel")
	params := map[string]interface{}{}
	params["organizationId"] = OrganizationId
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type":       "application/json",
		"accept":             "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
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

type CreatePlanEstimationOptions struct {
	Invites        []string
	CouponId       string
	enabledSetters map[string]bool
}

func (options CreatePlanEstimationOptions) New() *CreatePlanEstimationOptions {
	options.enabledSetters = map[string]bool{
		"Invites":  false,
		"CouponId": false,
	}
	return &options
}

type CreatePlanEstimationOption func(*CreatePlanEstimationOptions)

func (srv *Organizations) WithCreatePlanEstimationInvites(v []string) CreatePlanEstimationOption {
	return func(o *CreatePlanEstimationOptions) {
		o.Invites = v
		o.enabledSetters["Invites"] = true
	}
}
func (srv *Organizations) WithCreatePlanEstimationCouponId(v string) CreatePlanEstimationOption {
	return func(o *CreatePlanEstimationOptions) {
		o.CouponId = v
		o.enabledSetters["CouponId"] = true
	}
}

// CreatePlanEstimation create a billing plan estimation for upgrading or
// downgrading an organization plan.
func (srv *Organizations) CreatePlanEstimation(OrganizationId string, BillingPlan string, optionalSetters ...CreatePlanEstimationOption) (*models.EstimationPlanChange, error) {
	r := strings.NewReplacer("{organizationId}", OrganizationId)
	path := r.Replace("/organizations/{organizationId}/plan/estimations")
	options := CreatePlanEstimationOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["organizationId"] = OrganizationId
	params["billingPlan"] = BillingPlan
	if options.enabledSetters["Invites"] {
		params["invites"] = options.Invites
	}
	if options.enabledSetters["CouponId"] {
		params["couponId"] = options.CouponId
	}
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type":       "application/json",
		"accept":             "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.EstimationPlanChange{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.EstimationPlanChange
	parsed, ok := resp.Result.(models.EstimationPlanChange)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// ListRegions get all available regions for an organization.
func (srv *Organizations) ListRegions(OrganizationId string) (*models.ConsoleRegionList, error) {
	r := strings.NewReplacer("{organizationId}", OrganizationId)
	path := r.Replace("/organizations/{organizationId}/regions")
	params := map[string]interface{}{}
	params["organizationId"] = OrganizationId
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"accept":             "application/json",
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

type GetScopesOptions struct {
	ProjectId      string
	enabledSetters map[string]bool
}

func (options GetScopesOptions) New() *GetScopesOptions {
	options.enabledSetters = map[string]bool{
		"ProjectId": false,
	}
	return &options
}

type GetScopesOption func(*GetScopesOptions)

func (srv *Organizations) WithGetScopesProjectId(v string) GetScopesOption {
	return func(o *GetScopesOptions) {
		o.ProjectId = v
		o.enabledSetters["ProjectId"] = true
	}
}

// GetScopes get Scopes
func (srv *Organizations) GetScopes(OrganizationId string, optionalSetters ...GetScopesOption) (*models.Roles, error) {
	r := strings.NewReplacer("{organizationId}", OrganizationId)
	path := r.Replace("/organizations/{organizationId}/roles")
	options := GetScopesOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["organizationId"] = OrganizationId
	if options.enabledSetters["ProjectId"] {
		params["projectId"] = options.ProjectId
	}
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"accept":             "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Roles{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Roles
	parsed, ok := resp.Result.(models.Roles)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// SetBillingTaxId set an organization's billing tax ID.
func (srv *Organizations) SetBillingTaxId(OrganizationId string, TaxId string) (*models.Organization, error) {
	r := strings.NewReplacer("{organizationId}", OrganizationId)
	path := r.Replace("/organizations/{organizationId}/taxId")
	params := map[string]interface{}{}
	params["organizationId"] = OrganizationId
	params["taxId"] = TaxId
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type":       "application/json",
		"accept":             "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
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

type GetUsageOptions struct {
	StartDate      string
	EndDate        string
	enabledSetters map[string]bool
}

func (options GetUsageOptions) New() *GetUsageOptions {
	options.enabledSetters = map[string]bool{
		"StartDate": false,
		"EndDate":   false,
	}
	return &options
}

type GetUsageOption func(*GetUsageOptions)

func (srv *Organizations) WithGetUsageStartDate(v string) GetUsageOption {
	return func(o *GetUsageOptions) {
		o.StartDate = v
		o.enabledSetters["StartDate"] = true
	}
}
func (srv *Organizations) WithGetUsageEndDate(v string) GetUsageOption {
	return func(o *GetUsageOptions) {
		o.EndDate = v
		o.enabledSetters["EndDate"] = true
	}
}

// GetUsage get the usage data for an organization.
func (srv *Organizations) GetUsage(OrganizationId string, optionalSetters ...GetUsageOption) (*models.UsageOrganization, error) {
	r := strings.NewReplacer("{organizationId}", OrganizationId)
	path := r.Replace("/organizations/{organizationId}/usage")
	options := GetUsageOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["organizationId"] = OrganizationId
	if options.enabledSetters["StartDate"] {
		params["startDate"] = options.StartDate
	}
	if options.enabledSetters["EndDate"] {
		params["endDate"] = options.EndDate
	}
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"accept":             "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.UsageOrganization{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.UsageOrganization
	parsed, ok := resp.Result.(models.UsageOrganization)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type ValidatePaymentOptions struct {
	Invites        []string
	enabledSetters map[string]bool
}

func (options ValidatePaymentOptions) New() *ValidatePaymentOptions {
	options.enabledSetters = map[string]bool{
		"Invites": false,
	}
	return &options
}

type ValidatePaymentOption func(*ValidatePaymentOptions)

func (srv *Organizations) WithValidatePaymentInvites(v []string) ValidatePaymentOption {
	return func(o *ValidatePaymentOptions) {
		o.Invites = v
		o.enabledSetters["Invites"] = true
	}
}

// ValidatePayment validate payment for team after creation or upgrade.
func (srv *Organizations) ValidatePayment(OrganizationId string, optionalSetters ...ValidatePaymentOption) (*models.Organization, error) {
	r := strings.NewReplacer("{organizationId}", OrganizationId)
	path := r.Replace("/organizations/{organizationId}/validate")
	options := ValidatePaymentOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["organizationId"] = OrganizationId
	if options.enabledSetters["Invites"] {
		params["invites"] = options.Invites
	}
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type":       "application/json",
		"accept":             "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
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
