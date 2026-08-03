package models

import (
    "encoding/json"
    "errors"
)

// VcsNamespace Model
type VcsNamespace struct {
    // VCS (Version Control System) namespace ID.
    Id string `json:"$id"`
    // VCS (Version Control System) namespace display name.
    Name string `json:"name"`
    // VCS (Version Control System) namespace path, used to filter repositories by
    // namespace.
    Path string `json:"path"`
    // Namespace type. Either the user's personal namespace or a
    // group/organization.
    Type string `json:"type"`
    // Namespace avatar URL.
    AvatarUrl string `json:"avatarUrl"`

    // Used by Decode() method
    data []byte
}

func (model VcsNamespace) New(data []byte) *VcsNamespace {
    model.data = data
    return &model
}

func (model *VcsNamespace) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}