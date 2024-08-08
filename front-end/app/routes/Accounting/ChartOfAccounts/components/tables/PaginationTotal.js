import React from 'react';
import PropTypes from 'prop-types';

export const PaginationTotal = ({ from, to, size }) => (
    <span className="small ml-2">
        Showing { from } to { to } of { size } Results
    </span>
);
PaginationTotal.propTypes = {
    from: PropTypes.number,
    to: PropTypes.number,
    size: PropTypes.number,
};