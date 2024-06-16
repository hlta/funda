import React, { createContext, useReducer, useEffect } from 'react';
import PropTypes from 'prop-types';
import * as authService from '../services/authService';

const initialState = {
    isAuthenticated: false,
    user: null,
    loading: true,
};

const authReducer = (state, action) => {
    switch (action.type) {
        case 'LOGIN':
            return {
                ...state,
                isAuthenticated: true,
                user: action.payload,
                loading: false,
            };
        case 'LOGOUT':
            return {
                ...state,
                isAuthenticated: false,
                user: null,
                loading: false,
            };
        case 'SET_LOADING':
            return {
                ...state,
                loading: action.payload,
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
        const { isAuthenticated, user } = await authService.checkAuth();
        if (isAuthenticated) {
            dispatch({ type: 'LOGIN', payload: user });
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
    };

    const logout = async () => {
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
