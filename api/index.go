package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var app *gin.Engine

type Song struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Album  string `json:"album"`
	Year   string `json:"year"`
	Label  string `json:"label"`
}

var songs = map[string][]Song{
	"songs": {
		{
			"1",
			"Give It Away",
			"Red Hot Chili Peppers",
			"Blood Sugar Sex Magik",
			"1991",
			"Warner Bros Records",
		},
		{
			"2",
			"Song 2",
			"Blur",
			"Blur",
			"1997",
			"Food Records",
		},
	},
}

func myRoute(r *gin.RouterGroup) {
	r.GET("/songs", func(c *gin.Context) {
		c.JSON(http.StatusOK, songs)
	})
	r.GET("/song/:id", func(c *gin.Context) {
		id := c.Param("id")

		for _, s := range songs["songs"] {
			if s.Id == id {
				c.JSON(http.StatusOK, gin.H{
					"song": s,
				})
				return
			}
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("no song with id %s", id),
		})
	})
}

func init() {
	app = gin.New()
	r := app.Group("/api")

	myRoute(r)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	w.Header().Set("Content-Type", "application/json")
	app.ServeHTTP(w, r)
}
