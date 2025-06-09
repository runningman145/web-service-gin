package main

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
)

// declare an album struct, used to store an album in memory
// album represents data about a record album
type album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	// gin.context carries request details, validates and serializes JSON etc
	c.IndentedJSON(http.StatusOK, albums)
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	// initialize a Gin router
	router := gin.Default()
	router.GET("/albums", getAlbums) // associating GET with albums path
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.DELETE("/albums/:id", deleteAlbumById)

	router.Run("localhost:8080")
}

// postAlbums adds an album from JSON received in the request body
func postAlbums(c *gin.Context) {
	var newAlbum album

	// call bindJSON to bind the received JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumById locates the album whose value ID matches the id
// the id param sent by the client, then return the album as a response
func getAlbumByID(c *gin.Context) {
	// using c.Param to retrieve the id path parameter from the url
	id := c.Param("id")

	// loop over the list of albums looking for the album whose id values matches
	for _, a := range albums {
		if a.ID == id {
			// if found, serialize to JSON and return a response with 200OK
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found!"})
}

// deleteAlbum by id, used to delete an album whose id value matches the id
func deleteAlbumById(c *gin.Context) {
	// retrieve the id path parameter from the url
	id := c.Param("id")

	// loop over the list of albums in order to get the particular album id
	for _, a := range albums {
		if a.ID == id {
			// if the album is found, serialize and return a 200 OK for album deleted successfully
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No item to be deleted"})
}
