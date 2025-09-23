package site

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
	Code        UnsetMode
	Domain      UnsetMode
	ObjType     UnsetMode
	Name        UnsetMode
	Description UnsetMode
	Bookmark    UnsetMode
	Langs       UnsetMode
	Apps        UnsetMode
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
func WithCodeUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Code = m
	}
}
func WithDomainUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Domain = m
	}
}
func WithObjTypeUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.ObjType = m
	}
}
func WithNameUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Name = m
	}
}
func WithDescriptionUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Description = m
	}
}
func WithBookmarkUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Bookmark = m
	}
}
func WithLangsUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Langs = m
	}
}
func WithAppsUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Apps = m
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
func GetUpdateDocument(obj *Site, opts ...UnsetOption) UpdateDocument {

	uo := &UnsetOptions{DefaultMode: KeepCurrent}
	for _, o := range opts {
		o(uo)
	}

	ud := UpdateDocument{}
	ud.setOrUnsetCode(obj.Code, uo.ResolveUnsetMode(uo.Code))
	ud.setOrUnsetDomain(obj.Domain, uo.ResolveUnsetMode(uo.Domain))
	ud.setOrUnsetObj_type(obj.ObjType, uo.ResolveUnsetMode(uo.ObjType))
	ud.setOrUnsetName(obj.Name, uo.ResolveUnsetMode(uo.Name))
	ud.setOrUnsetDescription(obj.Description, uo.ResolveUnsetMode(uo.Description))
	ud.setOrUnsetBookmark(obj.Bookmark, uo.ResolveUnsetMode(uo.Bookmark))
	ud.setOrUnsetLangs(obj.Langs, uo.ResolveUnsetMode(uo.Langs))
	ud.setOrUnsetApps(obj.Apps, uo.ResolveUnsetMode(uo.Apps))
	ud.setOrUnsetSys_info(&obj.SysInfo, uo.ResolveUnsetMode(uo.SysInfo))

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

// SetCode No Remarks
func (ud *UpdateDocument) SetCode(p string) *UpdateDocument {
	mName := fmt.Sprintf(CodeFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetCode No Remarks
func (ud *UpdateDocument) UnsetCode() *UpdateDocument {
	mName := fmt.Sprintf(CodeFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetCode No Remarks
func (ud *UpdateDocument) setOrUnsetCode(p string, um UnsetMode) {
	if p != "" {
		ud.SetCode(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetCode()
		case SetData2Default:
			ud.UnsetCode()
		}
	}
}

func UpdateWithCode(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetCode(p)
		} else {
			ud.UnsetCode()
		}
	}
}

// @tpm-schematics:start-region("code-field-update-section")
// @tpm-schematics:end-region("code-field-update-section")

// SetDomain No Remarks
func (ud *UpdateDocument) SetDomain(p string) *UpdateDocument {
	mName := fmt.Sprintf(DomainFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetDomain No Remarks
func (ud *UpdateDocument) UnsetDomain() *UpdateDocument {
	mName := fmt.Sprintf(DomainFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetDomain No Remarks
func (ud *UpdateDocument) setOrUnsetDomain(p string, um UnsetMode) {
	if p != "" {
		ud.SetDomain(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetDomain()
		case SetData2Default:
			ud.UnsetDomain()
		}
	}
}

func UpdateWithDomain(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetDomain(p)
		} else {
			ud.UnsetDomain()
		}
	}
}

// @tpm-schematics:start-region("domain-field-update-section")
// @tpm-schematics:end-region("domain-field-update-section")

// SetObj_type No Remarks
func (ud *UpdateDocument) SetObj_type(p string) *UpdateDocument {
	mName := fmt.Sprintf(ObjTypeFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetObj_type No Remarks
func (ud *UpdateDocument) UnsetObj_type() *UpdateDocument {
	mName := fmt.Sprintf(ObjTypeFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetObj_type No Remarks
func (ud *UpdateDocument) setOrUnsetObj_type(p string, um UnsetMode) {
	if p != "" {
		ud.SetObj_type(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetObj_type()
		case SetData2Default:
			ud.UnsetObj_type()
		}
	}
}

func UpdateWithObj_type(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetObj_type(p)
		} else {
			ud.UnsetObj_type()
		}
	}
}

// @tpm-schematics:start-region("obj-type-field-update-section")
// @tpm-schematics:end-region("obj-type-field-update-section")

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

// SetBookmark No Remarks
func (ud *UpdateDocument) SetBookmark(p bool) *UpdateDocument {
	mName := fmt.Sprintf(BookmarkFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetBookmark No Remarks
func (ud *UpdateDocument) UnsetBookmark() *UpdateDocument {
	mName := fmt.Sprintf(BookmarkFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetBookmark No Remarks
func (ud *UpdateDocument) setOrUnsetBookmark(p bool, um UnsetMode) {
	if p {
		ud.SetBookmark(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetBookmark()
		case SetData2Default:
			ud.UnsetBookmark()
		}
	}
}

func UpdateWithBookmark(p bool) UpdateOption {
	return func(ud *UpdateDocument) {
		if p {
			ud.SetBookmark(p)
		} else {
			ud.UnsetBookmark()
		}
	}
}

// @tpm-schematics:start-region("bookmark-field-update-section")
// @tpm-schematics:end-region("bookmark-field-update-section")

// SetLangs No Remarks
func (ud *UpdateDocument) SetLangs(p string) *UpdateDocument {
	mName := fmt.Sprintf(LangsFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetLangs No Remarks
func (ud *UpdateDocument) UnsetLangs() *UpdateDocument {
	mName := fmt.Sprintf(LangsFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetLangs No Remarks
func (ud *UpdateDocument) setOrUnsetLangs(p string, um UnsetMode) {
	if p != "" {
		ud.SetLangs(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetLangs()
		case SetData2Default:
			ud.UnsetLangs()
		}
	}
}

func UpdateWithLangs(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetLangs(p)
		} else {
			ud.UnsetLangs()
		}
	}
}

// @tpm-schematics:start-region("langs-field-update-section")
// @tpm-schematics:end-region("langs-field-update-section")

// SetApps No Remarks
func (ud *UpdateDocument) SetApps(p []commons.App) *UpdateDocument {
	mName := fmt.Sprintf(AppsFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetApps No Remarks
func (ud *UpdateDocument) UnsetApps() *UpdateDocument {
	mName := fmt.Sprintf(AppsFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetApps No Remarks - here2
func (ud *UpdateDocument) setOrUnsetApps(p []commons.App, um UnsetMode) {
	if len(p) > 0 {
		ud.SetApps(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetApps()
		case SetData2Default:
			ud.UnsetApps()
		}
	}
}

func UpdateWithApps(p []commons.App) UpdateOption {
	return func(ud *UpdateDocument) {
		if len(p) > 0 {
			ud.SetApps(p)
		} else {
			ud.UnsetApps()
		}
	}
}

// @tpm-schematics:start-region("apps-field-update-section")
// @tpm-schematics:end-region("apps-field-update-section")

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
