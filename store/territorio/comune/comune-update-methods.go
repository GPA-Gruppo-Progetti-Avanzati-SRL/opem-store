package comune

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
	DefaultMode   UnsetMode
	OId           UnsetMode
	Bid           UnsetMode
	Et            UnsetMode
	Name          UnsetMode
	Cap1          UnsetMode
	Cap2          UnsetMode
	Provincia     UnsetMode
	Nazione       UnsetMode
	CodeIstat     UnsetMode
	CodeCatastale UnsetMode
	Cab           UnsetMode
	SysInfo       UnsetMode
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
func WithEtUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Et = m
	}
}
func WithNameUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Name = m
	}
}
func WithCap1UnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Cap1 = m
	}
}
func WithCap2UnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Cap2 = m
	}
}
func WithProvinciaUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Provincia = m
	}
}
func WithNazioneUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Nazione = m
	}
}
func WithCodeIstatUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.CodeIstat = m
	}
}
func WithCodeCatastaleUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.CodeCatastale = m
	}
}
func WithCabUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Cab = m
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
func GetUpdateDocument(obj *Comune, opts ...UnsetOption) UpdateDocument {

	uo := &UnsetOptions{DefaultMode: KeepCurrent}
	for _, o := range opts {
		o(uo)
	}

	ud := UpdateDocument{}
	ud.setOrUnset_bid(obj.Bid, uo.ResolveUnsetMode(uo.Bid))
	ud.setOrUnset_et(obj.Et, uo.ResolveUnsetMode(uo.Et))
	ud.setOrUnsetName(obj.Name, uo.ResolveUnsetMode(uo.Name))
	ud.setOrUnsetCap1(obj.Cap1, uo.ResolveUnsetMode(uo.Cap1))
	ud.setOrUnsetCap2(obj.Cap2, uo.ResolveUnsetMode(uo.Cap2))
	ud.setOrUnsetProvincia(&obj.Provincia, uo.ResolveUnsetMode(uo.Provincia))
	ud.setOrUnsetNazione(&obj.Nazione, uo.ResolveUnsetMode(uo.Nazione))
	ud.setOrUnsetCode_istat(obj.CodeIstat, uo.ResolveUnsetMode(uo.CodeIstat))
	ud.setOrUnsetCode_catastale(obj.CodeCatastale, uo.ResolveUnsetMode(uo.CodeCatastale))
	ud.setOrUnsetCab(obj.Cab, uo.ResolveUnsetMode(uo.Cab))
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

// SetCap1 No Remarks
func (ud *UpdateDocument) SetCap1(p string) *UpdateDocument {
	mName := fmt.Sprintf(Cap1FieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetCap1 No Remarks
func (ud *UpdateDocument) UnsetCap1() *UpdateDocument {
	mName := fmt.Sprintf(Cap1FieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetCap1 No Remarks
func (ud *UpdateDocument) setOrUnsetCap1(p string, um UnsetMode) {
	if p != "" {
		ud.SetCap1(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetCap1()
		case SetData2Default:
			ud.UnsetCap1()
		}
	}
}

func UpdateWithCap1(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetCap1(p)
		} else {
			ud.UnsetCap1()
		}
	}
}

// @tpm-schematics:start-region("cap1-field-update-section")
// @tpm-schematics:end-region("cap1-field-update-section")

// SetCap2 No Remarks
func (ud *UpdateDocument) SetCap2(p string) *UpdateDocument {
	mName := fmt.Sprintf(Cap2FieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetCap2 No Remarks
func (ud *UpdateDocument) UnsetCap2() *UpdateDocument {
	mName := fmt.Sprintf(Cap2FieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetCap2 No Remarks
func (ud *UpdateDocument) setOrUnsetCap2(p string, um UnsetMode) {
	if p != "" {
		ud.SetCap2(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetCap2()
		case SetData2Default:
			ud.UnsetCap2()
		}
	}
}

func UpdateWithCap2(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetCap2(p)
		} else {
			ud.UnsetCap2()
		}
	}
}

// @tpm-schematics:start-region("cap2-field-update-section")
// @tpm-schematics:end-region("cap2-field-update-section")

// SetProvincia No Remarks
func (ud *UpdateDocument) SetProvincia(p *commons.BidTextPair) *UpdateDocument {
	mName := fmt.Sprintf(ProvinciaFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetProvincia No Remarks
func (ud *UpdateDocument) UnsetProvincia() *UpdateDocument {
	mName := fmt.Sprintf(ProvinciaFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetProvincia No Remarks - here2
func (ud *UpdateDocument) setOrUnsetProvincia(p *commons.BidTextPair, um UnsetMode) {
	if p != nil && !p.IsZero() {
		ud.SetProvincia(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetProvincia()
		case SetData2Default:
			ud.UnsetProvincia()
		}
	}
}

func UpdateWithProvincia(p *commons.BidTextPair) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != nil && !p.IsZero() {
			ud.SetProvincia(p)
		} else {
			ud.UnsetProvincia()
		}
	}
}

// @tpm-schematics:start-region("provincia-field-update-section")
// @tpm-schematics:end-region("provincia-field-update-section")

// SetNazione No Remarks
func (ud *UpdateDocument) SetNazione(p *commons.BidTextPair) *UpdateDocument {
	mName := fmt.Sprintf(NazioneFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetNazione No Remarks
func (ud *UpdateDocument) UnsetNazione() *UpdateDocument {
	mName := fmt.Sprintf(NazioneFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetNazione No Remarks - here2
func (ud *UpdateDocument) setOrUnsetNazione(p *commons.BidTextPair, um UnsetMode) {
	if p != nil && !p.IsZero() {
		ud.SetNazione(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetNazione()
		case SetData2Default:
			ud.UnsetNazione()
		}
	}
}

func UpdateWithNazione(p *commons.BidTextPair) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != nil && !p.IsZero() {
			ud.SetNazione(p)
		} else {
			ud.UnsetNazione()
		}
	}
}

// @tpm-schematics:start-region("nazione-field-update-section")
// @tpm-schematics:end-region("nazione-field-update-section")

// SetCode_istat No Remarks
func (ud *UpdateDocument) SetCode_istat(p string) *UpdateDocument {
	mName := fmt.Sprintf(CodeIstatFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetCode_istat No Remarks
func (ud *UpdateDocument) UnsetCode_istat() *UpdateDocument {
	mName := fmt.Sprintf(CodeIstatFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetCode_istat No Remarks
func (ud *UpdateDocument) setOrUnsetCode_istat(p string, um UnsetMode) {
	if p != "" {
		ud.SetCode_istat(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetCode_istat()
		case SetData2Default:
			ud.UnsetCode_istat()
		}
	}
}

func UpdateWithCode_istat(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetCode_istat(p)
		} else {
			ud.UnsetCode_istat()
		}
	}
}

// @tpm-schematics:start-region("code-istat-field-update-section")
// @tpm-schematics:end-region("code-istat-field-update-section")

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

// SetCab No Remarks
func (ud *UpdateDocument) SetCab(p string) *UpdateDocument {
	mName := fmt.Sprintf(CabFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetCab No Remarks
func (ud *UpdateDocument) UnsetCab() *UpdateDocument {
	mName := fmt.Sprintf(CabFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetCab No Remarks
func (ud *UpdateDocument) setOrUnsetCab(p string, um UnsetMode) {
	if p != "" {
		ud.SetCab(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetCab()
		case SetData2Default:
			ud.UnsetCab()
		}
	}
}

func UpdateWithCab(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetCab(p)
		} else {
			ud.UnsetCab()
		}
	}
}

// @tpm-schematics:start-region("cab-field-update-section")
// @tpm-schematics:end-region("cab-field-update-section")

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
