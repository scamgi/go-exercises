package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // The PostgreSQL driver, blank import is used because it's only needed for its side effects (registering the driver).
)

// User represents the structure of our data in the database.
type User struct {
	ID   int
	Name string
	Age  int
}

func main() {
	// 1. Define the connection string.
	// This string contains the necessary information to connect to the PostgreSQL database
	// that was set up in the docker-compose.yml file.
	connStr := "user=myuser password=mypassword dbname=mydatabase host=localhost port=5432 sslmode=disable"

	// 2. Open a database connection pool.
	// sql.Open doesn't establish any connections to the database, nor does it validate driver parameters.
	// It simply prepares the database abstraction for later use.
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	// Defering db.Close() ensures that the connection pool is closed when the main function exits.
	defer db.Close()

	// 3. Verify the connection is alive.
	// db.Ping() is used to confirm that a connection to the database can be established.
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v. Is the Docker container running?", err)
	}
	fmt.Println("Successfully connected to the PostgreSQL database!")

	// 4. Create a table to store user data if it doesn't already exist.
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		age INT
	);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}
	fmt.Println("Table 'users' is ready.")

	// 5. Insert new data into the table.
	// Using placeholders ($1, $2) for values is a best practice to prevent SQL injection attacks.
	fmt.Println("Inserting new users...")
	usersToInsert := []User{
		{Name: "Charles Babbage", Age: 79},
		{Name: "Grace Hopper", Age: 85},
	}

	for _, user := range usersToInsert {
		// db.Exec is used for SQL statements that don't return rows (e.g., INSERT, UPDATE, DELETE).
		_, err := db.Exec("INSERT INTO users (name, age) VALUES ($1, $2)", user.Name, user.Age)
		if err != nil {
			log.Printf("Could not insert user %s: %v. It might already exist.", user.Name, err)
		}
	}
	fmt.Println("Users inserted.")

	// 6. Query the data from the table.
	fmt.Println("\n--- Querying all users from the database ---")
	// db.Query is used for SELECT statements that return multiple rows.
	rows, err := db.Query("SELECT id, name, age FROM users ORDER BY name")
	if err != nil {
		log.Fatalf("Error querying users: %v", err)
	}
	// It's crucial to close the rows when finished to release the database connection.
	defer rows.Close()

	var users []User
	// Iterate through the result set.
	for rows.Next() {
		var u User
		// rows.Scan copies the column values from the current row into the provided variables.
		if err := rows.Scan(&u.ID, &u.Name, &u.Age); err != nil {
			log.Fatalf("Error scanning row: %v", err)
		}
		users = append(users, u)
	}

	// Check for any errors that occurred during row iteration.
	if err = rows.Err(); err != nil {
		log.Fatalf("Error during row iteration: %v", err)
	}

	// 7. Print the retrieved results.
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", user.ID, user.Name, user.Age)
	}
}
