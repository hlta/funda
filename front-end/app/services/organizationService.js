// Define API endpoints as constants
const API_ENDPOINTS = {
    ORGANIZATIONS: '/api/organizations',
  };
  
  /**
   * Creates a new organization.
   * @param {Object} apiClient - The API client instance.
   * @param {Object} data - The data for the new organization.
   * @returns {Promise<Object>} The created organization data.
   */
  export const createOrganization = async (apiClient, data) => {
    const response = await apiClient.post(API_ENDPOINTS.ORGANIZATIONS, data);
    return response.data.data;
  };
  
  /**
   * Retrieves an organization by its ID.
   * @param {Object} apiClient - The API client instance.
   * @param {number} id - The ID of the organization.
   * @returns {Promise<Object>} The organization data.
   */
  export const getOrganization = async (apiClient, id) => {
    const response = await apiClient.get(`${API_ENDPOINTS.ORGANIZATIONS}/${id}`);
    return response.data.data;
  };
  
  /**
   * Updates an existing organization by its ID.
   * @param {Object} apiClient - The API client instance.
   * @param {number} id - The ID of the organization.
   * @param {Object} data - The updated data for the organization.
   * @returns {Promise<Object>} The updated organization data.
   */
  export const updateOrganization = async (apiClient, id, data) => {
    const response = await apiClient.put(`${API_ENDPOINTS.ORGANIZATIONS}/${id}`, data);
    return response.data.data;
  };