package vcs

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/models"
	"strings"
)

// Vcs service
type Vcs struct {
	client client.Client
}

func New(clt client.Client) *Vcs {
	return &Vcs{
		client: clt,
	}
}

type CreateRepositoryDetectionOptions struct {
	ProviderRootDirectory string
	enabledSetters        map[string]bool
}

func (options CreateRepositoryDetectionOptions) New() *CreateRepositoryDetectionOptions {
	options.enabledSetters = map[string]bool{
		"ProviderRootDirectory": false,
	}
	return &options
}

type CreateRepositoryDetectionOption func(*CreateRepositoryDetectionOptions)

func (srv *Vcs) WithCreateRepositoryDetectionProviderRootDirectory(v string) CreateRepositoryDetectionOption {
	return func(o *CreateRepositoryDetectionOptions) {
		o.ProviderRootDirectory = v
		o.enabledSetters["ProviderRootDirectory"] = true
	}
}

// CreateRepositoryDetection analyze a GitHub repository to automatically
// detect the programming language and runtime environment. This endpoint
// scans the repository's files and language statistics to determine the
// appropriate runtime settings for your function. The GitHub installation
// must be properly configured and the repository must be accessible through
// your installation for this endpoint to work.
func (srv *Vcs) CreateRepositoryDetection(InstallationId string, ProviderRepositoryId string, Type string, optionalSetters ...CreateRepositoryDetectionOption) (models.Model, error) {
	r := strings.NewReplacer("{installationId}", InstallationId)
	path := r.Replace("/vcs/github/installations/{installationId}/detections")
	options := CreateRepositoryDetectionOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["installationId"] = InstallationId
	params["providerRepositoryId"] = ProviderRepositoryId
	params["type"] = Type
	if options.enabledSetters["ProviderRootDirectory"] {
		params["providerRootDirectory"] = options.ProviderRootDirectory
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

		var response map[string]interface{}
		if err := json.Unmarshal(bytes, &response); err != nil {
			return nil, err
		}
		if fmt.Sprint(response["type"]) == "runtime" {
			parsed := models.DetectionRuntime{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["type"]) == "framework" {
			parsed := models.DetectionFramework{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}

		return nil, errors.New("unable to match response to any expected response model")
	}
	parsed, ok := resp.Result.(models.Model)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return parsed, nil

}

type ListRepositoriesOptions struct {
	Search         string
	Queries        []string
	enabledSetters map[string]bool
}

func (options ListRepositoriesOptions) New() *ListRepositoriesOptions {
	options.enabledSetters = map[string]bool{
		"Search":  false,
		"Queries": false,
	}
	return &options
}

type ListRepositoriesOption func(*ListRepositoriesOptions)

func (srv *Vcs) WithListRepositoriesSearch(v string) ListRepositoriesOption {
	return func(o *ListRepositoriesOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}
func (srv *Vcs) WithListRepositoriesQueries(v []string) ListRepositoriesOption {
	return func(o *ListRepositoriesOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}

// ListRepositories get a list of GitHub repositories available through your
// installation. This endpoint returns repositories with their basic
// information, detected runtime environments, and latest push dates. You can
// optionally filter repositories using a search term. Each repository's
// runtime is automatically detected based on its contents and language
// statistics. The GitHub installation must be properly configured for this
// endpoint to work.
func (srv *Vcs) ListRepositories(InstallationId string, Type string, optionalSetters ...ListRepositoriesOption) (models.Model, error) {
	r := strings.NewReplacer("{installationId}", InstallationId)
	path := r.Replace("/vcs/github/installations/{installationId}/providerRepositories")
	options := ListRepositoriesOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["installationId"] = InstallationId
	params["type"] = Type
	if options.enabledSetters["Search"] {
		params["search"] = options.Search
	}
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

		var response map[string]interface{}
		if err := json.Unmarshal(bytes, &response); err != nil {
			return nil, err
		}
		if fmt.Sprint(response["type"]) == "runtime" {
			parsed := models.ProviderRepositoryRuntimeList{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["type"]) == "framework" {
			parsed := models.ProviderRepositoryFrameworkList{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}

		return nil, errors.New("unable to match response to any expected response model")
	}
	parsed, ok := resp.Result.(models.Model)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return parsed, nil

}

type CreateRepositoryOptions struct {
	ProviderNamespace string
	enabledSetters    map[string]bool
}

func (options CreateRepositoryOptions) New() *CreateRepositoryOptions {
	options.enabledSetters = map[string]bool{
		"ProviderNamespace": false,
	}
	return &options
}

type CreateRepositoryOption func(*CreateRepositoryOptions)

func (srv *Vcs) WithCreateRepositoryProviderNamespace(v string) CreateRepositoryOption {
	return func(o *CreateRepositoryOptions) {
		o.ProviderNamespace = v
		o.enabledSetters["ProviderNamespace"] = true
	}
}

// CreateRepository create a new GitHub repository through your installation.
// This endpoint allows you to create either a public or private repository by
// specifying a name and visibility setting. The repository will be created
// under your GitHub user account or organization, depending on your
// installation type. The GitHub installation must be properly configured and
// have the necessary permissions for repository creation.
func (srv *Vcs) CreateRepository(InstallationId string, Name string, Private bool, optionalSetters ...CreateRepositoryOption) (*models.ProviderRepository, error) {
	r := strings.NewReplacer("{installationId}", InstallationId)
	path := r.Replace("/vcs/github/installations/{installationId}/providerRepositories")
	options := CreateRepositoryOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["installationId"] = InstallationId
	params["name"] = Name
	params["private"] = Private
	if options.enabledSetters["ProviderNamespace"] {
		params["providerNamespace"] = options.ProviderNamespace
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

		parsed := models.ProviderRepository{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ProviderRepository
	parsed, ok := resp.Result.(models.ProviderRepository)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetRepository get detailed information about a specific GitHub repository
// from your installation. This endpoint returns repository details including
// its ID, name, visibility status, organization, and latest push date. The
// GitHub installation must be properly configured and have access to the
// requested repository for this endpoint to work.
func (srv *Vcs) GetRepository(InstallationId string, ProviderRepositoryId string) (*models.ProviderRepository, error) {
	r := strings.NewReplacer("{installationId}", InstallationId, "{providerRepositoryId}", ProviderRepositoryId)
	path := r.Replace("/vcs/github/installations/{installationId}/providerRepositories/{providerRepositoryId}")
	params := map[string]interface{}{}
	params["installationId"] = InstallationId
	params["providerRepositoryId"] = ProviderRepositoryId
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

		parsed := models.ProviderRepository{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ProviderRepository
	parsed, ok := resp.Result.(models.ProviderRepository)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type ListRepositoryBranchesOptions struct {
	Search         string
	Queries        []string
	enabledSetters map[string]bool
}

func (options ListRepositoryBranchesOptions) New() *ListRepositoryBranchesOptions {
	options.enabledSetters = map[string]bool{
		"Search":  false,
		"Queries": false,
	}
	return &options
}

type ListRepositoryBranchesOption func(*ListRepositoryBranchesOptions)

func (srv *Vcs) WithListRepositoryBranchesSearch(v string) ListRepositoryBranchesOption {
	return func(o *ListRepositoryBranchesOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}
func (srv *Vcs) WithListRepositoryBranchesQueries(v []string) ListRepositoryBranchesOption {
	return func(o *ListRepositoryBranchesOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}

// ListRepositoryBranches get a list of branches from a GitHub repository in
// your installation. This endpoint supports filtering by a search term and
// pagination using query strings such as `Query.limit()`, `Query.offset()`,
// `Query.cursorAfter()`, and `Query.cursorBefore()`. It returns branch names
// along with the total number of matches. The GitHub installation must be
// properly configured and have access to the requested repository for this
// endpoint to work.
func (srv *Vcs) ListRepositoryBranches(InstallationId string, ProviderRepositoryId string, optionalSetters ...ListRepositoryBranchesOption) (*models.BranchList, error) {
	r := strings.NewReplacer("{installationId}", InstallationId, "{providerRepositoryId}", ProviderRepositoryId)
	path := r.Replace("/vcs/github/installations/{installationId}/providerRepositories/{providerRepositoryId}/branches")
	options := ListRepositoryBranchesOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["installationId"] = InstallationId
	params["providerRepositoryId"] = ProviderRepositoryId
	if options.enabledSetters["Search"] {
		params["search"] = options.Search
	}
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

		parsed := models.BranchList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.BranchList
	parsed, ok := resp.Result.(models.BranchList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type GetRepositoryContentsOptions struct {
	ProviderRootDirectory string
	ProviderReference     string
	enabledSetters        map[string]bool
}

func (options GetRepositoryContentsOptions) New() *GetRepositoryContentsOptions {
	options.enabledSetters = map[string]bool{
		"ProviderRootDirectory": false,
		"ProviderReference":     false,
	}
	return &options
}

type GetRepositoryContentsOption func(*GetRepositoryContentsOptions)

func (srv *Vcs) WithGetRepositoryContentsProviderRootDirectory(v string) GetRepositoryContentsOption {
	return func(o *GetRepositoryContentsOptions) {
		o.ProviderRootDirectory = v
		o.enabledSetters["ProviderRootDirectory"] = true
	}
}
func (srv *Vcs) WithGetRepositoryContentsProviderReference(v string) GetRepositoryContentsOption {
	return func(o *GetRepositoryContentsOptions) {
		o.ProviderReference = v
		o.enabledSetters["ProviderReference"] = true
	}
}

// GetRepositoryContents get a list of files and directories from a GitHub
// repository connected to your project. This endpoint returns the contents of
// a specified repository path, including file names, sizes, and whether each
// item is a file or directory. The GitHub installation must be properly
// configured and the repository must be accessible through your installation
// for this endpoint to work.
func (srv *Vcs) GetRepositoryContents(InstallationId string, ProviderRepositoryId string, optionalSetters ...GetRepositoryContentsOption) (*models.VcsContentList, error) {
	r := strings.NewReplacer("{installationId}", InstallationId, "{providerRepositoryId}", ProviderRepositoryId)
	path := r.Replace("/vcs/github/installations/{installationId}/providerRepositories/{providerRepositoryId}/contents")
	options := GetRepositoryContentsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["installationId"] = InstallationId
	params["providerRepositoryId"] = ProviderRepositoryId
	if options.enabledSetters["ProviderRootDirectory"] {
		params["providerRootDirectory"] = options.ProviderRootDirectory
	}
	if options.enabledSetters["ProviderReference"] {
		params["providerReference"] = options.ProviderReference
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

		parsed := models.VcsContentList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.VcsContentList
	parsed, ok := resp.Result.(models.VcsContentList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// UpdateExternalDeployments authorize and create deployments for a GitHub
// pull request in your project. This endpoint allows external contributions
// by creating deployments from pull requests, enabling preview environments
// for code review. The pull request must be open and not previously
// authorized. The GitHub installation must be properly configured and have
// access to both the repository and pull request for this endpoint to work.
func (srv *Vcs) UpdateExternalDeployments(InstallationId string, RepositoryId string, ProviderPullRequestId string) (*interface{}, error) {
	r := strings.NewReplacer("{installationId}", InstallationId, "{repositoryId}", RepositoryId)
	path := r.Replace("/vcs/github/installations/{installationId}/repositories/{repositoryId}")
	params := map[string]interface{}{}
	params["installationId"] = InstallationId
	params["repositoryId"] = RepositoryId
	params["providerPullRequestId"] = ProviderPullRequestId
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

type ListInstallationsOptions struct {
	Queries        []string
	Search         string
	Total          bool
	enabledSetters map[string]bool
}

func (options ListInstallationsOptions) New() *ListInstallationsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Search":  false,
		"Total":   false,
	}
	return &options
}

type ListInstallationsOption func(*ListInstallationsOptions)

func (srv *Vcs) WithListInstallationsQueries(v []string) ListInstallationsOption {
	return func(o *ListInstallationsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Vcs) WithListInstallationsSearch(v string) ListInstallationsOption {
	return func(o *ListInstallationsOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}
func (srv *Vcs) WithListInstallationsTotal(v bool) ListInstallationsOption {
	return func(o *ListInstallationsOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}

// ListInstallations list all VCS installations configured for the current
// project. This endpoint returns a list of installations including their
// provider, organization, and other configuration details.
func (srv *Vcs) ListInstallations(optionalSetters ...ListInstallationsOption) (*models.InstallationList, error) {
	path := "/vcs/installations"
	options := ListInstallationsOptions{}.New()
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
		"accept":             "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.InstallationList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.InstallationList
	parsed, ok := resp.Result.(models.InstallationList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetInstallation get a VCS installation by its unique ID. This endpoint
// returns the installation's details including its provider, organization,
// and configuration.
func (srv *Vcs) GetInstallation(InstallationId string) (*models.Installation, error) {
	r := strings.NewReplacer("{installationId}", InstallationId)
	path := r.Replace("/vcs/installations/{installationId}")
	params := map[string]interface{}{}
	params["installationId"] = InstallationId
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

		parsed := models.Installation{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Installation
	parsed, ok := resp.Result.(models.Installation)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// DeleteInstallation delete a VCS installation by its unique ID. This
// endpoint removes the installation and all its associated repositories from
// the project.
func (srv *Vcs) DeleteInstallation(InstallationId string) (*interface{}, error) {
	r := strings.NewReplacer("{installationId}", InstallationId)
	path := r.Replace("/vcs/installations/{installationId}")
	params := map[string]interface{}{}
	params["installationId"] = InstallationId
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type":       "application/json",
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

type ListNamespacesOptions struct {
	Search         string
	Queries        []string
	enabledSetters map[string]bool
}

func (options ListNamespacesOptions) New() *ListNamespacesOptions {
	options.enabledSetters = map[string]bool{
		"Search":  false,
		"Queries": false,
	}
	return &options
}

type ListNamespacesOption func(*ListNamespacesOptions)

func (srv *Vcs) WithListNamespacesSearch(v string) ListNamespacesOption {
	return func(o *ListNamespacesOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}
func (srv *Vcs) WithListNamespacesQueries(v []string) ListNamespacesOption {
	return func(o *ListNamespacesOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}

// ListNamespaces list provider namespaces available to a VCS installation.
// This can include the user personal namespace and any groups or
// organizations the installation can browse.
func (srv *Vcs) ListNamespaces(InstallationId string, optionalSetters ...ListNamespacesOption) (*models.VcsNamespaceList, error) {
	r := strings.NewReplacer("{installationId}", InstallationId)
	path := r.Replace("/vcs/installations/{installationId}/namespaces")
	options := ListNamespacesOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["installationId"] = InstallationId
	if options.enabledSetters["Search"] {
		params["search"] = options.Search
	}
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

		parsed := models.VcsNamespaceList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.VcsNamespaceList
	parsed, ok := resp.Result.(models.VcsNamespaceList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
