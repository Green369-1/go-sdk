package teams

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemWithTeam_PatchRequestBody 
type ItemWithTeam_PatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The description of the team.
    description *string
    // The name of the team.
    name *string
    // The ID of a team to set as the parent team.
    parent_team_id *int32
}
// NewItemWithTeam_PatchRequestBody instantiates a new ItemWithTeam_PatchRequestBody and sets the default values.
func NewItemWithTeam_PatchRequestBody()(*ItemWithTeam_PatchRequestBody) {
    m := &ItemWithTeam_PatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemWithTeam_PatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemWithTeam_PatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemWithTeam_PatchRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemWithTeam_PatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDescription gets the description property value. The description of the team.
func (m *ItemWithTeam_PatchRequestBody) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemWithTeam_PatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
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
    res["parent_team_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentTeamId(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name of the team.
func (m *ItemWithTeam_PatchRequestBody) GetName()(*string) {
    return m.name
}
// GetParentTeamId gets the parent_team_id property value. The ID of a team to set as the parent team.
func (m *ItemWithTeam_PatchRequestBody) GetParentTeamId()(*int32) {
    return m.parent_team_id
}
// Serialize serializes information the current object
func (m *ItemWithTeam_PatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("parent_team_id", m.GetParentTeamId())
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
func (m *ItemWithTeam_PatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDescription sets the description property value. The description of the team.
func (m *ItemWithTeam_PatchRequestBody) SetDescription(value *string)() {
    m.description = value
}
// SetName sets the name property value. The name of the team.
func (m *ItemWithTeam_PatchRequestBody) SetName(value *string)() {
    m.name = value
}
// SetParentTeamId sets the parent_team_id property value. The ID of a team to set as the parent team.
func (m *ItemWithTeam_PatchRequestBody) SetParentTeamId(value *int32)() {
    m.parent_team_id = value
}
// ItemWithTeam_PatchRequestBodyable 
type ItemWithTeam_PatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetName()(*string)
    GetParentTeamId()(*int32)
    SetDescription(value *string)()
    SetName(value *string)()
    SetParentTeamId(value *int32)()
}
