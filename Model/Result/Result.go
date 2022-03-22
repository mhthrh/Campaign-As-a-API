package Result

import (
	"github.com/mhthrh/Campaign/Utilitys/JsonUtil"
	"net/http"
	"time"
)

type Response struct {
	Header struct {
		ID     string
		Time   time.Time
		Status struct {
			Stat int
			Desc string
		}
	}
	Body struct {
		Code   int
		Result string
	}
}

func New(id string, stat, code int, desc, result string) *Response {
	r := new(Response)
	r.Header.ID = id
	r.Header.Time = time.Now()
	r.Header.Status.Stat = stat
	r.Header.Status.Desc = desc
	r.Body.Code = code
	r.Body.Result = result
	return r
}
func (r *Response) SendResponse(w *http.ResponseWriter) {
	(*w).Header().Set("Result", JsonUtil.New(nil, nil).Struct2Json(r.Header))
	(*w).Write([]byte(r.Body.Result))
	(*w).WriteHeader(r.Body.Code)
}
