package usage

import (
	"encoding/json"
	"errors"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/models"
	"strings"
)

// Usage service
type Usage struct {
	client client.Client
}

func New(clt client.Client) *Usage {
	return &Usage{
		client: clt,
	}
}

type ListEventsOptions struct {
	Queries        []string
	Interval       string
	Dimensions     []string
	StartAt        string
	EndAt          string
	OrderBy        string
	OrderDir       string
	Limit          int
	Offset         int
	enabledSetters map[string]bool
}

func (options ListEventsOptions) New() *ListEventsOptions {
	options.enabledSetters = map[string]bool{
		"Queries":    false,
		"Interval":   false,
		"Dimensions": false,
		"StartAt":    false,
		"EndAt":      false,
		"OrderBy":    false,
		"OrderDir":   false,
		"Limit":      false,
		"Offset":     false,
	}
	return &options
}

type ListEventsOption func(*ListEventsOptions)

func (srv *Usage) WithListEventsQueries(v []string) ListEventsOption {
	return func(o *ListEventsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Usage) WithListEventsInterval(v string) ListEventsOption {
	return func(o *ListEventsOptions) {
		o.Interval = v
		o.enabledSetters["Interval"] = true
	}
}
func (srv *Usage) WithListEventsDimensions(v []string) ListEventsOption {
	return func(o *ListEventsOptions) {
		o.Dimensions = v
		o.enabledSetters["Dimensions"] = true
	}
}
func (srv *Usage) WithListEventsStartAt(v string) ListEventsOption {
	return func(o *ListEventsOptions) {
		o.StartAt = v
		o.enabledSetters["StartAt"] = true
	}
}
func (srv *Usage) WithListEventsEndAt(v string) ListEventsOption {
	return func(o *ListEventsOptions) {
		o.EndAt = v
		o.enabledSetters["EndAt"] = true
	}
}
func (srv *Usage) WithListEventsOrderBy(v string) ListEventsOption {
	return func(o *ListEventsOptions) {
		o.OrderBy = v
		o.enabledSetters["OrderBy"] = true
	}
}
func (srv *Usage) WithListEventsOrderDir(v string) ListEventsOption {
	return func(o *ListEventsOptions) {
		o.OrderDir = v
		o.enabledSetters["OrderDir"] = true
	}
}
func (srv *Usage) WithListEventsLimit(v int) ListEventsOption {
	return func(o *ListEventsOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *Usage) WithListEventsOffset(v int) ListEventsOption {
	return func(o *ListEventsOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}

// ListEvents aggregate usage event metrics. `metrics[]` (1-10) is required;
// the response always contains one entry per requested metric, each with its
// own `points[]` time series.
//
// **Two response shapes**:
// - Omit `interval` for a flat top-N table — one point per dimension
// combination, no time axis. Useful for "top 10 paths by bandwidth in the
// last 7 days".
// - Pass `interval` (`1m`, `15m`, `30m`, `1h`, `1d`) for a time series —
// one point per (time bucket × dimension combination).
//
// `dimensions[]` breaks each point down by one or more attributes (service,
// path, status, country, …). `queries[]` filters the underlying events
// using the standard Utopia query syntax — `equal("path",
// ["/v1/storage/files"])`, `equal("resourceType", ["bucket"])`,
// `equal("resourceId", ["abc123"])`, `startsWith("path", ["/v1/storage"])`,
// `equal("status", ["200", "201"])`, `isNotNull("resourceId")`. Supported
// attributes: see `queries[]` param. Supported methods: `equal`, `notEqual`,
// `contains`, `startsWith`, `endsWith`, `isNull`, `isNotNull`. Pass multiple
// metrics to render stacked charts in one round-trip.
// `orderBy=value`+`orderDir=desc`+`limit=N` returns the top-N by aggregated
// value. When `startAt` is omitted, the default window adapts to `interval`
// (or 7d when interval is omitted).
func (srv *Usage) ListEvents(Metrics []string, optionalSetters ...ListEventsOption) (*models.UsageEventList, error) {
	path := "/usage/events"
	options := ListEventsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["metrics"] = Metrics
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["Interval"] {
		params["interval"] = options.Interval
	}
	if options.enabledSetters["Dimensions"] {
		params["dimensions"] = options.Dimensions
	}
	if options.enabledSetters["StartAt"] {
		params["startAt"] = options.StartAt
	}
	if options.enabledSetters["EndAt"] {
		params["endAt"] = options.EndAt
	}
	if options.enabledSetters["OrderBy"] {
		params["orderBy"] = options.OrderBy
	}
	if options.enabledSetters["OrderDir"] {
		params["orderDir"] = options.OrderDir
	}
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

		parsed := models.UsageEventList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.UsageEventList
	parsed, ok := resp.Result.(models.UsageEventList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type ListGaugesOptions struct {
	Queries        []string
	Interval       string
	Dimensions     []string
	StartAt        string
	EndAt          string
	OrderBy        string
	OrderDir       string
	Limit          int
	Offset         int
	enabledSetters map[string]bool
}

func (options ListGaugesOptions) New() *ListGaugesOptions {
	options.enabledSetters = map[string]bool{
		"Queries":    false,
		"Interval":   false,
		"Dimensions": false,
		"StartAt":    false,
		"EndAt":      false,
		"OrderBy":    false,
		"OrderDir":   false,
		"Limit":      false,
		"Offset":     false,
	}
	return &options
}

type ListGaugesOption func(*ListGaugesOptions)

func (srv *Usage) WithListGaugesQueries(v []string) ListGaugesOption {
	return func(o *ListGaugesOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Usage) WithListGaugesInterval(v string) ListGaugesOption {
	return func(o *ListGaugesOptions) {
		o.Interval = v
		o.enabledSetters["Interval"] = true
	}
}
func (srv *Usage) WithListGaugesDimensions(v []string) ListGaugesOption {
	return func(o *ListGaugesOptions) {
		o.Dimensions = v
		o.enabledSetters["Dimensions"] = true
	}
}
func (srv *Usage) WithListGaugesStartAt(v string) ListGaugesOption {
	return func(o *ListGaugesOptions) {
		o.StartAt = v
		o.enabledSetters["StartAt"] = true
	}
}
func (srv *Usage) WithListGaugesEndAt(v string) ListGaugesOption {
	return func(o *ListGaugesOptions) {
		o.EndAt = v
		o.enabledSetters["EndAt"] = true
	}
}
func (srv *Usage) WithListGaugesOrderBy(v string) ListGaugesOption {
	return func(o *ListGaugesOptions) {
		o.OrderBy = v
		o.enabledSetters["OrderBy"] = true
	}
}
func (srv *Usage) WithListGaugesOrderDir(v string) ListGaugesOption {
	return func(o *ListGaugesOptions) {
		o.OrderDir = v
		o.enabledSetters["OrderDir"] = true
	}
}
func (srv *Usage) WithListGaugesLimit(v int) ListGaugesOption {
	return func(o *ListGaugesOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *Usage) WithListGaugesOffset(v int) ListGaugesOption {
	return func(o *ListGaugesOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}

// ListGauges aggregate usage gauge snapshots. Gauges are point-in-time values
// (storage totals, resource counts, …); each point carries the latest
// snapshot in its interval via `argMax(value, time)`. `metrics[]` (1-10) is
// required; the response always contains one entry per requested metric, each
// with its own `points[]` time series.
//
// A metric with no stored samples in the window returns an empty `points[]`.
// A metric that really did read zero returns a point whose `value` is `0`, so
// "no such series" and "a genuine zero" are different answers.
//
// **Two response shapes**:
// - Omit `interval` for a flat top-N table — `argMax(value, time)` per
// dimension combination over the whole window, no time axis. Useful for "top
// 10 resources by current storage".
// - Pass `interval` (`1m`, `15m`, `30m`, `1h`, `1d`) for a time series —
// one snapshot per (time bucket × dimension combination).
//
// `dimensions[]` breaks each point down further. Supported on gauges:
// `resourceId`, `teamId`, `service`, `resourceType`, `ordinal`. `service` and
// `resourceType` enable per-service / per-resource-type panels (e.g.
// storage-by-service: group `files.storage`, `deployments.storage`,
// `builds.storage`, `databases.storage` by `service`). `ordinal` separates
// per-node series for multi-node resources such as dedicated databases. It is
// a stable per-node identity, not a role — ordinal 0 is the first member
// created, and a failover can leave the primary on any ordinal, so read the
// role from the database's replicas endpoint rather than inferring it here.
// `queries[]` filters the underlying rows using the standard Utopia query
// syntax — `equal("resourceType", ["bucket"])`, `equal("resourceId",
// ["abc123"])`, `equal("teamId", ["team_x"])`, `equal("ordinal", ["0"])`,
// `isNotNull("teamId")`. Supported attributes: see `queries[]` param.
// Supported methods: `equal`, `notEqual`, `isNull`, `isNotNull`. Pass
// multiple metrics to render stacked charts in one round-trip.
// `orderBy=value`+`orderDir=desc`+`limit=N` returns the top-N. When `startAt`
// is omitted, the default window adapts to interval (or 7d when interval is
// omitted).
func (srv *Usage) ListGauges(Metrics []string, optionalSetters ...ListGaugesOption) (*models.UsageGaugeList, error) {
	path := "/usage/gauges"
	options := ListGaugesOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["metrics"] = Metrics
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["Interval"] {
		params["interval"] = options.Interval
	}
	if options.enabledSetters["Dimensions"] {
		params["dimensions"] = options.Dimensions
	}
	if options.enabledSetters["StartAt"] {
		params["startAt"] = options.StartAt
	}
	if options.enabledSetters["EndAt"] {
		params["endAt"] = options.EndAt
	}
	if options.enabledSetters["OrderBy"] {
		params["orderBy"] = options.OrderBy
	}
	if options.enabledSetters["OrderDir"] {
		params["orderDir"] = options.OrderDir
	}
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

		parsed := models.UsageGaugeList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.UsageGaugeList
	parsed, ok := resp.Result.(models.UsageGaugeList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
