package models

import (
	"github.com/nikola43/ethfaucet"
)

type Client struct {
	base.CustomGormModel
	//ClinicID              uint                   `gorm:"type:INTEGER NULL; DEFAULT:NULL" json:"clinic_id" xml:"clinic_id" form:"clinic_id"`
	Email                 string                      `gorm:"index; unique; type:varchar(64) not null" json:"email"`

}
