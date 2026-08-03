package models

import (
    "encoding/json"
    "errors"
)

// Credit Model
type Credit struct {
    // Credit ID.
    Id string `json:"$id"`
    // Credit creation time in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Credit update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Credit permissions. [Learn more about permissions](/docs/permissions).
    Permissions []string `json:"$permissions"`
    // coupon ID
    CouponId string `json:"couponId"`
    // ID of the User.
    UserId string `json:"userId"`
    // ID of the Team.
    TeamId string `json:"teamId"`
    // Provided credit amount
    Credits float64 `json:"credits"`
    // Provided credit amount
    Total float64 `json:"total"`
    // Credit expiration time in ISO 8601 format.
    Expiration string `json:"expiration"`
    // Status of the credit. Can be one of `disabled`, `active` or `expired`.
    Status string `json:"status"`

    // Used by Decode() method
    data []byte
}

func (model Credit) New(data []byte) *Credit {
    model.data = data
    return &model
}

func (model *Credit) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}