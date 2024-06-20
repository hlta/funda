import { useContext, useEffect, useState } from 'react';
import { AuthContext } from '../contexts/AuthContext';

export const useOrganizations = () => {
    const { organizations, selectedOrg, switchOrganization } = useContext(AuthContext);
    const [orgs, setOrgs] = useState(organizations);
    const [selected, setSelected] = useState(selectedOrg);

    useEffect(() => {
        setOrgs(organizations);
    }, [organizations]);

    useEffect(() => {
        setSelected(selectedOrg);
    }, [selectedOrg]);

    const switchOrg = async (orgId) => {
        await switchOrganization(orgId);
    };

    return {
        orgs,
        selected,
        switchOrg,
    };
};
