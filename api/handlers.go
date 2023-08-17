package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (s *Server) handleGetHome(c echo.Context) error {
    return c.Render(http.StatusOK, "index", echo.Map{})
}

func (s *Server) handleGetTest(c echo.Context) error {
    htmlStr := ""
    for i := 0; i < 3; i++ {
        htmlStr += "<p>Element " + strconv.Itoa(i) + "</p>"
    }
    return c.HTML(http.StatusOK, htmlStr)
}
