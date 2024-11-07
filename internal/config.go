package internal

import (
	"github.com/danargh/apis-perizinan-app/pkg/env"
	"github.com/joho/godotenv"
)

type Config struct {
	BaseURL         string
	HTTPPort        int
	CookieSecretKey string
	DBDsn           string
	Automigrate     bool
	JWTSecretKey    string
}

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		BaseURL:         env.GetString("BASE_URL", "http://localhost:8000"),
		HTTPPort:        env.GetInt("HTTP_PORT", 8000),
		CookieSecretKey: env.GetString("COOKIE_SECRET_KEY", "osxcxvka3ke4335avd6zpaf2lnkdwk5t"),
		DBDsn:           env.GetString("DB_DSN", "user:pass@localhost:port/db"),
		Automigrate:     env.GetBool("DB_AUTOMIGRATE", true),
		JWTSecretKey:    env.GetString("JWT_SECRET_KEY", "sfxfvu37244ljpzpmfmqvvgcft6d33kb"),
	}
	// cfg.baseURL = env.GetString("BASE_URL", "http://localhost:8000")
	// cfg.httpPort = env.GetInt("HTTP_PORT", 8000)
	// cfg.basicAuth.username = env.GetString("BASIC_AUTH_USERNAME", "admin")
	// cfg.basicAuth.hashedPassword = env.GetString("BASIC_AUTH_HASHED_PASSWORD", "$2a$10$jRb2qniNcoCyQM23T59RfeEQUbgdAXfR6S0scynmKfJa5Gj3arGJa")
	// cfg.cookieSecretKey = env.GetString("COOKIE_SECRET_KEY", "osxcxvka3ke4335avd6zpaf2lnkdwk5t")
	// cfg.dbDsn = env.GetString("DB_DSN", "user:pass@localhost:port/db")
	// cfg.automigrate = env.GetBool("DB_AUTOMIGRATE", true)
	// cfg.jwtSecretKey = env.GetString("JWT_SECRET_KEY", "sfxfvu37244ljpzpmfmqvvgcft6d33kb")

	return cfg, err
}
