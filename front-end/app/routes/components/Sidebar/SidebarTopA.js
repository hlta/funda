import React from "react";
import { Link } from "react-router-dom";
import { Sidebar, UncontrolledButtonDropdown, DropdownToggle, DropdownMenu, DropdownItem } from "./../../../components";
import { useAuth } from "../../../hooks/useAuth";
import Avatar from 'react-avatar';

const SidebarTopA = () => {
  const { performLogout, user } = useAuth();

  const handleLogout = async () => {
    await performLogout();
  };

  const fullName = user ? `${user.firstName} ${user.lastName}` : 'User';

  return (
    <React.Fragment>
      {/* START: Sidebar Default */}
      <Sidebar.HideSlim>
        <Sidebar.Section className="pt-0">
          <Link to="/" className="d-block">
            <Sidebar.HideSlim>
              <Avatar
                name={fullName}
                size="50"
                round={true}
                color="#007bff"
                fgColor="#ffffff"
              />
            </Sidebar.HideSlim>
          </Link>

          <UncontrolledButtonDropdown>
            <DropdownToggle
              color="link"
              className="pl-0 pb-0 btn-profile sidebar__link"
            >
              {fullName}
              <i className="fa fa-angle-down ml-2"></i>
            </DropdownToggle>
            <DropdownMenu persist>
              <DropdownItem header>
                {fullName}
              </DropdownItem>
              <DropdownItem divider />
              <DropdownItem tag={Link} to="/apps/profile-details">
                My Profile
              </DropdownItem>
              <DropdownItem tag={Link} to="/apps/settings-edit">
                Settings
              </DropdownItem>
              <DropdownItem tag={Link} to="/apps/billing-edit">
                Billings
              </DropdownItem>
              <DropdownItem divider />
              <DropdownItem onClick={handleLogout}>
                <i className="fa fa-fw fa-sign-out mr-2"></i>
                Sign Out
              </DropdownItem>
            </DropdownMenu>
          </UncontrolledButtonDropdown>
        </Sidebar.Section>
      </Sidebar.HideSlim>
      {/* END: Sidebar Default */}

      {/* START: Sidebar Slim */}
      <Sidebar.ShowSlim>
        <Sidebar.Section>
          <Avatar
            name={fullName}
            size="30"
            round={true}
            color="#007bff"
            fgColor="#ffffff"
          />
        </Sidebar.Section>
      </Sidebar.ShowSlim>
      {/* END: Sidebar Slim */}
    </React.Fragment>
  );
};

export { SidebarTopA };
