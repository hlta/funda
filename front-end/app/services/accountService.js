// Define API endpoints as constants
const API_ENDPOINTS = {
    ACCOUNTS: '/api/accounts',
  };
  
  /**
   * Creates a new account.
   * @param {Object} apiClient - The API client instance.
   * @param {Object} data - The data for the new account.
   * @returns {Promise<Object>} The created account data.
   */
  export const createAccount = async (apiClient, data) => {
    const response = await apiClient.post(API_ENDPOINTS.ACCOUNTS, data);
    return response.data.data;
  };
  
  /**
   * Retrieves an account by its ID.
   * @param {Object} apiClient - The API client instance.
   * @param {number} id - The ID of the account.
   * @returns {Promise<Object>} The account data.
   */
  export const getAccount = async (apiClient, id) => {
    const response = await apiClient.get(`${API_ENDPOINTS.ACCOUNTS}/${id}`);
    return response.data.data;
  };
  
  /**
   * Retrieves all accounts.
   * @param {Object} apiClient - The API client instance.
   * @returns {Promise<Array>} The list of accounts.
   */
  export const getAllAccounts = async (apiClient) => {
    const response = await apiClient.get(API_ENDPOINTS.ACCOUNTS);
    return response.data.data;
  };
  
  /**
   * Updates an existing account by its ID.
   * @param {Object} apiClient - The API client instance.
   * @param {number} id - The ID of the account.
   * @param {Object} data - The updated data for the account.
   * @returns {Promise<Object>} The updated account data.
   */
  export const updateAccount = async (apiClient, id, data) => {
    const response = await apiClient.put(`${API_ENDPOINTS.ACCOUNTS}/${id}`, data);
    return response.data.data;
  };
  
  /**
   * Deletes an account by its ID.
   * @param {Object} apiClient - The API client instance.
   * @param {number} id - The ID of the account.
   * @returns {Promise<void>}
   */
  export const deleteAccount = async (apiClient, id) => {
    await apiClient.delete(`${API_ENDPOINTS.ACCOUNTS}/${id}`);
  };
  