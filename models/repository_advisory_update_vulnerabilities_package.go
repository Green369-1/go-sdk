package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RepositoryAdvisoryUpdate_vulnerabilities_package the name of the package affected by the vulnerability.
type RepositoryAdvisoryUpdate_vulnerabilities_package struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The package's language or package management ecosystem.
    ecosystem *SecurityAdvisoryEcosystems
    // The unique package name within its ecosystem.
    name *string
}
// NewRepositoryAdvisoryUpdate_vulnerabilities_package instantiates a new repositoryAdvisoryUpdate_vulnerabilities_package and sets the default values.
func NewRepositoryAdvisoryUpdate_vulnerabilities_package()(*RepositoryAdvisoryUpdate_vulnerabilities_package) {
    m := &RepositoryAdvisoryUpdate_vulnerabilities_package{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRepositoryAdvisoryUpdate_vulnerabilities_packageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRepositoryAdvisoryUpdate_vulnerabilities_packageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoryAdvisoryUpdate_vulnerabilities_package(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RepositoryAdvisoryUpdate_vulnerabilities_package) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEcosystem gets the ecosystem property value. The package's language or package management ecosystem.
func (m *RepositoryAdvisoryUpdate_vulnerabilities_package) GetEcosystem()(*SecurityAdvisoryEcosystems) {
    return m.ecosystem
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RepositoryAdvisoryUpdate_vulnerabilities_package) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ecosystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSecurityAdvisoryEcosystems)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEcosystem(val.(*SecurityAdvisoryEcosystems))
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
    return res
}
// GetName gets the name property value. The unique package name within its ecosystem.
func (m *RepositoryAdvisoryUpdate_vulnerabilities_package) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *RepositoryAdvisoryUpdate_vulnerabilities_package) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetEcosystem() != nil {
        cast := (*m.GetEcosystem()).String()
        err := writer.WriteStringValue("ecosystem", &cast)
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RepositoryAdvisoryUpdate_vulnerabilities_package) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEcosystem sets the ecosystem property value. The package's language or package management ecosystem.
func (m *RepositoryAdvisoryUpdate_vulnerabilities_package) SetEcosystem(value *SecurityAdvisoryEcosystems)() {
    m.ecosystem = value
}
// SetName sets the name property value. The unique package name within its ecosystem.
func (m *RepositoryAdvisoryUpdate_vulnerabilities_package) SetName(value *string)() {
    m.name = value
}
// RepositoryAdvisoryUpdate_vulnerabilities_packageable 
type RepositoryAdvisoryUpdate_vulnerabilities_packageable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEcosystem()(*SecurityAdvisoryEcosystems)
    GetName()(*string)
    SetEcosystem(value *SecurityAdvisoryEcosystems)()
    SetName(value *string)()
}
