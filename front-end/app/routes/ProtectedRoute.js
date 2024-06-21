import React from 'react';
import PropTypes from 'prop-types';
import { Route, Redirect } from 'react-router-dom';
import { useRolesAndPermissions } from '../hooks/useRolesAndPermissions';
import { useAuth } from '../hooks/useAuth';

const ProtectedRoute = ({ component: Component, requiredPermissions, ...rest }) => {
    const { isAuthenticated } = useAuth();
    const { hasPermission } = useRolesAndPermissions();

    return (
        <Route
            {...rest}
            render={(props) =>
                isAuthenticated ? (
                    <Component {...props} />
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
