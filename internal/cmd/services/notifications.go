package services

import (
	"github.com/spf13/cobra"

	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/notifications"

	"github.com/ChiragAgg5k/appwrite-cli-go/internal/app"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/query"
)

// NewNotificationsCommand builds the `notifications` command tree.
func NewNotificationsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "notifications",
		Short: "The Notifications service allows you to read and manage your Appwrite Console notifications.",
	}

	cmd.AddCommand(newNotificationsListCommand())
	cmd.AddCommand(newNotificationsUpdateCommand())

	return cmd
}

func newNotificationsListCommand() *cobra.Command {
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
		Use:   "list",
		Short: "Get the list of notifications for the currently logged in console user. Use queries to filter the results by attributes such as read status, view timestamps, or creation date.\n",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := app.ClientForConsole()
			if err != nil {
				return err
			}
			service := notifications.New(client)

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
			options := []notifications.ListOption{}
			if cmd.Flags().Changed("queries") {
				options = append(options, service.WithListQueries(queries))
			}

			result, err := service.List(options...)
			if err != nil {
				return err
			}

			return app.Render(result)
		},
	}

	cmd.Flags().StringArrayVar(&queries, "queries", nil, "Array of query strings generated using the Query class provided by the SDK. Learn more about queries (https://appwrite.io/docs/queries). Maximum of 100 queries are allowed, each 4096 characters long. You may filter on the following attributes: read, type, channel, messageId, projectId, resourceType, resourceId, parentResourceType, parentResourceId, firstSeen, lastSeen")
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

func newNotificationsUpdateCommand() *cobra.Command {
	var notificationId string
	var read bool

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a notification by its unique ID. Use the `read` parameter to mark the notification as read or unread.\n",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := app.ClientForConsole()
			if err != nil {
				return err
			}
			service := notifications.New(client)

			result, err := service.Update(notificationId, read)
			if err != nil {
				return err
			}

			return app.Render(result)
		},
	}

	cmd.Flags().StringVar(&notificationId, "notification-id", "", "Notification ID.")
	_ = cmd.MarkFlagRequired("notification-id")
	cmd.Flags().BoolVar(&read, "read", false, "Notification read status.")
	_ = cmd.MarkFlagRequired("read")
	return cmd
}
