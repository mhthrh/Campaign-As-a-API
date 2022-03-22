package Controller

import (
	"fmt"
	"github.com/mhthrh/Campaign/Model/Campaign"
	"github.com/mhthrh/Campaign/Model/Event"
	"github.com/mhthrh/Campaign/Utilitys/ConfigUtil"
	"github.com/mhthrh/Campaign/Utilitys/DbUtil/DbPool"
	"github.com/mhthrh/Campaign/Utilitys/ValidationUtil"
	"github.com/sirupsen/logrus"
)

type Key struct{}
type CampaignHandler struct {
	Campaign []Campaign.Campaign
	Event    []Event.Event
}
type Controller struct {
	l    *logrus.Entry
	v    *ValidationUtil.Validation
	db   *DbPool.DBs
	cam  *CampaignHandler
	Conf *ConfigUtil.Config
	hash *string
}

var InvalidPath = fmt.Errorf("invalid Path, path must be /ViewControler/[id]")

type GenericError struct {
	Message string `json:"message"`
}

type ValidationError struct {
	Messages []string `json:"messages"`
}

func New(l *logrus.Entry, v *ValidationUtil.Validation, db *DbPool.DBs, c *CampaignHandler, con *ConfigUtil.Config, hsh *string) *Controller {
	return &Controller{l, v, db, c, con, hsh}
}
