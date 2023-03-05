package api

import (
	"net/http"
	"encoding/json"

	s"github.com/aliziyacevik/exchanger/internal/service"

)


type Handler interface {
	Get(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
}

type handler struct {
	service		s.Service
	mux		http.Handler
}

func (h *handler) Post(w http.ResponseWriter, r*http.Request) {
           	var q s.Query

                 err := json.NewDecoder(r.Body).Decode(&q)
                 if err != nil {
                        http.Error(w, err.Error(), http.StatusBadRequest)
                    }
                 res, err := h.service.Convert(q)
                 resJson, err := json.Marshal(res)
                 if err != nil {
                     http.Error(w, err.Error(), http.StatusInternalServerError)
                     }
                 w.Write([]byte(resJson))
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ratatata"))
}


func NewHandler(service s.Service, h http.Handler) Handler {
	return &handler{
		service:	service,
		mux:		h,
	}
}



