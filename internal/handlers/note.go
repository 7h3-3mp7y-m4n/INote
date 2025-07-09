package handlers

import (
	"7h3-3mp7y-m4n/INote/internal/db"
	gen "7h3-3mp7y-m4n/INote/internal/db/generated"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type Note struct {
	Content   string    `json:"content"`
	ExpiresAt time.Time `json:"expires_at"`
	Format    string    `json:"format"`
}

func CreateNoteHandler(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input Note
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Basic validation for format (optional)
		if input.Format == "" {
			input.Format = "plain"
		}

		note, err := db.Queries.CreateNote(c.Request.Context(), gen.CreateNoteParams{
			Content:   input.Content,
			ExpiresAt: pgtype.Timestamptz{Time: input.ExpiresAt, Valid: true},
			Format:    pgtype.Text{String: input.Format, Valid: true},
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, note)
	}
}

func GetNoteHandler(db *db.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		slug := ctx.Param("slug")
		note, err := db.Queries.GetNoteBySlug(ctx.Request.Context(), slug)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "note not found"})
			return
		}
		ctx.JSON(http.StatusOK, note)
	}
}
