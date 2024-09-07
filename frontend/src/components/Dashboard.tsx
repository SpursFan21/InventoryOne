// src/components/Dashboard.tsx
import React from 'react';
import { Link } from 'react-router-dom';
import './Dashboard.css';

const Dashboard: React.FC = () => {
  return (
    <div className="dashboard">
      <h1>Service Dashboard</h1>
      <div className="button-container">
        <Link to="/inventory" className="dashboard-button">Inventory Service</Link>
        <Link to="/notifications" className="dashboard-button">Notification Service</Link>
        <Link to="/orders" className="dashboard-button">Order Service</Link>
        <Link to="/shipments" className="dashboard-button">Shipment Service</Link>
      </div>
    </div>
  );
}

export default Dashboard;
