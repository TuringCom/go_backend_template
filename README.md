# Turing Back End Challenge

To complete this challenge, you need to ensure all route returns a similar response object as described in our API guide.
To achieve this goal

- You will have to fix the existing bugs
- Implement the incomplete functions,
- Add test cases for the main functions of the system.
- Use Dockerfile to the root of the project to run the app in docker environment

## Getting started

### Prerequisites

In order to install and run this project locally, you would need to have the following installed on you local machine.

- [**Golang**](https://golang.org/doc/install)
- [**MySQL**](https://www.mysql.com/downloads/)
- [**Govendor**]([https://https://bundler.io/](https://github.com/kardianos/govendor))

### Installation

- Clone this repository

- Navigate to the project directory
...
- Create a MySQL database and run the `sql` file in the database directory to migrate the database

```sh
  mysql -u <db_user_name> -D <db_name> -p < db/dump.sql
```

- Run `go run main.go` to start the app in development

## Docker

- Build image

`docker build -t go_application .`

- Run container
`docker run -it -p 8000:8000 -v $(pwd):/backend go_application`

## Request and Response Object API guide for all Endpoints

- Check [here](https://docs.google.com/document/d/1J12z1vPo8S5VEmcHGNejjJBOcqmPrr6RSQNdL58qJyE/edit?usp=sharing)
- Visit `http://127.0.0.1:8000`