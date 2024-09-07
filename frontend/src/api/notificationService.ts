import axios from 'axios';

const API_URL = 'http://localhost:8081';

export const fetchNotifications = () => {
  return axios.get(`${API_URL}/notifications`);
};

export const createNotification = (notification: { message: string }) => {
  return axios.post(`${API_URL}/notifications`, notification);
};
