import React, { forwardRef } from 'react';
import PropTypes from 'prop-types';
import classNames from 'classnames';
import { CustomInput as RSCustomInput } from 'reactstrap';

const CustomInput = forwardRef(({ className, label, ...otherProps }, ref) => {
    const inputClass = classNames(className, {
        'custom-control-empty': !label
    });

    return (
        <RSCustomInput className={inputClass} label={label} innerRef={ref} {...otherProps} />
    );
});

CustomInput.displayName = 'CustomInput';

CustomInput.propTypes = {
    className: PropTypes.string,
    label: PropTypes.node,
    ...RSCustomInput.propTypes,
};

export { CustomInput };
