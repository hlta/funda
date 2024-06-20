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
        credentials: 'include', // Important for sending cookies
    });
    const data = await handleError(response);
    return data.user; // Adjust based on your backend response
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
        credentials: 'include', // Important for sending cookies
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
            return {
                isAuthenticated: true,
                user: data.data.user, // Adjust based on your backend response
                roles: data.data.roles, // Adjust based on your backend response
                permissions: data.data.permissions, // Adjust based on your backend response
            };
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
    return data.data; // Adjust based on your backend response
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
        token: data.token, // Adjust based on your backend response
        roles: data.roles, // Adjust based on your backend response
        permissions: data.permissions, // Adjust based on your backend response
    };
};
