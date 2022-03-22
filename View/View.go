package View

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mhthrh/Campaign/Controller"
	"github.com/mhthrh/Campaign/Model/Campaign"
	"github.com/mhthrh/Campaign/Model/Event"
	"github.com/mhthrh/Campaign/Utilitys/ConfigUtil"
	"github.com/mhthrh/Campaign/Utilitys/DbUtil/DbPool"
	"github.com/mhthrh/Campaign/Utilitys/ValidationUtil"
	"github.com/sirupsen/logrus"
	"net/http"
)

func RunApiOnRouter(sm *mux.Router, log *logrus.Entry, db *DbPool.DBs, config *ConfigUtil.Config) error {
	ca := Controller.CampaignHandler{
		Campaign: nil,
		Event:    nil,
	}
	h := ""
	validation := ValidationUtil.NewValidation()
	database := db.Pull()
	cpms := Campaign.LoadCampaign(database.Db)
	if cpms == nil {
		return fmt.Errorf("no Campaign declared")
	}

	ca.Campaign = cpms

	ents := Event.LoadEvents(database.Db)
	if ents == nil {
		return fmt.Errorf("no Events declared")
	}
	ca.Event = ents
	db.Push(database)
	ph := Controller.New(log, validation, db, &ca, config, &h)
	sm.Use(ph.HttpMiddleware)

	//ID use as password
	getR := sm.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/scores/{id:[0-9]+}", ph.LoadChanges)

	postR := sm.Methods(http.MethodPost).Subrouter()
	postR.HandleFunc("/scores", ph.AddScore)
	postR.HandleFunc("/campaign", ph.AddCampaign)
	postR.HandleFunc("/event", ph.AddEvent)
	return nil

}
