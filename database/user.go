package database

import (
	"fmt"

	"github.com/desmos-labs/plutus/types"
)

func (db *Database) SetUserPreferences(user *types.User, preferences *types.UserPreferences) error {
	// Insert the user
	userID, err := db.storeUSer(user)
	if err != nil {
		return err
	}

	stmt := `
INSERT INTO user_preferences (user_id, currency)
VALUES ($1, $2) ON CONFLICT (user_id) DO UPDATE 
	SET currency = excluded.currency`
	_, err = db.sql.Exec(stmt, user, userID, preferences.Currency)
	return err
}

type userPreferencesRow struct {
	UserID   uint64 `db:"user_id"`
	Currency string `db:"currency"`
}

func (db *Database) GetUserPreferences(desmosAddress string) (*types.UserPreferences, error) {
	stmt := `
SELECT user_preferences.* FROM user_preferences JOIN user_account on user_account.id = user_preferences.user_id
WHERE user_account.desmos_address = $1`

	var rows []userPreferencesRow
	err := db.sql.Select(&rows, stmt, desmosAddress)
	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return types.DefaultPreferences(), nil
	}

	if len(rows) > 1 {
		return nil, fmt.Errorf("multiple preferences for user with address %s", desmosAddress)
	}

	return types.NewUserPreferences(rows[0].Currency), nil
}
