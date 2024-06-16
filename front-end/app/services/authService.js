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
    return data.user; 
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
        if (response.ok) {
            const data = await response.json();
            return { isAuthenticated: true, user: data.user }; 
        } else {
            return { isAuthenticated: false, user: null };
        }
    } catch (error) {
        return { isAuthenticated: false, user: null };
    }
};
