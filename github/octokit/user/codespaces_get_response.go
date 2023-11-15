package user

import (
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CodespacesGetResponse 
type CodespacesGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The codespaces property
    codespaces []i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Codespaceable
    // The total_count property
    total_count *int32
}
// NewCodespacesGetResponse instantiates a new CodespacesGetResponse and sets the default values.
func NewCodespacesGetResponse()(*CodespacesGetResponse) {
    m := &CodespacesGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodespacesGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCodespacesGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodespacesGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CodespacesGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCodespaces gets the codespaces property value. The codespaces property
func (m *CodespacesGetResponse) GetCodespaces()([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Codespaceable) {
    return m.codespaces
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CodespacesGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["codespaces"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateCodespaceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Codespaceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Codespaceable)
                }
            }
            m.SetCodespaces(res)
        }
        return nil
    }
    res["total_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalCount(val)
        }
        return nil
    }
    return res
}
// GetTotalCount gets the total_count property value. The total_count property
func (m *CodespacesGetResponse) GetTotalCount()(*int32) {
    return m.total_count
}
// Serialize serializes information the current object
func (m *CodespacesGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCodespaces() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCodespaces()))
        for i, v := range m.GetCodespaces() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("codespaces", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_count", m.GetTotalCount())
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
func (m *CodespacesGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCodespaces sets the codespaces property value. The codespaces property
func (m *CodespacesGetResponse) SetCodespaces(value []i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Codespaceable)() {
    m.codespaces = value
}
// SetTotalCount sets the total_count property value. The total_count property
func (m *CodespacesGetResponse) SetTotalCount(value *int32)() {
    m.total_count = value
}
// CodespacesGetResponseable 
type CodespacesGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCodespaces()([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Codespaceable)
    GetTotalCount()(*int32)
    SetCodespaces(value []i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Codespaceable)()
    SetTotalCount(value *int32)()
}