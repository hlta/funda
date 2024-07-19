import config from '../config';

const API_URL = config.apiUrl;

const handleError = async (response) => {
    const data = await response.json();
    if (!response.ok) {
        const error = new Error(data.message || 'Request failed');
        error.response = data;
        throw error;
    }
    return data;
};

export const login = async (credentials) => {
    const response = await fetch(`${API_URL}/login`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(credentials),
        credentials: 'include', 
    });
    const data = await handleError(response);
    return data.data; 
};

export const register = async (userData) => {
    const response = await fetch(`${API_URL}/signup`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(userData),
    });
    return handleError(response);
};

export const logout = async () => {
    const response = await fetch(`${API_URL}/logout`, {
        method: 'POST',
        credentials: 'include',
    });
    return handleError(response);
};

export const checkAuth = async () => {
    try {
        const response = await fetch(`${API_URL}/auth/check`, {
            credentials: 'include',
        });
        const data = await response.json();
        if (response.ok) {
            return data.data;
        } else {
            return { isAuthenticated: false, user: null };
        }
    } catch (error) {
        return { isAuthenticated: false, user: null };
    }
};

export const getUserOrganizations = async () => {
    const response = await fetch(`${API_URL}/auth/orgs`, {
        credentials: 'include',
    });
    const data = await handleError(response);
    return data.data; 
};

export const switchOrganization = async (orgId) => {
    const response = await fetch(`${API_URL}/auth/switch-org`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ org_id: orgId }),
        credentials: 'include',
    });
    const data = await handleError(response);
    return {
        token: data.data.token,
        roles: data.data.roles, 
        permissions: data.data.permissions, 
    };
};
