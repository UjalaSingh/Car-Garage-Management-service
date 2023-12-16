package migrations

import (
	"gofr.dev/pkg/datastore"
	"gofr.dev/pkg/log"
)

type K20231113171002 struct {
}

func (k K20231113171002) Up(d *datastore.DataStore, logger log.Logger) error {
	if _, err := d.DB().Exec(createCarTable); err != nil {
		return err
	}

	return nil
}

func (k K20231113171002) Down(d *datastore.DataStore, logger log.Logger) error {
	if _, err := d.DB().Exec(dropCarTable); err != nil {
		return err
	}

	return nil
}
