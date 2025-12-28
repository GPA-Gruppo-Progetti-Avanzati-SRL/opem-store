package box

import "go.mongodb.org/mongo-driver/v2/bson"
import "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/commons"

// @tpm-schematics:start-region("top-file-section")
const (
	EntityType   = "box"
	CollectionId = "box"

	SupplyTypeNotKnown   = "not_known"
	SupplyTypeAnonymous  = "anonymous"
	SupplyTypeNominative = "nominative"

	SupplyType
	StatusDaGenerare                         = "CX" // value: "magazzino da generare in stato Confermato"
	StatusInAttesaDiInvio                    = "CH" // value: in attesa di invio - Stato non esistente in PCWEB
	StatusInAttesaDiConferma                 = "C1" // value: in attesa di conferma
	StatusConfermatoDisponibilePerLaConsegna = "CO" // Confermato (Item disponibili per la consegna)
	EventReintegroMagazzino                  = "reintegro-magazzino"
	EventRichiestaCarte                      = "richiesta-carte"
)

// @tpm-schematics:end-region("top-file-section")

type Box struct {
	OId           bson.ObjectID       `json:"_id,omitempty" bson:"_id,omitempty" yaml:"_id,omitempty"`
	Domain        string              `json:"domain,omitempty" bson:"domain,omitempty" yaml:"domain,omitempty"`
	Site          string              `json:"site,omitempty" bson:"site,omitempty" yaml:"site,omitempty"`
	Bid           string              `json:"_bid,omitempty" bson:"_bid,omitempty" yaml:"_bid,omitempty"`
	Et            string              `json:"_et,omitempty" bson:"_et,omitempty" yaml:"_et,omitempty"`
	Magazzino     commons.BidTextPair `json:"magazzino,omitempty" bson:"magazzino,omitempty" yaml:"magazzino,omitempty"`
	Prodotto      commons.BidTextPair `json:"prodotto,omitempty" bson:"prodotto,omitempty" yaml:"prodotto,omitempty"`
	FocalPoint    commons.BidTextPair `json:"focal_point,omitempty" bson:"focal_point,omitempty" yaml:"focal_point,omitempty"`
	SupplyType    string              `json:"supply_type,omitempty" bson:"supply_type,omitempty" yaml:"supply_type,omitempty"`
	Info          Info                `json:"info,omitempty" bson:"info,omitempty" yaml:"info,omitempty"`
	Status        Status              `json:"status,omitempty" bson:"status,omitempty" yaml:"status,omitempty"`
	Recipient     commons.Address     `json:"recipient,omitempty" bson:"recipient,omitempty" yaml:"recipient,omitempty"`
	CardBidsRange commons.ValueRange  `json:"card_bids_range,omitempty" bson:"card_bids_range,omitempty" yaml:"card_bids_range,omitempty"`
	Events        []commons.Event     `json:"events,omitempty" bson:"events,omitempty" yaml:"events,omitempty"`
	Notes         []commons.Note      `json:"notes,omitempty" bson:"notes,omitempty" yaml:"notes,omitempty"`
	Activities    commons.Activities  `json:"activities,omitempty" bson:"activities,omitempty" yaml:"activities,omitempty"`
	SysInfo       commons.SysInfo     `json:"sys_info,omitempty" bson:"sys_info,omitempty" yaml:"sys_info,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s Box) IsZero() bool {
	return s.OId == bson.NilObjectID && s.Domain == "" && s.Site == "" && s.Bid == "" && s.Et == "" && s.Magazzino.IsZero() && s.Prodotto.IsZero() && s.FocalPoint.IsZero() && s.SupplyType == "" && s.Info.IsZero() && s.Status.IsZero() && s.Recipient.IsZero() && s.CardBidsRange.IsZero() && len(s.Events) == 0 && len(s.Notes) == 0 && s.Activities.IsZero() && s.SysInfo.IsZero()
}

type QueryResult struct {
	Records int   `json:"records,omitempty" bson:"records,omitempty" yaml:"records,omitempty"`
	Data    []Box `json:"data,omitempty" bson:"data,omitempty" yaml:"data,omitempty"`
}

// @tpm-schematics:start-region("bottom-file-section")

func (s Box) WithNewTodoActivity(activity commons.Activity) (commons.Activities, bool) {
	for _, a := range s.Activities.Logs {
		if a.Id == activity.Id {
			return s.Activities, false
		}
	}

	for _, a := range s.Activities.Todos {
		if a.Id == activity.Id {
			return s.Activities, false
		}
	}

	s.Activities.Todos = append(s.Activities.Todos, activity)
	return s.Activities, true
}

// WithActivityDone il booleano in questo caso ritorna il fatto se lo stato dell'attività e' stato modificato o meno.
func (s Box) WithActivityDone(activity commons.Activity) (commons.Activities, bool) {

	ndx := -1
	for i := 0; i < len(s.Activities.Todos); i++ {
		if s.Activities.Todos[i].Id == activity.Id {
			ndx = i
			break
		}
	}

	if ndx >= 0 {
		if len(s.Activities.Todos) <= 1 {
			s.Activities.Todos = nil
		} else {
			s.Activities.Todos = append(s.Activities.Todos[:ndx], s.Activities.Todos[ndx+1:]...)
		}
	}

	s.Activities.Logs = append(s.Activities.Logs, activity)
	return s.Activities, false
}

// @tpm-schematics:end-region("bottom-file-section")
