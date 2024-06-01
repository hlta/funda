import React from 'react';
import PropTypes from 'prop-types';
import { Link } from 'react-router-dom';

import { LogoThemed } from './../LogoThemed/LogoThemed';

const HeaderAuth = (props) => (
    <div className="mb-4">
        <div className="mb-4 text-center">
            <Link to="/" className="d-inline-block">
                {
                    props.icon ? (
                        <i className={ `fa fa-${ props.icon } fa-3x ${ props.iconClassName }` }></i>
                    ) : (
                        <LogoThemed checkBackground height="30" />
                    )
                }
            </Link>
        </div>
       
    </div>
)
HeaderAuth.propTypes = {
    icon: PropTypes.node,
    iconClassName: PropTypes.node,
};

export { HeaderAuth };
