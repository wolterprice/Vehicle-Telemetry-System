#pragma once

#include <Arduino.h>
#include "sensors.h"

struct Telemetry {
  float speed;
  float rpm;
  float temperature;
  float acceleration;
};

Telemetry buildTelemetry(const SensorReadings& readings);
String telemetryToJson(const Telemetry& telemetry);

