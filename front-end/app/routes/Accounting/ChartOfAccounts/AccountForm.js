import React, { useState, useEffect } from 'react';
import { useHistory, useParams } from 'react-router-dom';
import { Form, FormGroup, Label, Input, Button } from 'reactstrap';
import axios from 'axios';

const AccountForm = () => {
  const { id } = useParams();
  const history = useHistory();
  const [account, setAccount] = useState({
    id: 0,
    code: 0,
    name: '',
    type: '',
    tax_rate: '',
    balance: 0,
    org_name: '',
    ytd: 0
  });

  useEffect(() => {
    if (id) {
      const fetchAccount = async () => {
        const response = await axios.get(`/api/accounts/${id}`);
        setAccount(response.data.data);
      };

      fetchAccount();
    }
  }, [id]);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setAccount(prevAccount => ({ ...prevAccount, [name]: value }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    if (id) {
      await axios.put(`/api/accounts/${id}`, account);
    } else {
      await axios.post('/api/accounts', account);
    }

    history.push('/accounting/chart-of-accounts');
  };

  return (
    <div>
      <h2>{id ? 'Edit Account' : 'Create Account'}</h2>
      <Form onSubmit={handleSubmit}>
        <FormGroup>
          <Label for="code">Code</Label>
          <Input type="number" name="code" id="code" value={account.code} onChange={handleChange} />
        </FormGroup>
        <FormGroup>
          <Label for="name">Name</Label>
          <Input type="text" name="name" id="name" value={account.name} onChange={handleChange} />
        </FormGroup>
        <FormGroup>
          <Label for="type">Type</Label>
          <Input type="text" name="type" id="type" value={account.type} onChange={handleChange} />
        </FormGroup>
        <FormGroup>
          <Label for="tax_rate">Tax Rate</Label>
          <Input type="text" name="tax_rate" id="tax_rate" value={account.tax_rate} onChange={handleChange} />
        </FormGroup>
        <FormGroup>
          <Label for="balance">Balance</Label>
          <Input type="number" name="balance" id="balance" value={account.balance} onChange={handleChange} />
        </FormGroup>
        <FormGroup>
          <Label for="org_name">Organization Name</Label>
          <Input type="text" name="org_name" id="org_name" value={account.org_name} onChange={handleChange} />
        </FormGroup>
        <FormGroup>
          <Label for="ytd">YTD</Label>
          <Input type="number" name="ytd" id="ytd" value={account.ytd} onChange={handleChange} />
        </FormGroup>
        <Button type="submit" color="primary">{id ? 'Update' : 'Create'}</Button>
      </Form>
    </div>
  );
};

export default AccountForm;
