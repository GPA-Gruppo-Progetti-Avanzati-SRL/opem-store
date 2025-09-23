package file

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
	Fn          UnsetMode
	Descr       UnsetMode
	Role        UnsetMode
	EntRefs     UnsetMode
	Metadata    UnsetMode
	Vrnts       UnsetMode
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
func WithFnUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Fn = m
	}
}
func WithDescrUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Descr = m
	}
}
func WithRoleUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Role = m
	}
}
func WithEntRefsUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.EntRefs = m
	}
}
func WithMetadataUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Metadata = m
	}
}
func WithVrntsUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Vrnts = m
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
func GetUpdateDocument(obj *File, opts ...UnsetOption) UpdateDocument {

	uo := &UnsetOptions{DefaultMode: KeepCurrent}
	for _, o := range opts {
		o(uo)
	}

	ud := UpdateDocument{}
	ud.setOrUnsetFn(obj.Fn, uo.ResolveUnsetMode(uo.Fn))
	ud.setOrUnsetDescr(obj.Descr, uo.ResolveUnsetMode(uo.Descr))
	ud.setOrUnsetRole(obj.Role, uo.ResolveUnsetMode(uo.Role))
	ud.setOrUnsetEnt_refs(obj.EntRefs, uo.ResolveUnsetMode(uo.EntRefs))
	ud.setOrUnsetMetadata(obj.Metadata, uo.ResolveUnsetMode(uo.Metadata))
	ud.setOrUnsetVrnts(obj.Vrnts, uo.ResolveUnsetMode(uo.Vrnts))
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

// SetFn No Remarks
func (ud *UpdateDocument) SetFn(p string) *UpdateDocument {
	mName := fmt.Sprintf(FnFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetFn No Remarks
func (ud *UpdateDocument) UnsetFn() *UpdateDocument {
	mName := fmt.Sprintf(FnFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetFn No Remarks
func (ud *UpdateDocument) setOrUnsetFn(p string, um UnsetMode) {
	if p != "" {
		ud.SetFn(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetFn()
		case SetData2Default:
			ud.UnsetFn()
		}
	}
}

func UpdateWithFn(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetFn(p)
		} else {
			ud.UnsetFn()
		}
	}
}

// @tpm-schematics:start-region("fn-field-update-section")
// @tpm-schematics:end-region("fn-field-update-section")

// SetDescr No Remarks
func (ud *UpdateDocument) SetDescr(p string) *UpdateDocument {
	mName := fmt.Sprintf(DescrFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetDescr No Remarks
func (ud *UpdateDocument) UnsetDescr() *UpdateDocument {
	mName := fmt.Sprintf(DescrFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetDescr No Remarks
func (ud *UpdateDocument) setOrUnsetDescr(p string, um UnsetMode) {
	if p != "" {
		ud.SetDescr(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetDescr()
		case SetData2Default:
			ud.UnsetDescr()
		}
	}
}

func UpdateWithDescr(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetDescr(p)
		} else {
			ud.UnsetDescr()
		}
	}
}

// @tpm-schematics:start-region("descr-field-update-section")
// @tpm-schematics:end-region("descr-field-update-section")

// SetRole No Remarks
func (ud *UpdateDocument) SetRole(p string) *UpdateDocument {
	mName := fmt.Sprintf(RoleFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetRole No Remarks
func (ud *UpdateDocument) UnsetRole() *UpdateDocument {
	mName := fmt.Sprintf(RoleFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetRole No Remarks
func (ud *UpdateDocument) setOrUnsetRole(p string, um UnsetMode) {
	if p != "" {
		ud.SetRole(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetRole()
		case SetData2Default:
			ud.UnsetRole()
		}
	}
}

func UpdateWithRole(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetRole(p)
		} else {
			ud.UnsetRole()
		}
	}
}

// @tpm-schematics:start-region("role-field-update-section")
// @tpm-schematics:end-region("role-field-update-section")

// SetEnt_refs No Remarks
func (ud *UpdateDocument) SetEnt_refs(p []EntRefStruct) *UpdateDocument {
	mName := fmt.Sprintf(EntRefsFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetEnt_refs No Remarks
func (ud *UpdateDocument) UnsetEnt_refs() *UpdateDocument {
	mName := fmt.Sprintf(EntRefsFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetEnt_refs No Remarks - here2
func (ud *UpdateDocument) setOrUnsetEnt_refs(p []EntRefStruct, um UnsetMode) {
	if len(p) > 0 {
		ud.SetEnt_refs(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetEnt_refs()
		case SetData2Default:
			ud.UnsetEnt_refs()
		}
	}
}

func UpdateWithEnt_refs(p []EntRefStruct) UpdateOption {
	return func(ud *UpdateDocument) {
		if len(p) > 0 {
			ud.SetEnt_refs(p)
		} else {
			ud.UnsetEnt_refs()
		}
	}
}

// @tpm-schematics:start-region("ent-refs-field-update-section")
// @tpm-schematics:end-region("ent-refs-field-update-section")

// SetMetadata No Remarks
func (ud *UpdateDocument) SetMetadata(p bson.M) *UpdateDocument {
	mName := fmt.Sprintf(MetadataFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetMetadata No Remarks
func (ud *UpdateDocument) UnsetMetadata() *UpdateDocument {
	mName := fmt.Sprintf(MetadataFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetMetadata No Remarks
func (ud *UpdateDocument) setOrUnsetMetadata(p bson.M, um UnsetMode) {
	if len(p) != 0 {
		ud.SetMetadata(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetMetadata()
		case SetData2Default:
			ud.UnsetMetadata()
		}
	}
}

func UpdateWithMetadata(p bson.M) UpdateOption {
	return func(ud *UpdateDocument) {
		if len(p) != 0 {
			ud.SetMetadata(p)
		} else {
			ud.UnsetMetadata()
		}
	}
}

// @tpm-schematics:start-region("metadata-field-update-section")
// @tpm-schematics:end-region("metadata-field-update-section")

// SetVrnts No Remarks
func (ud *UpdateDocument) SetVrnts(p []commons.FileVariant) *UpdateDocument {
	mName := fmt.Sprintf(VrntsFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetVrnts No Remarks
func (ud *UpdateDocument) UnsetVrnts() *UpdateDocument {
	mName := fmt.Sprintf(VrntsFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetVrnts No Remarks - here2
func (ud *UpdateDocument) setOrUnsetVrnts(p []commons.FileVariant, um UnsetMode) {
	if len(p) > 0 {
		ud.SetVrnts(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetVrnts()
		case SetData2Default:
			ud.UnsetVrnts()
		}
	}
}

func UpdateWithVrnts(p []commons.FileVariant) UpdateOption {
	return func(ud *UpdateDocument) {
		if len(p) > 0 {
			ud.SetVrnts(p)
		} else {
			ud.UnsetVrnts()
		}
	}
}

// @tpm-schematics:start-region("vrnts-field-update-section")
// @tpm-schematics:end-region("vrnts-field-update-section")

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
