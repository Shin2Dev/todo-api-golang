// PERAN SEBAGAI MODEL
package models

import "gorm.io/gorm"

// STRUCT = MODEL
type Todo struct {
	// STRUKTUR MODEL GOLANG -> FIELD | TYPE | MAPPING JSON
	// ID        int    `json: "id"`
	gorm.Model        // OTOMATIS MENAMBAHKAN KOLOM (ID, CREATED_AT, UPDATED_AT, DELETED_AT)
	Title      string `json: "title" binding: "required, min = 3"`
	Completed  bool   `json: "completed"`
}
