package person

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
	FirstNameFieldName                 = "first_name"
	LastNameFieldName                  = "last_name"
	CfFieldName                        = "cf"
	MaleFemaleFieldName                = "male_female"
	BirthFieldName                     = "birth"
	Birth_DateFieldName                = "birth.date"
	Birth_CountryFieldName             = "birth.country"
	Birth_Country_BidFieldName         = "birth.country.bid"
	Birth_County_BidFieldName          = "birth.county.bid"
	Birth_Townhall_BidFieldName        = "birth.townhall.bid"
	IdDocument_Townhall_BidFieldName   = "id_document.townhall.bid"
	IdDocument_County_BidFieldName     = "id_document.county.bid"
	Country_BidFieldName               = "country.bid"
	Addresses_I_County_BidFieldName    = "addresses.%d.county.bid"
	Addresses_County_BidFieldName      = "addresses.county.bid"
	Addresses_I_Townhall_BidFieldName  = "addresses.%d.townhall.bid"
	Addresses_Townhall_BidFieldName    = "addresses.townhall.bid"
	Birth_Country_TextFieldName        = "birth.country.text"
	Birth_County_TextFieldName         = "birth.county.text"
	Birth_Townhall_TextFieldName       = "birth.townhall.text"
	IdDocument_Townhall_TextFieldName  = "id_document.townhall.text"
	IdDocument_County_TextFieldName    = "id_document.county.text"
	Country_TextFieldName              = "country.text"
	Addresses_I_County_TextFieldName   = "addresses.%d.county.text"
	Addresses_County_TextFieldName     = "addresses.county.text"
	Addresses_I_Townhall_TextFieldName = "addresses.%d.townhall.text"
	Addresses_Townhall_TextFieldName   = "addresses.townhall.text"
	Birth_CountyFieldName              = "birth.county"
	Birth_TownhallFieldName            = "birth.townhall"
	IdDocumentFieldName                = "id_document"
	IdDocument_TypeFieldName           = "id_document.type"
	IdDocument_CodeFieldName           = "id_document.code"
	IdDocument_IssueDateFieldName      = "id_document.issue_date"
	IdDocument_TownhallFieldName       = "id_document.townhall"
	IdDocument_CountyFieldName         = "id_document.county"
	IdDocument_IssuerFieldName         = "id_document.issuer"
	IdDocument_ExpiresAtFieldName      = "id_document.expires_at"
	CountryFieldName                   = "country"
	AddressesFieldName                 = "addresses"
	Addresses_IFieldName               = "addresses.%d"
	SysInfoFieldName                   = "sys_info"
	SysInfo_StatusFieldName            = "sys_info.status"
	SysInfo_CreatedAtFieldName         = "sys_info.created_at"
	SysInfo_ModifiedAtFieldName        = "sys_info.modified_at"
)

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
