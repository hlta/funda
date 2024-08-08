import { useContext, useCallback, useState } from 'react';
import { AuthContext } from '../contexts/AuthContext';
import { createApiClient } from '../services/apiClient';
import * as accountService from '../services/accountService';

export const useAccounts = () => {
  const { token } = useContext(AuthContext);
  const apiClient = createApiClient(token);
  const [accounts, setAccounts] = useState([]);
  const [loading, setLoading] = useState(true);

  const fetchAccounts = useCallback(async () => {
    setLoading(true);
    try {
      const data = await accountService.getAllAccounts(apiClient);
      setAccounts(data);
    } catch (error) {
      console.error('Error fetching accounts:', error);
    } finally {
      setLoading(false);
    }
  }, [apiClient]);

  const createAccount = useCallback(async (data) => {
    const newAccount = await accountService.createAccount(apiClient, data);
    setAccounts((prevAccounts) => [...prevAccounts, newAccount]);
    return newAccount;
  }, [apiClient]);

  const getAccount = useCallback(async (id) => {
    return await accountService.getAccount(apiClient, id);
  }, [apiClient]);

  const updateAccount = useCallback(async (id, data) => {
    const updatedAccount = await accountService.updateAccount(apiClient, id, data);
    setAccounts((prevAccounts) =>
      prevAccounts.map((account) =>
        account.id === id ? updatedAccount : account
      )
    );
    return updatedAccount;
  }, [apiClient]);

  const deleteAccount = useCallback(async (id) => {
    await accountService.deleteAccount(apiClient, id);
    setAccounts((prevAccounts) =>
      prevAccounts.filter((account) => account.id !== id)
    );
  }, [apiClient]);

  return {
    accounts,
    loading,
    fetchAccounts,
    createAccount,
    getAccount,
    updateAccount,
    deleteAccount,
  };
};
