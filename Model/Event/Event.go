package Event

import (
	"database/sql"
	"fmt"
	"github.com/mhthrh/Campaign/Utilitys/DbUtil/PgSql"
)

type Event struct {
	ID          int
	CampaignId  int
	Event       string
	Point       int
	Description string
}

func LoadEvents(db *sql.DB) []Event {

	var event Event
	var Events []Event
	rows, err := PgSql.RunQuery(db, fmt.Sprintf("SELECT \"ID\", \"CampaignId\", \"Event\", \"Point\", \"Description\" FROM public.\"Events_View\""))
	if err != nil {
		return nil
	}

	for rows.Next() {
		rows.Scan(&event.ID, &event.CampaignId, &event.Event, event.Point, &event.Description)
		Events = append(Events, event)
	}

	return Events
}

func AddEvents(db *sql.DB) error {
	return nil
}
