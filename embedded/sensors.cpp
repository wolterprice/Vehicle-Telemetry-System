#include "sensors.h"

// Toggle this to use simulated values if sensors are not available.
#ifndef USE_SIMULATION
#define USE_SIMULATION true
#endif

static const float kWheelCircumferenceMeters = 2.05f; // typical passenger car tire
static const float kPulsesPerRevolution = 2.0f;        // hall sensor magnets

void initSensors() {
#if !USE_SIMULATION
  // TODO: Initialize I2C and real sensors here.
  // Example: Wire.begin(); mpu.initialize();
#endif
}

static SensorReadings simulateReadings(float t) {
  SensorReadings r{};
  r.rpm = 1500.0f + 1200.0f * sin(t * 0.5f);
  if (r.rpm < 700.0f) r.rpm = 700.0f;
  float speedMps = (r.rpm / 60.0f) * kWheelCircumferenceMeters;
  r.speed = speedMps * 3.6f;
  r.temperature = 80.0f + 10.0f * sin(t * 0.2f);
  r.acceleration = 0.2f + 0.4f * sin(t * 0.9f);
  return r;
}

SensorReadings readSensors(float dtSeconds) {
  (void)dtSeconds;
  float t = millis() / 1000.0f;

#if USE_SIMULATION
  return simulateReadings(t);
#else
  SensorReadings r{};
  // TODO: Replace with actual sensor reads.
  r.rpm = 0.0f;
  r.speed = 0.0f;
  r.temperature = 0.0f;
  r.acceleration = 0.0f;
  return r;
#endif
}

