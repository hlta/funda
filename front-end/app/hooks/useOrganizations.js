import { useContext, useCallback } from 'react';
import { AuthContext } from '../contexts/AuthContext';
import * as organizationService from '../services/organizationService';

export const useOrganizations = () => {
    const { organizations, selectedOrg, switchOrganization, addOrganization } = useContext(AuthContext);

    const switchOrg = useCallback(async (orgId) => {
        await switchOrganization(orgId);
    }, [switchOrganization]);

    const addOrg = useCallback(async (data) => {
        const newOrg = await organizationService.createOrganization(data);
        await addOrganization(newOrg);
    }, [addOrganization]);

    const getOrg = useCallback(async (id) => {
        return await organizationService.getOrganization(id);
    }, []);

    const updateOrg = useCallback(async (id, data) => {
        return await organizationService.updateOrganization(id, data);
    }, []);

    return {
        orgs: organizations,
        selected: selectedOrg,
        switchOrg,
        addOrg,
        getOrg,
        updateOrg,
    };
};
