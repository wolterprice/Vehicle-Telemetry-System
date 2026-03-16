package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"vehicle-telemetry-system/backend/models"
)

type TelemetryHandler struct {
	db *sql.DB
}

func NewTelemetryHandler(db *sql.DB) *TelemetryHandler {
	return &TelemetryHandler{db: db}
}

func (h *TelemetryHandler) PostTelemetry(c *gin.Context) {
	var input models.TelemetryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}
	if err := input.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.db.Exec(
		`INSERT INTO telemetry (speed, rpm, temperature, acceleration)
		 VALUES (?, ?, ?, ?)`,
		input.Speed, input.RPM, input.Temperature, input.Acceleration,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db insert failed"})
		return
	}

	id, _ := res.LastInsertId()
	record, err := h.fetchByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db read failed"})
		return
	}

	c.JSON(http.StatusCreated, record)
}

func (h *TelemetryHandler) GetLatest(c *gin.Context) {
	row := h.db.QueryRow(
		`SELECT id, speed, rpm, temperature, acceleration, timestamp
		 FROM telemetry ORDER BY id DESC LIMIT 1`,
	)

	var record models.TelemetryRecord
	if err := row.Scan(&record.ID, &record.Speed, &record.RPM, &record.Temperature, &record.Acceleration, &record.Timestamp); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "no telemetry yet"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db read failed"})
		return
	}
	c.JSON(http.StatusOK, record)
}

func (h *TelemetryHandler) GetHistory(c *gin.Context) {
	limit := 200
	if val := c.Query("limit"); val != "" {
		if parsed, err := strconv.Atoi(val); err == nil && parsed > 0 && parsed <= 1000 {
			limit = parsed
		}
	}

	rows, err := h.db.Query(
		`SELECT id, speed, rpm, temperature, acceleration, timestamp
		 FROM telemetry ORDER BY id DESC LIMIT ?`, limit,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db read failed"})
		return
	}
	defer rows.Close()

	records := make([]models.TelemetryRecord, 0, limit)
	for rows.Next() {
		var record models.TelemetryRecord
		if err := rows.Scan(&record.ID, &record.Speed, &record.RPM, &record.Temperature, &record.Acceleration, &record.Timestamp); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "db scan failed"})
			return
		}
		records = append(records, record)
	}

	// Reverse to return ascending time order.
	for i, j := 0, len(records)-1; i < j; i, j = i+1, j-1 {
		records[i], records[j] = records[j], records[i]
	}

	c.JSON(http.StatusOK, records)
}

func (h *TelemetryHandler) fetchByID(id int64) (models.TelemetryRecord, error) {
	row := h.db.QueryRow(
		`SELECT id, speed, rpm, temperature, acceleration, timestamp
		 FROM telemetry WHERE id = ?`, id,
	)

	var record models.TelemetryRecord
	if err := row.Scan(&record.ID, &record.Speed, &record.RPM, &record.Temperature, &record.Acceleration, &record.Timestamp); err != nil {
		return models.TelemetryRecord{}, err
	}
	return record, nil
}
