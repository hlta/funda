import { useContext } from 'react';
import { AuthContext } from '../contexts/AuthContext';
import * as authService from '../services/authService';

export const useAuth = () => {
    const { isAuthenticated, user, login, logout } = useContext(AuthContext);

    const performLogin = async (credentials) => {
        const user = await authService.login(credentials);
        login(user);
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
        user,
        performLogin,
        performLogout,
        performRegister,
    };
};
