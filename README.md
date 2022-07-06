## Demo configs for Envoy, Kong, KrakenD, Nginx

Each gateway is mounted to the localhost ports 8000-8003
They are configured with a self signed certificate and accepting https.

There are 2 services mounted to the gateways.

`/dogs` endpoint will route the request to the dog service that returns a random breed. The `/cats` implement the feline counterpart.

## Requirements

- Docker
- Docker compose

## How to use
- After checkout run `docker-compose up`
- `curl -kv https://localhost:8000/dogs` (try different ports for the different gateways, or the /cats endpoint for a cat.
