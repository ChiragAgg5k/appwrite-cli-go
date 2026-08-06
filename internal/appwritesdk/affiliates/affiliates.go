package affiliates

import (
	"encoding/json"
	"errors"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/models"
	"strings"
)

// Affiliates service
type Affiliates struct {
	client client.Client
}

func New(clt client.Client) *Affiliates {
	return &Affiliates{
		client: clt,
	}
}

type ListLinksOptions struct {
	Queries        []string
	enabledSetters map[string]bool
}

func (options ListLinksOptions) New() *ListLinksOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
	}
	return &options
}

type ListLinksOption func(*ListLinksOptions)

func (srv *Affiliates) WithListLinksQueries(v []string) ListLinksOption {
	return func(o *ListLinksOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}

// ListLinks list affiliate links for the current account.
func (srv *Affiliates) ListLinks(optionalSetters ...ListLinksOption) (*models.AffiliateLinkList, error) {
	path := "/affiliates/links"
	options := ListLinksOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
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

		parsed := models.AffiliateLinkList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AffiliateLinkList
	parsed, ok := resp.Result.(models.AffiliateLinkList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateLinkOptions struct {
	Name           string
	enabledSetters map[string]bool
}

func (options CreateLinkOptions) New() *CreateLinkOptions {
	options.enabledSetters = map[string]bool{
		"Name": false,
	}
	return &options
}

type CreateLinkOption func(*CreateLinkOptions)

func (srv *Affiliates) WithCreateLinkName(v string) CreateLinkOption {
	return func(o *CreateLinkOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}

// CreateLink create a shareable affiliate link for the current account. Every
// console user is automatically in the affiliates program.
func (srv *Affiliates) CreateLink(LinkId string, optionalSetters ...CreateLinkOption) (*models.AffiliateLink, error) {
	path := "/affiliates/links"
	options := CreateLinkOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["linkId"] = LinkId
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
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

		parsed := models.AffiliateLink{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AffiliateLink
	parsed, ok := resp.Result.(models.AffiliateLink)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetLink get a single affiliate link owned by the current account.
func (srv *Affiliates) GetLink(LinkId string) (*models.AffiliateLink, error) {
	r := strings.NewReplacer("{linkId}", LinkId)
	path := r.Replace("/affiliates/links/{linkId}")
	params := map[string]interface{}{}
	params["linkId"] = LinkId
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

		parsed := models.AffiliateLink{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AffiliateLink
	parsed, ok := resp.Result.(models.AffiliateLink)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// DeleteLink delete an affiliate link owned by the current account. Existing
// referrals and rewards keep their stored link IDs for history.
func (srv *Affiliates) DeleteLink(LinkId string) (*interface{}, error) {
	r := strings.NewReplacer("{linkId}", LinkId)
	path := r.Replace("/affiliates/links/{linkId}")
	params := map[string]interface{}{}
	params["linkId"] = LinkId
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

type ListReferralsOptions struct {
	Queries        []string
	enabledSetters map[string]bool
}

func (options ListReferralsOptions) New() *ListReferralsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
	}
	return &options
}

type ListReferralsOption func(*ListReferralsOptions)

func (srv *Affiliates) WithListReferralsQueries(v []string) ListReferralsOption {
	return func(o *ListReferralsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}

// ListReferrals list referrals attributed to the current account's affiliate
// links. Responses include privacy-safe metadata only (truncated user ID and
// signup country), never email or name. Referrals are created automatically
// on signup when the invite cookie is present.
func (srv *Affiliates) ListReferrals(optionalSetters ...ListReferralsOption) (*models.AffiliateReferralList, error) {
	path := "/affiliates/referrals"
	options := ListReferralsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
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

		parsed := models.AffiliateReferralList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AffiliateReferralList
	parsed, ok := resp.Result.(models.AffiliateReferralList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type ListRewardsOptions struct {
	Queries        []string
	enabledSetters map[string]bool
}

func (options ListRewardsOptions) New() *ListRewardsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
	}
	return &options
}

type ListRewardsOption func(*ListRewardsOptions)

func (srv *Affiliates) WithListRewardsQueries(v []string) ListRewardsOption {
	return func(o *ListRewardsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}

// ListRewards list rewards earned by the current account from affiliate link
// conversions.
func (srv *Affiliates) ListRewards(optionalSetters ...ListRewardsOption) (*models.AffiliateRewardList, error) {
	path := "/affiliates/rewards"
	options := ListRewardsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
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

		parsed := models.AffiliateRewardList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AffiliateRewardList
	parsed, ok := resp.Result.(models.AffiliateRewardList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// UpdateReward claim a pending affiliate reward by setting its status to
// `claimed`. Creates organization credits for the target organization. The
// current user must be an owner of that organization.
func (srv *Affiliates) UpdateReward(RewardId string, Status string, OrganizationId string) (*models.AffiliateReward, error) {
	r := strings.NewReplacer("{rewardId}", RewardId)
	path := r.Replace("/affiliates/rewards/{rewardId}")
	params := map[string]interface{}{}
	params["rewardId"] = RewardId
	params["status"] = Status
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

		parsed := models.AffiliateReward{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AffiliateReward
	parsed, ok := resp.Result.(models.AffiliateReward)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
