const API_URL = 'http://localhost:8080'; // Base URL for your API

export const login = async (credentials) => {
    const response = await fetch(`${API_URL}/login`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(credentials),
    });
    const data = await response.json(); // Always parse JSON since backend sends JSON responses
    if (!response.ok) {
        throw new Error(data.error || 'Login failed. Please check your credentials and try again.');
    }
    return data.token; // Assuming the response will always contain a token when successful
};

export const register = async (userData) => {
    const response = await fetch(`${API_URL}/signup`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(userData),
    });
    const data = await response.json(); // Always parse JSON since backend sends JSON responses
    if (!response.ok) {
        throw new Error(data.error || 'Registration failed. Please try different credentials.');
    }
    return data.message; // Use the message from the successful registration
};
