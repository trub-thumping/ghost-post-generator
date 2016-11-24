ghost-post-generator
==

| who       | what |
|-----------|------|
| dockerhub | https://hub.docker.com/r/jspc/ghost-post-generator/   |
| circleci  | https://circleci.com/gh/jspc/ghost-post-generator   |
| licence   | MIT   |


Preseed a ghost blog (via the database) with some test posts. This makes development of themes and tooling simpler. At least for me.

Building
--

```bash
$ go get
$ CGO_ENABLED=0 go build -a -installsuffix cgo
$ docker build -t jspc/ghost-blog-generator .
```

Running
--

This lives in docker hub at the above link. Usage can be seen as per:

```
$ docker run jspc/ghost-post-generator -h
Usage of /ghost-post-generator:
  -d string
        Ghost's database name (default "ghostDB")
  -e    Empty the posts table prior to seeding
  -n int
        Number of posts to seed (default 25)
  -u string
        Username and password for mySQL (default "ghost:ghost")
```

This tool massively relies on the fact that the user:

  * Is running a containerised mysql; and
  * Links to the mysql container using the link name 'mysql'

Thus:

```
$ docker run --link some-mysql:mysql jspc/ghost-post-generator -d ghost -u ghost:ghostpassword
2016/11/24 14:06:58 Connecting to "ghost:ghostpassword@tcp(172.17.0.5:3306)/ghost"
```

Licence
--

MIT License

Copyright (c) 2016 jspc

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
