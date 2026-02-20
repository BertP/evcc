package whitegood

import (
	"fmt"
	"sync"
	"time"

	"github.com/evcc-io/evcc/api"
	"github.com/evcc-io/evcc/core/site"
	"github.com/evcc-io/evcc/util"
)

// Appliance represents a tracked whitegood unit
type Appliance struct {
	Name    string
	Device  api.Whitegood
	Running bool
}

// Coordinator runs a background loop checking grid surplus
type Coordinator struct {
	log        *util.Logger
	site       site.API
	appliances []*Appliance
	mu         sync.Mutex
}

// NewCoordinator creates a new Whitegoods orchestrator
func NewCoordinator(site site.API, whitegoods []api.Whitegood, names []string) *Coordinator {
	log := util.NewLogger("whitegoods")

	var apps []*Appliance
	for i, dev := range whitegoods {
		name := fmt.Sprintf("whitegood-%d", i)
		if i < len(names) {
			name = names[i]
		}
		apps = append(apps, &Appliance{
			Name:   name,
			Device: dev,
		})
	}

	return &Coordinator{
		log:        log,
		site:       site,
		appliances: apps,
	}
}

// Run initiates the background loop
func (c *Coordinator) Run(stopC <-chan struct{}) {
	c.log.INFO.Println("starting whitegoods coordinator")
	go c.loop(stopC)
}

func (c *Coordinator) loop(stopC <-chan struct{}) {
	for {
		select {
		case <-stopC:
			return
		case <-time.After(1 * time.Minute):
			c.checkSurplus()
		}
	}
}

func (c *Coordinator) checkSurplus() {
	c.mu.Lock()
	defer c.mu.Unlock()

	grid := c.site.GetGridPower()
	if grid >= -100 { // require at least 100W surplus (grid power negative)
		return
	}

	surplus := -grid

	for _, app := range c.appliances {
		state, err := app.Device.Status()
		if err != nil {
			c.log.ERROR.Printf("appliance %s status error: %v", app.Name, err)
			continue
		}

		if state == api.WhitegoodStateIdle && surplus >= app.Device.RequiredPower() {
			c.log.INFO.Printf("starting appliance %s (surplus: %.0fW, required: %.0fW)", app.Name, surplus, app.Device.RequiredPower())
			if err := app.Device.Start(); err != nil {
				c.log.ERROR.Printf("failed to start %s: %v", app.Name, err)
			} else {
				// Prevent immediate retriggers and account for the shifted load
				surplus -= app.Device.RequiredPower()
			}
		}
	}

	// Publish current states to the UI
	c.site.Publish("whitegoods", c.Appliances())
}

// Appliances returns the current list of appliances and states
func (c *Coordinator) Appliances() []map[string]any {
	c.mu.Lock()
	defer c.mu.Unlock()

	var res []map[string]any
	for _, app := range c.appliances {
		state, _ := app.Device.Status() // ignore error for API presentation
		res = append(res, map[string]any{
			"name":          app.Name,
			"status":        state,
			"requiredPower": app.Device.RequiredPower(),
		})
	}

	return res
}

// StartAppliance manually overrides the load shifting to start an appliance right now
func (c *Coordinator) StartAppliance(name string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, app := range c.appliances {
		if app.Name == name {
			return app.Device.Start()
		}
	}

	return fmt.Errorf("appliance not found: %s", name)
}
