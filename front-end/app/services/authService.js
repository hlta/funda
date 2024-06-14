const API_URL = 'http://localhost:8080';

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
    return handleError(response);
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
