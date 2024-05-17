package db

import (
	"database/sql"
	"fmt"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type postgresDb struct {
	name    		string
	uri    			string
	isSslDisabled 	bool
	db  			*sql.DB
	logger 			zerolog.Logger
}

func NewPostgresDb(name, uri string, isSslDisabled bool) *postgresDb {
	return &postgresDb{
		name: name,
		uri: uri,
		isSslDisabled: isSslDisabled,
		logger: log.With().Str("actor", "db/postgres").Logger(),
	}
}

func (p *postgresDb) Connect() {
	log.Info().Msg("Connecting to database")

	if p.name == "" {
		p.logger.
			Fatal().
			Str("status", "missing parameters").
			Str("reason", "name is required").
			Send()
	}
	if p.uri == "" {
		p.logger.
			Fatal().
			Str("dbName", p.name).
			Str("status", "missing parameters").
			Str("reason", "uri is required").
			Send()
	}

	connectionUri := fmt.Sprintf("%s/%s", p.uri, p.name)

	if (p.isSslDisabled) {
		connectionUri = fmt.Sprintf("%s?sslmode=disable", connectionUri)
	}

	db, err := sql.Open("postgres", connectionUri)

	if err != nil {
		log.Info().Msg("Failed to connect to database")
		log.Fatal().Err(err).Msg("Failed to connect to database")
		panic(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal().Err(err).Msg("Failed to ping database")
	}

	p.db = db

	p.logger.Info().Msg("Connected to database")
}

func (p *postgresDb) Disconnect() {
	log.Info().Msg("Disconnecting from database")
	err := p.db.Close()

	if err != nil {
		p.logger.Error().Err(err).Str("status", "Failed to disconnect gracefully from database").Send()
	}

	p.logger.Info().Str("status", "disconnected successfully").Send()
}