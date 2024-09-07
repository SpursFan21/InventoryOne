import axios from 'axios';

const API_URL = 'http://localhost:8083';

export const fetchShipments = () => {
  return axios.get(`${API_URL}/shipments`);
};

export const createShipment = (shipment: { details: string }) => {
  return axios.post(`${API_URL}/shipments`, shipment);
};

export const trackShipment = (id: string) => {
  return axios.get(`${API_URL}/shipments/${id}`);
};
