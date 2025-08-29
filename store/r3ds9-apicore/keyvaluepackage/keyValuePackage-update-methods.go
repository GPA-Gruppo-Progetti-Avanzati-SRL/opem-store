package keyvaluepackage

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"

	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/commons"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

func UpdateMethodsGoInfo() string {
	i := fmt.Sprintf("tpm_morphia query filter support generated for %s package on %s", "author", time.Now().String())
	return i
}

type UnsetMode int64

const (
	UnSpecified     UnsetMode = 0
	KeepCurrent               = 1
	UnsetData                 = 2
	SetData2Default           = 3
)

type UnsetOption func(uopt *UnsetOptions)

type UnsetOptions struct {
	DefaultMode UnsetMode
	OId         UnsetMode
	Name        UnsetMode
	Scope       UnsetMode
	ObjType     UnsetMode
	Category    UnsetMode
	Issystem    UnsetMode
	Description UnsetMode
	Inherited   UnsetMode
	Properties  UnsetMode
	SysInfo     UnsetMode
}

func (uo *UnsetOptions) ResolveUnsetMode(um UnsetMode) UnsetMode {
	if um == UnSpecified {
		um = uo.DefaultMode
	}

	return um
}

func WithDefaultUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.DefaultMode = m
	}
}
func WithOIdUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.OId = m
	}
}
func WithNameUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Name = m
	}
}
func WithScopeUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Scope = m
	}
}
func WithObjTypeUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.ObjType = m
	}
}
func WithCategoryUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Category = m
	}
}
func WithIssystemUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Issystem = m
	}
}
func WithDescriptionUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Description = m
	}
}
func WithInheritedUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Inherited = m
	}
}
func WithPropertiesUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Properties = m
	}
}
func WithSysInfoUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.SysInfo = m
	}
}

type UpdateOption func(ud *UpdateDocument)
type UpdateOptions []UpdateOption

// GetUpdateDocumentFromOptions convenience method to create an update document from single updates instead of a whole object
func GetUpdateDocumentFromOptions(opts ...UpdateOption) UpdateDocument {
	ud := UpdateDocument{}
	for _, o := range opts {
		o(&ud)
	}

	return ud
}

// GetUpdateDocument
// Convenience method to create an Update Document from the values of the top fields of the object. The convenience is in the handling
// the unset because if I pass an empty struct to the update it generates an empty object anyway in the db. Handling the unset eliminates
// the issue and delete an existing value without creating an empty struct.
func GetUpdateDocument(obj *KeyValuePackage, opts ...UnsetOption) UpdateDocument {

	uo := &UnsetOptions{DefaultMode: KeepCurrent}
	for _, o := range opts {
		o(uo)
	}

	ud := UpdateDocument{}
	ud.setOrUnsetName(obj.Name, uo.ResolveUnsetMode(uo.Name))
	ud.setOrUnsetScope(obj.Scope, uo.ResolveUnsetMode(uo.Scope))
	ud.setOrUnsetObjType(obj.ObjType, uo.ResolveUnsetMode(uo.ObjType))
	ud.setOrUnsetCategory(obj.Category, uo.ResolveUnsetMode(uo.Category))
	ud.setOrUnsetIssystem(obj.Issystem, uo.ResolveUnsetMode(uo.Issystem))
	ud.setOrUnsetDescription(obj.Description, uo.ResolveUnsetMode(uo.Description))
	ud.setOrUnsetInherited(obj.Inherited, uo.ResolveUnsetMode(uo.Inherited))
	ud.setOrUnsetProperties(obj.Properties, uo.ResolveUnsetMode(uo.Properties))
	ud.setOrUnsetSysInfo(&obj.SysInfo, uo.ResolveUnsetMode(uo.SysInfo))

	return ud
}

// SetOId No Remarks
func (ud *UpdateDocument) SetOId(p primitive.ObjectID) *UpdateDocument {
	mName := fmt.Sprintf(OIdFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetOId No Remarks
func (ud *UpdateDocument) UnsetOId() *UpdateDocument {
	mName := fmt.Sprintf(OIdFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetOId No Remarks
func (ud *UpdateDocument) setOrUnsetOId(p primitive.ObjectID, um UnsetMode) {
	if !p.IsZero() {
		ud.SetOId(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetOId()
		case SetData2Default:
			ud.UnsetOId()
		}
	}
}

// @tpm-schematics:start-region("o-id-field-update-section")
// @tpm-schematics:end-region("o-id-field-update-section")

// SetName No Remarks
func (ud *UpdateDocument) SetName(p string) *UpdateDocument {
	mName := fmt.Sprintf(NameFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetName No Remarks
func (ud *UpdateDocument) UnsetName() *UpdateDocument {
	mName := fmt.Sprintf(NameFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetName No Remarks
func (ud *UpdateDocument) setOrUnsetName(p string, um UnsetMode) {
	if p != "" {
		ud.SetName(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetName()
		case SetData2Default:
			ud.UnsetName()
		}
	}
}

func UpdateWithName(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetName(p)
		} else {
			ud.UnsetName()
		}
	}
}

// @tpm-schematics:start-region("name-field-update-section")
// @tpm-schematics:end-region("name-field-update-section")

// SetScope No Remarks
func (ud *UpdateDocument) SetScope(p string) *UpdateDocument {
	mName := fmt.Sprintf(ScopeFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetScope No Remarks
func (ud *UpdateDocument) UnsetScope() *UpdateDocument {
	mName := fmt.Sprintf(ScopeFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetScope No Remarks
func (ud *UpdateDocument) setOrUnsetScope(p string, um UnsetMode) {
	if p != "" {
		ud.SetScope(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetScope()
		case SetData2Default:
			ud.UnsetScope()
		}
	}
}

func UpdateWithScope(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetScope(p)
		} else {
			ud.UnsetScope()
		}
	}
}

// @tpm-schematics:start-region("scope-field-update-section")
// @tpm-schematics:end-region("scope-field-update-section")

// SetObjType No Remarks
func (ud *UpdateDocument) SetObjType(p string) *UpdateDocument {
	mName := fmt.Sprintf(ObjTypeFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetObjType No Remarks
func (ud *UpdateDocument) UnsetObjType() *UpdateDocument {
	mName := fmt.Sprintf(ObjTypeFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetObjType No Remarks
func (ud *UpdateDocument) setOrUnsetObjType(p string, um UnsetMode) {
	if p != "" {
		ud.SetObjType(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetObjType()
		case SetData2Default:
			ud.UnsetObjType()
		}
	}
}

func UpdateWithObjType(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetObjType(p)
		} else {
			ud.UnsetObjType()
		}
	}
}

// @tpm-schematics:start-region("obj-type-field-update-section")
// @tpm-schematics:end-region("obj-type-field-update-section")

// SetCategory No Remarks
func (ud *UpdateDocument) SetCategory(p string) *UpdateDocument {
	mName := fmt.Sprintf(CategoryFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetCategory No Remarks
func (ud *UpdateDocument) UnsetCategory() *UpdateDocument {
	mName := fmt.Sprintf(CategoryFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetCategory No Remarks
func (ud *UpdateDocument) setOrUnsetCategory(p string, um UnsetMode) {
	if p != "" {
		ud.SetCategory(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetCategory()
		case SetData2Default:
			ud.UnsetCategory()
		}
	}
}

func UpdateWithCategory(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetCategory(p)
		} else {
			ud.UnsetCategory()
		}
	}
}

// @tpm-schematics:start-region("category-field-update-section")
// @tpm-schematics:end-region("category-field-update-section")

// SetIssystem No Remarks
func (ud *UpdateDocument) SetIssystem(p bool) *UpdateDocument {
	mName := fmt.Sprintf(IssystemFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetIssystem No Remarks
func (ud *UpdateDocument) UnsetIssystem() *UpdateDocument {
	mName := fmt.Sprintf(IssystemFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetIssystem No Remarks
func (ud *UpdateDocument) setOrUnsetIssystem(p bool, um UnsetMode) {
	if p {
		ud.SetIssystem(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetIssystem()
		case SetData2Default:
			ud.UnsetIssystem()
		}
	}
}

func UpdateWithIssystem(p bool) UpdateOption {
	return func(ud *UpdateDocument) {
		if p {
			ud.SetIssystem(p)
		} else {
			ud.UnsetIssystem()
		}
	}
}

// @tpm-schematics:start-region("issystem-field-update-section")
// @tpm-schematics:end-region("issystem-field-update-section")

// SetDescription No Remarks
func (ud *UpdateDocument) SetDescription(p string) *UpdateDocument {
	mName := fmt.Sprintf(DescriptionFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetDescription No Remarks
func (ud *UpdateDocument) UnsetDescription() *UpdateDocument {
	mName := fmt.Sprintf(DescriptionFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetDescription No Remarks
func (ud *UpdateDocument) setOrUnsetDescription(p string, um UnsetMode) {
	if p != "" {
		ud.SetDescription(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetDescription()
		case SetData2Default:
			ud.UnsetDescription()
		}
	}
}

func UpdateWithDescription(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetDescription(p)
		} else {
			ud.UnsetDescription()
		}
	}
}

// @tpm-schematics:start-region("description-field-update-section")
// @tpm-schematics:end-region("description-field-update-section")

// SetInherited No Remarks
func (ud *UpdateDocument) SetInherited(p bool) *UpdateDocument {
	mName := fmt.Sprintf(InheritedFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetInherited No Remarks
func (ud *UpdateDocument) UnsetInherited() *UpdateDocument {
	mName := fmt.Sprintf(InheritedFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetInherited No Remarks
func (ud *UpdateDocument) setOrUnsetInherited(p bool, um UnsetMode) {
	if p {
		ud.SetInherited(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetInherited()
		case SetData2Default:
			ud.UnsetInherited()
		}
	}
}

func UpdateWithInherited(p bool) UpdateOption {
	return func(ud *UpdateDocument) {
		if p {
			ud.SetInherited(p)
		} else {
			ud.UnsetInherited()
		}
	}
}

// @tpm-schematics:start-region("inherited-field-update-section")
// @tpm-schematics:end-region("inherited-field-update-section")

// SetProperties No Remarks
func (ud *UpdateDocument) SetProperties(p []KeyValue) *UpdateDocument {
	mName := fmt.Sprintf(PropertiesFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetProperties No Remarks
func (ud *UpdateDocument) UnsetProperties() *UpdateDocument {
	mName := fmt.Sprintf(PropertiesFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetProperties No Remarks - here2
func (ud *UpdateDocument) setOrUnsetProperties(p []KeyValue, um UnsetMode) {
	if len(p) > 0 {
		ud.SetProperties(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetProperties()
		case SetData2Default:
			ud.UnsetProperties()
		}
	}
}

func UpdateWithProperties(p []KeyValue) UpdateOption {
	return func(ud *UpdateDocument) {
		if len(p) > 0 {
			ud.SetProperties(p)
		} else {
			ud.UnsetProperties()
		}
	}
}

// @tpm-schematics:start-region("properties-field-update-section")
// @tpm-schematics:end-region("properties-field-update-section")

// SetSysInfo No Remarks
func (ud *UpdateDocument) SetSysInfo(p *commons.SysInfo) *UpdateDocument {
	mName := fmt.Sprintf(SysInfoFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetSysInfo No Remarks
func (ud *UpdateDocument) UnsetSysInfo() *UpdateDocument {
	mName := fmt.Sprintf(SysInfoFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetSysInfo No Remarks - here2
func (ud *UpdateDocument) setOrUnsetSysInfo(p *commons.SysInfo, um UnsetMode) {
	if p != nil && !p.IsZero() {
		ud.SetSysInfo(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetSysInfo()
		case SetData2Default:
			ud.UnsetSysInfo()
		}
	}
}

func UpdateWithSysInfo(p *commons.SysInfo) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != nil && !p.IsZero() {
			ud.SetSysInfo(p)
		} else {
			ud.UnsetSysInfo()
		}
	}
}

// @tpm-schematics:start-region("sys-info-field-update-section")
// @tpm-schematics:end-region("sys-info-field-update-section")

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
