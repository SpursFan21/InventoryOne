import React, { useState, useEffect } from 'react';
import { fetchInventory, createInventoryItem } from '../api/inventoryService';

const Inventory: React.FC = () => {
  const [inventory, setInventory] = useState<{ id: string; name: string }[]>([]);
  const [newItem, setNewItem] = useState<string>('');

  useEffect(() => {
    fetchInventory().then(response => setInventory(response.data));
  }, []);

  const handleAddItem = () => {
    createInventoryItem({ name: newItem }).then(() => {
      fetchInventory().then(response => setInventory(response.data));
      setNewItem('');
    });
  };

  return (
    <div>
      <h1>Inventory</h1>
      <input
        value={newItem}
        onChange={(e) => setNewItem(e.target.value)}
        placeholder="New Inventory Item"
      />
      <button onClick={handleAddItem}>Add Item</button>
      <ul>
        {inventory.map(item => <li key={item.id}>{item.name}</li>)}
      </ul>
    </div>
  );
};

export default Inventory;
