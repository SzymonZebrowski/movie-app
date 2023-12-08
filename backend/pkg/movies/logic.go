package movies

import (
	"net/http"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"movielist-app/pkg/models"
)

func ListMoviesHandler(c* gin.Context) {
	mov, err := models.ListMovies()

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
	}

	if mov == nil {
	 c.AbortWithStatus(http.StatusNotFound)
	} else {
	 c.JSON(http.StatusOK, mov)
	}
}

func CreateMovieHandler(c* gin.Context) {
	var mov models.Movie

	if err := c.BindJSON(&mov); err != nil {
	 c.AbortWithStatus(http.StatusBadRequest)
	}

	if err := models.CreateMovie(mov); err != nil {
		log.Error(err.Error())

		c.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
	}

	c.JSON(http.StatusCreated, mov)
}

func GetMovieHandler(c* gin.Context) {
	id := c.Param("id")

	mov, err := models.GetMovie(id)
	if err != nil {
		log.Error(err.Error())

		c.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
	}

	if mov == nil {
	 c.AbortWithStatus(http.StatusNotFound)
	} else {
	 c.JSON(http.StatusOK, mov)
	}
}
