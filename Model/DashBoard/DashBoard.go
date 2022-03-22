package DashBoard

import (
	"database/sql"
	"fmt"
	"github.com/mhthrh/Campaign/Model/Campaign"
	"github.com/mhthrh/Campaign/Utilitys/DbUtil/PgSql"
	"time"
)

type Dashboard struct {
	Duration  time.Duration
	Campaigns []Campaign.Campaign
	Result    []struct {
		UserName string
		Score    int
	}
}

func ReadResult(d *sql.DB, duration []int) []Dashboard {
	var dash Dashboard
	var dashs []Dashboard
	rows, err := PgSql.RunQuery(d, fmt.Sprintf("select * from scores d inner join events e on d.eventid=e.id and group by campain, date, user order by sum(point) in duration "))
	if err != nil {
		return nil
	}
	for rows.Next() {
		rows.Scan()
		dashs = append(dashs, dash)
	}
	return dashs
}
