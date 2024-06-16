// src/index.js
import '@babel/polyfill';
import React from 'react';
import { render } from 'react-dom';
import App from './components/App';
import { AuthProvider } from './contexts/AuthContext';

render(
    <AuthProvider>
        <App />
    </AuthProvider>,
    document.querySelector('#root')
);
