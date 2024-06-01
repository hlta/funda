import React from 'react';
import PropTypes from 'prop-types';

const FooterText = ({ term, privacy, helps }) => (
    <div className="d-flex justify-content-between">
        <a
            href="/"
            target="_blank"
            rel="noopener noreferrer"
            className="text-decoration-none flex-grow-1"
        >
            {term}
        </a>
        <a
            href="/"
            target="_blank"
            rel="noopener noreferrer"
            className="text-decoration-none mx-5"
        >
            {privacy}
        </a>
        <a
            href="/"
            target="_blank"
            rel="noopener noreferrer"
            className="text-decoration-none flex-grow-1 text-end"
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
