import React from 'react';
import { CSVExport } from 'react-bootstrap-table2-toolkit';

const { ExportCSVButton } = CSVExport;

const ExportButton = (props) => (
  <ExportCSVButton {...props} className="btn btn-sm btn-outline-secondary">
    Export
  </ExportCSVButton>
);

export default ExportButton;
