package Campaign

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/mhthrh/Campaign/Utilitys/DbUtil/PgSql"
	"time"
)

type Campaign struct {
	ID        int
	Name      string
	StartDate time.Time
	EndDate   time.Time
}

func LoadCampaign(db *sql.DB) []Campaign {

	var campaign Campaign
	var Campaigns []Campaign
	rows, err := PgSql.RunQuery(db, fmt.Sprintf("SELECT \"ID\", \"Name\", \"StartDate\", \"EndDate\" FROM public.\"Campaign_View\""))
	if err != nil {
		return nil
	}

	for rows.Next() {
		rows.Scan(&campaign.ID, &campaign.Name, &campaign.StartDate, &campaign.EndDate)
		Campaigns = append(Campaigns, campaign)
	}
	return Campaigns
}

func AddCampaign(db *sql.DB) error {

	return errors.New("not implemented")
}
