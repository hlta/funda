import React, { useEffect, useState } from 'react';
import { Table, Input, Button } from 'reactstrap';
import axios from 'axios';
import { useHistory } from 'react-router-dom';

const AccountList = () => {
  const [accounts, setAccounts] = useState([]);
  const [searchTerm, setSearchTerm] = useState('');
  const history = useHistory();

  useEffect(() => {
    const fetchAccounts = async () => {
      const response = await axios.get('/api/accounts');
      setAccounts(response.data.data);
    };
    fetchAccounts();
  }, []);

  const handleSearch = (event) => {
    setSearchTerm(event.target.value);
  };

  const filteredAccounts = accounts.filter(account =>
    account.name.toLowerCase().includes(searchTerm.toLowerCase())
  );

  return (
    <div>
      <Input 
        type="text" 
        placeholder="Search accounts..." 
        value={searchTerm} 
        onChange={handleSearch} 
        className="mb-3"
      />
      <Table striped>
        <thead>
          <tr>
            <th>ID</th>
            <th>Code</th>
            <th>Name</th>
            <th>Type</th>
            <th>Tax Rate</th>
            <th>Balance</th>
            <th>Org Name</th>
            <th>YTD</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          {filteredAccounts.map(account => (
            <tr key={account.id}>
              <td>{account.id}</td>
              <td>{account.code}</td>
              <td>{account.name}</td>
              <td>{account.type}</td>
              <td>{account.tax_rate}</td>
              <td>{account.balance}</td>
              <td>{account.org_name}</td>
              <td>{account.ytd}</td>
              <td>
                <Button color="info" size="sm" onClick={() => history.push(`/accounting/chart-of-accounts/${account.id}`)}>View</Button>{' '}
                <Button color="primary" size="sm" onClick={() => history.push(`/accounting/chart-of-accounts/${account.id}/edit`)}>Edit</Button>
              </td>
            </tr>
          ))}
        </tbody>
      </Table>
      <Button color="primary" onClick={() => history.push('/accounting/chart-of-accounts/new')}>Add Account</Button>
    </div>
  );
};

export default AccountList;
