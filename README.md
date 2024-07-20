# Go OIDC Keycloak

This project is a simple example how to use Keycloak as an OIDC provider for your Go project.

## Features

- Go backend with an section for authenticated users
- Keycloak with postgres database
- Auto-import of configured Keycloak realm
- Configuration via .env file

## Requirements

- Docker
- Docker Compose
- Go

## Setup

1. Run `docker-compose up` and wait until Postgres and Keycloak are successfully started
2. Login to your Keycloak with the URL `http://localhost:8080`. The default admin username and password are `admin` and `admin1234`
3. Switch to the `application` realm on the top left
4. Go to `Clients` > `application` > `Credentials` and click on `Regenerate secret`
5. Copy the `example.env` to `.env` and insert your just generated secret as `OIDC_CLIENT_SECRET`
6. Go to `Users` and click `Create user`
7. Create a user and set a password under the tab `Credentials` after creation of the user

### Create a new realm and client

The following steps are necessary if you want to create a new realm. Please keep in mind that you have to modify the configuration accordingly to use your new realm.

1. Click `Create realm` on the top left realm menu
2. Set a realm name and click `Create`
3. Go to `Clients` and choose `Create client`
4. Set a `Client ID` and click `Next`
5. Activate `Client authentication` and click `Next`
6. Set your valid redirect URIs. This can be `*` for any URI. But it's not recommended for production!
7. Click `Save to create your new client`

### Export realm settings

You can export your realm settings at anytime for the auto-import process. Just go into the `Realm settings` of your Keycloak realm and choose the `Partial export` action on the top right. The users and client secrets are not included in the realm export.

## Usage

1. Open `http://localhost:8081` in your browser
2. Login with your Keycloak user
3. Copy your `access_token`
4. Call `http://localhost:8081` with your `access_token` as bearer token in Postman or curl
5. Now you should get the string `authenticated`. That means you're successfully authenticated! 

## Credits

See https://stackoverflow.com/questions/48855122/keycloak-adaptor-for-golang-application for further information.

## Disclaimer

This project is just an example to show how to use Keycloak as an OIDC provider for Go. Don't use this for production. Just use it as inspiration. No matter what you're doing: Change the default credentials I used in the `docker-compose.yml`!