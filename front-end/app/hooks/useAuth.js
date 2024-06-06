import { useContext } from 'react';
import { AuthContext } from '../contexts/AuthContext';
import * as authService from '../services/authService';

export const useAuth = () => {
    const { isAuthenticated, token, login, logout } = useContext(AuthContext);

    const performLogin = async (credentials) => {
        const token = await authService.login(credentials);
        login(token);

    };

    const performLogout = () => {
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
        token,
        performLogin,
        performLogout,
        performRegister,
    };
};
