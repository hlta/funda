import { useContext } from 'react';
import { AuthContext } from '../contexts/AuthContext';
import * as authService from '../services/authService';

export const useAuth = () => {
    const { isAuthenticated, login, logout } = useContext(AuthContext);

    const performLogin = async (credentials) => {
        try {
            await authService.login(credentials);
            login();
        } catch (error) {
            throw error;
        }
    };

    const performLogout = async () => {
        await authService.logout();
        logout();
    };

    const performRegister = async (userData) => {
        try {
            await authService.register(userData);
            await performLogin({ email: userData.email, password: userData.password });
        } catch (error) {
            throw error;
        }
    };

    return {
        isAuthenticated,
        performLogin,
        performLogout,
        performRegister,
    };
};
