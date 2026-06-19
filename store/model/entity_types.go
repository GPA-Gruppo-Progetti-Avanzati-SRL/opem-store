package model

// Costanti del campo _et (entity type) usate come discriminatore nei documenti MongoDB
// e nei file YAML del driver static.
const (
	EntityTypeUser     = "user"
	EntityTypeOIDCUser = "oidc-user"
	EntityTypeDomain   = "domain"
	EntityTypeSite     = "site"
	EntityTypeKV       = "kv"
	EntityTypeSequence = "sequence"

	EntityTypeRoleCaps = "role-caps"
	EntityTypeCapGroup = "cap-group"
	EntityTypeCapDef   = "cap-def"
)
