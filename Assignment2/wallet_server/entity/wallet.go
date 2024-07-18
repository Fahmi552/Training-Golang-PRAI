package entity

// type Wallet struct {
// 	UserID  int32   `gorm:"primary_key"`
// 	Balance float32 `gorm:"type:numeric(15,2);default:0.00"`
// }

type Wallet struct {
	UserID  int     `json:"user_id" gorm:"primaryKey;column:user_id"`
	Balance float64 `json:"balance" gorm:"column:balance;type:numeric(15,2);default:0.00"`
}

func (Wallet) TableName() string {
	return "twallets"
}
