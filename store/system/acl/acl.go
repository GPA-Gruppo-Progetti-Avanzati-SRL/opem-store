// Package acl contiene i modelli e le operazioni di persistenza per il sistema
// ACL + Capabilities del front-door.
//
// Tre tipi di documento vivono nella stessa collection (opem_acl):
//
//   - cap-def   : catalogo delle capabilities (definito una volta per app+id)
//   - cap-group : raggruppamento di IDs di cap-def (riutilizzabile)
//   - role-caps : assegnazione di cap-groups/IDs a un ruolo (contestualizzata per domain/site/app)
package acl

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/commons"
)

// ── Costanti ──────────────────────────────────────────────────────────────────

const (
	// CollectionId è l'ID logico della collection nel linked-service MongoDB.
	// Corrisponde a: collections: [ { id: "acl", name: "opem_acl" } ]
	CollectionId = "acl"

	// Tipi di documento (_et)
	EntityTypeCapDef   = "cap-def"
	EntityTypeCapGroup = "cap-group"
	EntityTypeRoleCaps = "role-caps"

	// Categorie di capability
	CategoryUI  = "ui"
	CategoryAPI = "api"
)

// ── CapDef ────────────────────────────────────────────────────────────────────

// CapDef è la definizione canonica di una capability.
// Esiste una sola istanza per (app, _id) nella collection.
// I campi UI e API sono mutuamente esclusivi (discriminati da Category).
//
// Naming convention _id: "<resource>:<azione>.<categoria>"
//   es.  "sim:list.ui"   → capability lato UI
//        "sim:list.api"  → capability lato API
//
// Il campo _id è l'unica chiave: il vecchio campo "id" (bson:"id") è stato rimosso.
type CapDef struct {
	OId      string          `json:"_id,omitempty"      bson:"_id,omitempty"` // stringa, es. "sim:list.ui"
	Et       string          `json:"_et,omitempty"      bson:"_et,omitempty"`
	App      string          `json:"app"                bson:"app"`
	Category string          `json:"category"           bson:"category"` // "ui" | "api"
	SysInfo  commons.SysInfo `json:"sys_info,omitempty" bson:"sys_info,omitempty"`

	// ── Campi UI (Category == "ui") ──────────────────────────────────────────
	Label    string `json:"label,omitempty"    bson:"label,omitempty"`
	Endpoint string `json:"endpoint,omitempty" bson:"endpoint,omitempty"`
	Icon     string `json:"icon,omitempty"     bson:"icon,omitempty"`
	Order    int    `json:"order,omitempty"    bson:"order,omitempty"`
	Visible  *bool  `json:"visible,omitempty"  bson:"visible,omitempty"`

	// ── Campi API (Category == "api") ────────────────────────────────────────
	Mapping string `json:"mapping,omitempty" bson:"mapping,omitempty"` // nome del mapping nel proxy-config
	Method  string `json:"method,omitempty"  bson:"method,omitempty"`  // GET | POST | * ecc.
	Path    string `json:"path,omitempty"    bson:"path,omitempty"`    // path glob (es. /simulazioni/*)
}

func (c CapDef) IsZero() bool {
	return c.OId == "" && c.App == ""
}

// ── CapGroup ──────────────────────────────────────────────────────────────────

// CapGroup raggruppa IDs di cap-def per riuso nelle assegnazioni role-caps.
// Non ridefinisce le capabilities: contiene solo riferimenti (IDs).
// app="*" rende il gruppo cross-app.
//
// Naming convention _id: uguale al campo code (es. "sim-viewer").
// Il code è l'unica chiave di lookup: _id = code garantisce l'unicità a livello globale.
type CapGroup struct {
	OId          string          `json:"_id,omitempty"         bson:"_id,omitempty"` // stringa, es. "sim-viewer"
	Et           string          `json:"_et,omitempty"         bson:"_et,omitempty"`
	Code         string          `json:"code"                  bson:"code"`
	App          string          `json:"app"                   bson:"app"`
	Description  string          `json:"description,omitempty" bson:"description,omitempty"`
	Capabilities []string        `json:"capabilities"          bson:"capabilities"` // lista di cap-def _id
	SysInfo      commons.SysInfo `json:"sys_info,omitempty"    bson:"sys_info,omitempty"`
}

func (c CapGroup) IsZero() bool {
	return c.Code == "" && c.App == ""
}

// ── RoleCapsEntry ─────────────────────────────────────────────────────────────

// RoleCapsEntry assegna cap-groups e/o IDs individuali di cap-def a un ruolo.
//
// domain, site, app sono opzionali: se assenti il documento vale per qualsiasi
// valore del campo (semantica "wildcard implicita").
// Più documenti per lo stesso ruolo vengono trovati e le capabilities risultanti
// sono sempre l'UNIONE (nessuna negazione).
//
// Naming convention _id:
//   - globale (nessun domain/site):   "<role>"                  → "FA_admin"
//   - contestuale (domain+site):      "<role>:<domain>:<site>"  → "FA_admin:leas:CC"
type RoleCapsEntry struct {
	OId          string          `json:"_id,omitempty"      bson:"_id,omitempty"` // stringa, es. "FA_admin" o "FA_admin:leas:CC"
	Et           string          `json:"_et,omitempty"      bson:"_et,omitempty"`
	Role         string          `json:"role"               bson:"role"`
	Domain       string          `json:"domain,omitempty"   bson:"domain,omitempty"`
	Site         string          `json:"site,omitempty"     bson:"site,omitempty"`
	App          string          `json:"app,omitempty"      bson:"app,omitempty"`
	CapGroups    []string        `json:"cap_groups"         bson:"cap_groups"`   // codici di cap-group (_id del cap-group)
	Capabilities []string        `json:"capabilities"       bson:"capabilities"` // _id diretti di cap-def
	SysInfo      commons.SysInfo `json:"sys_info,omitempty" bson:"sys_info,omitempty"`
}

func (r RoleCapsEntry) IsZero() bool {
	return r.Role == ""
}

