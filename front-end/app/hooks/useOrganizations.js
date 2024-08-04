import { useContext, useCallback } from 'react';
import { AuthContext } from '../contexts/AuthContext';
import { createApiClient } from '../services/apiClient';
import * as organizationService from '../services/organizationService';

export const useOrganizations = () => {
    const { token, organizations, selectedOrg, switchOrganization, addOrganization } = useContext(AuthContext);
    const apiClient = createApiClient(token);

    const switchOrg = useCallback(async (orgId) => {
        await switchOrganization(orgId);
    }, [switchOrganization]);

    const addOrg = useCallback(async (data) => {
        const newOrg = await organizationService.createOrganization(apiClient, data);
        await addOrganization(newOrg);
        return newOrg;
    }, [addOrganization, apiClient]);

    const getOrg = useCallback(async (id) => {
        return await organizationService.getOrganization(apiClient, id);
    }, [apiClient]);

    const updateOrg = useCallback(async (id, data) => {
        return await organizationService.updateOrganization(apiClient, id, data);
    }, [apiClient]);

    return {
        organizations,
        selectedOrg,
        addOrg,
        getOrg,
        switchOrg,
        updateOrg,
    };
};
