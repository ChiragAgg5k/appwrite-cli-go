package services

import (
	"github.com/spf13/cobra"

	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/vcs"

	"github.com/ChiragAgg5k/appwrite-cli-go/internal/app"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/query"
)

// NewVcsCommand builds the `vcs` command tree.
func NewVcsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vcs",
		Short: "The VCS service allows you to interact with providers like GitHub, GitLab etc.",
	}

	cmd.AddCommand(newVcsCreateRepositoryDetectionCommand())
	cmd.AddCommand(newVcsListRepositoriesCommand())
	cmd.AddCommand(newVcsCreateRepositoryCommand())
	cmd.AddCommand(newVcsGetRepositoryCommand())
	cmd.AddCommand(newVcsListRepositoryBranchesCommand())
	cmd.AddCommand(newVcsGetRepositoryContentsCommand())
	cmd.AddCommand(newVcsUpdateExternalDeploymentsCommand())
	cmd.AddCommand(newVcsListInstallationsCommand())
	cmd.AddCommand(newVcsGetInstallationCommand())
	cmd.AddCommand(newVcsDeleteInstallationCommand())
	cmd.AddCommand(newVcsListNamespacesCommand())

	return cmd
}

func newVcsCreateRepositoryDetectionCommand() *cobra.Command {
	var installationId string
	var providerRepositoryId string
	var typeArg string
	var providerRootDirectory string

	cmd := &cobra.Command{
		Use:   "create-repository-detection",
		Short: "Analyze a GitHub repository to automatically detect the programming language and runtime environment. This endpoint scans the repository's files and language statistics to determine the appropriate runtime settings for your function. The GitHub installation must be properly configured and the repository must be accessible through your installation for this endpoint to work.",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := app.ClientForProject("")
			if err != nil {
				return err
			}
			service := vcs.New(client)

			// An unset flag must be omitted, not sent as its zero value: the
			// TypeScript passes undefined and the SDK drops it.
			options := []vcs.CreateRepositoryDetectionOption{}
			if cmd.Flags().Changed("provider-root-directory") {
				options = append(options, service.WithCreateRepositoryDetectionProviderRootDirectory(providerRootDirectory))
			}

			result, err := service.CreateRepositoryDetection(installationId, providerRepositoryId, typeArg, options...)
			if err != nil {
				return err
			}

			return app.Render(result)
		},
	}

	cmd.Flags().StringVar(&installationId, "installation-id", "", "Installation Id")
	_ = cmd.MarkFlagRequired("installation-id")
	cmd.Flags().StringVar(&providerRepositoryId, "provider-repository-id", "", "Repository Id")
	_ = cmd.MarkFlagRequired("provider-repository-id")
	cmd.Flags().StringVar(&typeArg, "type", "", "Detector type. Must be one of the following: runtime, framework")
	_ = cmd.MarkFlagRequired("type")
	cmd.Flags().StringVar(&providerRootDirectory, "provider-root-directory", "", "Path to Root Directory")
	return cmd
}

func newVcsListRepositoriesCommand() *cobra.Command {
	var installationId string
	var typeArg string
	var search string
	var queries []string
	var filter []string
	var where []string
	var sortAsc []string
	var sortDesc []string
	var limit int
	var offset int
	var cursorAfter string
	var cursorBefore string

	cmd := &cobra.Command{
		Use:   "list-repositories",
		Short: "Get a list of GitHub repositories available through your installation. This endpoint returns repositories with their basic information, detected runtime environments, and latest push dates. You can optionally filter repositories using a search term. Each repository's runtime is automatically detected based on its contents and language statistics. The GitHub installation must be properly configured for this endpoint to work.",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := app.ClientForProject("")
			if err != nil {
				return err
			}
			service := vcs.New(client)

			parsedFilter, err := query.ParseFilters(filter)
			if err != nil {
				return err
			}
			parsedWhere, err := query.ParseFilters(where)
			if err != nil {
				return err
			}

			queries, err := query.Build(query.Options{
				Queries:      queries,
				Filter:       parsedFilter,
				Where:        parsedWhere,
				SortAsc:      sortAsc,
				SortDesc:     sortDesc,
				Limit:        app.FlagInt(cmd, "limit", limit),
				Offset:       app.FlagInt(cmd, "offset", offset),
				CursorAfter:  app.FlagString(cmd, "cursor-after", cursorAfter),
				CursorBefore: app.FlagString(cmd, "cursor-before", cursorBefore),
			})
			if err != nil {
				return err
			}

			// An unset flag must be omitted, not sent as its zero value: the
			// TypeScript passes undefined and the SDK drops it.
			options := []vcs.ListRepositoriesOption{}
			if cmd.Flags().Changed("search") {
				options = append(options, service.WithListRepositoriesSearch(search))
			}
			if cmd.Flags().Changed("queries") {
				options = append(options, service.WithListRepositoriesQueries(queries))
			}

			result, err := service.ListRepositories(installationId, typeArg, options...)
			if err != nil {
				return err
			}

			return app.Render(result)
		},
	}

	cmd.Flags().StringVar(&installationId, "installation-id", "", "Installation Id")
	_ = cmd.MarkFlagRequired("installation-id")
	cmd.Flags().StringVar(&typeArg, "type", "", "Detector type. Must be one of the following: runtime, framework")
	_ = cmd.MarkFlagRequired("type")
	cmd.Flags().StringVar(&search, "search", "", "Search term to filter your list results. Max length: 256 chars.")
	cmd.Flags().StringArrayVar(&queries, "queries", nil, "Array of query strings generated using the Query class provided by the SDK. Learn more about queries (https://appwrite.io/docs/queries). Only supported methods are limit, offset, and equal on namespace.")
	cmd.Flags().StringArrayVar(&filter, "filter", nil, "Filter using a simple comparison expression. Repeat for multiple filters. Supports field=value, field!=value, field>value, field>=value, field<value, and field<=value.")
	cmd.Flags().StringArrayVar(&where, "where", nil, "Deprecated. Use --filter instead. Filter using a simple comparison expression. Repeat for multiple filters.")
	cmd.Flags().StringArrayVar(&sortAsc, "sort-asc", nil, "Sort results by an attribute in ascending order. Repeat for multiple sort fields.")
	cmd.Flags().StringArrayVar(&sortDesc, "sort-desc", nil, "Sort results by an attribute in descending order. Repeat for multiple sort fields.")
	cmd.Flags().IntVar(&limit, "limit", 0, "Maximum number of results to return.")
	cmd.Flags().IntVar(&offset, "offset", 0, "Number of results to skip.")
	cmd.Flags().StringVar(&cursorAfter, "cursor-after", "", "Return results after this cursor ID.")
	cmd.Flags().StringVar(&cursorBefore, "cursor-before", "", "Return results before this cursor ID.")
	return cmd
}

func newVcsCreateRepositoryCommand() *cobra.Command {
	var installationId string
	var name string
	var xprivate bool
	var providerNamespace string

	cmd := &cobra.Command{
		Use:   "create-repository",
		Short: "Create a new GitHub repository through your installation. This endpoint allows you to create either a public or private repository by specifying a name and visibility setting. The repository will be created under your GitHub user account or organization, depending on your installation type. The GitHub installation must be properly configured and have the necessary permissions for repository creation.",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := app.ClientForProject("")
			if err != nil {
				return err
			}
			service := vcs.New(client)

			// An unset flag must be omitted, not sent as its zero value: the
			// TypeScript passes undefined and the SDK drops it.
			options := []vcs.CreateRepositoryOption{}
			if cmd.Flags().Changed("provider-namespace") {
				options = append(options, service.WithCreateRepositoryProviderNamespace(providerNamespace))
			}

			result, err := service.CreateRepository(installationId, name, xprivate, options...)
			if err != nil {
				return err
			}

			return app.Render(result)
		},
	}

	cmd.Flags().StringVar(&installationId, "installation-id", "", "Installation Id")
	_ = cmd.MarkFlagRequired("installation-id")
	cmd.Flags().StringVar(&name, "name", "", "Repository name (slug)")
	_ = cmd.MarkFlagRequired("name")
	cmd.Flags().BoolVar(&xprivate, "xprivate", false, "Mark repository public or private")
	_ = cmd.MarkFlagRequired("xprivate")
	cmd.Flags().StringVar(&providerNamespace, "provider-namespace", "", "Namespace of the git repository. Defaults to the installation's own namespace.")
	return cmd
}

func newVcsGetRepositoryCommand() *cobra.Command {
	var installationId string
	var providerRepositoryId string

	cmd := &cobra.Command{
		Use:   "get-repository",
		Short: "Get detailed information about a specific GitHub repository from your installation. This endpoint returns repository details including its ID, name, visibility status, organization, and latest push date. The GitHub installation must be properly configured and have access to the requested repository for this endpoint to work.",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := app.ClientForProject("")
			if err != nil {
				return err
			}
			service := vcs.New(client)

			result, err := service.GetRepository(installationId, providerRepositoryId)
			if err != nil {
				return err
			}

			return app.Render(result)
		},
	}

	cmd.Flags().StringVar(&installationId, "installation-id", "", "Installation Id")
	_ = cmd.MarkFlagRequired("installation-id")
	cmd.Flags().StringVar(&providerRepositoryId, "provider-repository-id", "", "Repository Id")
	_ = cmd.MarkFlagRequired("provider-repository-id")
	return cmd
}

func newVcsListRepositoryBranchesCommand() *cobra.Command {
	var installationId string
	var providerRepositoryId string
	var search string
	var queries []string
	var filter []string
	var where []string
	var sortAsc []string
	var sortDesc []string
	var limit int
	var offset int
	var cursorAfter string
	var cursorBefore string

	cmd := &cobra.Command{
		Use:   "list-repository-branches",
		Short: "Get a list of branches from a GitHub repository in your installation. This endpoint supports filtering by a search term and pagination using query strings such as `Query.limit()`, `Query.offset()`, `Query.cursorAfter()`, and `Query.cursorBefore()`. It returns branch names along with the total number of matches. The GitHub installation must be properly configured and have access to the requested repository for this endpoint to work.\n",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := app.ClientForProject("")
			if err != nil {
				return err
			}
			service := vcs.New(client)

			parsedFilter, err := query.ParseFilters(filter)
			if err != nil {
				return err
			}
			parsedWhere, err := query.ParseFilters(where)
			if err != nil {
				return err
			}

			queries, err := query.Build(query.Options{
				Queries:      queries,
				Filter:       parsedFilter,
				Where:        parsedWhere,
				SortAsc:      sortAsc,
				SortDesc:     sortDesc,
				Limit:        app.FlagInt(cmd, "limit", limit),
				Offset:       app.FlagInt(cmd, "offset", offset),
				CursorAfter:  app.FlagString(cmd, "cursor-after", cursorAfter),
				CursorBefore: app.FlagString(cmd, "cursor-before", cursorBefore),
			})
			if err != nil {
				return err
			}

			// An unset flag must be omitted, not sent as its zero value: the
			// TypeScript passes undefined and the SDK drops it.
			options := []vcs.ListRepositoryBranchesOption{}
			if cmd.Flags().Changed("search") {
				options = append(options, service.WithListRepositoryBranchesSearch(search))
			}
			if cmd.Flags().Changed("queries") {
				options = append(options, service.WithListRepositoryBranchesQueries(queries))
			}

			result, err := service.ListRepositoryBranches(installationId, providerRepositoryId, options...)
			if err != nil {
				return err
			}

			return app.Render(result)
		},
	}

	cmd.Flags().StringVar(&installationId, "installation-id", "", "Installation Id")
	_ = cmd.MarkFlagRequired("installation-id")
	cmd.Flags().StringVar(&providerRepositoryId, "provider-repository-id", "", "Repository Id")
	_ = cmd.MarkFlagRequired("provider-repository-id")
	cmd.Flags().StringVar(&search, "search", "", "Search term to filter your list results. Max length: 256 chars.")
	cmd.Flags().StringArrayVar(&queries, "queries", nil, "Array of query strings generated using the Query class provided by the SDK. Learn more about queries (https://appwrite.io/docs/queries). Only supported methods are limit, offset, cursorAfter, and cursorBefore")
	cmd.Flags().StringArrayVar(&filter, "filter", nil, "Filter using a simple comparison expression. Repeat for multiple filters. Supports field=value, field!=value, field>value, field>=value, field<value, and field<=value.")
	cmd.Flags().StringArrayVar(&where, "where", nil, "Deprecated. Use --filter instead. Filter using a simple comparison expression. Repeat for multiple filters.")
	cmd.Flags().StringArrayVar(&sortAsc, "sort-asc", nil, "Sort results by an attribute in ascending order. Repeat for multiple sort fields.")
	cmd.Flags().StringArrayVar(&sortDesc, "sort-desc", nil, "Sort results by an attribute in descending order. Repeat for multiple sort fields.")
	cmd.Flags().IntVar(&limit, "limit", 0, "Maximum number of results to return.")
	cmd.Flags().IntVar(&offset, "offset", 0, "Number of results to skip.")
	cmd.Flags().StringVar(&cursorAfter, "cursor-after", "", "Return results after this cursor ID.")
	cmd.Flags().StringVar(&cursorBefore, "cursor-before", "", "Return results before this cursor ID.")
	return cmd
}

func newVcsGetRepositoryContentsCommand() *cobra.Command {
	var installationId string
	var providerRepositoryId string
	var providerRootDirectory string
	var providerReference string

	cmd := &cobra.Command{
		Use:   "get-repository-contents",
		Short: "Get a list of files and directories from a GitHub repository connected to your project. This endpoint returns the contents of a specified repository path, including file names, sizes, and whether each item is a file or directory. The GitHub installation must be properly configured and the repository must be accessible through your installation for this endpoint to work.",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := app.ClientForProject("")
			if err != nil {
				return err
			}
			service := vcs.New(client)

			// An unset flag must be omitted, not sent as its zero value: the
			// TypeScript passes undefined and the SDK drops it.
			options := []vcs.GetRepositoryContentsOption{}
			if cmd.Flags().Changed("provider-root-directory") {
				options = append(options, service.WithGetRepositoryContentsProviderRootDirectory(providerRootDirectory))
			}
			if cmd.Flags().Changed("provider-reference") {
				options = append(options, service.WithGetRepositoryContentsProviderReference(providerReference))
			}

			result, err := service.GetRepositoryContents(installationId, providerRepositoryId, options...)
			if err != nil {
				return err
			}

			return app.Render(result)
		},
	}

	cmd.Flags().StringVar(&installationId, "installation-id", "", "Installation Id")
	_ = cmd.MarkFlagRequired("installation-id")
	cmd.Flags().StringVar(&providerRepositoryId, "provider-repository-id", "", "Repository Id")
	_ = cmd.MarkFlagRequired("provider-repository-id")
	cmd.Flags().StringVar(&providerRootDirectory, "provider-root-directory", "", "Path to get contents of nested directory")
	cmd.Flags().StringVar(&providerReference, "provider-reference", "", "Git reference (branch, tag, commit) to get contents from")
	return cmd
}

func newVcsUpdateExternalDeploymentsCommand() *cobra.Command {
	var installationId string
	var repositoryId string
	var providerPullRequestId string

	cmd := &cobra.Command{
		Use:   "update-external-deployments",
		Short: "Authorize and create deployments for a GitHub pull request in your project. This endpoint allows external contributions by creating deployments from pull requests, enabling preview environments for code review. The pull request must be open and not previously authorized. The GitHub installation must be properly configured and have access to both the repository and pull request for this endpoint to work.",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := app.ClientForProject("")
			if err != nil {
				return err
			}
			service := vcs.New(client)

			result, err := service.UpdateExternalDeployments(installationId, repositoryId, providerPullRequestId)
			if err != nil {
				return err
			}

			return app.Render(result)
		},
	}

	cmd.Flags().StringVar(&installationId, "installation-id", "", "Installation Id")
	_ = cmd.MarkFlagRequired("installation-id")
	cmd.Flags().StringVar(&repositoryId, "repository-id", "", "VCS Repository Id")
	_ = cmd.MarkFlagRequired("repository-id")
	cmd.Flags().StringVar(&providerPullRequestId, "provider-pull-request-id", "", "GitHub Pull Request Id")
	_ = cmd.MarkFlagRequired("provider-pull-request-id")
	return cmd
}

func newVcsListInstallationsCommand() *cobra.Command {
	var queries []string
	var search string
	var total bool
	var filter []string
	var where []string
	var sortAsc []string
	var sortDesc []string
	var limit int
	var offset int
	var cursorAfter string
	var cursorBefore string

	cmd := &cobra.Command{
		Use:   "list-installations",
		Short: "List all VCS installations configured for the current project. This endpoint returns a list of installations including their provider, organization, and other configuration details.\n",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := app.ClientForProject("")
			if err != nil {
				return err
			}
			service := vcs.New(client)

			parsedFilter, err := query.ParseFilters(filter)
			if err != nil {
				return err
			}
			parsedWhere, err := query.ParseFilters(where)
			if err != nil {
				return err
			}

			queries, err := query.Build(query.Options{
				Queries:      queries,
				Filter:       parsedFilter,
				Where:        parsedWhere,
				SortAsc:      sortAsc,
				SortDesc:     sortDesc,
				Limit:        app.FlagInt(cmd, "limit", limit),
				Offset:       app.FlagInt(cmd, "offset", offset),
				CursorAfter:  app.FlagString(cmd, "cursor-after", cursorAfter),
				CursorBefore: app.FlagString(cmd, "cursor-before", cursorBefore),
			})
			if err != nil {
				return err
			}

			// An unset flag must be omitted, not sent as its zero value: the
			// TypeScript passes undefined and the SDK drops it.
			options := []vcs.ListInstallationsOption{}
			if cmd.Flags().Changed("queries") {
				options = append(options, service.WithListInstallationsQueries(queries))
			}
			if cmd.Flags().Changed("search") {
				options = append(options, service.WithListInstallationsSearch(search))
			}
			if cmd.Flags().Changed("total") {
				options = append(options, service.WithListInstallationsTotal(total))
			}

			result, err := service.ListInstallations(options...)
			if err != nil {
				return err
			}

			return app.Render(result)
		},
	}

	cmd.Flags().StringArrayVar(&queries, "queries", nil, "Array of query strings generated using the Query class provided by the SDK. Learn more about queries (https://appwrite.io/docs/queries). Maximum of 100 queries are allowed, each 4096 characters long. You may filter on the following attributes: provider, organization")
	cmd.Flags().StringVar(&search, "search", "", "Search term to filter your list results. Max length: 256 chars.")
	cmd.Flags().BoolVar(&total, "total", false, "When set to false, the total count returned will be 0 and will not be calculated.")
	cmd.Flags().Lookup("total").NoOptDefVal = "true"
	cmd.Flags().StringArrayVar(&filter, "filter", nil, "Filter using a simple comparison expression. Repeat for multiple filters. Supports field=value, field!=value, field>value, field>=value, field<value, and field<=value.")
	cmd.Flags().StringArrayVar(&where, "where", nil, "Deprecated. Use --filter instead. Filter using a simple comparison expression. Repeat for multiple filters.")
	cmd.Flags().StringArrayVar(&sortAsc, "sort-asc", nil, "Sort results by an attribute in ascending order. Repeat for multiple sort fields.")
	cmd.Flags().StringArrayVar(&sortDesc, "sort-desc", nil, "Sort results by an attribute in descending order. Repeat for multiple sort fields.")
	cmd.Flags().IntVar(&limit, "limit", 0, "Maximum number of results to return.")
	cmd.Flags().IntVar(&offset, "offset", 0, "Number of results to skip.")
	cmd.Flags().StringVar(&cursorAfter, "cursor-after", "", "Return results after this cursor ID.")
	cmd.Flags().StringVar(&cursorBefore, "cursor-before", "", "Return results before this cursor ID.")
	return cmd
}

func newVcsGetInstallationCommand() *cobra.Command {
	var installationId string

	cmd := &cobra.Command{
		Use:   "get-installation",
		Short: "Get a VCS installation by its unique ID. This endpoint returns the installation's details including its provider, organization, and configuration. ",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := app.ClientForProject("")
			if err != nil {
				return err
			}
			service := vcs.New(client)

			result, err := service.GetInstallation(installationId)
			if err != nil {
				return err
			}

			return app.Render(result)
		},
	}

	cmd.Flags().StringVar(&installationId, "installation-id", "", "Installation Id")
	_ = cmd.MarkFlagRequired("installation-id")
	return cmd
}

func newVcsDeleteInstallationCommand() *cobra.Command {
	var installationId string

	cmd := &cobra.Command{
		Use:   "delete-installation",
		Short: "Delete a VCS installation by its unique ID. This endpoint removes the installation and all its associated repositories from the project.",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := app.ClientForProject("")
			if err != nil {
				return err
			}
			service := vcs.New(client)

			result, err := service.DeleteInstallation(installationId)
			if err != nil {
				return err
			}

			return app.Render(result)
		},
	}

	cmd.Flags().StringVar(&installationId, "installation-id", "", "Installation Id")
	_ = cmd.MarkFlagRequired("installation-id")
	return cmd
}

func newVcsListNamespacesCommand() *cobra.Command {
	var installationId string
	var search string
	var queries []string
	var limit int
	var offset int

	cmd := &cobra.Command{
		Use:   "list-namespaces",
		Short: "List provider namespaces available to a VCS installation. This can include the user personal namespace and any groups or organizations the installation can browse.",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := app.ClientForProject("")
			if err != nil {
				return err
			}
			service := vcs.New(client)

			queries, err := query.Build(query.Options{
				Queries: queries,
				Limit:   app.FlagInt(cmd, "limit", limit),
				Offset:  app.FlagInt(cmd, "offset", offset),
			})
			if err != nil {
				return err
			}

			// An unset flag must be omitted, not sent as its zero value: the
			// TypeScript passes undefined and the SDK drops it.
			options := []vcs.ListNamespacesOption{}
			if cmd.Flags().Changed("search") {
				options = append(options, service.WithListNamespacesSearch(search))
			}
			if cmd.Flags().Changed("queries") {
				options = append(options, service.WithListNamespacesQueries(queries))
			}

			result, err := service.ListNamespaces(installationId, options...)
			if err != nil {
				return err
			}

			return app.Render(result)
		},
	}

	cmd.Flags().StringVar(&installationId, "installation-id", "", "Installation Id")
	_ = cmd.MarkFlagRequired("installation-id")
	cmd.Flags().StringVar(&search, "search", "", "Search term to filter your list results. Max length: 256 chars.")
	cmd.Flags().StringArrayVar(&queries, "queries", nil, "Array of query strings generated using the Query class provided by the SDK. Learn more about queries (https://appwrite.io/docs/queries). Only supported methods are limit and offset")
	cmd.Flags().IntVar(&limit, "limit", 0, "Maximum number of results to return.")
	cmd.Flags().IntVar(&offset, "offset", 0, "Number of results to skip.")
	return cmd
}
