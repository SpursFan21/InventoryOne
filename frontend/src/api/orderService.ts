import axios from 'axios';

const API_URL = 'http://localhost:8082';

export const fetchOrders = () => {
  return axios.get(`${API_URL}/orders`);
};

export const createOrder = (order: { details: string }) => {
  return axios.post(`${API_URL}/orders`, order);
};
