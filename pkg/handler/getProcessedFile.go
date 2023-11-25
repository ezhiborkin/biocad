package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// @Summary Get processed files
// @Description Get processed files with pagination
// @Tags APIs
// @Produce json
// @Param page query int false "Page number for pagination (default is 1)"
// @Param limit query int false "Number of items to show per page (default is 10)"
// @Success 200 {array} unit.ProcessedFile
// @Failure 500 {object} er.errorResponse
// @Router /processedfiles [get]
func (h *Handler) getProcessedFiles(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, _ := strconv.Atoi(pageStr)
	if page <= 0 {
		page = 1
	}

	limit, _ := strconv.Atoi(limitStr)
	if limit <= 0 {
		limit = 10
	}

	skip := (page - 1) * limit

	// filter := bson.M{"filepath": ""}

	options := options.Find().SetSkip(int64(skip)).SetLimit(int64(limit))

	files, err := h.services.GetProcessedFileS(options)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	}

	c.JSON(http.StatusOK, files)
}
