package q01projectinit

import "homework04/internal/blog"

// BuildApp implements Q1: initialize module dependencies and bootstrap app.
func BuildApp(dsn string) (*blog.App, error) {
	return blog.NewSQLiteApp(blog.Config{
		DSN:       dsn,
		JWTSecret: "homework04-q1-secret",
	})
}
