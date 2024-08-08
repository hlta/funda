import React from 'react';
import PropTypes from 'prop-types';
import {
    Button
} from './../../../../../components';

export const ExportButton = ({children, onExport, ...props}) => {
    return (
        <Button
            { ...props }
            onClick={() => { onExport() }}
        >
            { children }
        </Button>
    );
}

ExportButton.propTypes = {
    size: PropTypes.string,
    outline: PropTypes.bool,
    onExport: PropTypes.func,
    children: PropTypes.node
}

ExportButton.defaultProps = {
    size: 'sm',
    outline: true
}
