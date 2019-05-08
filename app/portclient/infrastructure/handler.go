package infrastructure

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bkatrenko/ports/app/portclient/domain/service"
	"github.com/gorilla/mux"
	wrapper "github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// APIHandler contains router and client app service to handle incoming requests
type APIHandler struct {
	router  *mux.Router
	service *service.ClientPortsService
}

// NewAPIHandler is constructor for handlers structure which should init routes for some version
func NewAPIHandler(service *service.ClientPortsService) (*APIHandler, error) {
	router := mux.NewRouter()

	handler := &APIHandler{router: router}
	handler.addV1Routes()
	handler.service = service

	return handler, nil
}

func (h *APIHandler) addV1Routes() {
	r := h.router.PathPrefix("/v1").Subrouter()
	r.HandleFunc("/ports", h.getPortsV1).Methods("GET")
}

func (h *APIHandler) getPortsV1(w http.ResponseWriter, r *http.Request) {
	offset, err := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 64)
	if err != nil {
		h.writeError(w, http.StatusBadRequest, "bad offset")
		return
	}

	limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 64)
	if err != nil {
		h.writeError(w, http.StatusBadRequest, "bad limit")
		return
	}

	ports, err := h.service.RetrievePorts(offset, limit)
	if err != nil {
		log.Error(err)
		h.writeError(w, http.StatusInternalServerError, "internal error")
		return
	}

	h.writeJSON(w, http.StatusOK, ports)
}

func (h *APIHandler) writeError(w http.ResponseWriter, code int, err string) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(code)
	w.Write([]byte(fmt.Sprintf("{\"error\": \"%s\"}", err)))
}

func (h *APIHandler) writeJSON(w http.ResponseWriter, conte int, object interface{}) error {
	w.Header().Add("Content-type", "application/json")
	if err := json.NewEncoder(w).Encode(object); err != nil {
		return wrapper.Wrap(err, "error while write json")
	}

	return nil
}
