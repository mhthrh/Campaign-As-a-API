package Controller

import (
	"context"
	"errors"
	"fmt"
	"github.com/mhthrh/Campaign/Model/Input"
	"net/http"
	"time"
)

func (b *Controller) HttpMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fErr := func(err error, i int, in interface{}) {
			start := time.Now()
			b.l.WithFields(map[string]interface{}{
				"method":     r.Method,
				"path":       r.URL,
				"status":     nil,
				"latency_ns": time.Since(start).Nanoseconds(),
			}).Info("request details")
			//DashBoard.New(0, -1, http.StatusBadRequest, "UnSuccess", JsonUtil.New(nil, nil).Struct2Json(&in)).SendResponse(&w)
		}
		b.l.WithFields(map[string]interface{}{
			"method":  r.Method,
			"path":    r.URL,
			"status":  nil,
			"Message": "received request ",
		})
		fNext := func(in interface{}) {
			ctx := context.WithValue(r.Context(), Key{}, in)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		}

		if r.Host != fmt.Sprintf("%s:%d", b.Conf.Server.IP, b.Conf.Server.Port) {
			err := errors.New("access denied")
			fErr(err, http.StatusForbidden, GenericError{Message: err.Error()})
			return
		}
		if r.Method == http.MethodGet {
			fNext(nil)
		}

		inp, err := Input.ToInput(r.Body)
		if err != nil {
			fErr(err, http.StatusBadRequest, GenericError{Message: err.Error()})
			return
		}

		errs := b.v.Validate(inp)
		if len(errs) != 0 {
			fErr(errors.New("validation issue"), http.StatusUnprocessableEntity, ValidationError{Messages: errs.Errors()})
			return
		}
		fNext(inp)
	})
}
