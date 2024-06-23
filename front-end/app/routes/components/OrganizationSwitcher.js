import React, { useState } from 'react';
import { UncontrolledButtonDropdown, DropdownToggle, DropdownMenu, DropdownItem } from 'reactstrap';
import { useOrganizations } from '../../../hooks/useOrganizations';
import classNames from 'classnames';
import PropTypes from 'prop-types';

const OrganizationSwitcher = ({ down, className, sidebar }) => {
  const { orgs, selected, switchOrg } = useOrganizations();
  const [dropdownOpen, setDropdownOpen] = useState(false);

  const toggle = () => setDropdownOpen(prevState => !prevState);

  const handleOrgSwitch = async (org) => {
    await switchOrg(org.id);
  };

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
        {selected ? selected.name : 'Select Organization'}
        <i className={`fa ${down ? "fa-angle-down" : "fa-angle-up"} ml-2`}></i>
      </DropdownToggle>
      <DropdownMenu>
        {orgs.map(org => (
          <DropdownItem key={org.id} onClick={() => handleOrgSwitch(org)}>
            {org.name}
            {selected && selected.id === org.id && (
              <i className="fa fa-fw fa-check text-success ml-auto align-self-center pl-3" />
            )}
          </DropdownItem>
        ))}
        <DropdownItem divider />
        <DropdownItem onClick={() => console.log('Add a new organization')}>
          <i className="fa fa-fw fa-plus text-primary mr-2" />
          Add a new organization
        </DropdownItem>
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
