package nazione

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"

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
	DefaultMode   UnsetMode
	OId           UnsetMode
	Et            UnsetMode
	Code          UnsetMode
	Name          UnsetMode
	CodeUic       UnsetMode
	CodeIso3      UnsetMode
	CodeCatastale UnsetMode
	Status        UnsetMode
	Order         UnsetMode
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
func WithEtUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Et = m
	}
}
func WithCodeUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Code = m
	}
}
func WithNameUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Name = m
	}
}
func WithCodeUicUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.CodeUic = m
	}
}
func WithCodeIso3UnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.CodeIso3 = m
	}
}
func WithCodeCatastaleUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.CodeCatastale = m
	}
}
func WithStatusUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Status = m
	}
}
func WithOrderUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Order = m
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
func GetUpdateDocument(obj *Nazione, opts ...UnsetOption) UpdateDocument {

	uo := &UnsetOptions{DefaultMode: KeepCurrent}
	for _, o := range opts {
		o(uo)
	}

	ud := UpdateDocument{}
	ud.setOrUnset_et(obj.Et, uo.ResolveUnsetMode(uo.Et))
	ud.setOrUnsetCode(obj.Code, uo.ResolveUnsetMode(uo.Code))
	ud.setOrUnsetName(obj.Name, uo.ResolveUnsetMode(uo.Name))
	ud.setOrUnsetCode_uic(obj.CodeUic, uo.ResolveUnsetMode(uo.CodeUic))
	ud.setOrUnsetCode_iso3(obj.CodeIso3, uo.ResolveUnsetMode(uo.CodeIso3))
	ud.setOrUnsetCode_catastale(obj.CodeCatastale, uo.ResolveUnsetMode(uo.CodeCatastale))
	ud.setOrUnsetStatus(obj.Status, uo.ResolveUnsetMode(uo.Status))
	ud.setOrUnsetOrder(obj.Order, uo.ResolveUnsetMode(uo.Order))

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

// SetCode_uic No Remarks
func (ud *UpdateDocument) SetCode_uic(p string) *UpdateDocument {
	mName := fmt.Sprintf(CodeUicFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetCode_uic No Remarks
func (ud *UpdateDocument) UnsetCode_uic() *UpdateDocument {
	mName := fmt.Sprintf(CodeUicFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetCode_uic No Remarks
func (ud *UpdateDocument) setOrUnsetCode_uic(p string, um UnsetMode) {
	if p != "" {
		ud.SetCode_uic(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetCode_uic()
		case SetData2Default:
			ud.UnsetCode_uic()
		}
	}
}

func UpdateWithCode_uic(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetCode_uic(p)
		} else {
			ud.UnsetCode_uic()
		}
	}
}

// @tpm-schematics:start-region("code-uic-field-update-section")
// @tpm-schematics:end-region("code-uic-field-update-section")

// SetCode_iso3 No Remarks
func (ud *UpdateDocument) SetCode_iso3(p string) *UpdateDocument {
	mName := fmt.Sprintf(CodeIso3FieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetCode_iso3 No Remarks
func (ud *UpdateDocument) UnsetCode_iso3() *UpdateDocument {
	mName := fmt.Sprintf(CodeIso3FieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetCode_iso3 No Remarks
func (ud *UpdateDocument) setOrUnsetCode_iso3(p string, um UnsetMode) {
	if p != "" {
		ud.SetCode_iso3(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetCode_iso3()
		case SetData2Default:
			ud.UnsetCode_iso3()
		}
	}
}

func UpdateWithCode_iso3(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetCode_iso3(p)
		} else {
			ud.UnsetCode_iso3()
		}
	}
}

// @tpm-schematics:start-region("code-iso3-field-update-section")
// @tpm-schematics:end-region("code-iso3-field-update-section")

// SetCode_catastale No Remarks
func (ud *UpdateDocument) SetCode_catastale(p string) *UpdateDocument {
	mName := fmt.Sprintf(CodeCatastaleFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetCode_catastale No Remarks
func (ud *UpdateDocument) UnsetCode_catastale() *UpdateDocument {
	mName := fmt.Sprintf(CodeCatastaleFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetCode_catastale No Remarks
func (ud *UpdateDocument) setOrUnsetCode_catastale(p string, um UnsetMode) {
	if p != "" {
		ud.SetCode_catastale(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetCode_catastale()
		case SetData2Default:
			ud.UnsetCode_catastale()
		}
	}
}

func UpdateWithCode_catastale(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetCode_catastale(p)
		} else {
			ud.UnsetCode_catastale()
		}
	}
}

// @tpm-schematics:start-region("code-catastale-field-update-section")
// @tpm-schematics:end-region("code-catastale-field-update-section")

// SetStatus No Remarks
func (ud *UpdateDocument) SetStatus(p string) *UpdateDocument {
	mName := fmt.Sprintf(StatusFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetStatus No Remarks
func (ud *UpdateDocument) UnsetStatus() *UpdateDocument {
	mName := fmt.Sprintf(StatusFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetStatus No Remarks
func (ud *UpdateDocument) setOrUnsetStatus(p string, um UnsetMode) {
	if p != "" {
		ud.SetStatus(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetStatus()
		case SetData2Default:
			ud.UnsetStatus()
		}
	}
}

func UpdateWithStatus(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetStatus(p)
		} else {
			ud.UnsetStatus()
		}
	}
}

// @tpm-schematics:start-region("status-field-update-section")
// @tpm-schematics:end-region("status-field-update-section")

// SetOrder No Remarks
func (ud *UpdateDocument) SetOrder(p int32) *UpdateDocument {
	mName := fmt.Sprintf(OrderFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetOrder No Remarks
func (ud *UpdateDocument) UnsetOrder() *UpdateDocument {
	mName := fmt.Sprintf(OrderFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetOrder No Remarks
func (ud *UpdateDocument) setOrUnsetOrder(p int32, um UnsetMode) {
	if p != 0 {
		ud.SetOrder(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetOrder()
		case SetData2Default:
			ud.UnsetOrder()
		}
	}
}

func UpdateWithOrder(p int32) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != 0 {
			ud.SetOrder(p)
		} else {
			ud.UnsetOrder()
		}
	}
}

// @tpm-schematics:start-region("order-field-update-section")
// @tpm-schematics:end-region("order-field-update-section")

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
