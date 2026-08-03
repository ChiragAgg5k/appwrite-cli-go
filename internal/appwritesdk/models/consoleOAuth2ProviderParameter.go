package models

import (
    "encoding/json"
    "errors"
)

// ConsoleOAuth2ProviderParameter Model
type ConsoleOAuth2ProviderParameter struct {
    // Parameter ID. Maps to the request body field used by the project OAuth2
    // update endpoint (e.g. `clientId`, `appKey`, `tenant`).
    Id string `json:"$id"`
    // Verbose, user-facing parameter name as shown in the provider's own
    // dashboard. Includes alternate names when the provider exposes more than
    // one.
    Name string `json:"name"`
    // Example value for this parameter.
    Example string `json:"example"`
    // Optional hint for this parameter, typically calling out a common wrong
    // value. Empty string when no hint is set.
    Hint string `json:"hint"`

    // Used by Decode() method
    data []byte
}

func (model ConsoleOAuth2ProviderParameter) New(data []byte) *ConsoleOAuth2ProviderParameter {
    model.data = data
    return &model
}

func (model *ConsoleOAuth2ProviderParameter) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}