package pgx_demo

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
)

func TestPgx(t *testing.T) {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	//#   - Connection parameters:
	//#     (webui)    http://127.0.0.1:18080/demologin?password=demo78698&username=demo
	//#     (sql)      postgresql://demo:demo78698@127.0.0.1:25432/movr?sslmode=require&sslrootcert=%2Ftmp%2Fdemo1079327321%2Fca.crt
	//#     (sql/jdbc) jdbc:postgresql://127.0.0.1:25432/movr?password=demo78698&sslmode=require&sslrootcert=%2Ftmp%2Fdemo1079327321%2Fca.crt&user=demo
	//#     (sql/unix) postgresql://demo:demo78698@/movr?host=%2Ftmp%2Fdemo1079327321&port=25432
	os.Setenv("DATABASE_URL", "postgres://demo:demo78698@192.168.100.117:25432/movr")
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var name string
	var weight int64
	err = conn.QueryRow(context.Background(), "select name, weight from widgets where id=$1", 42).Scan(&name, &weight)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(name, weight)
}

func TestAA(t *testing.T) {
	os.Setenv("DATABASE_URL", "aa")
}
