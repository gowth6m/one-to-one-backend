package one_to_one

import (
	"net/http"
	"one-to-one/internal/api"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OneToOneHandler struct {
	Repo OneToOneRepository
}

func NewOneToOneHandler(repo OneToOneRepository) *OneToOneHandler {
	return &OneToOneHandler{Repo: repo}
}

// @Summary Create a new weekly report
// @Description Create a new weekly report
// @Tags one-to-one
// @Accept json
// @Produce json
// @Param report body CreateWeeklyReportRequest true "Weekly report object to be created"
// @Success 201 {object} WeeklyReportResponse "Weekly report created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request format or parameters"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /one-to-one/create [post]
func (h *OneToOneHandler) CreateWeeklyReport(c *gin.Context) {
	var reqPayload CreateWeeklyReportRequest
	if err := c.ShouldBindJSON(&reqPayload); err != nil {
		api.Error(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	userID, err := primitive.ObjectIDFromHex(c.GetString("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	createdReport, err := h.Repo.CreateWeeklyReport(c.Request.Context(), reqPayload, userID)
	if err != nil {
		api.Error(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	api.Success(c, http.StatusCreated, "Created weekly report successfully", createdReport)
}

// @Summary Get all weekly reports for a reportee
// @Description Get all weekly reports
// @Tags one-to-one
// @Accept json
// @Produce json
// @Success 200 {array} WeeklyReportResponse "List of weekly reports"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /one-to-one/reportee/all [get]
func (h *OneToOneHandler) GetAllWeeklyReportsForReportee(c *gin.Context) {

	userID, err := primitive.ObjectIDFromHex(c.GetString("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	reports, err := h.Repo.GetAllWeeklyReports(c.Request.Context(), userID, true)
	if err != nil {
		api.Error(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	api.Success(c, http.StatusOK, "Fetched all weekly reports successfully", reports)
}

// @Summary Update a weekly report for a reportee
// @Description Update a weekly report
// @Tags one-to-one
// @Accept json
// @Produce json
// @Param report body UpdateWeeklyReportRequest true "Weekly report object to be updated"
// @Success 200 {object} WeeklyReportResponse "Weekly report updated successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request format or parameters"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /one-to-one/reportee/update [put]
func (h *OneToOneHandler) UpdateWeeklyReportForReportee(c *gin.Context) {
	var reqPayload UpdateWeeklyReportRequest
	if err := c.ShouldBindJSON(&reqPayload); err != nil {
		api.Error(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	userID, err := primitive.ObjectIDFromHex(c.GetString("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	updatedReport, err := h.Repo.UpdateWeeklyReport(c.Request.Context(), reqPayload, userID, true)
	if err != nil {
		api.Error(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	api.Success(c, http.StatusOK, "Updated weekly report successfully", updatedReport)
}

// @Summary Get a weekly report by week and year for a reportee
// @Description Get a weekly report by week and year for a reportee
// @Tags one-to-one
// @Accept json
// @Produce json
// @Param week query int true "Week number"
// @Param year query int true "Year"
// @Success 200 {object} WeeklyReportResponse "Weekly report"
// @Failure 400 {object} map[string]interface{} "Invalid request format or parameters"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /one-to-one/reportee [get]
func (h *OneToOneHandler) GetWeeklyReportByWeekAndYearForReportee(c *gin.Context) {
	weekStr := c.Query("week")
	yearStr := c.Query("year")

	week, errWeek := strconv.Atoi(weekStr)
	year, errYear := strconv.Atoi(yearStr)

	if errWeek != nil || errYear != nil {
		week, year = GetCurrentWeekAndYear()
	}

	userID, err := primitive.ObjectIDFromHex(c.GetString("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	report, err := h.Repo.GetWeeklyReportByWeekAndYear(c.Request.Context(), week, year, userID, true)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "No weekly report found"})
		} else {
			api.Error(c, http.StatusInternalServerError, "Error fetching weekly report", nil)
		}
		return
	}

	api.Success(c, http.StatusOK, "Fetched weekly report successfully", report)
}

// @Summary Get all weekly reports for a reportTo
// @Description Get all weekly reports
// @Tags one-to-one
// @Accept json
// @Produce json
// @Success 200 {array} WeeklyReportResponse "List of weekly reports"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /one-to-one/report-to/all [get]
func (h *OneToOneHandler) GetAllWeeklyReportsForReportTo(c *gin.Context) {

	userID, err := primitive.ObjectIDFromHex(c.GetString("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	reports, err := h.Repo.GetAllWeeklyReports(c.Request.Context(), userID, false)
	if err != nil {
		api.Error(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	api.Success(c, http.StatusOK, "Fetched all weekly reports successfully", reports)
}

// @Summary Update a weekly report for a reportTo
// @Description Update a weekly report
// @Tags one-to-one
// @Accept json
// @Produce json
// @Param report body UpdateWeeklyReportRequest true "Weekly report object to be updated"
// @Success 200 {object} WeeklyReportResponse "Weekly report updated successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request format or parameters"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /one-to-one/report-to/update [put]
func (h *OneToOneHandler) UpdateWeeklyReportForReportTo(c *gin.Context) {
	var reqPayload UpdateWeeklyReportRequest
	if err := c.ShouldBindJSON(&reqPayload); err != nil {
		api.Error(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	userID, err := primitive.ObjectIDFromHex(c.GetString("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	updatedReport, err := h.Repo.UpdateWeeklyReport(c.Request.Context(), reqPayload, userID, false)
	if err != nil {
		api.Error(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	api.Success(c, http.StatusOK, "Updated weekly report successfully", updatedReport)
}

// @Summary Get a weekly report by week and year for a reportTo
// @Description Get a weekly report by week and year for a reportTo
// @Tags one-to-one
// @Accept json
// @Produce json
// @Param week query int true "Week number"
// @Param year query int true "Year"
// @Success 200 {object} WeeklyReportResponse "Weekly report"
// @Failure 400 {object} map[string]interface{} "Invalid request format or parameters"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /one-to-one/report-to [get]
func (h *OneToOneHandler) GetWeeklyReportByWeekAndYearForReportTo(c *gin.Context) {
	weekStr := c.Query("week")
	yearStr := c.Query("year")

	week, errWeek := strconv.Atoi(weekStr)
	year, errYear := strconv.Atoi(yearStr)

	if errWeek != nil || errYear != nil {
		week, year = GetCurrentWeekAndYear()
	}

	userID, err := primitive.ObjectIDFromHex(c.GetString("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	report, err := h.Repo.GetWeeklyReportByWeekAndYear(c.Request.Context(), week, year, userID, false)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "No weekly report found"})
		} else {
			api.Error(c, http.StatusInternalServerError, "Error fetching weekly report", nil)
		}
		return
	}

	api.Success(c, http.StatusOK, "Fetched weekly report successfully", report)
}
