package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemActionsWorkflowsItemDispatchesPostRequestBody_inputs input keys and values configured in the workflow file. The maximum number of properties is 10. Any default properties configured in the workflow file will be used when `inputs` are omitted.
type ItemItemActionsWorkflowsItemDispatchesPostRequestBody_inputs struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
}
// NewItemItemActionsWorkflowsItemDispatchesPostRequestBody_inputs instantiates a new ItemItemActionsWorkflowsItemDispatchesPostRequestBody_inputs and sets the default values.
func NewItemItemActionsWorkflowsItemDispatchesPostRequestBody_inputs()(*ItemItemActionsWorkflowsItemDispatchesPostRequestBody_inputs) {
    m := &ItemItemActionsWorkflowsItemDispatchesPostRequestBody_inputs{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemActionsWorkflowsItemDispatchesPostRequestBody_inputsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemActionsWorkflowsItemDispatchesPostRequestBody_inputsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemActionsWorkflowsItemDispatchesPostRequestBody_inputs(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemActionsWorkflowsItemDispatchesPostRequestBody_inputs) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemActionsWorkflowsItemDispatchesPostRequestBody_inputs) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *ItemItemActionsWorkflowsItemDispatchesPostRequestBody_inputs) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemActionsWorkflowsItemDispatchesPostRequestBody_inputs) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// ItemItemActionsWorkflowsItemDispatchesPostRequestBody_inputsable 
type ItemItemActionsWorkflowsItemDispatchesPostRequestBody_inputsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
