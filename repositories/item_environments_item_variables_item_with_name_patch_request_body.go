package repositories

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemEnvironmentsItemVariablesItemWithNamePatchRequestBody 
type ItemEnvironmentsItemVariablesItemWithNamePatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The name of the variable.
    name *string
    // The value of the variable.
    value *string
}
// NewItemEnvironmentsItemVariablesItemWithNamePatchRequestBody instantiates a new ItemEnvironmentsItemVariablesItemWithNamePatchRequestBody and sets the default values.
func NewItemEnvironmentsItemVariablesItemWithNamePatchRequestBody()(*ItemEnvironmentsItemVariablesItemWithNamePatchRequestBody) {
    m := &ItemEnvironmentsItemVariablesItemWithNamePatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemEnvironmentsItemVariablesItemWithNamePatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemEnvironmentsItemVariablesItemWithNamePatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemEnvironmentsItemVariablesItemWithNamePatchRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemEnvironmentsItemVariablesItemWithNamePatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemEnvironmentsItemVariablesItemWithNamePatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name of the variable.
func (m *ItemEnvironmentsItemVariablesItemWithNamePatchRequestBody) GetName()(*string) {
    return m.name
}
// GetValue gets the value property value. The value of the variable.
func (m *ItemEnvironmentsItemVariablesItemWithNamePatchRequestBody) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *ItemEnvironmentsItemVariablesItemWithNamePatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemEnvironmentsItemVariablesItemWithNamePatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetName sets the name property value. The name of the variable.
func (m *ItemEnvironmentsItemVariablesItemWithNamePatchRequestBody) SetName(value *string)() {
    m.name = value
}
// SetValue sets the value property value. The value of the variable.
func (m *ItemEnvironmentsItemVariablesItemWithNamePatchRequestBody) SetValue(value *string)() {
    m.value = value
}
// ItemEnvironmentsItemVariablesItemWithNamePatchRequestBodyable 
type ItemEnvironmentsItemVariablesItemWithNamePatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    GetValue()(*string)
    SetName(value *string)()
    SetValue(value *string)()
}
