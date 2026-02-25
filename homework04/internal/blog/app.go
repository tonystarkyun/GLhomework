package blog

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const userIDContextKey = "userID"

// App contains runtime dependencies.
type App struct {
	DB        *gorm.DB
	Router    *gin.Engine
	JWTSecret []byte
	Logger    *log.Logger
}

// Claims are JWT claims for authenticated users.
type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// Config controls app construction.
type Config struct {
	DSN       string
	JWTSecret string
	Logger    *log.Logger
}

// NewSQLiteApp builds the full blog app backed by sqlite.
func NewSQLiteApp(cfg Config) (*App, error) {
	if cfg.DSN == "" {
		return nil, fmt.Errorf("dsn is required")
	}
	if cfg.JWTSecret == "" {
		return nil, fmt.Errorf("jwt secret is required")
	}

	db, err := gorm.Open(sqlite.Open(cfg.DSN), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}
	if err := db.AutoMigrate(&User{}, &Post{}, &Comment{}); err != nil {
		return nil, fmt.Errorf("auto migrate: %w", err)
	}

	logger := cfg.Logger
	if logger == nil {
		logger = log.New(io.Discard, "", log.LstdFlags)
	}

	gin.SetMode(gin.ReleaseMode)
	app := &App{
		DB:        db,
		JWTSecret: []byte(cfg.JWTSecret),
		Logger:    logger,
	}
	app.Router = app.newRouter()
	return app, nil
}

func (a *App) newRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(a.requestLogMiddleware())

	r.POST("/register", a.handleRegister)
	r.POST("/login", a.handleLogin)

	r.GET("/posts", a.handleListPosts)
	r.GET("/posts/:id", a.handleGetPost)
	r.GET("/posts/:id/comments", a.handleListComments)

	auth := r.Group("/")
	auth.Use(a.authMiddleware())
	auth.POST("/posts", a.handleCreatePost)
	auth.PUT("/posts/:id", a.handleUpdatePost)
	auth.DELETE("/posts/:id", a.handleDeletePost)
	auth.POST("/posts/:id/comments", a.handleCreateComment)

	return r
}

func (a *App) requestLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		a.Logger.Printf("request method=%s path=%s status=%d latency=%s", c.Request.Method, c.Request.URL.Path, c.Writer.Status(), time.Since(start).String())
	}
}

func (a *App) authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if !strings.HasPrefix(header, "Bearer ") {
			a.abortError(c, http.StatusUnauthorized, "missing or invalid authorization header")
			return
		}

		tokenString := strings.TrimPrefix(header, "Bearer ")
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return a.JWTSecret, nil
		})
		if err != nil || !token.Valid {
			a.abortError(c, http.StatusUnauthorized, "invalid or expired token")
			return
		}

		c.Set(userIDContextKey, claims.UserID)
		c.Next()
	}
}

func userIDFromContext(c *gin.Context) uint {
	v, ok := c.Get(userIDContextKey)
	if !ok {
		return 0
	}
	id, _ := v.(uint)
	return id
}

func parseIDParam(c *gin.Context, key string) (uint, bool) {
	raw := c.Param(key)
	n, err := strconv.ParseUint(raw, 10, 64)
	if err != nil || n == 0 {
		return 0, false
	}
	return uint(n), true
}

func (a *App) abortError(c *gin.Context, status int, message string) {
	a.Logger.Printf("error status=%d method=%s path=%s msg=%s", status, c.Request.Method, c.Request.URL.Path, message)
	c.AbortWithStatusJSON(status, gin.H{"error": message})
}

func hashPassword(password string) (string, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(h), nil
}

func comparePassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func (a *App) newToken(u User) (string, error) {
	now := time.Now()
	claims := &Claims{
		UserID:   u.ID,
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(a.JWTSecret)
}
