package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (h *Handler) getProcessedData(c *gin.Context) {
	unitGUID := c.Query("unit_guid")
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
	filter := bson.M{"unitguid": unitGUID}

	// Define options for pagination
	options := options.Find().SetSkip(int64(skip)).SetLimit(int64(limit))

	units, err := h.services.GetProcessedDataS(filter, options)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	}

	c.JSON(http.StatusOK, units)
}
