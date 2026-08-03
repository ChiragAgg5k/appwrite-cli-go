package models

import (
    "encoding/json"
    "errors"
)

// Coupon Model
type Coupon struct {
    // coupon ID
    Id string `json:"$id"`
    // coupon ID
    Code string `json:"code"`
    // Provided credit amount
    Credits float64 `json:"credits"`
    // Coupon expiration time in ISO 8601 format.
    Expiration string `json:"expiration"`
    // Credit validity in days.
    Validity int `json:"validity"`
    // Campaign the coupon is associated with`.
    Campaign string `json:"campaign"`
    // Status of the coupon. Can be one of `disabled`, `active` or `expired`.
    Status string `json:"status"`
    // If the coupon is only valid for new organizations or not.
    OnlyNewOrgs bool `json:"onlyNewOrgs"`

    // Used by Decode() method
    data []byte
}

func (model Coupon) New(data []byte) *Coupon {
    model.data = data
    return &model
}

func (model *Coupon) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}