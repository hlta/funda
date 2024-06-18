import React from 'react';
import { Link } from 'react-router-dom';
import PropTypes from 'prop-types';
import { useAuth } from './../../hooks/useAuth';
import {
    NavItem,
    NavLink
} from './../../components';

const NavbarUser = (props) => {
    const { performLogout } = useAuth();

    const handleLogout = async (e) => {
        e.preventDefault();
        await performLogout();
    };

    return (
        <NavItem { ...props }>
            <NavLink tag={ Link } to="#" onClick={handleLogout}>
                <i className="fa fa-power-off"></i>
            </NavLink>
        </NavItem>
    );
};

NavbarUser.propTypes = {
    className: PropTypes.string,
    style: PropTypes.object
};

export { NavbarUser };

