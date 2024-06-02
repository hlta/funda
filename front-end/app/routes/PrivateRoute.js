import React from 'react';
import { Route, Redirect } from 'react-router-dom';
import { useAuth } from '../hooks/useAuth';  // Adjusted to use the hook from the hooks directory

const PrivateRoute = ({ component: Component, ...rest }) => {
    const { isAuthenticated } = useAuth(); // This will now correctly use the custom hook

    return (
        <Route
            {...rest}
            render={props =>
                isAuthenticated ? (
                    <Component {...props} />
                ) : (
                    <Redirect to="/login" />
                )
            }
        />
    );
};

export default PrivateRoute;


