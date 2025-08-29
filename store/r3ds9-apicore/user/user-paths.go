package user

// This file contains the paths to the field in the generated entity.
// A path is a string with all the identifiers from the root document to the single leaves.
// In case of maps and arrays place holder for the key (%s) or the index %d have been provided.

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

const (
	OIdFieldName                     = "_id"
	NicknameFieldName                = "nickname"
	ObjTypeFieldName                 = "objType"
	FirstnameFieldName               = "firstname"
	LastnameFieldName                = "lastname"
	EmailFieldName                   = "email"
	PasswordFieldName                = "password"
	RolesFieldName                   = "roles"
	Roles_IFieldName                 = "roles.%d"
	SysInfoFieldName                 = "sysInfo"
	SysInfo_StatusFieldName          = "sysInfo.status"
	SysInfo_CreatedatFieldName       = "sysInfo.createdat"
	SysInfo_ModifiedatFieldName      = "sysInfo.modifiedat"
	ProfilePictureFieldName          = "profilePicture"
	ProfilePicture_OIdFieldName      = "profilePicture._id"
	ProfilePicture_SrcSetFieldName   = "profilePicture.srcSet"
	ProfilePicture_SrcSet_IFieldName = "profilePicture.srcSet.%d"
)

// @tpm-schematics:start-region("bottom-file-section")
// @tpm-schematics:end-region("bottom-file-section")
