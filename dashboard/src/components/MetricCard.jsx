import React from "react";

export default function MetricCard({ label, value, unit, accent }) {
  const display = Number.isFinite(value) ? value.toFixed(1) : "--";
  return (
    <div className="card">
      <h3>{label}</h3>
      <div className="value mono" style={{ color: accent }}>
        {display}
        <span className="unit">{unit}</span>
      </div>
    </div>
  );
}

