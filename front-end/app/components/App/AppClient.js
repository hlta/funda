import React from 'react';
import { hot } from 'react-hot-loader';
import { BrowserRouter as Router } from 'react-router-dom';
import AppLayout from './../../layout/default';
import { AuthProvider } from '../../contexts/AuthContext';
import { RoutedContent } from './../../routes';


const basePath = process.env.BASE_PATH || '/';

const AppClient = () => {
    return (
        <AuthProvider>
            <Router basename={basePath}>
                <AppLayout>
                    <RoutedContent />
                </AppLayout>
            </Router>
        </AuthProvider>
    );
}

export default hot(module)(AppClient);
