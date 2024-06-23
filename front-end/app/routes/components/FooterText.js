import React from 'react';
import PropTypes from 'prop-types';

const FooterText = ({ term, privacy, helps }) => (
    <div className="d-flex justify-content-between">
        <a
            href="/"
            target="_blank"
            rel="noopener noreferrer"
            className="text-decoration-none"
        >
            {term}
        </a>
        <a
            href="/"
            target="_blank"
            rel="noopener noreferrer"
            className="text-decoration-none"
        >
            {privacy}
        </a>
        <a
            href="/"
            target="_blank"
            rel="noopener noreferrer"
            className="text-decoration-none"
        >
            {helps}
        </a>
    </div>
);

FooterText.propTypes = {
    term: PropTypes.node,
    privacy: PropTypes.node,
    helps: PropTypes.node,
};

FooterText.defaultProps = {
    term: "Terms of Use",
    privacy: "Privacy",
    helps: "Help Center",
};

export { FooterText };
