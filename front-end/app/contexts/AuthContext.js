import React, { createContext, useReducer, useEffect } from 'react';
import { useHistory } from 'react-router-dom';
import axios from 'axios';

const API_URL = 'http://localhost:8080';

export const AuthContext = createContext();

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

export const AuthProvider = ({ children }) => {
    const [state, dispatch] = useReducer(authReducer, initialState);
    const history = useHistory();

    const checkAuth = async () => {
        try {
            await axios.get(`${API_URL}/auth/check`, { withCredentials: true });
            dispatch({ type: 'LOGIN' });
        } catch (error) {
            dispatch({ type: 'LOGOUT' });
        }
    };

    useEffect(() => {
        checkAuth();
    }, []);

    const login = async (credentials) => {
        try {
            await axios.post(`${API_URL}/login`, credentials, { withCredentials: true });
            dispatch({ type: 'LOGIN' });
            history.push('/');
        } catch (error) {
            console.error('Login failed:', error);
        }
    };

    const logout = async () => {
        try {
            await axios.post(`${API_URL}/logout`, {}, { withCredentials: true });
            dispatch({ type: 'LOGOUT' });
            history.push('/login');
        } catch (error) {
            console.error('Logout failed:', error);
        }
    };

    return (
        <AuthContext.Provider value={{ ...state, login, logout }}>
            {children}
        </AuthContext.Provider>
    );
};
