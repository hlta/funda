import { useContext } from 'react';
import { AuthContext } from '../contexts/AuthContext';
import * as authService from '../services/authService';

export const useAuth = () => {
    const { isAuthenticated, login, logout } = useContext(AuthContext);

    const performLogin = async (credentials) => {

        await authService.login(credentials);
        login();
        
    };

    const performLogout = async () => {
        await authService.logout();
        logout();
    };

    const performRegister = async (userData) => {
       
        await authService.register(userData);
        await performLogin({ email: userData.email, password: userData.password });
       
    };

    return {
        isAuthenticated,
        performLogin,
        performLogout,
        performRegister,
    };
};
