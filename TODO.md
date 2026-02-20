# TODO


# ✅ TODO: Miele 3rd-Party API Integration

## Step 1: Establish API Connection

- [x] Establish a connection to the **Miele 3rd-Party API**
- [x] Retrieve the **access token**
- [x] Implement **access token refresh**

## Authorization & Configuration

- [x] Follow the authorization concept defined in `mieleApiAuthorization.md`
- [x] Store **Client ID** and **Client Secret** in a **separate configuration file**
- [x] Ensure **Client ID** and **Client Secret** are **not** stored in `evcc.yaml`

## Integration Design

- [x] Implement the API connection as a **separate integration**
- [x] Ensure the integration can:
  - [x] Retrieve the access token
  - [x] Refresh the access token

## UI / Integration Card

- [x] Display the current **connection status** on the integration card
- [x] Provide an option to **disconnect** the API connection via the integration card


## Step 2: Build Miele Appliance Konfiguration Card Connection
