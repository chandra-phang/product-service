# product-service

**Microservice for Managing Product Operations**

## Getting Started

### 1. Set Up Environment Variables

Create a `.env` file with the following configurations:

```bash
DB_HOST=localhost
DB_PORT=3306
DB_USER={{YOUR-DB-USER}}
DB_PASSWORD={{YOUR-DB-PASSWORD}}
DB_NAME=product_svc
```

### 2. Create a new database:

```sql
CREATE DATABASE product_svc;
```

### 3. Run Database Migrations

Execute the SQL commands in `db/migrations` to set up the database schema.

### 4. Run the Application

Launch the application using the following command:

```bash
go run main.go
```

### 5. Access the Server

The server will be accessible at [http://localhost:8080](http://localhost:8080).

## Contributing

We welcome contributions! Feel free to submit issues, feature requests, or pull requests.

## License

This project is licensed under the [MIT License](LICENSE).
