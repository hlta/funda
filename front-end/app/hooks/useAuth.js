import { useContext } from 'react';
import { AuthContext } from '../contexts/AuthContext';
import * as authService from '../services/authService';

export const useAuth = () => {
    const { isAuthenticated, token, login, logout } = useContext(AuthContext);

    const performLogin = async (credentials) => {
        try {
            const token = await authService.login(credentials);
            login(token);
        } catch (error) {
            console.error('Login failed:', error);
        }
    };

    const performLogout = () => {
        logout();
    };

    const performRegister = async (userData) => {
        try {
            const token = await authService.register(userData);
            login(token);
        } catch (error) {
            console.error('Registration failed:', error);
        }
    };

    return {
        isAuthenticated,
        token,
        performLogin,
        performLogout,
        performRegister,
    };
};
