// DropdownProfile.js
import React from 'react';
import PropTypes from "prop-types";
import { Link } from "react-router-dom";
import { useAuth } from './../../../hooks/useAuth';
import { PROFILE_ROUTE, SETTINGS_ROUTE, BILLINGS_ROUTE } from "./../../../constants/routes";
import { DropdownMenu, DropdownItem } from "./../../../components";

const DropdownProfile = (props) => {
  const { performLogout, user } = useAuth();

  const handleLogout = async () => {
    await performLogout();
  };
  const fullName = user ? `${user.firstName} ${user.lastName}` : 'User';

  return (
    <React.Fragment>
      <DropdownMenu right={props.right}>
        <DropdownItem header>
          {fullName}
        </DropdownItem>
        <DropdownItem divider />
        <DropdownItem tag={Link} to={PROFILE_ROUTE}>
          My Profile
        </DropdownItem>
        <DropdownItem tag={Link} to={SETTINGS_ROUTE}>
          Settings
        </DropdownItem>
        <DropdownItem tag={Link} to={BILLINGS_ROUTE}>
          Billings
        </DropdownItem>
        <DropdownItem divider />
        <DropdownItem onClick={handleLogout}>
          <i className="fa fa-fw fa-sign-out mr-2"></i>
          Sign Out
        </DropdownItem>
      </DropdownMenu>
    </React.Fragment>
  );
};

DropdownProfile.propTypes = {
  position: PropTypes.string,
  right: PropTypes.bool,
};

DropdownProfile.defaultProps = {
  position: "",
};

export { DropdownProfile };
