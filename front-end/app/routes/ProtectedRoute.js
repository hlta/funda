import React from 'react';
import PropTypes from 'prop-types';
import { Route, Redirect } from 'react-router-dom';
import { useRolesAndPermissions } from '../hooks/useRolesAndPermissions';
import { useAuth } from '../hooks/useAuth';
import NoPermission from './components/NoPermission'; 

const ProtectedRoute = ({ component: Component, requiredPermissions, ...rest }) => {
    const { isAuthenticated, loading } = useAuth();
    const { hasPermission } = useRolesAndPermissions();

    if (loading) {
        return <div>Loading...</div>;
    }

    return (
        <Route
            {...rest}
            render={(props) =>
                isAuthenticated ? (
                    hasPermission(requiredPermissions) ? (
                        <Component {...props} />
                    ) : (
                        <NoPermission />
                    )
                ) : (
                    <Redirect to="/login" />
                )
            }
        />
    );
};

ProtectedRoute.propTypes = {
    component: PropTypes.elementType.isRequired,
    requiredPermissions: PropTypes.arrayOf(PropTypes.string).isRequired,
};

export default ProtectedRoute;
