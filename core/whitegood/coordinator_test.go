package whitegood

import (
	"testing"

	"github.com/evcc-io/evcc/api"
	"github.com/evcc-io/evcc/core/site"
	"github.com/evcc-io/evcc/util"
	"github.com/stretchr/testify/assert"
)

type mockSite struct {
	site.API
	gridPower float64
}

func (s *mockSite) GetGridPower() float64 {
	return s.gridPower
}

type mockAppliance struct {
	status  api.WhitegoodState
	started bool
	power   float64
}

func (m *mockAppliance) Status() (api.WhitegoodState, error) {
	return m.status, nil
}

func (m *mockAppliance) Start() error {
	m.started = true
	m.status = api.WhitegoodStateRunning
	return nil
}

func (m *mockAppliance) RequiredPower() float64 {
	return m.power
}

func TestCoordinatorSurplus(t *testing.T) {
	site := &mockSite{gridPower: -1500} // exactly 1500W surplus

	app1 := &mockAppliance{status: api.WhitegoodStateIdle, power: 1000}
	app2 := &mockAppliance{status: api.WhitegoodStateIdle, power: 800}

	c := &Coordinator{
		site: site,
		log:  util.NewLogger("test"),
	}

	// Create mock appliances
	c.appliances = append(c.appliances, &Appliance{Name: "app1", Device: app1})
	c.appliances = append(c.appliances, &Appliance{Name: "app2", Device: app2})

	// Before: none running
	assert.False(t, app1.started)
	assert.False(t, app2.started)

	// Run checkSurplus
	c.checkSurplus()

	// Should start the first appliance because 1500W >= 1000W
	assert.True(t, app1.started)

	// App1 takes 1000W, leaving 500W surplus.
	// App2 requires 800W, so it should NOT start!
	assert.False(t, app2.started)
}
