package main

import (
	"flag"
	"os"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rs/zerolog/log"
)

func main() {
	migrationsDir := flag.String("dir", "migrations", "migrations directory")
	databaseURL := flag.String("database", "", "database url")
	action := flag.String("action", "up", "action: up|down|version|force")
	steps := flag.Int("steps", 0, "number of steps for down")
	force := flag.Int("force", -1, "force version (used with action=force)")
	flag.Parse()

	if *databaseURL == "" {
		*databaseURL = os.Getenv("MIGRATE_DATABASE_URL")
	}
	if *databaseURL == "" {
		log.Fatal().Msg("MIGRATE_DATABASE_URL is not set")
	}

	m, err := migrate.New("file://"+*migrationsDir, *databaseURL)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create migrator")
	}
	defer m.Close()

	switch *action {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal().Err(err).Msg("migration up failed")
		}
	case "down":
		if *steps > 0 {
			if err := m.Steps(-*steps); err != nil && err != migrate.ErrNoChange {
				log.Fatal().Err(err).Msg("migration down failed")
			}
			break
		}
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal().Err(err).Msg("migration down failed")
		}
	case "version":
		v, dirty, err := m.Version()
		if err != nil && err != migrate.ErrNilVersion {
			log.Fatal().Err(err).Msg("migration version failed")
		}
		log.Info().Str("version", strconv.FormatUint(uint64(v), 10)).Bool("dirty", dirty).Msg("migration version")
	case "force":
		if *force < 0 {
			log.Fatal().Msg("force requires -force version")
		}
		if err := m.Force(*force); err != nil {
			log.Fatal().Err(err).Msg("migration force failed")
		}
	default:
		log.Fatal().Str("action", *action).Msg("unsupported action")
	}
}
