#pragma once

#include <Arduino.h>

struct SensorReadings {
  float speed;
  float rpm;
  float temperature;
  float acceleration;
};

void initSensors();
SensorReadings readSensors(float dtSeconds);

