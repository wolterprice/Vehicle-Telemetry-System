#include <Arduino.h>
#include "sensors.h"
#include "telemetry.h"

// Transport configuration
#define USE_WIFI false

#if USE_WIFI
#include <WiFi.h>
#endif

static const unsigned long kSampleIntervalMs = 1000;
static unsigned long lastSampleMs = 0;

#if USE_WIFI
const char* kSsid = "YOUR_WIFI_SSID";
const char* kPassword = "YOUR_WIFI_PASSWORD";
const char* kServerHost = "192.168.1.100";
const int kServerPort = 8080;
#endif

void setup() {
  Serial.begin(115200);
  initSensors();

#if USE_WIFI
  WiFi.begin(kSsid, kPassword);
  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
  }
#endif
}

static void sendTelemetry(const String& payload) {
#if USE_WIFI
  WiFiClient client;
  if (!client.connect(kServerHost, kServerPort)) {
    return;
  }

  client.println("POST /telemetry HTTP/1.1");
  client.print("Host: ");
  client.println(kServerHost);
  client.println("Content-Type: application/json");
  client.print("Content-Length: ");
  client.println(payload.length());
  client.println();
  client.print(payload);
  client.stop();
#else
  Serial.println(payload);
#endif
}

void loop() {
  unsigned long now = millis();
  if (now - lastSampleMs < kSampleIntervalMs) {
    return;
  }

  float dtSeconds = (now - lastSampleMs) / 1000.0f;
  lastSampleMs = now;

  SensorReadings readings = readSensors(dtSeconds);
  Telemetry telemetry = buildTelemetry(readings);
  String json = telemetryToJson(telemetry);
  sendTelemetry(json);
}

