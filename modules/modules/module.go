package modules

import (
	"github.com/forbole/juno/v4/modules"
	"github.com/forbole/juno/v4/types/config"

	"github.com/forbole/bdjuno/v4/database"
)

var (
	_ modules.Module                     = &Module{}
	_ modules.AdditionalOperationsModule = &Module{}
)

type Module struct {
	db  *database.Db
	cfg config.ChainConfig
}

// NewModule returns a new Module instance
func NewModule(cfg config.ChainConfig, db *database.Db) *Module {
	return &Module{
		cfg: cfg,
		db:  db,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "modules"
}
