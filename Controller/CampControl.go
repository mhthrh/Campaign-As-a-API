package Controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mhthrh/Campaign/Model/Campaign"
	"github.com/mhthrh/Campaign/Model/DashBoard"
	"github.com/mhthrh/Campaign/Model/Event"
	"github.com/mhthrh/Campaign/Model/Input"
	"github.com/mhthrh/Campaign/Model/Result"
	"github.com/mhthrh/Campaign/Utilitys/CryptoUtil"
	"github.com/mhthrh/Campaign/Utilitys/DbUtil/PgSql"
	"github.com/mhthrh/Campaign/Utilitys/HttpUtil"
	"github.com/mhthrh/Campaign/Utilitys/JsonUtil"
	"io"
	"net/http"
)

func (b *Controller) AddScore(rw http.ResponseWriter, r *http.Request) {
	find := false
	inp := r.Context().Value(Key{}).(*Input.Input)
	for _, i2 := range b.cam.Campaign {
		if i2.ID == inp.CampaignId {
			find = !find
			continue
		}
	}
	if !find {
		Result.New((*inp).UserName, 0, http.StatusOK, "campaign not found", "").SendResponse(&rw)
	}
	find = !find

	for _, i3 := range b.cam.Event {
		if i3.ID == inp.EventId && i3.CampaignId == inp.CampaignId {
			find = !find
			continue
		}
	}
	if !find {
		Result.New((*inp).UserName, 0, http.StatusOK, "event not found", "").SendResponse(&rw)
	}
	db := b.db.Pull()
	_, err := PgSql.ExecuteCommand(fmt.Sprintf("INSERT INTO public.\"Scores\"(\"ID\", \"CampaignId\", \"EventId\", \"DateTime\", \"User\", \"UserEmail\")VALUES ('%d', '%d', '%d', '%s', '%s', '%s')", inp.Id, inp.CampaignId, inp.EventId, "Date", inp.UserName, inp.Email), db.Db)
	b.db.Push(db)

	if err != nil {
		Result.New((*inp).UserName, 0, http.StatusOK, "score not registered", "").SendResponse(&rw)
	}

	go func(h string) {
		duration := []int{0, 15, 30}
		db := b.db.Pull()
		dash := DashBoard.ReadResult(db.Db, duration)
		k := CryptoUtil.NewKey()
		k.Text = JsonUtil.New(nil, nil).Struct2Json(dash)
		k.Sha256()
		result := k.Result
		if h != result {
			http.Post("https:// host:8585/dashboard", "application/json", func() io.Reader {
				var data bytes.Buffer
				json.NewEncoder(&data).Encode(dash)
				return &data
			}())
			b.hash = &result
		}

	}(*b.hash)

}

func (b *Controller) LoadChanges(w http.ResponseWriter, r *http.Request) {
	id := HttpUtil.GetID(r)
	if id != 123456 {
		Result.New("", 0, http.StatusForbidden, "Access deny", "").SendResponse(&w)
	}
	db := b.db.Pull()
	b.cam.Campaign = Campaign.LoadCampaign(db.Db)
	b.cam.Event = Event.LoadEvents(db.Db)
	b.db.Push(db)

	Result.New("", 0, http.StatusOK, "Success", "").SendResponse(&w)

}

func (b *Controller) AddCampaign(w http.ResponseWriter, r *http.Request) {
	Result.New("", 0, http.StatusNotImplemented, "Not implemented", "").SendResponse(&w)
}

func (b *Controller) AddEvent(w http.ResponseWriter, r *http.Request) {
	Result.New("", 0, http.StatusNotImplemented, "Not implemented", "").SendResponse(&w)
}
