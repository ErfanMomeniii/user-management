# user management
A Simple user management service with Go

## Installation

Install project with docker compose

```bash
    docker compose build
    docker compose up -d
    docker compose exec user-management ./localizer migrate
```
also you can install without docker compose like below
```bash
    go build .
    ./user-management migrate
    ./user-management start
```

## Commands
- ### start
 ```bash
 start           for starting the application
 ```
#### Flags
```bash 
Flags:
  -h, --help                help for start

Global Flags:
  -c, --config    string    Config file path
```
- ### migrate

```bash
migrate         for running migration on the database
```

#### Flags
``` bash
Flags:
  -h, --help              help for migrate
  -p, --path    string    path to migrations directory (default "./migrations")
  -s, --steps   int       number of steps to migrate. positive steps for up and negative steps for down. zero to upgrade all.
  -t, --table   string    database table holding migrations (default "schema_migrations")

Global Flags:
  -c, --config string   Config file path
```
## Run Tests

```sh
go test ./... -v
```

## API Reference

#### Get all users filtered by the country that paginated

```http
  GET /v1/users?page=${page}&country=${country}
```

| QueryParameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `page`      | `int` | page of the user list|
| `country`      | `string` | **optional**. Country of the user|

#### Get user by id

```http
  GET /v1/user/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of user to fetch |

#### Save new user

```http
  Post /v1/user
```
example of a request body:
```sh
[
  {
    "first_name":"a",
    "last_name":"b",
    "nickname":"c",
    "password":"12345678",
    "email":"abc@def.com",
    "country":"UK"
    }
  ]
```

#### Delete user

```http
  DELETE /v1/user/${id}
```

#### Update user information

```http
  PUT /v1/user/${id}
```
`
Example of a request body:
```sh
[
  {
    "first_name":"a",
    "last_name":"b",
    "nickname":"c",
    "password":"12345678",
    "email":"abc@def.com",
    "country":"UK"
    }
  ]
```


