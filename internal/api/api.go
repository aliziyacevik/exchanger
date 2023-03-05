package api

import (
	"net/http"
	"encoding/json"
	"log"

	s"github.com/aliziyacevik/exchanger/internal/service"

)


type Handler interface {
	Get(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
}

type handler struct {
	service		s.Service
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

func AllowMethods(fn func(http.ResponseWriter, *http.Request), listOfAllowedMethods ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		isAllowed := false
		for _, method := range listOfAllowedMethods {
			if r.Method == method {
				isAllowed = true
				break
			}
		}

		if isAllowed == false {
			log.Println("Method not allowed")
			http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
			return
		}
	
	fn(w, r)
	}

}


func NewHandler(service s.Service) Handler {
	return &handler{
		service:	service,
	}
}



