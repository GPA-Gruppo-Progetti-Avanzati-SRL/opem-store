package person

import (
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"

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
	Domain      UnsetMode
	Site        UnsetMode
	Bid         UnsetMode
	Et          UnsetMode
	FirstName   UnsetMode
	LastName    UnsetMode
	Cf          UnsetMode
	MaleFemale  UnsetMode
	Birth       UnsetMode
	IdDocument  UnsetMode
	Country     UnsetMode
	Addresses   UnsetMode
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
func WithFirstNameUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.FirstName = m
	}
}
func WithLastNameUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.LastName = m
	}
}
func WithCfUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Cf = m
	}
}
func WithMaleFemaleUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.MaleFemale = m
	}
}
func WithBirthUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Birth = m
	}
}
func WithIdDocumentUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.IdDocument = m
	}
}
func WithCountryUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Country = m
	}
}
func WithAddressesUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Addresses = m
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
func GetUpdateDocument(obj *Person, opts ...UnsetOption) UpdateDocument {

	uo := &UnsetOptions{DefaultMode: KeepCurrent}
	for _, o := range opts {
		o(uo)
	}

	ud := UpdateDocument{}
	ud.setOrUnsetDomain(obj.Domain, uo.ResolveUnsetMode(uo.Domain))
	ud.setOrUnsetSite(obj.Site, uo.ResolveUnsetMode(uo.Site))
	ud.setOrUnset_bid(obj.Bid, uo.ResolveUnsetMode(uo.Bid))
	ud.setOrUnset_et(obj.Et, uo.ResolveUnsetMode(uo.Et))
	ud.setOrUnsetFirst_name(obj.FirstName, uo.ResolveUnsetMode(uo.FirstName))
	ud.setOrUnsetLast_name(obj.LastName, uo.ResolveUnsetMode(uo.LastName))
	ud.setOrUnsetCf(obj.Cf, uo.ResolveUnsetMode(uo.Cf))
	ud.setOrUnsetMale_female(obj.MaleFemale, uo.ResolveUnsetMode(uo.MaleFemale))
	ud.setOrUnsetBirth(&obj.Birth, uo.ResolveUnsetMode(uo.Birth))
	ud.setOrUnsetId_document(&obj.IdDocument, uo.ResolveUnsetMode(uo.IdDocument))
	ud.setOrUnsetCountry(&obj.Country, uo.ResolveUnsetMode(uo.Country))
	ud.setOrUnsetAddresses(obj.Addresses, uo.ResolveUnsetMode(uo.Addresses))
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

// SetFirst_name No Remarks
func (ud *UpdateDocument) SetFirst_name(p string) *UpdateDocument {
	mName := fmt.Sprintf(FirstNameFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetFirst_name No Remarks
func (ud *UpdateDocument) UnsetFirst_name() *UpdateDocument {
	mName := fmt.Sprintf(FirstNameFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetFirst_name No Remarks
func (ud *UpdateDocument) setOrUnsetFirst_name(p string, um UnsetMode) {
	if p != "" {
		ud.SetFirst_name(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetFirst_name()
		case SetData2Default:
			ud.UnsetFirst_name()
		}
	}
}

func UpdateWithFirst_name(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetFirst_name(p)
		} else {
			ud.UnsetFirst_name()
		}
	}
}

// @tpm-schematics:start-region("first-name-field-update-section")
// @tpm-schematics:end-region("first-name-field-update-section")

// SetLast_name No Remarks
func (ud *UpdateDocument) SetLast_name(p string) *UpdateDocument {
	mName := fmt.Sprintf(LastNameFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetLast_name No Remarks
func (ud *UpdateDocument) UnsetLast_name() *UpdateDocument {
	mName := fmt.Sprintf(LastNameFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetLast_name No Remarks
func (ud *UpdateDocument) setOrUnsetLast_name(p string, um UnsetMode) {
	if p != "" {
		ud.SetLast_name(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetLast_name()
		case SetData2Default:
			ud.UnsetLast_name()
		}
	}
}

func UpdateWithLast_name(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetLast_name(p)
		} else {
			ud.UnsetLast_name()
		}
	}
}

// @tpm-schematics:start-region("last-name-field-update-section")
// @tpm-schematics:end-region("last-name-field-update-section")

// SetCf No Remarks
func (ud *UpdateDocument) SetCf(p string) *UpdateDocument {
	mName := fmt.Sprintf(CfFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetCf No Remarks
func (ud *UpdateDocument) UnsetCf() *UpdateDocument {
	mName := fmt.Sprintf(CfFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetCf No Remarks
func (ud *UpdateDocument) setOrUnsetCf(p string, um UnsetMode) {
	if p != "" {
		ud.SetCf(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetCf()
		case SetData2Default:
			ud.UnsetCf()
		}
	}
}

func UpdateWithCf(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetCf(p)
		} else {
			ud.UnsetCf()
		}
	}
}

// @tpm-schematics:start-region("cf-field-update-section")
// @tpm-schematics:end-region("cf-field-update-section")

// SetMale_female No Remarks
func (ud *UpdateDocument) SetMale_female(p string) *UpdateDocument {
	mName := fmt.Sprintf(MaleFemaleFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetMale_female No Remarks
func (ud *UpdateDocument) UnsetMale_female() *UpdateDocument {
	mName := fmt.Sprintf(MaleFemaleFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetMale_female No Remarks
func (ud *UpdateDocument) setOrUnsetMale_female(p string, um UnsetMode) {
	if p != "" {
		ud.SetMale_female(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetMale_female()
		case SetData2Default:
			ud.UnsetMale_female()
		}
	}
}

func UpdateWithMale_female(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetMale_female(p)
		} else {
			ud.UnsetMale_female()
		}
	}
}

// @tpm-schematics:start-region("male-female-field-update-section")
// @tpm-schematics:end-region("male-female-field-update-section")

// SetBirth No Remarks
func (ud *UpdateDocument) SetBirth(p *BirthInfo) *UpdateDocument {
	mName := fmt.Sprintf(BirthFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetBirth No Remarks
func (ud *UpdateDocument) UnsetBirth() *UpdateDocument {
	mName := fmt.Sprintf(BirthFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetBirth No Remarks - here2
func (ud *UpdateDocument) setOrUnsetBirth(p *BirthInfo, um UnsetMode) {
	if p != nil && !p.IsZero() {
		ud.SetBirth(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetBirth()
		case SetData2Default:
			ud.UnsetBirth()
		}
	}
}

func UpdateWithBirth(p *BirthInfo) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != nil && !p.IsZero() {
			ud.SetBirth(p)
		} else {
			ud.UnsetBirth()
		}
	}
}

// @tpm-schematics:start-region("birth-field-update-section")
// @tpm-schematics:end-region("birth-field-update-section")

// SetId_document No Remarks
func (ud *UpdateDocument) SetId_document(p *IdentityCard) *UpdateDocument {
	mName := fmt.Sprintf(IdDocumentFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetId_document No Remarks
func (ud *UpdateDocument) UnsetId_document() *UpdateDocument {
	mName := fmt.Sprintf(IdDocumentFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetId_document No Remarks - here2
func (ud *UpdateDocument) setOrUnsetId_document(p *IdentityCard, um UnsetMode) {
	if p != nil && !p.IsZero() {
		ud.SetId_document(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetId_document()
		case SetData2Default:
			ud.UnsetId_document()
		}
	}
}

func UpdateWithId_document(p *IdentityCard) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != nil && !p.IsZero() {
			ud.SetId_document(p)
		} else {
			ud.UnsetId_document()
		}
	}
}

// @tpm-schematics:start-region("id-document-field-update-section")
// @tpm-schematics:end-region("id-document-field-update-section")

// SetCountry No Remarks
func (ud *UpdateDocument) SetCountry(p *commons.BidTextPair) *UpdateDocument {
	mName := fmt.Sprintf(CountryFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetCountry No Remarks
func (ud *UpdateDocument) UnsetCountry() *UpdateDocument {
	mName := fmt.Sprintf(CountryFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetCountry No Remarks - here2
func (ud *UpdateDocument) setOrUnsetCountry(p *commons.BidTextPair, um UnsetMode) {
	if p != nil && !p.IsZero() {
		ud.SetCountry(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetCountry()
		case SetData2Default:
			ud.UnsetCountry()
		}
	}
}

func UpdateWithCountry(p *commons.BidTextPair) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != nil && !p.IsZero() {
			ud.SetCountry(p)
		} else {
			ud.UnsetCountry()
		}
	}
}

// @tpm-schematics:start-region("country-field-update-section")
// @tpm-schematics:end-region("country-field-update-section")

// SetAddresses No Remarks
func (ud *UpdateDocument) SetAddresses(p []commons.Address) *UpdateDocument {
	mName := fmt.Sprintf(AddressesFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetAddresses No Remarks
func (ud *UpdateDocument) UnsetAddresses() *UpdateDocument {
	mName := fmt.Sprintf(AddressesFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetAddresses No Remarks - here2
func (ud *UpdateDocument) setOrUnsetAddresses(p []commons.Address, um UnsetMode) {
	if len(p) > 0 {
		ud.SetAddresses(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetAddresses()
		case SetData2Default:
			ud.UnsetAddresses()
		}
	}
}

func UpdateWithAddresses(p []commons.Address) UpdateOption {
	return func(ud *UpdateDocument) {
		if len(p) > 0 {
			ud.SetAddresses(p)
		} else {
			ud.UnsetAddresses()
		}
	}
}

// @tpm-schematics:start-region("addresses-field-update-section")
// @tpm-schematics:end-region("addresses-field-update-section")

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
