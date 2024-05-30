package dbmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ChargingDataModel = (*customChargingDataModel)(nil)

type (
	// ChargingDataModel is an interface to be customized, add more methods here,
	// and implement the added methods in customChargingDataModel.
	ChargingDataModel interface {
		chargingDataModel
		withSession(session sqlx.Session) ChargingDataModel
	}

	customChargingDataModel struct {
		*defaultChargingDataModel
	}
)

// NewChargingDataModel returns a model for the database table.
func NewChargingDataModel(conn sqlx.SqlConn) ChargingDataModel {
	return &customChargingDataModel{
		defaultChargingDataModel: newChargingDataModel(conn),
	}
}

func (m *customChargingDataModel) withSession(session sqlx.Session) ChargingDataModel {
	return NewChargingDataModel(sqlx.NewSqlConnFromSession(session))
}
