import React, { useEffect, useState } from 'react';
import { useParams, useHistory } from 'react-router-dom';
import axios from 'axios';
import { Button } from 'reactstrap';

const AccountDetail = () => {
  const { id } = useParams();
  const [account, setAccount] = useState(null);
  const history = useHistory();

  useEffect(() => {
    const fetchAccount = async () => {
      const response = await axios.get(`/api/accounts/${id}`);
      setAccount(response.data.data);
    };

    fetchAccount();
  }, [id]);

  if (!account) {
    return <div>Loading...</div>;
  }

  return (
    <div>
      <h2>Account Details</h2>
      <p>ID: {account.id}</p>
      <p>Code: {account.code}</p>
      <p>Name: {account.name}</p>
      <p>Type: {account.type}</p>
      <p>Tax Rate: {account.tax_rate}</p>
      <p>Balance: {account.balance}</p>
      <p>Organization: {account.org_name}</p>
      <p>YTD: {account.ytd}</p>
      <Button color="primary" onClick={() => history.push(`/accounting/chart-of-accounts/${id}/edit`)}>Edit</Button>
    </div>
  );
};

export default AccountDetail;
