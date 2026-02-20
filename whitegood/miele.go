package whitegood

import (
	"context"
	"errors"
	"fmt"

	"github.com/evcc-io/evcc/api"
	"github.com/evcc-io/evcc/util"
	"github.com/evcc-io/evcc/util/miele"
)

func init() {
	registry.AddCtx("miele", NewMieleFromConfig)
}

// Miele appliance wrapper
type Miele struct {
	log      *util.Logger
	deviceID string
}

// NewMieleFromConfig creates Miele whitegood from config
func NewMieleFromConfig(ctx context.Context, other map[string]any) (api.Whitegood, error) {
	var cc struct {
		Device string
	}

	if err := util.DecodeOther(other, &cc); err != nil {
		return nil, err
	}

	if cc.Device == "" {
		return nil, errors.New("missing device")
	}

	w := &Miele{
		log:      util.NewLogger(fmt.Sprintf("miele-%s", cc.Device)),
		deviceID: cc.Device,
	}

	return w, nil
}

func (m *Miele) Status() (api.WhitegoodState, error) {
	if miele.Instance == nil || !miele.Instance.IsConnected() {
		return api.WhitegoodStateError, errors.New("miele not connected")
	}

	// Just treating everything as idle for now until provider API is fleshed out
	return api.WhitegoodStateIdle, nil
}

func (m *Miele) Start() error {
	if miele.Instance == nil || !miele.Instance.IsConnected() {
		return errors.New("miele not connected")
	}

	return miele.Instance.StartDevice(context.Background(), m.deviceID)
}

func (m *Miele) RequiredPower() float64 {
	// A standard default threshold, e.g., 1kW for a washing machine cycle
	return 1000.0
}
