# Prerequisites

- [Installing docker](https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-on-ubuntu-22-04)
- [Installing docker-compose](https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-compose-on-ubuntu-22-04)

# Setup & Running

## 1. Clone the repo
```bash
git clone 
```
## 2. Go to the directory
```bash
cd gymshark
```
## 3. Run docker-compose to up the services
```bash
docker-compose up --build -d web
```
If you need to run the test, you can run
```bash
docker-compose up --build -d test
```

## 4. Validate the service is up
After executing the docker-compose command, your server will be up and running with port **8080**.

To Validate if it's running, run in your terminal:
```bash
curl http://localhost:8080/healthcheck
```
You should get ```OK``` in the response

# Request Format

## Endpoint
```
GET /store/packs
```
## Body
```json
{
    "num_of_items": 501 // integer
}
```

## Response
```json
/*
{ 
    pack1_size: num of pack1,
    pack2_size: num of pack2,
    ...
*/
{
    "250": 1,
    "500": 1
}
```

## Example
Run in your terminal:
```bash
curl http://localhost:8080/store/packs -d '{"num_of_items": 501}'
```
