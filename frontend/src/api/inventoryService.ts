import axios from 'axios';

const API_URL = 'http://localhost:8080';

export const fetchInventory = () => {
  return axios.get(`${API_URL}/inventory`);
};

export const createInventoryItem = (item: { name: string }) => {
  return axios.post(`${API_URL}/inventory`, item);
};
