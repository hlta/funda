import { useContext } from 'react';
import { useHistory } from 'react-router-dom';
import { AuthContext } from '../contexts/AuthContext';
import * as authService from '../services/authService';
import { LOGIN_ROUTE } from '../constants/routes';

export const useAuth = () => {
    const { loading, isAuthenticated, user, organizations, selectedOrg, roles, permissions, login, logout, switchOrganization } = useContext(AuthContext);
    const history = useHistory();

    const performLogin = async (credentials) => {
        const user = await authService.login(credentials);
        login(user);
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

    const switchOrg = async (orgId) => {
        await switchOrganization(orgId);
    };

    return {
        loading,
        isAuthenticated,
        user,
        organizations,
        selectedOrg,
        roles,
        permissions,
        performLogin,
        performLogout,
        performRegister,
        switchOrg,
    };
};
