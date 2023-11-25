package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// @Summary Get processed data
// @Description Get processed data based on unit GUID with pagination
// @Tags APIs
// @Produce json
// @Param unit_guid query string true "Unit GUID to filter processed data"
// @Param page query int false "Page number for pagination (default is 1)"
// @Param limit query int false "Number of items to show per page (default is 10)"
// @Success 200 {array} unit.Unit
// @Failure 500 {object} er.errorResponse
// @Router /processeddata [get]
func (h *Handler) getProcessedData(c *gin.Context) {
	unitGUID := c.Query("unit_guid")
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

	filter := bson.M{"unitguid": unitGUID}

	options := options.Find().SetSkip(int64(skip)).SetLimit(int64(limit))

	units, err := h.services.GetProcessedDataS(filter, options)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	}

	c.JSON(http.StatusOK, units)
}
