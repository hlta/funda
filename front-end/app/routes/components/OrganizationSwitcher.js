import React, { useState } from 'react';
import { UncontrolledButtonDropdown, DropdownToggle, DropdownMenu, DropdownItem } from 'reactstrap';
import classNames from 'classnames';
import PropTypes from 'prop-types';
import { useOrganizations } from '../../hooks/useOrganizations';

// Component to render individual organization items
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

// Component for the Add New Organization option
const AddNewOrganizationItem = ({ onAdd }) => (
  <DropdownItem key="add-org" onClick={onAdd}>
    <i className="fa fa-fw fa-plus text-primary mr-2" />
    Add a new organization
  </DropdownItem>
);

AddNewOrganizationItem.propTypes = {
  onAdd: PropTypes.func.isRequired,
};

// Main component
const OrganizationSwitcher = ({ down, className, sidebar }) => {
  const { orgs, selected, switchOrg } = useOrganizations();
  const [dropdownOpen, setDropdownOpen] = useState(false);

  const toggle = () => setDropdownOpen(prevState => !prevState);

  const handleOrgSwitch = async (org) => {
    await switchOrg(org.id);
  };

  const selectedOrg = orgs.find(org => org.id === selected);

  return (
    <UncontrolledButtonDropdown isOpen={dropdownOpen} toggle={toggle} direction={down ? 'down' : 'up'} className={className}>
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
        <AddNewOrganizationItem onAdd={() => console.log('Add a new organization')} />
      </DropdownMenu>
    </UncontrolledButtonDropdown>
  );
};

OrganizationSwitcher.propTypes = {
  down: PropTypes.bool,
  className: PropTypes.string,
  sidebar: PropTypes.bool
};

export default OrganizationSwitcher;
