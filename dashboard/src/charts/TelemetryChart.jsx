import React from "react";
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Tooltip,
  Legend
} from "chart.js";
import { Line } from "react-chartjs-2";

ChartJS.register(CategoryScale, LinearScale, PointElement, LineElement, Tooltip, Legend);

export default function TelemetryChart({ title, labels, series, color }) {
  const data = {
    labels,
    datasets: [
      {
        label: title,
        data: series,
        borderColor: color,
        backgroundColor: `${color}33`,
        fill: true,
        tension: 0.35,
        pointRadius: 0
      }
    ]
  };

  const options = {
    responsive: true,
    plugins: {
      legend: { display: false }
    },
    scales: {
      x: {
        ticks: { maxTicksLimit: 6 }
      }
    }
  };

  return (
    <div className="chart-shell">
      <div className="chart-title">{title}</div>
      <Line data={data} options={options} />
    </div>
  );
}

