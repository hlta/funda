import React from 'react';
import {
    Route,
    Switch,
    Redirect
} from 'react-router';

// ----------- Pages Imports ---------------
import Financial from './Dashboards/Financial';
import Error404 from './Pages/Error404';
import ForgotPassword from './Pages/ForgotPassword';
import Login from './Pages/Login';
import Register from './Pages/Register';
import ProtectedRoute from './ProtectedRoute';

// ----------- Layout Imports ---------------
import { DefaultNavbar } from './../layout/components/DefaultNavbar';
import { DefaultSidebar } from './../layout/components/DefaultSidebar';

// ----------- Accounting Imports ---------------
import { ChartOfAccounts, AccountDetail, AccountForm } from './Accounting/ChartOfAccounts';


//------ Route Definitions --------
// eslint-disable-next-line no-unused-vars
export const RoutedContent = () => {
    return (
            <Switch>
                <Redirect from="/" to="/dashboards/financial" exact />
                <ProtectedRoute requiredPermissions={['View Reports']} path="/dashboards/financial" exact component={Financial} />
                { /*    Accounting Routes    */ }
                <ProtectedRoute requiredPermissions={['View Reports']} path="/accounting/chart-of-accounts" exact component={ChartOfAccounts} />
                <ProtectedRoute requiredPermissions={['View Reports']} path="/accounting/chart-of-accounts/new" exact component={AccountForm} />
                <ProtectedRoute requiredPermissions={['View Reports']} path="/accounting/chart-of-accounts/:id/edit" exact component={AccountForm} />
                <ProtectedRoute requiredPermissions={['View Reports']} path="/accounting/chart-of-accounts/:id" exact component={AccountDetail} />

                { /*    Auth Routes    */ }
                <Route component={Error404} path="/error-404" />
                <Route component={ForgotPassword} path="/forgot-password" />
                <Route component={Login} path="/login" />
                <Route component={Register} path="/register" />

                { /*    404    */ }
                <Redirect to="/error-404" />
            </Switch>
    );
};

//------ Custom Layout Parts --------
export const RoutedNavbars = () => (
    <Switch>
        { /* Default Navbar: */}
        <Route component={DefaultNavbar} />
    </Switch>  
);

export const RoutedSidebars = () => (
    <Switch>
        <Route component={DefaultSidebar} />
    </Switch>
);
