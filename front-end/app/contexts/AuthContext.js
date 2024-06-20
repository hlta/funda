import React, { createContext, useReducer, useEffect } from 'react';
import PropTypes from 'prop-types';
import * as authService from '../services/authService';

const initialState = {
    isAuthenticated: false,
    user: null,
    loading: true,
    organizations: [],
    selectedOrg: null,
    roles: [],
    permissions: [],
};

const authReducer = (state, action) => {
    switch (action.type) {
        case 'LOGIN':
            return {
                ...state,
                isAuthenticated: true,
                user: action.payload.user,
                roles: action.payload.roles,
                permissions: action.payload.permissions,
                loading: false,
            };
        case 'LOGOUT':
            return {
                ...state,
                isAuthenticated: false,
                user: null,
                organizations: [],
                selectedOrg: null,
                roles: [],
                permissions: [],
                loading: false,
            };
        case 'SET_LOADING':
            return {
                ...state,
                loading: action.payload,
            };
        case 'SET_ORGANIZATIONS':
            return {
                ...state,
                organizations: action.payload,
            };
        case 'SWITCH_ORGANIZATION':
            return {
                ...state,
                selectedOrg: action.payload.orgId,
                roles: action.payload.roles,
                permissions: action.payload.permissions,
            };
        default:
            return state;
    }
};

export const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
    const [state, dispatch] = useReducer(authReducer, initialState);

    const checkAuth = async () => {
        dispatch({ type: 'SET_LOADING', payload: true });
        const { isAuthenticated, user, roles, permissions } = await authService.checkAuth();
        if (isAuthenticated) {
            dispatch({ type: 'LOGIN', payload: { user, roles, permissions } });
            const organizations = await authService.getUserOrganizations();
            dispatch({ type: 'SET_ORGANIZATIONS', payload: organizations });
        } else {
            dispatch({ type: 'LOGOUT' });
        }
        dispatch({ type: 'SET_LOADING', payload: false });
    };

    useEffect(() => {
        checkAuth();
    }, []);

    const login = async (user) => {
        dispatch({ type: 'LOGIN', payload: user });
        const organizations = await authService.getUserOrganizations();
        dispatch({ type: 'SET_ORGANIZATIONS', payload: organizations });
    };

    const logout = async () => {
        dispatch({ type: 'LOGOUT' });
    };

    const switchOrganization = async (orgId) => {
        const { token, roles, permissions } = await authService.switchOrganization(orgId);
        dispatch({ type: 'SWITCH_ORGANIZATION', payload: { orgId, roles, permissions } });
        // Optionally, refresh user data with new roles and permissions based on selected organization
        const { user } = await authService.checkAuth(); // Assuming checkAuth returns user data including roles and permissions
        dispatch({ type: 'LOGIN', payload: { user, roles, permissions } });
    };

    return (
        <AuthContext.Provider value={{ ...state, login, logout, switchOrganization }}>
            {children}
        </AuthContext.Provider>
    );
};

AuthProvider.propTypes = {
    children: PropTypes.node.isRequired,
};
