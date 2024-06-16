import React, { createContext, useReducer, useEffect } from 'react';
import PropTypes from 'prop-types'; 
import * as authService from '../services/authService';

const initialState = {
    isAuthenticated: false,
    user: null,
};

const authReducer = (state, action) => {
    switch (action.type) {
        case 'LOGIN':
            return {
                ...state,
                isAuthenticated: true,
                user: action.payload,
            };
        case 'LOGOUT':
            return {
                ...state,
                isAuthenticated: false,
                user: null,
            };
        default:
            return state;
    }
};

export const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
    const [state, dispatch] = useReducer(authReducer, initialState);

    const checkAuth = async () => {
        const { isAuthenticated, user } = await authService.checkAuth();
        if (isAuthenticated) {
            dispatch({ type: 'LOGIN', payload: user });
        } else {
            dispatch({ type: 'LOGOUT' });
        }
    };

    useEffect(() => {
        checkAuth();
    }, []);

    const login = async (credentials) => {
        const user = await authService.login(credentials);
        dispatch({ type: 'LOGIN', payload: user });
    };

    const logout = async () => {
        await authService.logout();
        dispatch({ type: 'LOGOUT' });
    };

    return (
        <AuthContext.Provider value={{ ...state, login, logout }}>
            {children}
        </AuthContext.Provider>
    );
};

AuthProvider.propTypes = {
    children: PropTypes.node.isRequired,
};
