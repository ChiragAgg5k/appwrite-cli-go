package appwrite

import (
	"time"

	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/account"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/activities"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/advisor"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/affiliates"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/apps"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/assistant"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/avatars"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/backups"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/client"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/console"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/databases"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/documentsdb"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/domains"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/embeddings"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/functions"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/graphql"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/locale"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/manager"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/messaging"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/migrations"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/mongo"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/mysql"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/notifications"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/oauth2"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/organization"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/organizations"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/postgresql"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/presences"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/project"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/projects"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/proxy"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/sites"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/storage"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/tablesdb"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/teams"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/tokens"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/usage"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/users"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/vcs"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/vectorsdb"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/waf"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/appwritesdk/webhooks"
)

func NewAccount(clt client.Client) *account.Account {
	return account.New(clt)
}
func NewActivities(clt client.Client) *activities.Activities {
	return activities.New(clt)
}
func NewAffiliates(clt client.Client) *affiliates.Affiliates {
	return affiliates.New(clt)
}
func NewApps(clt client.Client) *apps.Apps {
	return apps.New(clt)
}
func NewAvatars(clt client.Client) *avatars.Avatars {
	return avatars.New(clt)
}
func NewBackups(clt client.Client) *backups.Backups {
	return backups.New(clt)
}
func NewAssistant(clt client.Client) *assistant.Assistant {
	return assistant.New(clt)
}
func NewConsole(clt client.Client) *console.Console {
	return console.New(clt)
}
func NewDatabases(clt client.Client) *databases.Databases {
	return databases.New(clt)
}
func NewDocumentsDB(clt client.Client) *documentsdb.DocumentsDB {
	return documentsdb.New(clt)
}
func NewDomains(clt client.Client) *domains.Domains {
	return domains.New(clt)
}
func NewEmbeddings(clt client.Client) *embeddings.Embeddings {
	return embeddings.New(clt)
}
func NewFunctions(clt client.Client) *functions.Functions {
	return functions.New(clt)
}
func NewGraphql(clt client.Client) *graphql.Graphql {
	return graphql.New(clt)
}
func NewLocale(clt client.Client) *locale.Locale {
	return locale.New(clt)
}
func NewManager(clt client.Client) *manager.Manager {
	return manager.New(clt)
}
func NewMessaging(clt client.Client) *messaging.Messaging {
	return messaging.New(clt)
}
func NewMigrations(clt client.Client) *migrations.Migrations {
	return migrations.New(clt)
}
func NewMongo(clt client.Client) *mongo.Mongo {
	return mongo.New(clt)
}
func NewMysql(clt client.Client) *mysql.Mysql {
	return mysql.New(clt)
}
func NewNotifications(clt client.Client) *notifications.Notifications {
	return notifications.New(clt)
}
func NewOauth2(clt client.Client) *oauth2.Oauth2 {
	return oauth2.New(clt)
}
func NewOrganization(clt client.Client) *organization.Organization {
	return organization.New(clt)
}
func NewOrganizations(clt client.Client) *organizations.Organizations {
	return organizations.New(clt)
}
func NewPostgresql(clt client.Client) *postgresql.Postgresql {
	return postgresql.New(clt)
}
func NewPresences(clt client.Client) *presences.Presences {
	return presences.New(clt)
}
func NewProject(clt client.Client) *project.Project {
	return project.New(clt)
}
func NewProjects(clt client.Client) *projects.Projects {
	return projects.New(clt)
}
func NewProxy(clt client.Client) *proxy.Proxy {
	return proxy.New(clt)
}
func NewAdvisor(clt client.Client) *advisor.Advisor {
	return advisor.New(clt)
}
func NewSites(clt client.Client) *sites.Sites {
	return sites.New(clt)
}
func NewStorage(clt client.Client) *storage.Storage {
	return storage.New(clt)
}
func NewTablesDB(clt client.Client) *tablesdb.TablesDB {
	return tablesdb.New(clt)
}
func NewTeams(clt client.Client) *teams.Teams {
	return teams.New(clt)
}
func NewTokens(clt client.Client) *tokens.Tokens {
	return tokens.New(clt)
}
func NewUsage(clt client.Client) *usage.Usage {
	return usage.New(clt)
}
func NewUsers(clt client.Client) *users.Users {
	return users.New(clt)
}
func NewVcs(clt client.Client) *vcs.Vcs {
	return vcs.New(clt)
}
func NewVectorsDB(clt client.Client) *vectorsdb.VectorsDB {
	return vectorsdb.New(clt)
}
func NewWaf(clt client.Client) *waf.Waf {
	return waf.New(clt)
}
func NewWebhooks(clt client.Client) *webhooks.Webhooks {
	return webhooks.New(clt)
}

// NewClient initializes a new Appwrite client with a given timeout
func NewClient(optionalSetters ...client.ClientOption) client.Client {
	return client.New(optionalSetters...)
}

// Helper method to construct NewClient()
func WithEndpoint(endpoint string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Endpoint = endpoint
		return nil
	}
}

// Helper method to construct NewClient()
func WithTimeout(timeout time.Duration) client.ClientOption {
	return func(clt *client.Client) error {
		httpClient, err := client.GetDefaultClient(timeout)
		if err != nil {
			return err
		}

		clt.Timeout = timeout
		clt.Client = httpClient

		return nil
	}
}

// Helper method to construct NewClient()
func WithSelfSigned(status bool) client.ClientOption {
	return func(clt *client.Client) error {
		clt.SelfSigned = status
		return nil
	}
}

// Helper method to construct NewClient()
func WithChunkSize(size int64) client.ClientOption {
	return func(clt *client.Client) error {
		clt.ChunkSize = size
		return nil
	}
}

// Helper method to construct NewClient()
//
// Your project ID
func WithProject(value string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Config["project"] = value
		return nil
	}
}

// Helper method to construct NewClient()
//
// Your secret API key
func WithKey(value string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Config["key"] = value
		clt.Headers["X-Appwrite-Key"] = value
		return nil
	}
}

// Helper method to construct NewClient()
//
// Your organization ID
func WithOrganization(value string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Config["organization"] = value
		clt.Headers["X-Appwrite-Organization"] = value
		return nil
	}
}

// Helper method to construct NewClient()
//
// Your secret JSON Web Token
func WithJWT(value string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Config["jwt"] = value
		clt.Headers["X-Appwrite-JWT"] = value
		return nil
	}
}

// Helper method to construct NewClient()
//
// The OAuth access token to authenticate with
func WithBearer(value string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Config["bearer"] = value
		clt.Headers["Authorization"] = "Bearer " + value
		return nil
	}
}

// Helper method to construct NewClient()
func WithLocale(value string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Config["locale"] = value
		clt.Headers["X-Appwrite-Locale"] = value
		return nil
	}
}

// Helper method to construct NewClient()
func WithMode(value string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Config["mode"] = value
		clt.Headers["X-Appwrite-Mode"] = value
		return nil
	}
}

// Helper method to construct NewClient()
//
// The user cookie to authenticate with. Used by SDKs that forward an incoming Cookie header in server-side runtimes.
func WithCookie(value string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Config["cookie"] = value
		clt.Headers["Cookie"] = value
		return nil
	}
}

// Helper method to construct NewClient()
//
// The user session to authenticate with
func WithSession(value string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Config["session"] = value
		clt.Headers["X-Appwrite-Session"] = value
		return nil
	}
}

// Helper method to construct NewClient()
//
// Your secret dev API key
func WithDevKey(value string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Config["devkey"] = value
		clt.Headers["X-Appwrite-Dev-Key"] = value
		return nil
	}
}

// Helper method to construct NewClient()
//
// Impersonate a user by ID
func WithImpersonateUserId(value string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Config["impersonateuserid"] = value
		clt.Headers["X-Appwrite-Impersonate-User-Id"] = value
		return nil
	}
}

// Helper method to construct NewClient()
//
// Impersonate a user by email
func WithImpersonateUserEmail(value string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Config["impersonateuseremail"] = value
		clt.Headers["X-Appwrite-Impersonate-User-Email"] = value
		return nil
	}
}

// Helper method to construct NewClient()
//
// Impersonate a user by phone
func WithImpersonateUserPhone(value string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Config["impersonateuserphone"] = value
		clt.Headers["X-Appwrite-Impersonate-User-Phone"] = value
		return nil
	}
}

// Helper method to construct NewClient()
//
// The platform type (Appwrite or Imagine)
func WithPlatform(value string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Config["platform"] = value
		clt.Headers["X-Appwrite-Platform"] = value
		return nil
	}
}
