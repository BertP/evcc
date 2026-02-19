# evcc Deployment Guide for Ubuntu Server

This guide outlines the steps to deploy `evcc` on an Ubuntu Server environment, ensuring parity with the development setup.

## Prerequisites

- **Ubuntu Server**: Ensure you have SSH access and `sudo` privileges.
- **Docker & Docker Compose**: Installed on the server.
- **Git**: Installed to clone the repository.

## Deployment Steps

### 1. Clone the Repository

Clone the project repository to your desired location on the server (e.g., `/opt/evcc` or `~/evcc`).

```bash
git clone https://github.com/BertP/evcc.git -b master
cd evcc
```

### 2. Configure evcc using config wizard

If you haven't already configured `evcc.yaml`, you can copy the example or create a new one.

```bash
# To run the configuration wizard
docker run -v $(pwd)/evcc.yaml:/etc/evcc.yaml -it evcc/evcc:latest evcc configure
```

> **Note:** The repository includes a simulation `evcc.yaml` for testing purposes. For production, ensure you configure your actual hardware.

### 3. Start the Service

Use Docker Compose to build and start the application.

```bash
# Build (optional if using pre-built image) and start in detached mode
docker compose up -d --build
```

### 4. Verify Deployment

Access the evcc interface via your browser:

`http://<SERVER_IP>:7070`

Check the logs to ensure everything is running smoothly:

```bash
docker compose logs -f
```

## Maintenance

### Updating

To update to the latest version:

```bash
git pull
docker compose pull
docker compose up -d
```

### Troubleshooting

- **USB Devices**: If using USB adapters (RS485), ensure they are passed through to the container or use `network_mode: host`.
- **Permission Issues**: Check `PUID` and `PGID` in `docker-compose.yml` if you encounter file permission errors with the database or config.
