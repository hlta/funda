import React, { useState } from 'react';
import {
    UncontrolledButtonDropdown,
    DropdownToggle,
    DropdownMenu,
    DropdownItem,
    Modal,
    ModalHeader,
    ModalBody,
    Alert
} from 'reactstrap';
import PropTypes from 'prop-types';
import { useOrganizations } from '../../../hooks/useOrganizations';
import AddOrganizationForm from './AddOrganizationForm';

const OrganizationItem = ({ org, selected, onSelect }) => (
    <DropdownItem
        key={org.id}
        onClick={() => selected !== org.id && onSelect(org)}
    >
        {org.name}
        {selected === org.id && (
            <i className="fa fa-fw fa-check text-success ml-auto align-self-center pl-3" />
        )}
    </DropdownItem>
);

OrganizationItem.propTypes = {
    org: PropTypes.object.isRequired,
    selected: PropTypes.number,
    onSelect: PropTypes.func.isRequired,
};

const AddNewOrganizationItem = ({ onAdd }) => (
    <DropdownItem key="add-org" onClick={onAdd}>
        <i className="fa fa-fw fa-plus text-primary mr-2" />
        Add a new organization
    </DropdownItem>
);

AddNewOrganizationItem.propTypes = {
    onAdd: PropTypes.func.isRequired,
};

const OrganizationSwitcher = () => {
    const { orgs, selected, switchOrg, addOrg } = useOrganizations();
    const [modalOpen, setModalOpen] = useState(false);
    const [error, setError] = useState(null);

    const toggleModal = () => {
        setError(null); // Clear error message when toggling modal
        setModalOpen(prevState => !prevState);
    };

    const handleOrgSwitch = async (org) => {
        try {
            await switchOrg(org.id);
            setError(null); // Clear error message on success
        } catch (err) {
            setError('Failed to switch organization.');
        }
    };

    const handleAddOrganization = async (data) => {
        try {
            const response = await addOrg(data);
            toggleModal();
            await switchOrg(response.id);
            setError(null); // Clear error message on success
        } catch (err) {
            setError('Failed to add new organization.');
        }
    };

    const selectedOrgName = orgs.find(org => org.id === selected)?.name || 'Default Organization';

    return (
        <>
            <UncontrolledButtonDropdown>
                <DropdownToggle color="link" className="pl-0 pb-0 btn-profile sidebar__link">
                    {selectedOrgName}
                    <i className="fa fa-angle-down ml-2"></i>
                </DropdownToggle>
                <DropdownMenu persist>
                    {orgs.map(org => (
                        <OrganizationItem
                            key={org.id}
                            org={org}
                            selected={selected}
                            onSelect={handleOrgSwitch}
                        />
                    ))}
                    <DropdownItem key="divider" divider />
                    <AddNewOrganizationItem onAdd={toggleModal} />
                </DropdownMenu>
            </UncontrolledButtonDropdown>

            {error && <Alert color="danger">{error}</Alert>} {/* Display error message */}

            <Modal isOpen={modalOpen} toggle={toggleModal}>
                <ModalHeader toggle={toggleModal}>Add New Organization</ModalHeader>
                <ModalBody>
                    {error && <Alert color="danger">{error}</Alert>} {/* Display error message inside modal */}
                    <AddOrganizationForm onSubmit={handleAddOrganization} />
                </ModalBody>
            </Modal>
        </>
    );
};

OrganizationSwitcher.propTypes = {
    down: PropTypes.bool,
    className: PropTypes.string,
    sidebar: PropTypes.bool,
};

export default OrganizationSwitcher;
