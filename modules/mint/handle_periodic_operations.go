package mint

import (
	"github.com/forbole/callisto/v4/modules/utils"

	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"
)

// RegisterPeriodicOperations implements modules.PeriodicOperationsModule
func (m *Module) RegisterPeriodicOperations(scheduler *gocron.Scheduler) error {
	log.Debug().Str("module", "mint").Msg("setting up periodic tasks")

	// Setup a cron job to run every 1 second (testing)
	if _, err := scheduler.Every(1).Second().Do(func() {
		utils.WatchMethod(m.UpdateInflation)
	}); err != nil {
		return err
	}

	return nil
}

// updateInflation fetches from the REST APIs the latest value for the
// inflation, and saves it inside the database.
func (m *Module) UpdateInflation() error {
	log.Debug().
		Str("module", "mint").
		Str("operation", "inflation").
		Msg("getting inflation data")

	block, err := m.db.GetLastBlockHeightAndTimestamp()
	if err != nil {
		return err
	}

	// Get the inflation
	inflation, err := m.source.GetInflation(block.Height)
	if err != nil {
		return err
	}

	log.Debug().
		Str("module", "mint").
		Str("operation", "inflation").
		Int64("height", block.Height).
		Str("inflation", inflation.String()).
		Msg("storing inflation")

	return m.db.SaveInflation(inflation.String(), block.Height)
}
