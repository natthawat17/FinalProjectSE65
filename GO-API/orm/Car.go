package orm

import (
	"gorm.io/gorm"  // framwork ต่อกับ database ภาษา GO
)

type Car struct {
	gorm.Model
	Carname string
	Detail string
	Image string
}