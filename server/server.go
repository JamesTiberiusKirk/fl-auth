package server

import (
	"fl-auth/db"
	"fl-auth/server/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// ContextParams stores context parameters for server initialization
type ContextParams struct {
	DB *db.Client
}

// Init to initialize an Echo route
func Init(dbc *db.Client) *echo.Echo {
	e := echo.New()
	contextParams := ContextParams{DB: dbc}
	e.Use(
		createContext(contextParams),
		// middleware.Logger(),
		middleware.Gzip(),
		middleware.CORS(),
	)

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[FL-AUTH]: [HTTP]: ${method} ${status} ${uri} ${user_agent}\n",
	}))

	e = initRoutes(e)
	return e
}

func initRoutes(e *echo.Echo) *echo.Echo {

	e.POST("/verify_me", controllers.VerifyMe)
	e.POST("/register", controllers.Register)
	e.POST("/login", controllers.Login)

	return e
}

/* ContextObjects attaches backend clients to the API context. */
func createContext(contextParams ContextParams) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", contextParams.DB)
			return next(c)
		}
	}
}
