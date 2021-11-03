package category_nurzhas_store

import (
	"encoding/json"
	setdata_common "github.com/kirigaikabuto/setdata-common"
	"io/ioutil"
	"net/http"
)

type OrderHttpEndpoints interface {
	MakeCreateOrderEndpoint() func(w http.ResponseWriter, r *http.Request)
	MakeListOrderEndpoint() func(w http.ResponseWriter, r *http.Request)
}

type orderHttpEndpoints struct {
	ch setdata_common.CommandHandler
}

func NewOrderHttpEndpoints(ch setdata_common.CommandHandler) OrderHttpEndpoints {
	return &orderHttpEndpoints{ch: ch}
}

func (o *orderHttpEndpoints) MakeCreateOrderEndpoint() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		setupResponse(&w, r)
		cmd := &CreateOrderCommand{}
		dataBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		err = json.Unmarshal(dataBytes, &cmd)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		response, err := o.ch.ExecCommand(cmd)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		respondJSON(w, http.StatusCreated, response)
	}
}

func (o *orderHttpEndpoints) MakeListOrderEndpoint() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		setupResponse(&w, r)
		cmd := &ListOrderCommand{}
		response, err := o.ch.ExecCommand(cmd)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		respondJSON(w, http.StatusCreated, response)
	}
}
