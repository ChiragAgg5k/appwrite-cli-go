package models

import (
    "encoding/json"
    "errors"
)

// ConsoleOAuth2Provider Model
type ConsoleOAuth2Provider struct {
    // OAuth2 provider ID.
    Id string `json:"$id"`
    // List of parameters required to configure this OAuth2 provider.
    Parameters []ConsoleOAuth2ProviderParameter `json:"parameters"`

    // Used by Decode() method
    data []byte
}

func (model ConsoleOAuth2Provider) New(data []byte) *ConsoleOAuth2Provider {
    model.data = data
    return &model
}

func (model *ConsoleOAuth2Provider) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}