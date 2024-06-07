package main

import (
	"log"
	"net/http"
	"os"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id" binding:"required"`
	Title  string  `json:"title" binding:"required"`
	Artist string  `json:"artist" bindiing:"required"`
	Price  float64 `json:"price" binding:"required"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	file, err := os.OpenFile("logging.txt", os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	log.SetOutput(file)
	
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbum)
	router.GET("/albums/:id", getAlbumByID)
	router.DELETE("/albums/:id", deleteAlbumByID)
	router.POST("/albums/list", postAlbums)
	router.GET("/albums/search", searchAlbum)
	router.PUT("/albums/:id", updateAlbum)

	router.Run("localhost:8080")
}

// get all albums
func getAlbums(c *gin.Context) {
	log.Println("Getting all albums")
	c.IndentedJSON(http.StatusOK, albums)
}

// add new album
func postAlbum(c *gin.Context) {
	var newAlbum album

	// bind the request body data to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		log.Println("Bad request when intending to post a new album")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	log.Println("Posting album", newAlbum)
	// add the new album to the slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// get album by id
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, album := range albums {
		if album.ID == id {
			log.Println("Getting album with id", id)
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	log.Printf("Album with id %v not found\n", id)
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album with id " + id + " not found"})
}

// delete one album by id
func deleteAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for idx, album := range albums {
		if album.ID == id {
			log.Println("Deleting album with id", id)
			albums = append(albums[:idx], albums[idx+1:]... )
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Album with id " + id + " deleted successfully"})
			return
		}
	}
	log.Printf("Album with id %v not found\n", id)
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album with id " + id + " not found"})
}

// save many albums
func postAlbums(c *gin.Context) {
	var newAlbums []album 
	if err := c.BindJSON(&newAlbums); err != nil {
		log.Println("Bad request when intending to post a list of albums")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	log.Println("Posting albums", albums)
	albums = append(albums, newAlbums...)
	c.IndentedJSON(http.StatusOK, albums)
}

// search for an album using a keyword
func searchAlbum(c *gin.Context) {
	key := c.Query("key")
	result := make([]album, 0)

	log.Println("Searching an album with keyword", key)
	for _, album := range albums {
		if strings.Contains(album.Title, key) || strings.Contains(album.Artist, key) {
			result = append(result, album)
		}
	}
	c.IndentedJSON(http.StatusOK, result)
}

// update an album
func updateAlbum(c *gin.Context) {
	id := c.Param("id")
	var updateData album

	if err := c.BindJSON(&updateData); err != nil {
		log.Println("Bad request when intending to update album with id", id)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if id != updateData.ID {
		log.Printf("Can not update album with id %v. Path parameter must be equal to the updated data id\n", id)
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Can not update the id of the album"})
		return
	}
	index := slices.IndexFunc(albums, func(al album) bool {
		return al.ID == id
	})
	albums := slices.Replace(albums, index, index + 1, updateData)
	log.Println("Updating album with id", id)
	c.IndentedJSON(http.StatusOK, albums[index])
}