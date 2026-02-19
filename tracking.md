# Project Change Tracking

This document tracks all significant changes and commits made to the project, including timestamps.

| Timestamp        | Description                                     | Files Changed                                                     |
| :--------------- | :---------------------------------------------- | :---------------------------------------------------------------- |
| 2026-02-19 15:45 | Initialized project deployment files            | `evcc.yaml`, `docker-compose.yml`, `DEPLOYMENT.md`                |
| 2026-02-19 15:52 | Fixed git clone command in DEPLOYMENT.md        | `DEPLOYMENT.md`                                                   |
| 2026-02-19 15:55 | Committed deployment configuration              | `evcc.yaml`, `docker-compose.yml`, `DEPLOYMENT.md`, `tracking.md` |
| 2026-02-19 15:58 | Force added simulation evcc.yaml                | `evcc.yaml`                                                       |
| 2026-02-19 15:59 | Synced repository with remote                   | All committed files                                               |
| 2026-02-19 16:15 | Fixed database path in docker-compose.yml       | `docker-compose.yml`                                              |
| 2026-02-19 16:20 | Updated troubleshooting in DEPLOYMENT.md        | `DEPLOYMENT.md`                                                   |
| 2026-02-19 16:22 | Removed demo configuration from evcc.yaml       | `evcc.yaml`                                                       |
| 2026-02-19 16:45 | Created Miele OAuth2 provider client            | `provider/miele/client.go`                                        |
| 2026-02-19 16:50 | Created Miele credentials loader and controller | `provider/miele/credentials.go`, `util/miele/controller.go`       |
| 2026-02-19 16:55 | implemented auth flow in Miele controller       | `util/miele/controller.go`                                        |
| 2026-02-19 17:15 | Wired Miele integration in cmd/root.go          | `cmd/root.go`                                                     |
| 2026-02-19 17:20 | Fixed compilation for Miele wiring              | `cmd/setup.go`                                                    |
| 2026-02-19 17:35 | Added Miele Integration Card to UI              | `assets/js/views/Config.vue`                                      |
