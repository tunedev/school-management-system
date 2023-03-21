package migration

import (
	"database/sql"
	"fmt"
	"io/ioutil"

	"github.com/pkg/errors"
)

type Migration struct {
	DB     *sql.DB
	Script string
}

func (m *Migration) Up() error {
	// Read SQL migration script
	script, err := ioutil.ReadFile(m.Script)
	if err != nil {
		return errors.Wrap(err, "failed to read migration script file")
	}

	// Execute SQL migration script
	if _, err := m.DB.Exec(string(script)); err != nil {
		return errors.Wrap(err, "failed to execute migration script")
	}

	fmt.Println("Migration successful!")
	return nil
}

func (m *Migration) Down() error {
	// TODO: implement database rollback
	return nil
}
