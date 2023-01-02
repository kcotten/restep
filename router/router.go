package router

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	hr "github.com/julienschmidt/httprouter"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/gomodule/redigo/redis"
)

// Router struct
type Router struct {
	Router *hr.Router
	Port   string
	logger log.Logger

	pool              *redis.Pool
	intervalSecond    int
	requestCountLimit int64
	redisKey          string
}

// Initialize the router to return sample JSON
func (router *Router) Init() {
	// Router
	var ok bool
	router.Router = hr.New()
	router.Router.HandlerFunc("GET", "/info", router.HandleRest)
	if router.Port, ok = os.LookupEnv("ROUTER_PORT"); !ok {
		router.Port = "8000"
	}
	router.logger = *log.Default()
	router.logger.Println("Application initialized with port: ", router.Port)

	// Prometheus
	router.Router.Handler("GET", "/metrics", promhttp.Handler())
	promInit()

	// Redis
	router.intervalSecond = 60
	router.requestCountLimit = int64(3)
	router.redisKey = "api_request_limit"
	redisServer := flag.String("redisServer", ":6379", "")
	flag.Parse()
	router.pool = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", *redisServer) },
	}

	// Example
	ExampleRet = example()
}

// Run the router at the provided address
func (router *Router) Run(addr string) {}

// Router handler for incoming rest requests
func (router *Router) HandleRest(w http.ResponseWriter, r *http.Request) {
	var status string
	timer := getTimer(status)
	conn := router.pool.Get()

	defer func() {
		getRequestCounter.WithLabelValues(status).Inc()
		timer.ObserveDuration()
		conn.Close()
	}()

	// simple rate limiting
	count, err := redis.Int64(conn.Do("INCR", router.redisKey))
	if err != nil {
		panic(err)
	}

	if count == 1 {
		_, err := redis.Int(conn.Do("EXPIRE", router.redisKey, router.intervalSecond))
		if err != nil {
			panic(err)
		}
	}

	if count > router.requestCountLimit {
		router.logger.Println("limit exceeded")
		http.Error(w, "limited exceeded", http.StatusTooManyRequests)
		return
	}

	// handle request
	router.logger.Println("Handling request: ", r)
	err = json.NewEncoder(w).Encode(ExampleRet)
	if err != nil {
		panic(err)
	}
	status = "success"
}

// Example go<=>json for endpoint to return
type Info struct {
	Item     string `json:"item"`
	Quantity string `json:"quantity"`
}

// Example return value
var ExampleRet []Info

// Fill in the example struct with some arbitrary values
func example() []Info {
	Info := []Info{
		{Item: "item1", Quantity: "3"},
		{Item: "item2", Quantity: "7"},
	}
	return Info
}
