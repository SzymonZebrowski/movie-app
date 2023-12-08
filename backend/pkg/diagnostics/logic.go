package diagnostics

import (
	"os"
	"time"
	"fmt"
	"net/http"
	"net"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"


	"movielist-app/pkg/configuration"

)

func Readyz(c *gin.Context) {
	cfg := configuration.GetConfig()

	dbconn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", cfg.Database.Address, cfg.Database.Port))
	if dbconn != nil {
		defer dbconn.Close()
	}


	if err != nil {
		log.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}	

	if cfg.Cache.Enabled {
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", cfg.Cache.Address, cfg.Cache.Port))
		if conn != nil {
			defer conn.Close()
		}

		if err != nil {
			log.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

type info struct {
	Hostname string `yaml:"hostname"`
	Time time.Time `yaml:"time"`
}


func ServerInfo(c *gin.Context) {
	hostname, err := os.Hostname()
	if err != nil {
		log.Error(err)
	}
	c.JSON(http.StatusOK, &info{
		Hostname: hostname,
		Time: time.Now(),
	})
}
