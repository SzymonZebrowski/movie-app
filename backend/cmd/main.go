package main

import (
	"os"
	"fmt"
	"time"
	"flag"

	"github.com/chenyahui/gin-cache"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/go-redis/redis/v8"
	"github.com/ilyakaznacheev/cleanenv"
	log "github.com/sirupsen/logrus"

	"movielist-app/pkg/movies"
	"movielist-app/pkg/diagnostics"
	"movielist-app/pkg/configuration"
)

var (
	config *configuration.Config
)

type Args struct {
	ConfigPath string
}

func init() {
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.JSONFormatter{})
}

// ProcessArgs processes and handles CLI arguments
func ProcessArgs(cfg interface{}) Args {
	var a Args

	f := flag.NewFlagSet("MovieList App", 1)
	f.StringVar(&a.ConfigPath, "c", "config.yml", "Path to configuration file")

	fu := f.Usage
	f.Usage = func() {
		fu()
		envHelp, _ := cleanenv.GetDescription(cfg, nil)
		fmt.Fprintln(f.Output())
		fmt.Fprintln(f.Output(), envHelp)
	}

	f.Parse(os.Args[1:])

	return a
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(cors.Default())

	cfg := configuration.GetConfig()

	args := ProcessArgs(&cfg)

	if err := cleanenv.ReadConfig(args.ConfigPath, cfg); err != nil {
		log.Fatal(err)
	}

	
	if err := cleanenv.ReadEnv(cfg); err != nil {
		log.Fatal(err)
	}

	if cfg.Cache.Enabled {
		redisStore := persist.NewRedisStore(redis.NewClient(&redis.Options{
			Network: "tcp",
			Addr:    fmt.Sprintf("%s:%d", cfg.Cache.Address, cfg.Cache.Port),
		}))

		router.GET("/movies", cache.CacheByRequestURI(redisStore, 10*time.Second), movies.ListMoviesHandler)
		router.GET("/movies/:id", cache.CacheByRequestURI(redisStore, 10*time.Second),  movies.GetMovieHandler)
		router.POST("/movies", movies.CreateMovieHandler)
	
	} else {
		router.GET("/movies", movies.ListMoviesHandler)
		router.GET("/movies/:id", movies.GetMovieHandler)
		router.POST("/movies", movies.CreateMovieHandler)
	}
	
	router.GET("/readyz", diagnostics.Readyz)
	router.GET("/healthz", diagnostics.Readyz)
	router.GET("/server-info", diagnostics.ServerInfo)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(":8080")
}
