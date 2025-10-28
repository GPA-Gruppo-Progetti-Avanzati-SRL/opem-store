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
	BirthInfoFieldName                 = "birth_info"
	BirthInfo_DateFieldName            = "birth_info.date"
	BirthInfo_CountryFieldName         = "birth_info.country"
	BirthInfo_Country_BidFieldName     = "birth_info.country.bid"
	BirthInfo_County_BidFieldName      = "birth_info.county.bid"
	BirthInfo_Townhall_BidFieldName    = "birth_info.townhall.bid"
	Country_BidFieldName               = "country.bid"
	Addresses_I_County_BidFieldName    = "addresses.%d.county.bid"
	Addresses_County_BidFieldName      = "addresses.county.bid"
	Addresses_I_Townhall_BidFieldName  = "addresses.%d.townhall.bid"
	Addresses_Townhall_BidFieldName    = "addresses.townhall.bid"
	BirthInfo_Country_TextFieldName    = "birth_info.country.text"
	BirthInfo_County_TextFieldName     = "birth_info.county.text"
	BirthInfo_Townhall_TextFieldName   = "birth_info.townhall.text"
	Country_TextFieldName              = "country.text"
	Addresses_I_County_TextFieldName   = "addresses.%d.county.text"
	Addresses_County_TextFieldName     = "addresses.county.text"
	Addresses_I_Townhall_TextFieldName = "addresses.%d.townhall.text"
	Addresses_Townhall_TextFieldName   = "addresses.townhall.text"
	BirthInfo_CountyFieldName          = "birth_info.county"
	BirthInfo_TownhallFieldName        = "birth_info.townhall"
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
