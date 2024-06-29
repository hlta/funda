import React, { useState } from 'react';
import {
    UncontrolledButtonDropdown,
    DropdownToggle,
    DropdownMenu,
    DropdownItem,
    Modal,
    ModalHeader,
    ModalBody
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
    selected: PropTypes.string,
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

    const toggleModal = () => setModalOpen(prevState => !prevState);

    const handleOrgSwitch = async (org) => {
        await switchOrg(org.id);
    };

    const handleAddOrganization = async (data) => {
        await addOrg(data);
        toggleModal();
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

            <Modal isOpen={modalOpen} toggle={toggleModal}>
                <ModalHeader toggle={toggleModal}>Add New Organization</ModalHeader>
                <ModalBody>
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
