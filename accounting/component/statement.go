package component

import (
	"context"
	"database/sql"

	"opencsg.com/csghub-server/accounting/types"
	"opencsg.com/csghub-server/builder/store/database"
)

type AccountingStatementComponent struct {
	asms *database.AccountStatementStore
}

func NewAccountingStatement() *AccountingStatementComponent {
	asc := &AccountingStatementComponent{
		asms: database.NewAccountStatementStore(),
	}
	return asc
}

func (a *AccountingStatementComponent) AddNewStatement(ctx context.Context, event *types.ACC_EVENT, eventExtra *types.ACC_EVENT_EXTRA, changeValue float64) error {
	statement := database.AccountStatement{
		EventUUID:  event.Uuid,
		UserID:     event.UserID,
		Value:      event.Value,
		Scene:      a.getValidSceneType(event.Scene),
		OpUID:      event.OpUID,
		CustomerID: eventExtra.CustomerID,
		EventDate:  event.CreatedAt,
		Price:      eventExtra.CustomerPrice,
		PriceUnit:  eventExtra.PriceUnit,
	}
	if event.Scene == int(database.SceneStarship) {
		// starship token count
		statement.Consumption = event.Value
	} else if event.Scene == int(database.SceneModelInference) || event.Scene == int(database.SceneSpace) || event.Scene == int(database.SceneModelFinetune) {
		// time duration of csghub resource
		statement.Consumption = eventExtra.CustomerDuration
	} else {
		statement.Consumption = 0
	}
	return a.asms.Create(ctx, statement, changeValue)
}

func (a *AccountingStatementComponent) ListStatementByUserIDAndTime(ctx context.Context, userID, startTime, endTime string, per, page int) ([]database.AccountStatement, error) {
	return a.asms.ListByUserIDAndTime(ctx, userID, startTime, endTime, per, page)
}

func (a *AccountingStatementComponent) FindStatementByEventID(ctx context.Context, event *types.ACC_EVENT) (*database.AccountStatement, error) {
	statement, err := a.asms.GetByEventID(ctx, event.Uuid)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &statement, err
}

func (a *AccountingStatementComponent) getValidSceneType(scene int) database.SceneType {
	switch scene {
	case 0:
		return database.SceneReserve
	case 1:
		return database.ScenePortalCharge
	case 10:
		return database.SceneModelInference
	case 11:
		return database.SceneSpace
	case 12:
		return database.SceneModelFinetune
	case 20:
		return database.SceneStarship
	default:
		return database.SceneUnknow
	}
}
