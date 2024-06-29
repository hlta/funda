import { useContext } from 'react';
import { AuthContext } from '../contexts/AuthContext';

export const useRolesAndPermissions = () => {
    const { roles, permissions } = useContext(AuthContext);

    const hasPermission = (requiredPermissions) => {
        if (!requiredPermissions || requiredPermissions.length === 0) {
            return true;
        }
        return requiredPermissions.every((perm) => permissions.includes(perm));
    };

    return {
        roles,
        permissions,
        hasPermission,
    };
};
