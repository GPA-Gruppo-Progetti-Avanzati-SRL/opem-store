package prodotto

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
	DefaultMode  UnsetMode
	OId          UnsetMode
	Domain       UnsetMode
	Site         UnsetMode
	Bid          UnsetMode
	Et           UnsetMode
	Name         UnsetMode
	PrimaryFunct UnsetMode
	Code         UnsetMode
	ExpirationAt UnsetMode
	SysInfo      UnsetMode
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
func WithDomainUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Domain = m
	}
}
func WithSiteUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Site = m
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
func WithPrimaryFunctUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.PrimaryFunct = m
	}
}
func WithCodeUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Code = m
	}
}
func WithExpirationAtUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.ExpirationAt = m
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
func GetUpdateDocument(obj *Prodotto, opts ...UnsetOption) UpdateDocument {

	uo := &UnsetOptions{DefaultMode: KeepCurrent}
	for _, o := range opts {
		o(uo)
	}

	ud := UpdateDocument{}
	ud.setOrUnsetDomain(obj.Domain, uo.ResolveUnsetMode(uo.Domain))
	ud.setOrUnsetSite(obj.Site, uo.ResolveUnsetMode(uo.Site))
	ud.setOrUnset_bid(obj.Bid, uo.ResolveUnsetMode(uo.Bid))
	ud.setOrUnset_et(obj.Et, uo.ResolveUnsetMode(uo.Et))
	ud.setOrUnsetName(obj.Name, uo.ResolveUnsetMode(uo.Name))
	ud.setOrUnsetPrimary_funct(obj.PrimaryFunct, uo.ResolveUnsetMode(uo.PrimaryFunct))
	ud.setOrUnsetCode(obj.Code, uo.ResolveUnsetMode(uo.Code))
	ud.setOrUnsetExpiration_at(obj.ExpirationAt, uo.ResolveUnsetMode(uo.ExpirationAt))
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

// SetSite No Remarks
func (ud *UpdateDocument) SetSite(p string) *UpdateDocument {
	mName := fmt.Sprintf(SiteFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetSite No Remarks
func (ud *UpdateDocument) UnsetSite() *UpdateDocument {
	mName := fmt.Sprintf(SiteFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetSite No Remarks
func (ud *UpdateDocument) setOrUnsetSite(p string, um UnsetMode) {
	if p != "" {
		ud.SetSite(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetSite()
		case SetData2Default:
			ud.UnsetSite()
		}
	}
}

func UpdateWithSite(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetSite(p)
		} else {
			ud.UnsetSite()
		}
	}
}

// @tpm-schematics:start-region("site-field-update-section")
// @tpm-schematics:end-region("site-field-update-section")

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

// SetPrimary_funct No Remarks
func (ud *UpdateDocument) SetPrimary_funct(p string) *UpdateDocument {
	mName := fmt.Sprintf(PrimaryFunctFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetPrimary_funct No Remarks
func (ud *UpdateDocument) UnsetPrimary_funct() *UpdateDocument {
	mName := fmt.Sprintf(PrimaryFunctFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetPrimary_funct No Remarks
func (ud *UpdateDocument) setOrUnsetPrimary_funct(p string, um UnsetMode) {
	if p != "" {
		ud.SetPrimary_funct(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetPrimary_funct()
		case SetData2Default:
			ud.UnsetPrimary_funct()
		}
	}
}

func UpdateWithPrimary_funct(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetPrimary_funct(p)
		} else {
			ud.UnsetPrimary_funct()
		}
	}
}

// @tpm-schematics:start-region("primary-funct-field-update-section")
// @tpm-schematics:end-region("primary-funct-field-update-section")

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

// SetExpiration_at No Remarks
func (ud *UpdateDocument) SetExpiration_at(p string) *UpdateDocument {
	mName := fmt.Sprintf(ExpirationAtFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetExpiration_at No Remarks
func (ud *UpdateDocument) UnsetExpiration_at() *UpdateDocument {
	mName := fmt.Sprintf(ExpirationAtFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetExpiration_at No Remarks
func (ud *UpdateDocument) setOrUnsetExpiration_at(p string, um UnsetMode) {
	if p != "" {
		ud.SetExpiration_at(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetExpiration_at()
		case SetData2Default:
			ud.UnsetExpiration_at()
		}
	}
}

func UpdateWithExpiration_at(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetExpiration_at(p)
		} else {
			ud.UnsetExpiration_at()
		}
	}
}

// @tpm-schematics:start-region("expiration-at-field-update-section")
// @tpm-schematics:end-region("expiration-at-field-update-section")

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
