package api

import (
	"html/template"
	"log"

	"github.com/jbic9832/go-calorie-tracker/storage"
	"github.com/labstack/echo/v4"
)

type Server struct {
	listenAddr string
	echo       *echo.Echo
	store      *storage.Storage
}

func NewServer(listenAddr string) *Server {
	e := echo.New()
	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = t

    store, err := storage.NewPostgresStore()
    if err != nil {
        log.Fatal(err)
    }
    
	return &Server{
		listenAddr: listenAddr,
		echo:       e,
        store: store,
	}
}

func (s *Server) Start() error {
	s.echo.GET("/", s.handleGetHome)
	s.echo.GET("/hello", s.handleGetTest)
	return s.echo.Start(s.listenAddr)
}
