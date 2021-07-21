# User service
by me

## Development
Requirements:
- Go 1.16;
- Latest Docker engine;

Start docker-compose cluster with `make start-dev` and shut it down with `make stop-dev`. Changes in `internal/users` will force both grpc and http servers to reload, changes outside the source folder require restarting the cluster.

`make run-tests` will run unit tests locally (this requires having mongodb instance on localhost:27017, otherwise repository tests will fail)

## gRPC API
When **dev.env** is used for server configuration, [reflection API](https://github.com/grpc/grpc/blob/master/doc/server-reflection.md) will be enabled, so the endpoints can be discovered by compatible client tools (e.g. [grpcurl](https://github.com/fullstorydev/grpcurl)).

## HTTP API
### Create User
**POST** `/users`

Request body fields (all fields required):
- **name** - user's first name, string;
- **last_name** - user's last name, string;
- **age** - user's age, number, should be 16 or more;
Example:
```json
{
    "name": "John",
    "last_name": "Doe",
    "age": 30
}
```
Response types:
- **201** - OK;
```json
{
    "user_id": "some1234id5678"
}
```
- **400** - one or more request fields do not meet constraints;
- **500** - great, you it died and that's probably your fault;

### Get User By Id
**GET** `/users/{id}`
Example:
```sh
    curl -i \                                        
        -H "Content-Type: application/json" \
        -X POST -d '{"name":"John","lastName":"Doe","age":40}' \
        localhost:8080/users
```
Response types:
- **200** - OK;
```json
{
    "id": "some1234id5678",
    "name": "John",
    "last_name": "Doe",
    "age": 40,
    "recording_date": "05-05-2021T20:30:15"
}
```
- **404** - user with given ID not found;


### Get List Of Users
**GET** `/users`
Response:
- **200** - OK

