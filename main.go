package main

import (
    "flag"
    "fmt"
    "log"
    "math/rand"
    "os"
    "time"
)

var (
    connectionString string
    count *int
    databaseName *string
    empty *bool
    userString *string

    err error
)

func init() {
    userString = flag.String("u", "ghost:ghost", "Username and password for mySQL")
    databaseName = flag.String("d", "ghostDB", "Ghost's database name")
    count = flag.Int("n", 25, "Number of posts to seed")
    empty = flag.Bool("e", false,  "Empty the posts table prior to seeding")
}

func main() {
    rand.Seed( time.Now().UTC().UnixNano())
    flag.Parse()

    connectionString = inferConnectionString()
    log.Printf("Connecting to %q", connectionString)

    db := NewDB(connectionString, *empty)
    defer db.Database.Close()

    for i := 0; i < *count; i++ {
        if err = db.AddPost( NewGhostPost() ); err != nil {
            log.Fatal(err)
        }
    }
}

func inferConnectionString()string {
    return fmt.Sprintf("%s@tcp(%s:%s)/%s",
        *userString,
        os.Getenv("MYSQL_PORT_3306_TCP_ADDR"),
        os.Getenv("MYSQL_PORT_3306_TCP_PORT"),
        *databaseName,
    )

}
