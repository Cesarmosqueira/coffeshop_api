package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"

	r "github.com/Cesarmosqueira/coffeshop_api/pkg/domain/response"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	port   string
}

func NewServer() Server {
	gin.SetMode(gin.ReleaseMode)

	var server Server
	server.router = gin.New()
	server.router.Use(gin.Logger())
	server.router.Use(CorsConfig())

	server.router.NoRoute(func(c *gin.Context) {
		status := http.StatusNotFound
		c.JSON(r.NewResponse(status, "Endpoint not found"))
	})

	server.port = os.Getenv("PORT")
	if server.port == "" {
		server.port = "8000"
	}
	server.port = ":" + server.port

	return server
}

func (s *Server) EnableApi(fn interface{}) {
	function := reflect.ValueOf(fn)
	router := []reflect.Value{reflect.ValueOf(s.router)}
	apiValue := function.Call(router)[0]

	for i := 0; i < apiValue.NumMethod(); i++ {
		method := apiValue.Method(i)
		method.Call([]reflect.Value{})
		log.Printf("[%s] %s enabled.", apiValue.Type(), apiValue.Type().Method(i).Name)
	}
}

func (s *Server) Run() {
	log.Println("[Server] Running on port: " + s.port[1:])
	s.router.Run(s.port)
	fmt.Println("Bye")
}
