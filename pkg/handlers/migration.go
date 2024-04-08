package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MigrationHandler struct {
	db *gorm.DB
}

func NewMigrationHandler(db *gorm.DB) *MigrationHandler {
	return &MigrationHandler{db: db}
}

func (h *MigrationHandler) Migrate(ctx *gin.Context) {

	_ = h.db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	_ = h.db.Exec("ALTER TABLE IF EXISTS intern DROP CONSTRAINT IF EXISTS unique_phone_deleted_at")
	_ = h.db.Exec("ALTER TABLE IF EXISTS intern DROP CONSTRAINT IF EXISTS unique_dept_id_rank")
	models := []interface{}{
		// TO DEMO
	}
	for _, m := range models {
		err := h.db.AutoMigrate(m)
		if err != nil {
			_ = ctx.Error(err)
			return
		}
	}

	_ = h.db.Exec("ALTER TABLE intern ADD CONSTRAINT unique_phone_deleted_at UNIQUE(phone_number, deleted_at)")
	_ = h.db.Exec("ALTER TABLE intern ADD CONSTRAINT unique_dept_id_rank UNIQUE(dept_id, rank)")
}
