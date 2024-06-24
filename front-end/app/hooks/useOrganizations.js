import { useContext } from 'react';
import { AuthContext } from '../contexts/AuthContext';

export const useOrganizations = () => {
    const { organizations, selectedOrg, switchOrganization } = useContext(AuthContext);

    const switchOrg = async (orgId) => {
        await switchOrganization(orgId);
    };

    return {
        orgs: organizations,
        selected: selectedOrg,
        switchOrg,
    };
};
