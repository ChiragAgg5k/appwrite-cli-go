package models

import (
    "encoding/json"
    "errors"
)

// ConsoleVariables Model
type ConsoleVariables struct {
    // CNAME target for your Appwrite custom domains.
    APPDOMAINTARGETCNAME string `json:"_APP_DOMAIN_TARGET_CNAME"`
    // A target for your Appwrite custom domains.
    APPDOMAINTARGETA string `json:"_APP_DOMAIN_TARGET_A"`
    // Maximum build timeout in seconds.
    APPCOMPUTEBUILDTIMEOUT int `json:"_APP_COMPUTE_BUILD_TIMEOUT"`
    // AAAA target for your Appwrite custom domains.
    APPDOMAINTARGETAAAA string `json:"_APP_DOMAIN_TARGET_AAAA"`
    // CAA target for your Appwrite custom domains.
    APPDOMAINTARGETCAA string `json:"_APP_DOMAIN_TARGET_CAA"`
    // Maximum file size allowed for file upload in bytes.
    APPSTORAGELIMIT int `json:"_APP_STORAGE_LIMIT"`
    // Maximum file size allowed for deployment in bytes.
    APPCOMPUTESIZELIMIT int `json:"_APP_COMPUTE_SIZE_LIMIT"`
    // Defines if usage stats are enabled. This value is set to 'enabled' by
    // default, to disable the usage stats set the value to 'disabled'.
    APPUSAGESTATS string `json:"_APP_USAGE_STATS"`
    // Defines if VCS (Version Control System) is enabled.
    APPVCSENABLED bool `json:"_APP_VCS_ENABLED"`
    // List of configured VCS providers.
    APPVCSPROVIDERS []string `json:"_APP_VCS_PROVIDERS"`
    // Defines if main domain is configured. If so, custom domains can be created.
    APPDOMAINENABLED bool `json:"_APP_DOMAIN_ENABLED"`
    // Defines if AI assistant is enabled.
    APPASSISTANTENABLED bool `json:"_APP_ASSISTANT_ENABLED"`
    // A comma separated list of domains to use for site URLs.
    APPDOMAINSITES string `json:"_APP_DOMAIN_SITES"`
    // A domain to use for function URLs.
    APPDOMAINFUNCTIONS string `json:"_APP_DOMAIN_FUNCTIONS"`
    // Defines if HTTPS is enforced for all requests.
    APPOPTIONSFORCEHTTPS string `json:"_APP_OPTIONS_FORCE_HTTPS"`
    // Comma-separated list of nameservers.
    APPDOMAINSNAMESERVERS string `json:"_APP_DOMAINS_NAMESERVERS"`
    // Database adapter in use.
    APPDBADAPTER string `json:"_APP_DB_ADAPTER"`
    // Whether the database adapter supports relationships.
    SupportForRelationships bool `json:"supportForRelationships"`
    // Whether the database adapter supports operators.
    SupportForOperators bool `json:"supportForOperators"`
    // Whether the database adapter supports spatial attributes.
    SupportForSpatials bool `json:"supportForSpatials"`
    // Whether the database adapter supports spatial indexes on nullable columns.
    SupportForSpatialIndexNull bool `json:"supportForSpatialIndexNull"`
    // Whether the database adapter supports fulltext wildcard search.
    SupportForFulltextWildcard bool `json:"supportForFulltextWildcard"`
    // Whether the database adapter supports multiple fulltext indexes per
    // collection.
    SupportForMultipleFulltextIndexes bool `json:"supportForMultipleFulltextIndexes"`
    // Whether the database adapter supports resizing attributes.
    SupportForAttributeResizing bool `json:"supportForAttributeResizing"`
    // Whether the database adapter supports fixed schemas with row width limits.
    SupportForSchemas bool `json:"supportForSchemas"`
    // Maximum index length supported by the database adapter.
    MaxIndexLength int `json:"maxIndexLength"`
    // Whether the database adapter uses integer sequence IDs.
    SupportForIntegerIds bool `json:"supportForIntegerIds"`
    // Whether email verification for console users is required. Can be "true" or
    // "false".
    APPCONSOLEEMAILVERIFICATION string `json:"_APP_CONSOLE_EMAIL_VERIFICATION"`

    // Used by Decode() method
    data []byte
}

func (model ConsoleVariables) New(data []byte) *ConsoleVariables {
    model.data = data
    return &model
}

func (model *ConsoleVariables) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}