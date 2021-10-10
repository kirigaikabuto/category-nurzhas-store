package category_nurzhas_store

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	setdata_common "github.com/kirigaikabuto/setdata-common"
	"io"
	"net/http"
	"strings"
)

type HttpEndpoints interface {
	MakeUploadPricesFile() func(w http.ResponseWriter, r *http.Request)
}

type httpEndpoints struct {
	ch setdata_common.CommandHandler
}

func NewHttpEndpoints(ch setdata_common.CommandHandler) HttpEndpoints {
	return &httpEndpoints{ch: ch}
}

func (h *httpEndpoints) MakeUploadPricesFile() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		setupResponse(&w, r)
		cmd := &UploadPricesFileCommand{}
		cmd.Name = r.URL.Query().Get("name")
		if cmd.Name == "" {
			respondJSON(w, http.StatusInternalServerError, errors.New("please need name of file"))
			return
		}
		buf := bytes.NewBuffer(nil)
		file, header, err := r.FormFile("fileupload")
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, err.Error())
			return
		}
		name := strings.Split(header.Filename, ".")
		fmt.Printf("File name %s\n", name[0])
		_, err = io.Copy(buf, file)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, err.Error())
			return
		}
		err = file.Close()
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, err.Error())
			return
		}
		cmd.File = buf
		resp, err := h.ch.ExecCommand(cmd)

		if err != nil {
			respondJSON(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondJSON(w, http.StatusOK, resp)
	}
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Allow-Origin")
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}
