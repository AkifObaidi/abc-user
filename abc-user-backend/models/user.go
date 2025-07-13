package models

import (
    "time"
)

type User struct {
    ID    string `json:"id" bson:"_id,omitempty" gorm:"primaryKey"`
    Name  string `json:"name" bson:"name" gorm:"type:varchar(100);not null"`
    Email string `json:"email" bson:"email" gorm:"type:varchar(100);unique;not null"`
    Age   int    `json:"age" bson:"age" gorm:"not null"`
    CreatedAt time.Time `json:"createdAt" bson:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updatedAt" bson:"updated_at,omitempty"`
}

