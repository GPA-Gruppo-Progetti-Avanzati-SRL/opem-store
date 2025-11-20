package card

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
	DefaultMode           UnsetMode
	OId                   UnsetMode
	Domain                UnsetMode
	Site                  UnsetMode
	Bid                   UnsetMode
	Et                    UnsetMode
	CardNumber            UnsetMode
	CardType              UnsetMode
	Status                UnsetMode
	IdCardExt             UnsetMode
	EnvelopeNumber        UnsetMode
	Funct                 UnsetMode
	FocalPoint            UnsetMode
	Product               UnsetMode
	LayoutCode            UnsetMode
	CorporateCode         UnsetMode
	Box                   UnsetMode
	Holder                UnsetMode
	Apps                  UnsetMode
	Addresses             UnsetMode
	Events                UnsetMode
	ExpiresAt             UnsetMode
	IssueDate             UnsetMode
	IssueConfirmationDate UnsetMode
	ActDate               UnsetMode
	SysInfo               UnsetMode
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
func WithCardNumberUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.CardNumber = m
	}
}
func WithCardTypeUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.CardType = m
	}
}
func WithStatusUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Status = m
	}
}
func WithIdCardExtUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.IdCardExt = m
	}
}
func WithEnvelopeNumberUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.EnvelopeNumber = m
	}
}
func WithFunctUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Funct = m
	}
}
func WithFocalPointUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.FocalPoint = m
	}
}
func WithProductUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Product = m
	}
}
func WithLayoutCodeUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.LayoutCode = m
	}
}
func WithCorporateCodeUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.CorporateCode = m
	}
}
func WithBoxUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Box = m
	}
}
func WithHolderUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Holder = m
	}
}
func WithAppsUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Apps = m
	}
}
func WithAddressesUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Addresses = m
	}
}
func WithEventsUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Events = m
	}
}
func WithExpiresAtUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.ExpiresAt = m
	}
}
func WithIssueDateUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.IssueDate = m
	}
}
func WithIssueConfirmationDateUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.IssueConfirmationDate = m
	}
}
func WithActDateUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.ActDate = m
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
func GetUpdateDocument(obj *Card, opts ...UnsetOption) UpdateDocument {

	uo := &UnsetOptions{DefaultMode: KeepCurrent}
	for _, o := range opts {
		o(uo)
	}

	ud := UpdateDocument{}
	ud.setOrUnsetDomain(obj.Domain, uo.ResolveUnsetMode(uo.Domain))
	ud.setOrUnsetSite(obj.Site, uo.ResolveUnsetMode(uo.Site))
	ud.setOrUnset_bid(obj.Bid, uo.ResolveUnsetMode(uo.Bid))
	ud.setOrUnset_et(obj.Et, uo.ResolveUnsetMode(uo.Et))
	ud.setOrUnsetCard_number(&obj.CardNumber, uo.ResolveUnsetMode(uo.CardNumber))
	ud.setOrUnsetCard_type(obj.CardType, uo.ResolveUnsetMode(uo.CardType))
	ud.setOrUnsetStatus(obj.Status, uo.ResolveUnsetMode(uo.Status))
	ud.setOrUnsetId_card_ext(obj.IdCardExt, uo.ResolveUnsetMode(uo.IdCardExt))
	ud.setOrUnsetEnvelope_number(&obj.EnvelopeNumber, uo.ResolveUnsetMode(uo.EnvelopeNumber))
	ud.setOrUnsetFunct(obj.Funct, uo.ResolveUnsetMode(uo.Funct))
	ud.setOrUnsetFocal_point(&obj.FocalPoint, uo.ResolveUnsetMode(uo.FocalPoint))
	ud.setOrUnsetProduct(&obj.Product, uo.ResolveUnsetMode(uo.Product))
	ud.setOrUnsetLayout_code(obj.LayoutCode, uo.ResolveUnsetMode(uo.LayoutCode))
	ud.setOrUnsetCorporate_code(obj.CorporateCode, uo.ResolveUnsetMode(uo.CorporateCode))
	ud.setOrUnsetBox(&obj.Box, uo.ResolveUnsetMode(uo.Box))
	ud.setOrUnsetHolder(&obj.Holder, uo.ResolveUnsetMode(uo.Holder))
	ud.setOrUnsetApps(obj.Apps, uo.ResolveUnsetMode(uo.Apps))
	ud.setOrUnsetAddresses(obj.Addresses, uo.ResolveUnsetMode(uo.Addresses))
	ud.setOrUnsetEvents(obj.Events, uo.ResolveUnsetMode(uo.Events))
	ud.setOrUnsetExpires_at(obj.ExpiresAt, uo.ResolveUnsetMode(uo.ExpiresAt))
	ud.setOrUnsetIssue_date(obj.IssueDate, uo.ResolveUnsetMode(uo.IssueDate))
	ud.setOrUnsetIssue_confirmation_date(obj.IssueConfirmationDate, uo.ResolveUnsetMode(uo.IssueConfirmationDate))
	ud.setOrUnsetAct_date(obj.ActDate, uo.ResolveUnsetMode(uo.ActDate))
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

// SetCard_number No Remarks
func (ud *UpdateDocument) SetCard_number(p *commons.ValueTextPair) *UpdateDocument {
	mName := fmt.Sprintf(CardNumberFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetCard_number No Remarks
func (ud *UpdateDocument) UnsetCard_number() *UpdateDocument {
	mName := fmt.Sprintf(CardNumberFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetCard_number No Remarks - here2
func (ud *UpdateDocument) setOrUnsetCard_number(p *commons.ValueTextPair, um UnsetMode) {
	if p != nil && !p.IsZero() {
		ud.SetCard_number(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetCard_number()
		case SetData2Default:
			ud.UnsetCard_number()
		}
	}
}

func UpdateWithCard_number(p *commons.ValueTextPair) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != nil && !p.IsZero() {
			ud.SetCard_number(p)
		} else {
			ud.UnsetCard_number()
		}
	}
}

// @tpm-schematics:start-region("card-number-field-update-section")
// @tpm-schematics:end-region("card-number-field-update-section")

// SetCard_type No Remarks
func (ud *UpdateDocument) SetCard_type(p string) *UpdateDocument {
	mName := fmt.Sprintf(CardTypeFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetCard_type No Remarks
func (ud *UpdateDocument) UnsetCard_type() *UpdateDocument {
	mName := fmt.Sprintf(CardTypeFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetCard_type No Remarks
func (ud *UpdateDocument) setOrUnsetCard_type(p string, um UnsetMode) {
	if p != "" {
		ud.SetCard_type(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetCard_type()
		case SetData2Default:
			ud.UnsetCard_type()
		}
	}
}

func UpdateWithCard_type(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetCard_type(p)
		} else {
			ud.UnsetCard_type()
		}
	}
}

// @tpm-schematics:start-region("card-type-field-update-section")
// @tpm-schematics:end-region("card-type-field-update-section")

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

// SetId_card_ext No Remarks
func (ud *UpdateDocument) SetId_card_ext(p string) *UpdateDocument {
	mName := fmt.Sprintf(IdCardExtFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetId_card_ext No Remarks
func (ud *UpdateDocument) UnsetId_card_ext() *UpdateDocument {
	mName := fmt.Sprintf(IdCardExtFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetId_card_ext No Remarks
func (ud *UpdateDocument) setOrUnsetId_card_ext(p string, um UnsetMode) {
	if p != "" {
		ud.SetId_card_ext(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetId_card_ext()
		case SetData2Default:
			ud.UnsetId_card_ext()
		}
	}
}

func UpdateWithId_card_ext(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetId_card_ext(p)
		} else {
			ud.UnsetId_card_ext()
		}
	}
}

// @tpm-schematics:start-region("id-card-ext-field-update-section")
// @tpm-schematics:end-region("id-card-ext-field-update-section")

// SetEnvelope_number No Remarks
func (ud *UpdateDocument) SetEnvelope_number(p *commons.ValueTextPair) *UpdateDocument {
	mName := fmt.Sprintf(EnvelopeNumberFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetEnvelope_number No Remarks
func (ud *UpdateDocument) UnsetEnvelope_number() *UpdateDocument {
	mName := fmt.Sprintf(EnvelopeNumberFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetEnvelope_number No Remarks - here2
func (ud *UpdateDocument) setOrUnsetEnvelope_number(p *commons.ValueTextPair, um UnsetMode) {
	if p != nil && !p.IsZero() {
		ud.SetEnvelope_number(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetEnvelope_number()
		case SetData2Default:
			ud.UnsetEnvelope_number()
		}
	}
}

func UpdateWithEnvelope_number(p *commons.ValueTextPair) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != nil && !p.IsZero() {
			ud.SetEnvelope_number(p)
		} else {
			ud.UnsetEnvelope_number()
		}
	}
}

// @tpm-schematics:start-region("envelope-number-field-update-section")
// @tpm-schematics:end-region("envelope-number-field-update-section")

// SetFunct No Remarks
func (ud *UpdateDocument) SetFunct(p string) *UpdateDocument {
	mName := fmt.Sprintf(FunctFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetFunct No Remarks
func (ud *UpdateDocument) UnsetFunct() *UpdateDocument {
	mName := fmt.Sprintf(FunctFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetFunct No Remarks
func (ud *UpdateDocument) setOrUnsetFunct(p string, um UnsetMode) {
	if p != "" {
		ud.SetFunct(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetFunct()
		case SetData2Default:
			ud.UnsetFunct()
		}
	}
}

func UpdateWithFunct(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetFunct(p)
		} else {
			ud.UnsetFunct()
		}
	}
}

// @tpm-schematics:start-region("funct-field-update-section")
// @tpm-schematics:end-region("funct-field-update-section")

// SetFocal_point No Remarks
func (ud *UpdateDocument) SetFocal_point(p *commons.BidTextPair) *UpdateDocument {
	mName := fmt.Sprintf(FocalPointFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetFocal_point No Remarks
func (ud *UpdateDocument) UnsetFocal_point() *UpdateDocument {
	mName := fmt.Sprintf(FocalPointFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetFocal_point No Remarks - here2
func (ud *UpdateDocument) setOrUnsetFocal_point(p *commons.BidTextPair, um UnsetMode) {
	if p != nil && !p.IsZero() {
		ud.SetFocal_point(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetFocal_point()
		case SetData2Default:
			ud.UnsetFocal_point()
		}
	}
}

func UpdateWithFocal_point(p *commons.BidTextPair) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != nil && !p.IsZero() {
			ud.SetFocal_point(p)
		} else {
			ud.UnsetFocal_point()
		}
	}
}

// @tpm-schematics:start-region("focal-point-field-update-section")
// @tpm-schematics:end-region("focal-point-field-update-section")

// SetProduct No Remarks
func (ud *UpdateDocument) SetProduct(p *commons.BidTextPair) *UpdateDocument {
	mName := fmt.Sprintf(ProductFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetProduct No Remarks
func (ud *UpdateDocument) UnsetProduct() *UpdateDocument {
	mName := fmt.Sprintf(ProductFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetProduct No Remarks - here2
func (ud *UpdateDocument) setOrUnsetProduct(p *commons.BidTextPair, um UnsetMode) {
	if p != nil && !p.IsZero() {
		ud.SetProduct(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetProduct()
		case SetData2Default:
			ud.UnsetProduct()
		}
	}
}

func UpdateWithProduct(p *commons.BidTextPair) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != nil && !p.IsZero() {
			ud.SetProduct(p)
		} else {
			ud.UnsetProduct()
		}
	}
}

// @tpm-schematics:start-region("product-field-update-section")
// @tpm-schematics:end-region("product-field-update-section")

// SetLayout_code No Remarks
func (ud *UpdateDocument) SetLayout_code(p string) *UpdateDocument {
	mName := fmt.Sprintf(LayoutCodeFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetLayout_code No Remarks
func (ud *UpdateDocument) UnsetLayout_code() *UpdateDocument {
	mName := fmt.Sprintf(LayoutCodeFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetLayout_code No Remarks
func (ud *UpdateDocument) setOrUnsetLayout_code(p string, um UnsetMode) {
	if p != "" {
		ud.SetLayout_code(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetLayout_code()
		case SetData2Default:
			ud.UnsetLayout_code()
		}
	}
}

func UpdateWithLayout_code(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetLayout_code(p)
		} else {
			ud.UnsetLayout_code()
		}
	}
}

// @tpm-schematics:start-region("layout-code-field-update-section")
// @tpm-schematics:end-region("layout-code-field-update-section")

// SetCorporate_code No Remarks
func (ud *UpdateDocument) SetCorporate_code(p string) *UpdateDocument {
	mName := fmt.Sprintf(CorporateCodeFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetCorporate_code No Remarks
func (ud *UpdateDocument) UnsetCorporate_code() *UpdateDocument {
	mName := fmt.Sprintf(CorporateCodeFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetCorporate_code No Remarks
func (ud *UpdateDocument) setOrUnsetCorporate_code(p string, um UnsetMode) {
	if p != "" {
		ud.SetCorporate_code(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetCorporate_code()
		case SetData2Default:
			ud.UnsetCorporate_code()
		}
	}
}

func UpdateWithCorporate_code(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetCorporate_code(p)
		} else {
			ud.UnsetCorporate_code()
		}
	}
}

// @tpm-schematics:start-region("corporate-code-field-update-section")
// @tpm-schematics:end-region("corporate-code-field-update-section")

// SetBox No Remarks
func (ud *UpdateDocument) SetBox(p *commons.BidTextPair) *UpdateDocument {
	mName := fmt.Sprintf(BoxFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetBox No Remarks
func (ud *UpdateDocument) UnsetBox() *UpdateDocument {
	mName := fmt.Sprintf(BoxFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetBox No Remarks - here2
func (ud *UpdateDocument) setOrUnsetBox(p *commons.BidTextPair, um UnsetMode) {
	if p != nil && !p.IsZero() {
		ud.SetBox(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetBox()
		case SetData2Default:
			ud.UnsetBox()
		}
	}
}

func UpdateWithBox(p *commons.BidTextPair) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != nil && !p.IsZero() {
			ud.SetBox(p)
		} else {
			ud.UnsetBox()
		}
	}
}

// @tpm-schematics:start-region("box-field-update-section")
// @tpm-schematics:end-region("box-field-update-section")

// SetHolder No Remarks
func (ud *UpdateDocument) SetHolder(p *CardHolder) *UpdateDocument {
	mName := fmt.Sprintf(HolderFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetHolder No Remarks
func (ud *UpdateDocument) UnsetHolder() *UpdateDocument {
	mName := fmt.Sprintf(HolderFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetHolder No Remarks - here2
func (ud *UpdateDocument) setOrUnsetHolder(p *CardHolder, um UnsetMode) {
	if p != nil && !p.IsZero() {
		ud.SetHolder(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetHolder()
		case SetData2Default:
			ud.UnsetHolder()
		}
	}
}

func UpdateWithHolder(p *CardHolder) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != nil && !p.IsZero() {
			ud.SetHolder(p)
		} else {
			ud.UnsetHolder()
		}
	}
}

// @tpm-schematics:start-region("holder-field-update-section")
// @tpm-schematics:end-region("holder-field-update-section")

// SetApps No Remarks
func (ud *UpdateDocument) SetApps(p []CardApp) *UpdateDocument {
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
func (ud *UpdateDocument) setOrUnsetApps(p []CardApp, um UnsetMode) {
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

func UpdateWithApps(p []CardApp) UpdateOption {
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

// SetEvents No Remarks
func (ud *UpdateDocument) SetEvents(p []commons.Event) *UpdateDocument {
	mName := fmt.Sprintf(EventsFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetEvents No Remarks
func (ud *UpdateDocument) UnsetEvents() *UpdateDocument {
	mName := fmt.Sprintf(EventsFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetEvents No Remarks - here2
func (ud *UpdateDocument) setOrUnsetEvents(p []commons.Event, um UnsetMode) {
	if len(p) > 0 {
		ud.SetEvents(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetEvents()
		case SetData2Default:
			ud.UnsetEvents()
		}
	}
}

func UpdateWithEvents(p []commons.Event) UpdateOption {
	return func(ud *UpdateDocument) {
		if len(p) > 0 {
			ud.SetEvents(p)
		} else {
			ud.UnsetEvents()
		}
	}
}

// @tpm-schematics:start-region("events-field-update-section")

func UpdateWithAddEvent(p commons.Event) UpdateOption {
	return func(ud *UpdateDocument) {
		if !p.IsZero() {
			mName := fmt.Sprintf(EventsFieldName)
			ud.Push().Add(func() bson.E {
				return bson.E{Key: mName, Value: p}
			})
		}
	}
}

// @tpm-schematics:end-region("events-field-update-section")

// SetExpires_at No Remarks
func (ud *UpdateDocument) SetExpires_at(p bson.DateTime) *UpdateDocument {
	mName := fmt.Sprintf(ExpiresAtFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetExpires_at No Remarks
func (ud *UpdateDocument) UnsetExpires_at() *UpdateDocument {
	mName := fmt.Sprintf(ExpiresAtFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetExpires_at No Remarks
func (ud *UpdateDocument) setOrUnsetExpires_at(p bson.DateTime, um UnsetMode) {
	if p != 0 {
		ud.SetExpires_at(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetExpires_at()
		case SetData2Default:
			ud.UnsetExpires_at()
		}
	}
}

func UpdateWithExpires_at(p bson.DateTime) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != 0 {
			ud.SetExpires_at(p)
		} else {
			ud.UnsetExpires_at()
		}
	}
}

// @tpm-schematics:start-region("expires-at-field-update-section")
// @tpm-schematics:end-region("expires-at-field-update-section")

// SetIssue_date No Remarks
func (ud *UpdateDocument) SetIssue_date(p bson.DateTime) *UpdateDocument {
	mName := fmt.Sprintf(IssueDateFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetIssue_date No Remarks
func (ud *UpdateDocument) UnsetIssue_date() *UpdateDocument {
	mName := fmt.Sprintf(IssueDateFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetIssue_date No Remarks
func (ud *UpdateDocument) setOrUnsetIssue_date(p bson.DateTime, um UnsetMode) {
	if p != 0 {
		ud.SetIssue_date(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetIssue_date()
		case SetData2Default:
			ud.UnsetIssue_date()
		}
	}
}

func UpdateWithIssue_date(p bson.DateTime) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != 0 {
			ud.SetIssue_date(p)
		} else {
			ud.UnsetIssue_date()
		}
	}
}

// @tpm-schematics:start-region("issue-date-field-update-section")
// @tpm-schematics:end-region("issue-date-field-update-section")

// SetIssue_confirmation_date No Remarks
func (ud *UpdateDocument) SetIssue_confirmation_date(p bson.DateTime) *UpdateDocument {
	mName := fmt.Sprintf(IssueConfirmationDateFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetIssue_confirmation_date No Remarks
func (ud *UpdateDocument) UnsetIssue_confirmation_date() *UpdateDocument {
	mName := fmt.Sprintf(IssueConfirmationDateFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetIssue_confirmation_date No Remarks
func (ud *UpdateDocument) setOrUnsetIssue_confirmation_date(p bson.DateTime, um UnsetMode) {
	if p != 0 {
		ud.SetIssue_confirmation_date(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetIssue_confirmation_date()
		case SetData2Default:
			ud.UnsetIssue_confirmation_date()
		}
	}
}

func UpdateWithIssue_confirmation_date(p bson.DateTime) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != 0 {
			ud.SetIssue_confirmation_date(p)
		} else {
			ud.UnsetIssue_confirmation_date()
		}
	}
}

// @tpm-schematics:start-region("issue-confirmation-date-field-update-section")
// @tpm-schematics:end-region("issue-confirmation-date-field-update-section")

// SetAct_date No Remarks
func (ud *UpdateDocument) SetAct_date(p bson.DateTime) *UpdateDocument {
	mName := fmt.Sprintf(ActDateFieldName)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetAct_date No Remarks
func (ud *UpdateDocument) UnsetAct_date() *UpdateDocument {
	mName := fmt.Sprintf(ActDateFieldName)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetAct_date No Remarks
func (ud *UpdateDocument) setOrUnsetAct_date(p bson.DateTime, um UnsetMode) {
	if p != 0 {
		ud.SetAct_date(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetAct_date()
		case SetData2Default:
			ud.UnsetAct_date()
		}
	}
}

func UpdateWithAct_date(p bson.DateTime) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != 0 {
			ud.SetAct_date(p)
		} else {
			ud.UnsetAct_date()
		}
	}
}

// @tpm-schematics:start-region("act-date-field-update-section")
// @tpm-schematics:end-region("act-date-field-update-section")

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
