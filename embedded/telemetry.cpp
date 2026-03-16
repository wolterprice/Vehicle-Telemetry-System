#include "telemetry.h"
#include <ArduinoJson.h>

Telemetry buildTelemetry(const SensorReadings& readings) {
  Telemetry t{};
  t.speed = readings.speed;
  t.rpm = readings.rpm;
  t.temperature = readings.temperature;
  t.acceleration = readings.acceleration;
  return t;
}

String telemetryToJson(const Telemetry& telemetry) {
  StaticJsonDocument<256> doc;
  doc["speed"] = telemetry.speed;
  doc["rpm"] = telemetry.rpm;
  doc["temperature"] = telemetry.temperature;
  doc["acceleration"] = telemetry.acceleration;

  String out;
  serializeJson(doc, out);
  return out;
}

