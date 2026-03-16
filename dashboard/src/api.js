const DEFAULT_BASE = "http://localhost:8080";
const baseUrl = import.meta.env.VITE_API_URL || DEFAULT_BASE;

export async function fetchLatest() {
  const res = await fetch(`${baseUrl}/telemetry/latest`, { cache: "no-store" });
  if (!res.ok) {
    throw new Error("Failed to fetch latest telemetry");
  }
  return res.json();
}

export async function fetchHistory(limit = 200) {
  const res = await fetch(`${baseUrl}/telemetry/history?limit=${limit}`, {
    cache: "no-store"
  });
  if (!res.ok) {
    throw new Error("Failed to fetch telemetry history");
  }
  return res.json();
}

export function connectWebSocket(onMessage, onStatus) {
  const wsBase = baseUrl.replace("https://", "wss://").replace("http://", "ws://");
  const socket = new WebSocket(`${wsBase}/ws`);

  socket.onopen = () => onStatus?.("connected");
  socket.onclose = () => onStatus?.("disconnected");
  socket.onerror = () => onStatus?.("disconnected");
  socket.onmessage = (event) => {
    try {
      onMessage(JSON.parse(event.data));
    } catch (err) {
      // Ignore malformed messages.
    }
  };

  return socket;
}

