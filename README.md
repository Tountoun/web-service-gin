# SIMPLE REST API BASED ON ALBUM DATA

## Description
This project as the title said is a simple rest api built with gin a go based framework. The api serves album data which are stored in memory.

## Functionnalities
- Get all ablums stored in memory
- Save a new album
- Get one album by its ID
- Delete one album by its ID


## Testing
- Get all albums
```shell
$ curl http://localhost:8080/albums \
    --header \
    "Content-Type: application/json" \
    --request "GET"
```

- Post a new album
```shell
$ curl http://localhost:8080/albums \
    --include --header \
    "Content-Type: application/json" \
    --request "POST" --data \
    '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
``` 

- Get album with id 2
```shell
$ curl http://localhost:8080/albums/2
```

- Delete album with id 2
```shell
$ curl http://localhost:8080/albums/2 --request "DELETE"
```
