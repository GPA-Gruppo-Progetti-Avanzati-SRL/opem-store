package card

// This file contains the paths to the field in the generated entity.
// A path is a string with all the identifiers from the root document to the single leaves.
// In case of maps and arrays place holder for the key (%s) or the index %d have been provided.

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

const (
	OIdFieldName                       = "_id"
	DomainFieldName                    = "domain"
	SiteFieldName                      = "site"
	BidFieldName                       = "_bid"
	EtFieldName                        = "_et"
	CardNumberFieldName                = "card_number"
	CardNumber_ValueFieldName          = "card_number.value"
	EnvelopeNumber_ValueFieldName      = "envelope_number.value"
	Apps_I_AppNumber_ValueFieldName    = "apps.%d.app_number.value"
	Apps_AppNumber_ValueFieldName      = "apps.app_number.value"
	CardNumber_TextFieldName           = "card_number.text"
	EnvelopeNumber_TextFieldName       = "envelope_number.text"
	Apps_I_AppNumber_TextFieldName     = "apps.%d.app_number.text"
	Apps_AppNumber_TextFieldName       = "apps.app_number.text"
	CardTypeFieldName                  = "card_type"
	StatusFieldName                    = "status"
	IdCardExtFieldName                 = "id_card_ext"
	EnvelopeNumberFieldName            = "envelope_number"
	FunctFieldName                     = "funct"
	FocalPointFieldName                = "focal_point"
	FocalPoint_BidFieldName            = "focal_point.bid"
	Product_BidFieldName               = "product.bid"
	Box_BidFieldName                   = "box.bid"
	Addresses_I_Country_BidFieldName   = "addresses.%d.country.bid"
	Addresses_Country_BidFieldName     = "addresses.country.bid"
	Addresses_I_County_BidFieldName    = "addresses.%d.county.bid"
	Addresses_County_BidFieldName      = "addresses.county.bid"
	Addresses_I_Townhall_BidFieldName  = "addresses.%d.townhall.bid"
	Addresses_Townhall_BidFieldName    = "addresses.townhall.bid"
	FocalPoint_TextFieldName           = "focal_point.text"
	Product_TextFieldName              = "product.text"
	Box_TextFieldName                  = "box.text"
	Addresses_I_Country_TextFieldName  = "addresses.%d.country.text"
	Addresses_Country_TextFieldName    = "addresses.country.text"
	Addresses_I_County_TextFieldName   = "addresses.%d.county.text"
	Addresses_County_TextFieldName     = "addresses.county.text"
	Addresses_I_Townhall_TextFieldName = "addresses.%d.townhall.text"
	Addresses_Townhall_TextFieldName   = "addresses.townhall.text"
	ProductFieldName                   = "product"
	LayoutCodeFieldName                = "layout_code"
	CorporateCodeFieldName             = "corporate_code"
	BoxFieldName                       = "box"
	HolderFieldName                    = "holder"
	Holder_BidFieldName                = "holder.bid"
	Holder_EmbossingNameFieldName      = "holder.embossing_name"
	Holder_RegistrationIdFieldName     = "holder.registration_id"
	AppsFieldName                      = "apps"
	Apps_IFieldName                    = "apps.%d"
	AddressesFieldName                 = "addresses"
	Addresses_IFieldName               = "addresses.%d"
	EventsFieldName                    = "events"
	Events_IFieldName                  = "events.%d"
	ExpiresAtFieldName                 = "expires_at"
	IssueDateFieldName                 = "issue_date"
	IssueConfirmationDateFieldName     = "issue_confirmation_date"
	ActDateFieldName                   = "act_date"
	SysInfoFieldName                   = "sys_info"
	SysInfo_StatusFieldName            = "sys_info.status"
	SysInfo_CreatedAtFieldName         = "sys_info.created_at"
	SysInfo_ModifiedAtFieldName        = "sys_info.modified_at"
)

// @tpm-schematics:start-region("bottom-file-section")
const (
	AppsAppNumberFieldName = "apps.app_number.value"
)

// @tpm-schematics:end-region("bottom-file-section")
