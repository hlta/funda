import React from 'react';
import { RoutedContent } from './../../routes';
import { useAuth } from '../../hooks/useAuth';

const AppWithAuth = () => {
    const { loading } = useAuth();

    if (loading) {
        return <div>Loading...</div>;
    }
    return <RoutedContent />;
}

export default AppWithAuth;
