package keyvaluepackage

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/commons"
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
	Bid         UnsetMode
	Scope       UnsetMode
	Et          UnsetMode
	Category    UnsetMode
	IsSystem    UnsetMode
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
func WithBidUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Bid = m
	}
}
func WithScopeUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Scope = m
	}
}
func WithEtUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Et = m
	}
}
func WithCategoryUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Category = m
	}
}
func WithIsSystemUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.IsSystem = m
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
	ud.setOrUnset_bid(obj.Bid, uo.ResolveUnsetMode(uo.Bid))
	ud.setOrUnsetScope(obj.Scope, uo.ResolveUnsetMode(uo.Scope))
	ud.setOrUnset_et(obj.Et, uo.ResolveUnsetMode(uo.Et))
	ud.setOrUnsetCategory(obj.Category, uo.ResolveUnsetMode(uo.Category))
	ud.setOrUnsetIs_system(obj.IsSystem, uo.ResolveUnsetMode(uo.IsSystem))
	ud.setOrUnsetDescription(obj.Description, uo.ResolveUnsetMode(uo.Description))
	ud.setOrUnsetInherited(obj.Inherited, uo.ResolveUnsetMode(uo.Inherited))
	ud.setOrUnsetProperties(obj.Properties, uo.ResolveUnsetMode(uo.Properties))
	ud.setOrUnsetSys_info(&obj.SysInfo, uo.ResolveUnsetMode(uo.SysInfo))

	return ud
}

// SetOId No Remarks
func (ud *UpdateDocument) SetOId(p bson.ObjectID) *UpdateDocument {
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
func (ud *UpdateDocument) setOrUnsetOId(p bson.ObjectID, um UnsetMode) {
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

// Set_bid No Remarks
func (ud *UpdateDocument) Set_bid(p string) *UpdateDocument {
	mName := fmt.Sprintf(BidFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// Unset_bid No Remarks
func (ud *UpdateDocument) Unset_bid() *UpdateDocument {
	mName := fmt.Sprintf(BidFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnset_bid No Remarks
func (ud *UpdateDocument) setOrUnset_bid(p string, um UnsetMode) {
	if p != "" {
		ud.Set_bid(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.Unset_bid()
		case SetData2Default:
			ud.Unset_bid()
		}
	}
}

func UpdateWith_bid(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.Set_bid(p)
		} else {
			ud.Unset_bid()
		}
	}
}

// @tpm-schematics:start-region("-bid-field-update-section")
// @tpm-schematics:end-region("-bid-field-update-section")

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

// Set_et No Remarks
func (ud *UpdateDocument) Set_et(p string) *UpdateDocument {
	mName := fmt.Sprintf(EtFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// Unset_et No Remarks
func (ud *UpdateDocument) Unset_et() *UpdateDocument {
	mName := fmt.Sprintf(EtFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnset_et No Remarks
func (ud *UpdateDocument) setOrUnset_et(p string, um UnsetMode) {
	if p != "" {
		ud.Set_et(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.Unset_et()
		case SetData2Default:
			ud.Unset_et()
		}
	}
}

func UpdateWith_et(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.Set_et(p)
		} else {
			ud.Unset_et()
		}
	}
}

// @tpm-schematics:start-region("-et-field-update-section")
// @tpm-schematics:end-region("-et-field-update-section")

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

// SetIs_system No Remarks
func (ud *UpdateDocument) SetIs_system(p bool) *UpdateDocument {
	mName := fmt.Sprintf(IsSystemFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetIs_system No Remarks
func (ud *UpdateDocument) UnsetIs_system() *UpdateDocument {
	mName := fmt.Sprintf(IsSystemFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetIs_system No Remarks
func (ud *UpdateDocument) setOrUnsetIs_system(p bool, um UnsetMode) {
	if p {
		ud.SetIs_system(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetIs_system()
		case SetData2Default:
			ud.UnsetIs_system()
		}
	}
}

func UpdateWithIs_system(p bool) UpdateOption {
	return func(ud *UpdateDocument) {
		if p {
			ud.SetIs_system(p)
		} else {
			ud.UnsetIs_system()
		}
	}
}

// @tpm-schematics:start-region("is-system-field-update-section")
// @tpm-schematics:end-region("is-system-field-update-section")

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

// SetSys_info No Remarks
func (ud *UpdateDocument) SetSys_info(p *commons.SysInfo) *UpdateDocument {
	mName := fmt.Sprintf(SysInfoFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetSys_info No Remarks
func (ud *UpdateDocument) UnsetSys_info() *UpdateDocument {
	mName := fmt.Sprintf(SysInfoFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetSys_info No Remarks - here2
func (ud *UpdateDocument) setOrUnsetSys_info(p *commons.SysInfo, um UnsetMode) {
	if p != nil && !p.IsZero() {
		ud.SetSys_info(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetSys_info()
		case SetData2Default:
			ud.UnsetSys_info()
		}
	}
}

func UpdateWithSys_info(p *commons.SysInfo) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != nil && !p.IsZero() {
			ud.SetSys_info(p)
		} else {
			ud.UnsetSys_info()
		}
	}
}

// @tpm-schematics:start-region("sys-info-field-update-section")
// @tpm-schematics:end-region("sys-info-field-update-section")

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
