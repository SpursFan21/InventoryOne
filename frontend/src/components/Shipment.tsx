import React, { useState, useEffect } from 'react';
import { fetchShipments, createShipment, trackShipment } from '../api/shipmentService';

const Shipment: React.FC = () => {
  const [shipments, setShipments] = useState<{ id: string; details: string }[]>([]);
  const [newShipment, setNewShipment] = useState<string>('');
  const [trackId, setTrackId] = useState<string>('');
  const [trackedShipment, setTrackedShipment] = useState<any | null>(null);

  useEffect(() => {
    fetchShipments().then(response => setShipments(response.data));
  }, []);

  const handleAddShipment = () => {
    createShipment({ details: newShipment }).then(() => {
      fetchShipments().then(response => setShipments(response.data));
      setNewShipment('');
    });
  };

  const handleTrackShipment = () => {
    trackShipment(trackId).then(response => setTrackedShipment(response.data));
  };

  return (
    <div>
      <h1>Shipments</h1>
      <input
        value={newShipment}
        onChange={(e) => setNewShipment(e.target.value)}
        placeholder="New Shipment Details"
      />
      <button onClick={handleAddShipment}>Add Shipment</button>
      <input
        value={trackId}
        onChange={(e) => setTrackId(e.target.value)}
        placeholder="Track Shipment ID"
      />
      <button onClick={handleTrackShipment}>Track Shipment</button>
      <ul>
        {shipments.map(shipment => <li key={shipment.id}>{shipment.details}</li>)}
      </ul>
      {trackedShipment && <div>Tracked Shipment: {JSON.stringify(trackedShipment)}</div>}
    </div>
  );
};

export default Shipment;
