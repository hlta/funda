import { useContext } from 'react';
import { useHistory } from 'react-router-dom';
import { AuthContext } from '../contexts/AuthContext';
import * as authService from '../services/authService';
import { LOGIN_ROUTE } from '../constants/routes';

export const useAuth = () => {
    const { loading, isAuthenticated, user, login, logout } = useContext(AuthContext);
    const history = useHistory();

    const performLogin = async (credentials) => {
        await login(credentials);
    };

    const performLogout = async () => {
        await authService.logout();
        await logout();
        history.push(LOGIN_ROUTE);
    };

    const performRegister = async (userData) => {
        await authService.register(userData);
        await performLogin({ email: userData.email, password: userData.password });
    
    };


    return {
        loading,
        isAuthenticated,
        user,
        performLogin,
        performLogout,
        performRegister,
    };
};
