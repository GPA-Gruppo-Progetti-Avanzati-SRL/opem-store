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
	CardTypeFieldName                  = "card_type"
	StatusFieldName                    = "status"
	IdCardExtFieldName                 = "id_card_ext"
	EnvelopeNumberFieldName            = "envelope_number"
	FunctFieldName                     = "funct"
	FocalPointFieldName                = "focal_point"
	FocalPoint_BidFieldName            = "focal_point.bid"
	Product_BidFieldName               = "product.bid"
	Magazzino_BidFieldName             = "magazzino.bid"
	Addresses_I_County_BidFieldName    = "addresses.%d.county.bid"
	Addresses_County_BidFieldName      = "addresses.county.bid"
	Addresses_I_Townhall_BidFieldName  = "addresses.%d.townhall.bid"
	Addresses_Townhall_BidFieldName    = "addresses.townhall.bid"
	FocalPoint_TextFieldName           = "focal_point.text"
	Product_TextFieldName              = "product.text"
	Magazzino_TextFieldName            = "magazzino.text"
	Addresses_I_County_TextFieldName   = "addresses.%d.county.text"
	Addresses_County_TextFieldName     = "addresses.county.text"
	Addresses_I_Townhall_TextFieldName = "addresses.%d.townhall.text"
	Addresses_Townhall_TextFieldName   = "addresses.townhall.text"
	ProductFieldName                   = "product"
	MagazzinoFieldName                 = "magazzino"
	HolderFieldName                    = "holder"
	Holder_BidFieldName                = "holder.bid"
	Holder_EmbossingNameFieldName      = "holder.embossing_name"
	Holder_RegistrationIdFieldName     = "holder.registration_id"
	AppsFieldName                      = "apps"
	Apps_IFieldName                    = "apps.%d"
	AddressesFieldName                 = "addresses"
	Addresses_IFieldName               = "addresses.%d"
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
// @tpm-schematics:end-region("bottom-file-section")
