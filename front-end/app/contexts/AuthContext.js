import React, { createContext, useReducer, useEffect, useCallback } from 'react';
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

const actionTypes = {
    LOGIN: 'LOGIN',
    LOGOUT: 'LOGOUT',
    SET_LOADING: 'SET_LOADING',
    SET_ORGANIZATIONS: 'SET_ORGANIZATIONS',
    SWITCH_ORGANIZATION: 'SWITCH_ORGANIZATION',
};

const authReducer = (state, action) => {
    switch (action.type) {
        case actionTypes.LOGIN:
            return {
                ...state,
                isAuthenticated: !!action.payload.user,
                selectedOrg: action.payload.user.selectedOrg,
                user: action.payload.user || null,
                roles: action.payload.roles || [],
                permissions: action.payload.permissions || [],
                loading: false,
            };
        case actionTypes.LOGOUT:
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
        case actionTypes.SET_LOADING:
            return {
                ...state,
                loading: action.payload,
            };
        case actionTypes.SET_ORGANIZATIONS:
            return {
                ...state,
                organizations: action.payload || [],
            };
        case actionTypes.SWITCH_ORGANIZATION:
            return {
                ...state,
                selectedOrg: action.payload.orgId,
                roles: action.payload.roles || [],
                permissions: action.payload.permissions || [],
                loading: false,
            };
        default:
            return state;
    }
};

export const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
    const [state, dispatch] = useReducer(authReducer, initialState);

    const setLoading = useCallback((loading) => {
        dispatch({ type: actionTypes.SET_LOADING, payload: loading });
    }, []);

    const checkAuth = useCallback(async () => {
        setLoading(true);
        try {
            const { isAuthenticated, user, roles, permissions } = await authService.checkAuth();
            if (isAuthenticated) {
                dispatch({ type: actionTypes.LOGIN, payload: { user, roles, permissions } });
                const organizations = await authService.getUserOrganizations();
                dispatch({ type: actionTypes.SET_ORGANIZATIONS, payload: organizations });
            } else {
                dispatch({ type: actionTypes.LOGOUT });
            }
        } finally {
            setLoading(false);
        }
    }, [setLoading]);

    useEffect(() => {
        let isMounted = true;
        if (isMounted) {
            checkAuth();
        }
        return () => {
            isMounted = false;
        };
    }, [checkAuth]);

    const login = async (credentials) => {
        setLoading(true);
        try {
            const user = await authService.login(credentials);
            if (user) {
                dispatch({ type: actionTypes.LOGIN, payload: { user, roles: user.roles, permissions: user.permissions } });
                const organizations = await authService.getUserOrganizations();
                dispatch({ type: actionTypes.SET_ORGANIZATIONS, payload: organizations });
            }
        } finally {
            setLoading(false);
        }
    };

    const logout = async () => {
        setLoading(true);
        await authService.logout();
        dispatch({ type: actionTypes.LOGOUT });
        setLoading(false);
    };

    const switchOrganization = async (orgId) => {
        setLoading(true);
        try {
            const { roles, permissions } = await authService.switchOrganization(orgId);
            dispatch({ type: actionTypes.SWITCH_ORGANIZATION, payload: { orgId, roles, permissions } });
            const { user } = await authService.checkAuth(); 
            dispatch({ type: actionTypes.LOGIN, payload: { user, roles, permissions } });
        } catch (error) {
            console.error('Switch organization error:', error);
        } finally {
            setLoading(false);
        }
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
