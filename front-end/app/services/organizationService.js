export const createOrganization = async (apiClient, data) => {
    const response = await apiClient.post('/organizations', data);
    return response.data.data;
};

export const getOrganization = async (apiClient, id) => {
    const response = await apiClient.get(`/organizations/${id}`);
    return response.data.data;
};

export const updateOrganization = async (apiClient, id, data) => {
    const response = await apiClient.put(`/organizations/${id}`, data);
    return response.data.data;
};
