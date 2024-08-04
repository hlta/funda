import { useContext } from 'react';
import { AuthContext } from '../contexts/AuthContext';

export const useRolesAndPermissions = () => {
    const { roles, permissions } = useContext(AuthContext);

    const hasPermission = (requiredPermissions) => {
       return requiredPermissions != null;
    };

    return {
        roles,
        permissions,
        hasPermission,
    };
};
