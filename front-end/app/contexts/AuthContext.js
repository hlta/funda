import React, { createContext, useReducer, useEffect } from 'react';
import { useHistory } from 'react-router-dom';
import * as authService from '../services/authService';

const initialState = {
    isAuthenticated: false,
};

const authReducer = (state, action) => {
    switch (action.type) {
        case 'LOGIN':
            return {
                ...state,
                isAuthenticated: true,
            };
        case 'LOGOUT':
            return {
                ...state,
                isAuthenticated: false,
            };
        default:
            return state;
    }
};

export const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
    const [state, dispatch] = useReducer(authReducer, initialState);
    const history = useHistory();

    const checkAuth = async () => {
        const isAuthenticated = await authService.checkAuth();
        if (isAuthenticated) {
            dispatch({ type: 'LOGIN' });
        } else {
            dispatch({ type: 'LOGOUT' });
        }
    };

    useEffect(() => {
        checkAuth();
    }, []);

    const login = () => {
        dispatch({ type: 'LOGIN' });
    };

    const logout = () => {
        dispatch({ type: 'LOGOUT' });
    };

    return (
        <AuthContext.Provider value={{ ...state, login, logout }}>
            {children}
        </AuthContext.Provider>
    );
};
