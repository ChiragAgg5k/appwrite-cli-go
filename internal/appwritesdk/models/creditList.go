package models

import (
    "encoding/json"
    "errors"
)

// CreditList Model
type CreditList struct {
    // Credits
    Credits []Credit `json:"credits"`
    // Total number of credits
    Total int `json:"total"`
    // Total available credit balance in USD
    Available float64 `json:"available"`

    // Used by Decode() method
    data []byte
}

func (model CreditList) New(data []byte) *CreditList {
    model.data = data
    return &model
}

func (model *CreditList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}