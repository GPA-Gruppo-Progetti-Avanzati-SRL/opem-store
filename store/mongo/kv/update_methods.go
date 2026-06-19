package kv

import (
	"fmt"
	"time"

	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/commons"
	model "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

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
	return func(uopt *UnsetOptions) { uopt.DefaultMode = m }
}
func WithOIdUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) { uopt.OId = m }
}
func WithBidUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) { uopt.Bid = m }
}
func WithScopeUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) { uopt.Scope = m }
}
func WithEtUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) { uopt.Et = m }
}
func WithCategoryUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) { uopt.Category = m }
}
func WithIsSystemUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) { uopt.IsSystem = m }
}
func WithDescriptionUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) { uopt.Description = m }
}
func WithInheritedUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) { uopt.Inherited = m }
}
func WithPropertiesUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) { uopt.Properties = m }
}
func WithSysInfoUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) { uopt.SysInfo = m }
}

type UpdateOption func(ud *UpdateDocument)
type UpdateOptions []UpdateOption

func GetUpdateDocumentFromOptions(opts ...UpdateOption) UpdateDocument {
	ud := UpdateDocument{}
	for _, o := range opts {
		o(&ud)
	}
	return ud
}

func GetUpdateDocument(obj *model.KeyValuePackage, opts ...UnsetOption) UpdateDocument {
	uo := &UnsetOptions{DefaultMode: KeepCurrent}
	for _, o := range opts {
		o(uo)
	}
	ud := UpdateDocument{}
	ud.setOrUnset_bid(obj.Bid, uo.ResolveUnsetMode(uo.Bid))
	ud.setOrUnsetScope(obj.Scope, uo.ResolveUnsetMode(uo.Scope))
	ud.setOrUnset_et(obj.Et, uo.ResolveUnsetMode(uo.Et))
	ud.setOrUnsetCategory(obj.Category, uo.ResolveUnsetMode(uo.Category))
	ud.setOrUnsetIsSystem(obj.IsSystem, uo.ResolveUnsetMode(uo.IsSystem))
	ud.setOrUnsetDescription(obj.Description, uo.ResolveUnsetMode(uo.Description))
	ud.setOrUnsetInherited(obj.Inherited, uo.ResolveUnsetMode(uo.Inherited))
	ud.setOrUnsetProperties(obj.Properties, uo.ResolveUnsetMode(uo.Properties))
	ud.setOrUnsetSysInfo(&obj.SysInfo, uo.ResolveUnsetMode(uo.SysInfo))
	return ud
}

func (ud *UpdateDocument) SetOId(p bson.ObjectID) *UpdateDocument {
	mName := fmt.Sprintf(OIdFieldName)
	ud.Set().Add(func() bson.E { return bson.E{Key: mName, Value: p} })
	return ud
}

func (ud *UpdateDocument) UnsetOId() *UpdateDocument {
	mName := fmt.Sprintf(OIdFieldName)
	ud.Unset().Add(func() bson.E { return bson.E{Key: mName, Value: ""} })
	return ud
}

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

func (ud *UpdateDocument) Set_bid(p string) *UpdateDocument {
	mName := fmt.Sprintf(BidFieldName)
	ud.Set().Add(func() bson.E { return bson.E{Key: mName, Value: p} })
	return ud
}

func (ud *UpdateDocument) Unset_bid() *UpdateDocument {
	mName := fmt.Sprintf(BidFieldName)
	ud.Unset().Add(func() bson.E { return bson.E{Key: mName, Value: ""} })
	return ud
}

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

func (ud *UpdateDocument) SetScope(p string) *UpdateDocument {
	mName := fmt.Sprintf(ScopeFieldName)
	ud.Set().Add(func() bson.E { return bson.E{Key: mName, Value: p} })
	return ud
}

func (ud *UpdateDocument) UnsetScope() *UpdateDocument {
	mName := fmt.Sprintf(ScopeFieldName)
	ud.Unset().Add(func() bson.E { return bson.E{Key: mName, Value: ""} })
	return ud
}

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

func (ud *UpdateDocument) Set_et(p string) *UpdateDocument {
	mName := fmt.Sprintf(EtFieldName)
	ud.Set().Add(func() bson.E { return bson.E{Key: mName, Value: p} })
	return ud
}

func (ud *UpdateDocument) Unset_et() *UpdateDocument {
	mName := fmt.Sprintf(EtFieldName)
	ud.Unset().Add(func() bson.E { return bson.E{Key: mName, Value: ""} })
	return ud
}

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

func (ud *UpdateDocument) SetCategory(p string) *UpdateDocument {
	mName := fmt.Sprintf(CategoryFieldName)
	ud.Set().Add(func() bson.E { return bson.E{Key: mName, Value: p} })
	return ud
}

func (ud *UpdateDocument) UnsetCategory() *UpdateDocument {
	mName := fmt.Sprintf(CategoryFieldName)
	ud.Unset().Add(func() bson.E { return bson.E{Key: mName, Value: ""} })
	return ud
}

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

func (ud *UpdateDocument) SetIsSystem(p bool) *UpdateDocument {
	mName := fmt.Sprintf(IsSystemFieldName)
	ud.Set().Add(func() bson.E { return bson.E{Key: mName, Value: p} })
	return ud
}

func (ud *UpdateDocument) UnsetIsSystem() *UpdateDocument {
	mName := fmt.Sprintf(IsSystemFieldName)
	ud.Unset().Add(func() bson.E { return bson.E{Key: mName, Value: ""} })
	return ud
}

func (ud *UpdateDocument) setOrUnsetIsSystem(p bool, um UnsetMode) {
	if p {
		ud.SetIsSystem(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetIsSystem()
		case SetData2Default:
			ud.UnsetIsSystem()
		}
	}
}

func UpdateWithIsSystem(p bool) UpdateOption {
	return func(ud *UpdateDocument) {
		if p {
			ud.SetIsSystem(p)
		} else {
			ud.UnsetIsSystem()
		}
	}
}

func (ud *UpdateDocument) SetDescription(p string) *UpdateDocument {
	mName := fmt.Sprintf(DescriptionFieldName)
	ud.Set().Add(func() bson.E { return bson.E{Key: mName, Value: p} })
	return ud
}

func (ud *UpdateDocument) UnsetDescription() *UpdateDocument {
	mName := fmt.Sprintf(DescriptionFieldName)
	ud.Unset().Add(func() bson.E { return bson.E{Key: mName, Value: ""} })
	return ud
}

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

func (ud *UpdateDocument) SetInherited(p bool) *UpdateDocument {
	mName := fmt.Sprintf(InheritedFieldName)
	ud.Set().Add(func() bson.E { return bson.E{Key: mName, Value: p} })
	return ud
}

func (ud *UpdateDocument) UnsetInherited() *UpdateDocument {
	mName := fmt.Sprintf(InheritedFieldName)
	ud.Unset().Add(func() bson.E { return bson.E{Key: mName, Value: ""} })
	return ud
}

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

func (ud *UpdateDocument) SetProperties(p []model.KeyValue) *UpdateDocument {
	mName := fmt.Sprintf(PropertiesFieldName)
	ud.Set().Add(func() bson.E { return bson.E{Key: mName, Value: p} })
	return ud
}

func (ud *UpdateDocument) UnsetProperties() *UpdateDocument {
	mName := fmt.Sprintf(PropertiesFieldName)
	ud.Unset().Add(func() bson.E { return bson.E{Key: mName, Value: ""} })
	return ud
}

func (ud *UpdateDocument) setOrUnsetProperties(p []model.KeyValue, um UnsetMode) {
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

func UpdateWithProperties(p []model.KeyValue) UpdateOption {
	return func(ud *UpdateDocument) {
		if len(p) > 0 {
			ud.SetProperties(p)
		} else {
			ud.UnsetProperties()
		}
	}
}

func (ud *UpdateDocument) SetSysInfo(p *commons.SysInfo) *UpdateDocument {
	mName := fmt.Sprintf(SysInfoFieldName)
	ud.Set().Add(func() bson.E { return bson.E{Key: mName, Value: p} })
	return ud
}

func (ud *UpdateDocument) UnsetSysInfo() *UpdateDocument {
	mName := fmt.Sprintf(SysInfoFieldName)
	ud.Unset().Add(func() bson.E { return bson.E{Key: mName, Value: ""} })
	return ud
}

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

func (ud *UpdateDocument) SetSysinfoModifiedAtNow() *UpdateDocument {
	mName := fmt.Sprintf(SysInfo_ModifiedAtFieldName)
	ud.CurrentDate().Add(func() bson.E { return bson.E{Key: mName, Value: true} })
	return ud
}
