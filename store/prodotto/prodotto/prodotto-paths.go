package prodotto

// This file contains the paths to the field in the generated entity.
// A path is a string with all the identifiers from the root document to the single leaves.
// In case of maps and arrays place holder for the key (%s) or the index %d have been provided.

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

const (
	OIdFieldName                    = "_id"
	DomainFieldName                 = "domain"
	SiteFieldName                   = "site"
	BidFieldName                    = "_bid"
	EtFieldName                     = "_et"
	NameFieldName                   = "name"
	PrimaryFunctFieldName           = "primary_funct"
	ExpiresAtFieldName              = "expires_at"
	PersBureauFieldName             = "pers_bureau"
	HostProductFieldName            = "host_product"
	HostProduct_CodeFieldName       = "host_product.code"
	HostProduct_LayoutCodeFieldName = "host_product.layout_code"
	HostProduct_ReleaseFieldName    = "host_product.release"
	PropertiesFieldName             = "properties"
	AppsFieldName                   = "apps"
	Apps_IFieldName                 = "apps.%d"
	SysInfoFieldName                = "sys_info"
	SysInfo_StatusFieldName         = "sys_info.status"
	SysInfo_CreatedAtFieldName      = "sys_info.created_at"
	SysInfo_ModifiedAtFieldName     = "sys_info.modified_at"
)

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
