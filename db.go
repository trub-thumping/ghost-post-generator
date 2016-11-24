package main

import (
    "database/sql"
    "fmt"
    "time"

    _"github.com/go-sql-driver/mysql"   // We use the generic sql library to make adding sqlite support easier later
)

const (
    postsTable = "posts"
)

type DB struct {
    Database *sql.DB    // Pointer to database... cursor? Object? whatever
    empty bool          // when true, truncate the `posts' table first
}

func NewDB(cs string, empty bool)(db DB) {
    db.empty = empty

    if db.Database, err = sql.Open("mysql", cs); err != nil {
        panic(err.Error())
    }

    if err = db.Database.Ping(); err != nil {
        panic(err.Error())
    }

    if db.empty {
        db.truncate()
    }

    return
}

func (db DB)AddPost(gh GhostPost) error {
    return db.insert( gh.UUID, gh.Title, gh.Slug, gh.Post, gh.Timestamp)
}

// mysql> describe posts;
// +------------------+------------------+------+-----+---------+----------------+
// | Field            | Type             | Null | Key | Default | Extra          |
// +------------------+------------------+------+-----+---------+----------------+
// | id               | int(10) unsigned | NO   | PRI | NULL    | auto_increment |
// | uuid             | varchar(36)      | NO   |     | NULL    |                |
// | title            | varchar(150)     | NO   |     | NULL    |                |
// | slug             | varchar(150)     | NO   | UNI | NULL    |                |
// | markdown         | mediumtext       | YES  |     | NULL    |                |
// | mobiledoc        | longtext         | YES  |     | NULL    |                |
// | html             | mediumtext       | YES  |     | NULL    |                |
// | amp              | mediumtext       | YES  |     | NULL    |                |
// | image            | text             | YES  |     | NULL    |                |
// | featured         | tinyint(1)       | NO   |     | 0       |                |
// | page             | tinyint(1)       | NO   |     | 0       |                |
// | status           | varchar(150)     | NO   |     | draft   |                |
// | language         | varchar(6)       | NO   |     | en_US   |                |
// | visibility       | varchar(150)     | NO   |     | public  |                |
// | meta_title       | varchar(150)     | YES  |     | NULL    |                |
// | meta_description | varchar(200)     | YES  |     | NULL    |                |
// | author_id        | int(11)          | NO   |     | NULL    |                |
// | created_at       | datetime         | NO   |     | NULL    |                |
// | created_by       | int(11)          | NO   |     | NULL    |                |
// | updated_at       | datetime         | YES  |     | NULL    |                |
// | updated_by       | int(11)          | YES  |     | NULL    |                |
// | published_at     | datetime         | YES  |     | NULL    |                |
// | published_by     | int(11)          | YES  |     | NULL    |                |
// +------------------+------------------+------+-----+---------+----------------+

func (db DB)insert(uuid, title, slug, markdown string, timestamp time.Time ) (e error) {
    _,e = db.Database.Exec(`INSERT INTO posts (uuid, title, slug, markdown, html, author_id, created_at, created_by, published_at, published_by, status)
values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
        uuid, title, slug, markdown, fmt.Sprintf("<p>%s</p>", markdown), 1, timestamp, 1, time.Now(), 1, "published",
    )

    return
}

func (db DB)truncate() {
    if _,e := db.Database.Exec("TRUNCATE TABLE posts"); e != nil {
        panic(e.Error())
    }
}
