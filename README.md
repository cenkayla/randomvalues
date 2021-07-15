# Vote service

- Written in Go
- Stores data in PostgreSQL with gorm

# Quickstart
```shell
git clone https://github.com/cenkayla/randomvalues.git
cd randomvalues
go run main.go
```
# Usage

To create n digit random value
```shell
$ curl --request POST \
	--url "localhost:8080/api/generate" \
	--header "Content-Type: application/json" \
	--data "{\"number\":5}"
```
Response
```shell
"Succesfully generated a value"
```

To retrieve random value with uuid
```shell
$ curl --request POST \
	--url "localhost:8080/api/retrieve" \
	--header "Content-Type: application/json" \
	--data "{\"uuid\":\"3a0d0b41-8302-46b6-b55f-10379fdfba5c\"}"
```
Response
```shell
{
  "UUID": "3a0d0b41-8302-46b6-b55f-10379fdfba5c",
  "Name": "8WCU9"
}
```

