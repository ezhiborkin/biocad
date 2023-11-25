package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// @Summary Get errors data
// @Description Get errors data based on filename with pagination
// @Produce json
// @Param filename query string true "File name to filter errors"
// @Param page query int false "Page number for pagination (default is 1)"
// @Param limit query int false "Number of items to show per page (default is 10)"
// @Success 200 {array} er.ErrorOpenFile
// @Failure 500 {object} er.errorResponse
// @Router /api/errorsdata [get]
func (h *Handler) getErrorsData(c *gin.Context) {
	fileName := c.Query("filename")
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

	filter := bson.M{"filename": fileName}

	options := options.Find().SetSkip(int64(skip)).SetLimit(int64(limit))

	units, err := h.services.GetProcessedErrorS(filter, options)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	}

	c.JSON(http.StatusOK, units)
}
