# ShorlServer
A tiny server for creating short links


## To run this server
`git clone` this repo to local, then `go run .` in the root directory.

##Currently supported routes (APIs):
1. `GET` `/generate`?`long_link`=`{your link}`
2. `POST` `/generate`, with a request body of type `application/json`:
```json
[
 "https://pkg.go.dev/cloud.google.com/go/bigtable",
 "https://reqbin.com/",
 "https://www.youtube.com/",
 "https://www.bilibili.com/",
 "https://music.youtube.com/library/playlists"
]
```

reponse example:
```json
[{
    "short_link": "sho.rl/r/P7qOf",
    "long_link": "https://pkg.go.dev/cloud.google.com/go/bigtable"
}, {
    "short_link": "sho.rl/r/2RY91W",
    "long_link": "https://reqbin.com/"
}, {
    "short_link": "sho.rl/r/moaTq",
    "long_link": "https://www.youtube.com/"
}, {
    "short_link": "sho.rl/r/17xLbG",
    "long_link": "https://www.bilibili.com/"
}, {
    "short_link": "sho.rl/r/46gfj",
    "long_link": "https://music.youtube.com/library/playlists"
}]
```

3. `GET` `/r/{short_link}`, redirect to the original page given the short link.
