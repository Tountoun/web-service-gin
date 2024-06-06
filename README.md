# SIMPLE REST API BASED ON ALBUM DATA

## Description
This project as the title said is a simple rest api built with gin a go based framework. The api serves album data which are stored in memory.

## Functionnalities
- Get all ablums stored in memory
- Save a new album
- Get one album by its ID
- Delete one album by its ID
- Save a list of albums


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

- Save a list of albums
```shell
$ curl http://localhost:8080/albums/list \
    --request "POST" \
    --header "Content-Type:application/json" \
    --data '[{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99},{"id": "5","title": "The secrets of Holly","artist": "Mikey Holly","price": 32.2}, {"id": "6","title": "Hart of Me","artist": "Sam Lion","price": 50}]'
```