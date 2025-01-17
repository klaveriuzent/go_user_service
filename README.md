# go_user_service
This is a Golang project that provides a foundation for building a Back-End with Golang and Postgre. The project uses Golang version 1.18 and is integrated with Postgre.

## Branch
Always check other branch
### main
[main](https://github.com/klaveriuzent/go_user_service/tree/main) | Role-Based Access Control (RBAC) & No Swagger Integrated (recommend for fullstack)
### swagger
[swagger](https://github.com/klaveriuzent/go_user_service/tree/swagger) | Role-Based Access Control (RBAC) & Swagger Integrated (recommend for team base)
### Initial (Always UPDATE)
[Initial](https://github.com/klaveriuzent/go_user_service/tree/Initial) | Authentication Only & Swagger Integrated (recommend for other project with no rbac)

## Requirement
`PostgreSQL` and `golang` on your computer.

## Getting Started
To get started, clone the [**repository**](https://github.com/klaveriuzent/go_user_service) to your local machine and navigate to the project directory. Then, edit a `.env.local` file in the project directory and set the following environment variables:

```
DB_HOST=localhost
DB_PORT=5432
DB_NAME=<database_name>
DB_USER=<database_username>
DB_PASSWORD=<database_password>
```

Replace `<database_name>`, `<database_username>`, and `<database_password>` with the appropriate values for your PostgreSQL database.

Finally, run the following command to start the server:
```
go run main.go
```

The server will be running on
```
http://localhost:8000/swagger/index.html#/
```
and you can use Swagger to test the API endpoints.

## Database Management
To manage the database, you can use a database management tool such as DBeaver. Connect to the PostgreSQL database using the connection details specified in the `.env.local` file (other option you can use `.env` with edit `main.go`). Once connected, you can create tables, insert data, and perform other database operations as needed.

## Contributing
If you would like to contribute to this project, please fork the repository and submit a pull request. We welcome contributions of all kinds, including bug fixes, feature additions, and documentation improvements.
