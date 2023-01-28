package cockroach_demo

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/lib/pq"
)

func TestConnect(t *testing.T) {
	db, err := sql.Open("postgres",
		"postgresql://demo:demo1288@127.0.0.1:15432/movr?sslmode=require&sslrootcert=%2Fvar%2Ffolders%2Fpv%2Fgwhnhg8n2437nhmynvb_x_th0000gn%2FT%2Fdemo3486628319%2Fca.crt")
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}
	defer db.Close()

	if _, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS accounts (id INT PRIMARY KEY, balance INT)"); err != nil {
		log.Fatal(err)
	}

	//if _, err := db.Exec(
	//	"INSERT INTO accounts (id, balance) VALUES (1, 1000), (2, 250)"); err != nil {
	//	log.Fatal(err)
	//}

	rows, err := db.Query("SELECT id, balance FROM accounts")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	fmt.Println("Initial balances:")
	for rows.Next() {
		var id, balance int
		if err := rows.Scan(&id, &balance); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d %d\n", id, balance)
	}

}
