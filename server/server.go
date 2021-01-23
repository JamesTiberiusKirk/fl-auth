package server

import (
	"fl-auth/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Use(
		middleware.Logger(),
		middleware.Gzip(),
		middleware.CORS(),
	)

	e = initRoutes(e)
	return e
}

func initRoutes(e *echo.Echo) *echo.Echo {

	// g := e.Group("/admin")

	// g.Use(middleware.BasicAuth(flMiddleware.Auth))

	// g.POST("/users", controllers.AddUser)
	// e.GET("/users", controllers.GetUsers)
	// e.POST("/register", controllers.Register)
	// e.POST("/login", controllers.Login)

	e.GET("/", controllers.HelloWorld)

	return e
}

// ContextObjects attaches backend clients to the API context
// func createContext(contextParams ContextParams) echo.MiddlewareFunc {
// 	return func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return func(c echo.Context) error {
// 			c.Set("db", contextParams.DB)
// 			return next(c)
// 		}
// 	}
// }
