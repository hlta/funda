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
    token: null,
};

const actionTypes = {
    LOGIN: 'LOGIN',
    LOGOUT: 'LOGOUT',
    SET_LOADING: 'SET_LOADING',
    SET_ORGANIZATIONS: 'SET_ORGANIZATIONS',
    SWITCH_ORGANIZATION: 'SWITCH_ORGANIZATION',
    ADD_ORGANIZATION: 'ADD_ORGANIZATION',
};

const authReducer = (state, action) => {
    switch (action.type) {
        case actionTypes.LOGIN:
            return {
                ...state,
                isAuthenticated: !!action.payload.user,
                selectedOrg: action.payload.selectedOrg,
                user: action.payload.user || null,
                roles: action.payload.roles || [],
                token: action.payload.token || null,
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
                selectedOrg: action.payload.selectedOrg,
                roles: action.payload.roles || [],
                token: action.payload.token || null,
                permissions: action.payload.permissions || [],
                loading: false,
            };
        case actionTypes.ADD_ORGANIZATION:
            return {
                ...state,
                organizations: [...state.organizations, action.payload.org],
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
            const { user, roles, permissions, token, selectedOrg } = await authService.checkAuth();
            if (token) {
                dispatch({ type: actionTypes.LOGIN, payload: { user, roles, permissions, token, selectedOrg } });
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
        checkAuth();
    }, [checkAuth]);

    const login = async (credentials) => {
        setLoading(true);
        try {
            const authData = await authService.login(credentials);
            if (authData) {
                dispatch({
                    type: actionTypes.LOGIN,
                    payload: {
                        token: authData.token,
                        user: authData.user,
                        roles: authData.roles,
                        selectedOrg: authData.SelectedOrg,
                        permissions: authData.permissions,
                    },
                });
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
            const { roles, permissions, token } = await authService.switchOrganization(orgId);
            dispatch({
                type: actionTypes.SWITCH_ORGANIZATION,
                payload: { selectedOrg:orgId, roles, permissions, token },
            });
        } catch (error) {
            console.error('Switch organization error:', error);
        } finally {
            setLoading(false);
        }
    };

    const addOrganization = (data) => {
        dispatch({ type: actionTypes.ADD_ORGANIZATION, payload: { org: data } });
    };

    return (
        <AuthContext.Provider
            value={{ ...state, login, logout, switchOrganization, addOrganization }}
        >
            {children}
        </AuthContext.Provider>
    );
};

AuthProvider.propTypes = {
    children: PropTypes.node.isRequired,
};
