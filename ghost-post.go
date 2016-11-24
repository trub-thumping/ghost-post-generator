package main

import (
    "fmt"
    "time"

    "github.com/drhodes/golorem"
    "github.com/satori/go.uuid"
)

type GhostPost struct {
    UUID string
    Timestamp time.Time
    Title string
    Slug string
    Post string
}


func NewGhostPost() (gp GhostPost) {
    gp.UUID = fmt.Sprintf("%s", uuid.NewV4() )          // Fuck's sake
    gp.Title = lorem.Word(1,150)
    gp.Post = lorem.Paragraph(10, 100)
    gp.Slug = gp.UUID
    gp.Timestamp = MakeTimeStamp()

    return
}
