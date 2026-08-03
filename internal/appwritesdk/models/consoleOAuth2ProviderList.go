package models

import (
    "encoding/json"
    "errors"
)

// ConsoleOAuth2ProvidersList Model
type ConsoleOAuth2ProviderList struct {
    // Total number of OAuth2 providers exposed by the server.
    Total int `json:"total"`
    // List of OAuth2 providers, each with the parameters required to configure
    // it.
    OAuth2Providers []ConsoleOAuth2Provider `json:"oAuth2Providers"`

    // Used by Decode() method
    data []byte
}

func (model ConsoleOAuth2ProviderList) New(data []byte) *ConsoleOAuth2ProviderList {
    model.data = data
    return &model
}

func (model *ConsoleOAuth2ProviderList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}