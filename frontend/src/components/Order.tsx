import React, { useState, useEffect } from 'react';
import { fetchOrders, createOrder } from '../api/orderService';

const Order: React.FC = () => {
  const [orders, setOrders] = useState<{ id: string; details: string }[]>([]);
  const [newOrder, setNewOrder] = useState<string>('');

  useEffect(() => {
    fetchOrders().then(response => setOrders(response.data));
  }, []);

  const handleAddOrder = () => {
    createOrder({ details: newOrder }).then(() => {
      fetchOrders().then(response => setOrders(response.data));
      setNewOrder('');
    });
  };

  return (
    <div>
      <h1>Orders</h1>
      <input
        value={newOrder}
        onChange={(e) => setNewOrder(e.target.value)}
        placeholder="New Order Details"
      />
      <button onClick={handleAddOrder}>Add Order</button>
      <ul>
        {orders.map(order => <li key={order.id}>{order.details}</li>)}
      </ul>
    </div>
  );
};

export default Order;
