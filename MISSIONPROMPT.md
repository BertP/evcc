# MISSION PROMPT: evcc Development & Deployment

## 🎯 Mission Statement
To customize, maintain, and deploy a robust and extended instance of **evcc** (Electric Vehicle Charge Controller) to optimize energy usage, specifically focusing on [insert specific goal, e.g., integrating a specific inverter or solar setup].

The functionality should be extended to include the following:
* adding White goods appliances new device category
* implement te to treat a white good appliance a variable load
* implementing the load shifting logic
* implement the funtionality to control the white good appliances via a REST API
* take in account, that the REST API will deliver more then one white good appliance, so the load shifting logic should be able to handle multiple white good appliances
* implementing a seperate integration for the REST-basedMiele 3rd Party API
* 

## 🏗️ Technical Architecture
* **Core:** Go (Golang)
* **Frontend:** JavaScript / Vue.js
* **Infrastructure:** Docker Compose
* **Database:** SQLite / InfluxDB (Standard evcc stack)

## 💻 Environment & Workflow
* **Development OS:** Windows 11 (using Bash/Git Bash)
* **Deployment Target:** Ubuntu Server
* **Orchestration:** Docker Compose stack
* **Critical Requirement:** Maintain `LF` (Linux) line endings for all configuration and shell files to ensure compatibility between the Windows dev environment and the Ubuntu production server.

## 🛠️ Key Objectives
1.  **Local Development:** Configure and test the `evcc.yaml` setup on Windows 11.
2.  **Containerization:** Ensure the Dockerfile and `docker-compose.yml` are optimized for the Ubuntu target.
3.  **Integration:** [Optional: List your specific hardware here, e.g., Tesla Wall Connector, SMA Inverter, etc.]
4.  **Deployment:** Streamline the transition from the local Windows environment to the remote Ubuntu server.

## 📜 Rules of Engagement
* **Simplicity First:** Prefer standard evcc configurations over complex custom code where possible. Implement step by step.
* **Testing:** Test every step on the server side. No local testing. Tests will be done with real devices
* **Documentation:** All logic changes or custom meter/charger configurations must be documented in this repository.
* **Container Parity:** The local Docker environment must mirror the production Ubuntu environment as closely as possible.

## 