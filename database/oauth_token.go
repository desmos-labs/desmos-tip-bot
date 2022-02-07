package database

import (
	"fmt"
	"time"

	"github.com/desmos-labs/plutus/types"
)

// SaveUserData allows to store the given user data
func (db *Database) SaveUserData(
	user *types.User, serviceAccount *types.ServiceAccount, applications []*types.ApplicationAccount,
) error {
	// Insert the user account
	var userID uint64
	stmt := `INSERT INTO user_account (desmos_address) VALUES ($1) ON CONFLICT DO NOTHING RETURNING id`
	err := db.sql.QueryRow(stmt, user.DesmosAddress).Scan(&userID)
	if err != nil {
		return err
	}

	// Insert the service account
	var serviceID uint64
	stmt = `
INSERT INTO service_account (user_id, service, access_token, refresh_token, creation_time)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT ON CONSTRAINT unique_service_account DO UPDATE 
    SET access_token = excluded.access_token,
    	refresh_token = excluded.refresh_token,
		creation_time = excluded.creation_time
RETURNING id`
	err = db.sql.QueryRow(stmt,
		userID, serviceAccount.Service, serviceAccount.AccessToken, serviceAccount.RefreshToken, time.Now(),
	).Scan(&serviceID)
	if err != nil {
		return err
	}

	// Insert the application accounts
	for _, applicationAccount := range applications {
		stmt = `
INSERT INTO application_account (service_account_id, application, username) 
VALUES ($1, $2, $3)
ON CONFLICT DO NOTHING`
		_, err = db.sql.Exec(stmt, serviceID, applicationAccount.Application, applicationAccount.Username)
		if err != nil {
			return err
		}
	}

	return nil
}

// --------------------------------------------------------------------------------------------------------------------

type serviceAccountRow struct {
	ID           uint64    `db:"id"`
	UserID       uint64    `db:"user_id"`
	Service      string    `db:"service"`
	AccessToken  string    `db:"access_token"`
	RefreshToken string    `db:"refresh_token"`
	CreationTime time.Time `db:"creation_time"`
}

// GetServiceAccounts returns the accounts associated with the given Desmos address
func (db *Database) GetServiceAccounts(desmosAddress string) ([]*types.ServiceAccount, error) {
	stmt := `
SELECT service_account.* 
FROM service_account JOIN user_account on user_account.id = service_account.user_id
WHERE user_account.desmos_address = $1`

	var rows []serviceAccountRow
	err := db.sql.Select(&rows, stmt, desmosAddress)
	if err != nil {
		return nil, err
	}

	accounts := make([]*types.ServiceAccount, len(rows))
	for i, row := range rows {
		accounts[i] = types.NewServiceAccount(row.Service, row.AccessToken, row.RefreshToken)
	}

	return accounts, nil
}

// --------------------------------------------------------------------------------------------------------------------

type applicationAccountRow struct {
	ServiceAccountID uint64 `db:"service_account_id"`
	Application      string `db:"application"`
	Username         string `db:"username"`
}

func (db *Database) GetAppAccount(application, username string) (*types.ApplicationAccount, error) {
	stmt := `SELECT * FROM application_account WHERE application ILIKE $1 AND username ILIKE $2`

	var rows []applicationAccountRow
	err := db.sql.Select(&rows, stmt, application, username)
	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return nil, nil
	}

	if len(rows) > 1 {
		return nil, fmt.Errorf("application account should be only one, found %d", len(rows))
	}

	return types.NewApplicationAccount(rows[0].Application, rows[0].Username), nil

}
