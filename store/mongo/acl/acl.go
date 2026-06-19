// Package acl contiene i modelli e le operazioni di persistenza per il sistema
// ACL + Capabilities del front-door.
//
// Tre tipi di documento vivono nella stessa collection (opem_acl):
//
//   - cap-def   : catalogo delle capabilities (definito una volta per app+id)
//   - cap-group : raggruppamento di IDs di cap-def (riutilizzabile)
//   - role-caps : assegnazione di cap-groups/IDs a un ruolo (contestualizzata per domain/site/app)
package acl

const (
	CollectionId = "acl"

	EntityTypeCapDef   = "cap-def"
	EntityTypeCapGroup = "cap-group"
	EntityTypeRoleCaps = "role-caps"

	CategoryUI     = "ui"
	CategoryAPI    = "api"
	CategoryAction = "action"
)
