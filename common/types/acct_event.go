package types

import (
	"time"

	"github.com/google/uuid"
)

var (
	REASON_SUCCESS        = 0 // charge success
	REASON_INVALID_FORMAT = 1 // invalid event data format
	REASON_CHARGE_FAIL    = 2 // fail to charge user fee
	REASON_LACK_BALANCE   = 3 // balance <= 0
	REASON_DUPLICATED     = 4 // duplicated charge
)

// generate charge event from client
type ACC_EVENT struct {
	Uuid      uuid.UUID `json:"uuid"`
	UserID    string    `json:"user_id"`
	Value     float64   `json:"value"`
	ValueType int       `json:"value_type"`
	Scene     int       `json:"scene"`
	OpUID     int64     `json:"op_uid"`
	CreatedAt time.Time `json:"created_at"`
	Extra     string    `json:"extra"`
}

// notify response to client
type ACC_NOTIFY struct {
	Uuid       uuid.UUID `json:"uuid"`
	UserID     string    `json:"user_id"`
	CreatedAt  time.Time `json:"created_at"`
	ReasonCode int       `json:"reason_code"`
	ReasonMsg  string    `json:"reason_msg"`
}

type CONSUMER_INFO struct {
	ConsumerID    string `json:"customer_id"`
	ConsumerPrice string `json:"customer_price"`
	PriceUnit     string `json:"price_unit"`
	Duration      string `json:"customer_duration"`
}

var (
	SceneReserve        int = 0  // system reserve
	ScenePortalCharge   int = 1  // portal charge fee
	SceneModelInference int = 10 // model inference endpoint
	SceneSpace          int = 11 // csghub space
	SceneModelFinetune  int = 12 // model finetune
	SceneStarship       int = 20 // starship
	SceneUnknow         int = 99 // unknow
)
