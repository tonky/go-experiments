package main

import "log"
import "fmt"
import "math/rand"
import "time"
import "github.com/gocql/gocql"
import "net/http"

var session *gocql.Session
var rnd *rand.Rand

func main() {
	cluster := gocql.NewCluster("192.168.99.100")
	cluster.Keyspace = "test"
	session, _ = cluster.CreateSession()
	defer session.Close()

	rnd = rand.New(rand.NewSource(99))

	fmt.Println("Starting server on port 8080")

	http.HandleFunc("/cass", cass)
	http.HandleFunc("/hello", hello)
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func cass(w http.ResponseWriter, r *http.Request) {
	user := UserQ(session, "SELECT * from users")

	// sleep randomly for a second on 1% of connections
	if rnd.Intn(100) == 1 {
		fmt.Println("Got a sleepy head here!")
		time.Sleep(1 * time.Second)
	}

	fmt.Fprintf(w, "Hi there, user %s", user["name"])
}

// UserQ ...
func UserQ(session *gocql.Session, query string) map[string]interface{} {
	user := map[string]interface{}{"id": 0, "name": ""}

	session.Query(query).MapScan(user)

	return user
}
