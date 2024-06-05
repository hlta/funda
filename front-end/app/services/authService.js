const API_URL = 'http://localhost:8080'; // Base URL for your API

export const login = async (credentials) => {
    const response = await fetch(`${API_URL}/login`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(credentials),
    });
    if (!response.ok) {
        throw new Error('Failed to login');
    }
    const data = await response.json();
    return data.token;
};

export const register = async (userData) => {
    const response = await fetch(`${API_URL}/signup`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(userData),
    });
    if (!response.ok) {
        throw new Error('Failed to register');
    }
    const data = await response.json();
    return data.token;
};
