import React, { createContext, useReducer } from 'react';
import { useHistory } from 'react-router-dom';

export const AuthContext = createContext();

const initialState = {
    isAuthenticated: false,
    token: null,
};

const authReducer = (state, action) => {
    switch (action.type) {
        case 'LOGIN':
            return {
                ...state,
                isAuthenticated: true,
                token: action.payload.token,
            };
        case 'LOGOUT':
            return {
                ...state,
                isAuthenticated: false,
                token: null,
            };
        default:
            return state;
    }
};

export const AuthProvider = ({ children }) => {
    const [state, dispatch] = useReducer(authReducer, initialState);
    const history = useHistory();

    const login = (token) => {
        dispatch({ type: 'LOGIN', payload: { token } });
        history.push('/dashboard'); // Redirect to dashboard after login
    };

    const logout = () => {
        dispatch({ type: 'LOGOUT' });
        history.push('/login'); // Redirect to login page after logout
    };

    return (
        <AuthContext.Provider value={{ ...state, login, logout }}>
            {children}
        </AuthContext.Provider>
    );
};
