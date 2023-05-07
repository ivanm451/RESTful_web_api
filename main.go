package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// album represents data about a record album
type album struct  {
    ID    string `json:"id"`
    Title string `json:"title"`
    Artist string `json:"artist"`
    Price float64 `json:"price"`
}

// albums slice to seed record album data
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// assigns the handler function to an endpoint path
// also sets up an association in which getAlbums handles requests to the /albums endpoint path
func main() {
    router := gin.Default() // initialize gin router with Default
    router.GET("/albums", getAlbums) // use GET function to associate GET HTTP method and /albums path (only name not function return)with handler function
    router.GET("/albums/:id", getAlbumByID)
    router.POST("/albums", postAlbums) // POST method is now associated with postAlbums function
    router.Run("localhost:8080") // using RUN to attach router to an http.Server and start server
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context){
    c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context){
    var newAlbum album

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    // Add the new ablum to the slice
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value  matches the id
// parameter sent by the client, then returns that album as a response
func getAlbumByID(c *gin.Context){
    id := c.Param("id")

    // Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
