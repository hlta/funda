import { useContext } from 'react';
import { AuthContext } from '../contexts/AuthContext';

export const useOrganizations = () => {
    const { organizations, selectedOrg, switchOrganization, addOrganization } = useContext(AuthContext);

    const switchOrg = async (orgId) => {
        await switchOrganization(orgId);
    };

    const addOrg = async (data) => {
        await addOrganization(data);
    };

    return {
        orgs: organizations,
        selected: selectedOrg,
        switchOrg,
        addOrg,
    };
};
