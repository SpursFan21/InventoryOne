// src/App.tsx
import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Inventory from './components/Inventory';
import Notification from './components/Notification';
import Order from './components/Order';
import Shipment from './components/Shipment';
import Dashboard from './components/Dashboard';

const App: React.FC = () => {
  return (
    <Router>
      <div>
        <Routes>
          <Route path="/" element={<Dashboard />} />
          <Route path="/inventory" element={<Inventory />} />
          <Route path="/notifications" element={<Notification />} />
          <Route path="/orders" element={<Order />} />
          <Route path="/shipments" element={<Shipment />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
