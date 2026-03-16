import React, { useEffect, useRef, useState } from "react";
import MetricCard from "./components/MetricCard.jsx";
import TelemetryChart from "./charts/TelemetryChart.jsx";
import { connectWebSocket, fetchHistory, fetchLatest } from "./api.js";

const MAX_POINTS = 120;

function clampSeries(series) {
  if (series.length <= MAX_POINTS) return series;
  return series.slice(series.length - MAX_POINTS);
}

export default function App() {
  const [latest, setLatest] = useState(null);
  const [history, setHistory] = useState([]);
  const [status, setStatus] = useState("connecting");
  const wsRef = useRef(null);

  useEffect(() => {
    let isActive = true;

    async function loadInitial() {
      try {
        const [latestData, historyData] = await Promise.all([
          fetchLatest().catch(() => null),
          fetchHistory()
        ]);
        if (!isActive) return;
        if (latestData) setLatest(latestData);
        setHistory(historyData);
      } catch (err) {
        setStatus("disconnected");
      }
    }

    loadInitial();

    const interval = setInterval(async () => {
      try {
        const latestData = await fetchLatest();
        if (!isActive) return;
        setLatest(latestData);
      } catch (err) {
        setStatus("disconnected");
      }
    }, 2000);

    return () => {
      isActive = false;
      clearInterval(interval);
    };
  }, []);

  useEffect(() => {
    const socket = connectWebSocket(
      (message) => {
        setStatus("connected");
        setLatest(message);
        setHistory((prev) => clampSeries([...prev, message]));
      },
      (state) => setStatus(state)
    );
    wsRef.current = socket;
    return () => socket.close();
  }, []);

  const labels = history.map((item) => item.timestamp?.slice(11, 19) || "--");
  const speedSeries = history.map((item) => item.speed);
  const rpmSeries = history.map((item) => item.rpm);
  const accelSeries = history.map((item) => item.acceleration);

  return (
    <div className="app">
      <header>
        <div>
          <div className="title">Vehicle Telemetry System</div>
          <div className="subtitle">
            Live vehicle data feed with historical performance trends.
          </div>
        </div>
        <div className="status">
          <span className="dot" style={{ background: status === "connected" ? "#0f766e" : "#b45309" }} />
          {status}
        </div>
      </header>

      <div className="grid metrics">
        <MetricCard label="Speed" value={latest?.speed} unit="km/h" accent="#0f766e" />
        <MetricCard label="RPM" value={latest?.rpm} unit="rpm" accent="#b45309" />
        <MetricCard label="Temperature" value={latest?.temperature} unit="C" accent="#0284c7" />
        <MetricCard label="Acceleration" value={latest?.acceleration} unit="g" accent="#0f172a" />
      </div>

      <div className="grid charts">
        <TelemetryChart title="Speed vs Time" labels={labels} series={speedSeries} color="#0f766e" />
        <TelemetryChart title="RPM vs Time" labels={labels} series={rpmSeries} color="#b45309" />
        <TelemetryChart title="Acceleration vs Time" labels={labels} series={accelSeries} color="#0284c7" />
      </div>
    </div>
  );
}

