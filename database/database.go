package database

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type User struct {
	ID      string `gorm:"primaryKey"`
	Address string
}

type CommentMessage struct {
	MessageID                 uuid.UUID `gorm:"primaryKey"`
	Sent                      bool
	ProfileId                 decimal.Decimal
	PubId                     decimal.Decimal
	ContentURI                string
	ProfileIdPointed          decimal.Decimal
	PubIdPointed              decimal.Decimal
	CollectModule             string
	CollectModuleReturnData   string
	ReferenceModule           string
	ReferenceModuleReturnData string
	Timestamp                 decimal.Decimal
}

type FollowMessage struct {
	MessageID uuid.UUID `gorm:"primaryKey"`
	ProfileId decimal.Decimal
	FollowNFT string
	Timestamp decimal.Decimal
	Sent      bool
}

type Clone struct {
	Day        time.Time `gorm:"primaryKey"`
	Count      int
	Uniques    int
	Repository string `gorm:"primaryKey"`
}

type View struct {
	Day        time.Time `gorm:"primaryKey"`
	Count      int
	Uniques    int
	Repository string `gorm:"primaryKey"`
}

type Path struct {
	Path       string `gorm:"primaryKey"`
	Title      string
	Count      int
	Uniques    int
	Day        time.Time `gorm:"primaryKey"`
	Repository string    `gorm:"primaryKey"`
}

type Referral struct {
	Referrer   string `gorm:"primaryKey"`
	Count      int
	Uniques    int
	Day        time.Time `gorm:"primaryKey"`
	Repository string    `gorm:"primaryKey"`
}
