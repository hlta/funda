import React from 'react';
import { hot } from 'react-hot-loader';
import { BrowserRouter as Router } from 'react-router-dom';
import AppLayout from './../../layout/default';
import AppWithAuth from './AppWithAuth'; 
import { AuthProvider } from '../../contexts/AuthContext';

const basePath = process.env.BASE_PATH || '/';

const AppClient = () => {
    return (
        <AuthProvider>
            <Router basename={basePath}>
                <AppLayout>
                    <AppWithAuth />
                </AppLayout>
            </Router>
        </AuthProvider>
    );
}

export default hot(module)(AppClient);
