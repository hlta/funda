import apiClient from "../contexts/apiClient";

export const createOrganization = async (organization) => {
    const response = await apiClient.post('/organizations', organization);
    return response.data;
};

export const getOrganization = async (id) => {
    const response = await apiClient.get(`/organizations/${id}`);
    return response.data;
};

export const updateOrganization = async (id, organization) => {
    const response = await apiClient.put(`/organizations/${id}`, organization);
    return response.data;
};
