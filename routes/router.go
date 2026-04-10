package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/KicaiBadai/gin-firebase-backend/handlers"
	"github.com/KicaiBadai/gin-firebase-backend/middleware"
)

func SetupRouter() *gin.Engine {
	// gin.Default() sudah include Logger & Recovery middleware
	r := gin.Default()

	// ─── CORS Middleware ───────────────────────────────────
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// ─── Init handlers ─────────────────────────────────────
	authHandler := handlers.NewAuthHandler()
	productHandler := handlers.NewProductHandler()

	// ─── API v1 group ──────────────────────────────────────
	v1 := r.Group("/v1")
	{
		// Health check (public)
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status":  "ok",
				"service": "gin-firebase-backend",
			})
		})

		// ── Auth routes (public) ───────────────────────────
		auth := v1.Group("/auth")
		{
			// POST /v1/auth/verify-token
			auth.POST("/verify-token", authHandler.VerifyToken)
		}

		// ── Protected routes (butuh JWT) ──────────────────
		protected := v1.Group("")
		protected.Use(middleware.AuthMiddleware())
		{
			// ── Products ────────────────────────────────
			products := protected.Group("/products")
			{
				// GET /v1/products
				products.GET("", productHandler.GetAll)

				// GET /v1/products/:id
				products.GET("/:id", productHandler.GetByID)

				// ── Admin only ─────────────────────────
				adminProducts := products.Group("")
				adminProducts.Use(middleware.AdminOnly())
				{
					// POST /v1/products
					adminProducts.POST("", productHandler.Create)

					// PUT /v1/products/:id
					adminProducts.PUT("/:id", productHandler.Update)

					// DELETE /v1/products/:id
					adminProducts.DELETE("/:id", productHandler.Delete)
				}
			}
		}
	}

	return r
}