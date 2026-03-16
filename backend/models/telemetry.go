package models

import "errors"

type TelemetryInput struct {
	Speed        float64 `json:"speed"`
	RPM          float64 `json:"rpm"`
	Temperature  float64 `json:"temperature"`
	Acceleration float64 `json:"acceleration"`
}

func (t TelemetryInput) Validate() error {
	if t.Speed < 0 || t.Speed > 400 {
		return errors.New("speed out of range")
	}
	if t.RPM < 0 || t.RPM > 20000 {
		return errors.New("rpm out of range")
	}
	if t.Temperature < -40 || t.Temperature > 200 {
		return errors.New("temperature out of range")
	}
	if t.Acceleration < -10 || t.Acceleration > 10 {
		return errors.New("acceleration out of range")
	}
	return nil
}

type TelemetryRecord struct {
	ID           int64   `json:"id"`
	Timestamp    string  `json:"timestamp"`
	Speed        float64 `json:"speed"`
	RPM          float64 `json:"rpm"`
	Temperature  float64 `json:"temperature"`
	Acceleration float64 `json:"acceleration"`
}

