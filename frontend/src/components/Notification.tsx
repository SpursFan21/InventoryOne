import React, { useState, useEffect } from 'react';
import { fetchNotifications, createNotification } from '../api/notificationService';

const Notification: React.FC = () => {
  const [notifications, setNotifications] = useState<{ id: string; message: string }[]>([]);
  const [newNotification, setNewNotification] = useState<string>('');

  useEffect(() => {
    fetchNotifications().then(response => setNotifications(response.data));
  }, []);

  const handleAddNotification = () => {
    createNotification({ message: newNotification }).then(() => {
      fetchNotifications().then(response => setNotifications(response.data));
      setNewNotification('');
    });
  };

  return (
    <div>
      <h1>Notifications</h1>
      <input
        value={newNotification}
        onChange={(e) => setNewNotification(e.target.value)}
        placeholder="New Notification"
      />
      <button onClick={handleAddNotification}>Add Notification</button>
      <ul>
        {notifications.map(notification => <li key={notification.id}>{notification.message}</li>)}
      </ul>
    </div>
  );
};

export default Notification;
