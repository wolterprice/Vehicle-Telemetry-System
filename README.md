# Vehicle Telemetry System

Production-style, end-to-end telemetry pipeline that simulates automotive data flow from an embedded device to a backend API and a live web dashboard. The project is organized to resemble real-world telemetry stacks used in motorsports, EV development, and fleet analytics.

## Overview

Flow: Sensors → Arduino → Serial/WiFi → Backend API → SQLite → Dashboard

Key features:
- Embedded firmware simulates speed, RPM, temperature, and acceleration
- Backend validates and stores telemetry, serves REST + WebSocket
- Dashboard shows live metrics and historical charts

## Repository Structure

```
embedded/   Arduino firmware
backend/    Go API server
dashboard/  React dashboard (Vite)
```

## Dependencies

Embedded (Arduino):
- ArduinoJson
- Wire
- MPU6050 (optional, only if using real sensor)
- WiFi library (optional, if using WiFi transport)

Backend (Go):
- github.com/gin-gonic/gin
- github.com/gorilla/websocket
- github.com/mattn/go-sqlite3

Dashboard (Node):
- react, react-dom
- chart.js, react-chartjs-2
- vite, @vitejs/plugin-react

## Running The Project

### 1) Backend (Go)

```
cd backend
go mod download
go run .
```

The server starts on `http://localhost:8080`.

### 2) Dashboard (React + Vite)

```
cd dashboard
npm install
npm run dev
```

Open `http://localhost:5173`.

To point the dashboard to a different API:
```
VITE_API_URL=http://localhost:8080 npm run dev
```

### 3) Embedded (Arduino)

Open `embedded/main.ino` in Arduino IDE.
- Ensure `ArduinoJson` is installed.
- Optional: set `USE_WIFI` to `true` and configure SSID.
- Upload to device.

By default, telemetry is printed to Serial every 1 second.

## REST API

### POST /telemetry
```
curl -X POST http://localhost:8080/telemetry \
  -H "Content-Type: application/json" \
  -d "{\"speed\":82.4,\"rpm\":3100,\"temperature\":88.2,\"acceleration\":0.45}"
```

### GET /telemetry/latest
```
curl http://localhost:8080/telemetry/latest
```

### GET /telemetry/history?limit=100
```
curl "http://localhost:8080/telemetry/history?limit=100"
```

## WebSocket

Live updates are broadcast to:
```
ws://localhost:8080/ws
```

## Docker (Optional)

```
docker compose up --build
```

Dashboard is served at `http://localhost:5173` and API at `http://localhost:8080`.

## Makefile

```
make backend-dev
make dashboard-dev
make docker-up
```

