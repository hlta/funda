import React, { useState } from 'react';
import { UncontrolledButtonDropdown, DropdownToggle, DropdownMenu, DropdownItem, Modal, ModalHeader, ModalBody } from 'reactstrap';
import classNames from 'classnames';
import PropTypes from 'prop-types';
import { useOrganizations } from '../../../hooks/useOrganizations';
import AddOrganizationForm from './AddOrganizationForm';

const OrganizationItem = ({ org, selected, onSelect }) => (
    <DropdownItem key={org.id} onClick={() => onSelect(org)}>
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

const OrganizationSwitcher = ({ down, className, sidebar }) => {
    const { orgs, selected, switchOrg, addOrg } = useOrganizations();
    const [dropdownOpen, setDropdownOpen] = useState(false);
    const [modalOpen, setModalOpen] = useState(false);

    const toggleDropdown = () => setDropdownOpen(prevState => !prevState);
    const toggleModal = () => setModalOpen(prevState => !prevState);

    const handleOrgSwitch = async (org) => {
        await switchOrg(org.id);
    };

    const handleAddOrganization = async (data) => {
        await addOrg(data);
        toggleModal();
    };

    const selectedOrg = orgs.find(org => org.id === selected);

    return (
        <>
            <UncontrolledButtonDropdown isOpen={dropdownOpen} toggle={toggleDropdown} direction={down ? 'down' : 'up'} className={className}>
                <DropdownToggle
                    disabled={!orgs.length}
                    tag="a"
                    href="#"
                    className={classNames(
                        'btn-switch-version',
                        {
                            'sidebar__link': sidebar,
                        }
                    )}
                >
                    {selectedOrg ? selectedOrg.name : 'Select Organization'}
                    <i className={`fa ${down ? "fa-angle-down" : "fa-angle-up"} ml-2`}></i>
                </DropdownToggle>
                <DropdownMenu>
                    {orgs.map(org => (
                        <OrganizationItem key={org.id} org={org} selected={selected} onSelect={handleOrgSwitch} />
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
    sidebar: PropTypes.bool
};

export default OrganizationSwitcher;
