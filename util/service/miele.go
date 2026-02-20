package service

import (
	"context"
	"net/http"

	"github.com/evcc-io/evcc/server/service"
	"github.com/evcc-io/evcc/util/miele"
)

func init() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /devices", getMieleDevices)

	service.Register("miele", mux)
}

func getMieleDevices(w http.ResponseWriter, req *http.Request) {
	if miele.Instance == nil {
		jsonWrite(w, []string{})
		return
	}

	devices, err := miele.Instance.GetDevices(context.Background())
	if err != nil {
		jsonError(w, http.StatusInternalServerError, err)
		return
	}

	var res []string
	for _, dev := range devices {
		res = append(res, dev.Ident.DeviceSN)
	}

	jsonWrite(w, res)
}
