package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (h *Handler) getProcessedFiles(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	// Set default values for page and limit
	page, _ := strconv.Atoi(pageStr)
	if page <= 0 {
		page = 1
	}

	limit, _ := strconv.Atoi(limitStr)
	if limit <= 0 {
		limit = 10 // Set a default limit
	}

	// Calculate the skip value for pagination
	skip := (page - 1) * limit

	// Define the filter based on the unit_guid
	// filter := bson.M{"filepath": ""}

	// Define options for pagination
	options := options.Find().SetSkip(int64(skip)).SetLimit(int64(limit))

	files, err := h.services.GetProcessedFileS(options)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	}

	c.JSON(http.StatusOK, files)
}
