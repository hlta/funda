import React, { useEffect, useState } from 'react';
import PropTypes from 'prop-types';
import BootstrapTable from 'react-bootstrap-table-next';
import paginationFactory from 'react-bootstrap-table2-paginator';
import filterFactory, { Comparator } from 'react-bootstrap-table2-filter';
import ToolkitProvider from 'react-bootstrap-table2-toolkit';
import { Button, CustomInput, ButtonGroup } from 'reactstrap';
import axios from 'axios';
import { useHistory } from 'react-router-dom';
import ExportButton from './ExportButton';
import SearchBar from './SearchBar';
import PaginationPanel from './PaginationPanel';
import SizePerPageDropdown from './SizePerPageDropdown';
import PaginationTotal from './PaginationTotal';
import { buildCustomTextFilter, buildCustomSelectFilter, buildCustomNumberFilter } from './filters';
import './AccountList.css'; // Import custom CSS file

const sortCaret = (order) => {
  if (!order) return <i className="fa fa-fw fa-sort text-muted"></i>;
  if (order === 'asc') return <i className="fa fa-fw fa-sort-up text-muted"></i>;
  if (order === 'desc') return <i className="fa fa-fw fa-sort-down text-muted"></i>;
  return null;
};

const AccountList = () => {
  const [accounts, setAccounts] = useState([]);
  const [selected, setSelected] = useState([]);
  const history = useHistory();

  useEffect(() => {
    const fetchAccounts = async () => {
      const response = await axios.get('/api/accounts');
      setAccounts(response.data.data);
    };
    fetchAccounts();
  }, []);

  const handleSelect = (row, isSelected) => {
    if (isSelected) {
      setSelected([...selected, row.id]);
    } else {
      setSelected(selected.filter((itemId) => itemId !== row.id));
    }
  };

  const handleSelectAll = (isSelected, rows) => {
    if (isSelected) {
      setSelected(rows.map((row) => row.id));
    } else {
      setSelected([]);
    }
  };

  const handleAddRow = () => {
    history.push('/accounting/chart-of-accounts/new');
  };

  const handleDeleteRow = () => {
    setAccounts(accounts.filter((account) => !selected.includes(account.id)));
    setSelected([]);
  };

  const handleResetFilters = () => {
    codeFilter('');
    nameFilter('');
    typeFilter('');
    taxRateFilter('');
    balanceFilter('');
    ytdFilter('');
  };

  let codeFilter;
  let nameFilter;
  let typeFilter;
  let taxRateFilter;
  let balanceFilter;
  let ytdFilter;

  const columnDefs = [
    {
      dataField: 'id',
      text: 'Account ID',
      headerFormatter: (column) => (
        <React.Fragment>
          <span className="text-nowrap">{column.text}</span>
          <a
            href="javascript:;"
            className="d-block small text-decoration-none text-nowrap"
            onClick={handleResetFilters}
          >
            Reset Filters <i className="fa fa-times fa-fw text-danger"></i>
          </a>
        </React.Fragment>
      ),
    },
    {
      dataField: 'code',
      text: 'Code',
      sort: true,
      sortCaret,
      formatter: (cell) => <span className="text-inverse">{cell}</span>,
      ...buildCustomTextFilter({
        placeholder: 'Enter code...',
        getFilter: (filter) => {
          codeFilter = filter;
        },
      }),
    },
    {
      dataField: 'name',
      text: 'Name',
      sort: true,
      sortCaret,
      formatter: (cell) => <span className="text-inverse">{cell}</span>,
      ...buildCustomTextFilter({
        placeholder: 'Enter name...',
        getFilter: (filter) => {
          nameFilter = filter;
        },
      }),
    },
    {
      dataField: 'type',
      text: 'Type',
      sort: true,
      sortCaret,
      ...buildCustomSelectFilter({
        placeholder: 'Select type',
        options: [
          { value: 'Asset', label: 'Asset' },
          { value: 'Liability', label: 'Liability' },
          { value: 'Equity', label: 'Equity' },
          { value: 'Expense', label: 'Expense' },
          { value: 'Revenue', label: 'Revenue' },
        ],
        getFilter: (filter) => {
          typeFilter = filter;
        },
      }),
    },
    {
      dataField: 'tax_rate',
      text: 'Tax Rate',
      sort: true,
      sortCaret,
      ...buildCustomTextFilter({
        placeholder: 'Enter tax rate...',
        getFilter: (filter) => {
          taxRateFilter = filter;
        },
      }),
    },
    {
      dataField: 'balance',
      text: 'Balance',
      sort: true,
      sortCaret,
      formatter: (cell) => `$${cell.toFixed(2)}`,
      ...buildCustomNumberFilter({
        comparators: [Comparator.EQ, Comparator.GT, Comparator.LT],
        getFilter: (filter) => {
          balanceFilter = filter;
        },
      }),
    },
    {
      dataField: 'ytd',
      text: 'YTD',
      sort: true,
      sortCaret,
      formatter: (cell) => `$${cell.toFixed(2)}`,
      ...buildCustomNumberFilter({
        comparators: [Comparator.EQ, Comparator.GT, Comparator.LT],
        getFilter: (filter) => {
          ytdFilter = filter;
        },
      }),
    },
  ];

  const paginationDef = paginationFactory({
    paginationSize: 5,
    showTotal: true,
    pageListRenderer: (props) => (
      <PaginationPanel
        {...props}
        size="sm"
        className="ml-md-auto mt-2 mt-md-0"
      />
    ),
    sizePerPageRenderer: (props) => <SizePerPageDropdown {...props} />,
    paginationTotalRenderer: (from, to, size) => (
      <PaginationTotal {...{ from, to, size }} />
    ),
  });

  const selectRowConfig = {
    mode: 'checkbox',
    selected: selected,
    onSelect: handleSelect,
    onSelectAll: handleSelectAll,
    selectionRenderer: ({ mode, checked, disabled }) => (
      <CustomInput type={mode} checked={checked} disabled={disabled} />
    ),
    selectionHeaderRenderer: ({ mode, checked, indeterminate }) => (
      <CustomInput
        type={mode}
        checked={checked}
        innerRef={(el) => el && (el.indeterminate = indeterminate)}
      />
    ),
  };

  return (
    <ToolkitProvider
      keyField="id"
      data={accounts}
      columns={columnDefs}
      search
      exportCSV
    >
      {(props) => (
        <React.Fragment>
          <div className="d-flex justify-content-end align-items-center mb-2">
            <h6 className="my-0">Chart of Accounts</h6>
            <div className="d-flex ml-auto">
              <SearchBar className="mr-2" {...props.searchProps} />
              <ButtonGroup>
                <ExportButton {...props.csvProps}>Export</ExportButton>
                <Button
                  size="sm"
                  outline
                  onClick={handleDeleteRow}
                >
                  Delete
                </Button>
                <Button
                  size="sm"
                  outline
                  onClick={handleAddRow}
                >
                  <i className="fa fa-fw fa-plus"></i>
                </Button>
              </ButtonGroup>
            </div>
          </div>
          <div className="filter-container">
            <div className="filter-box">
              {columnDefs.map((column, idx) => (
                column.filter && (
                  <div key={idx} className="filter-item">
                    {column.filter}
                  </div>
                )
              ))}
            </div>
          </div>
          <BootstrapTable
            classes="table-responsive"
            pagination={paginationDef}
            filter={filterFactory()}
            selectRow={selectRowConfig}
            bordered={false}
            responsive
            {...props.baseProps}
          />
        </React.Fragment>
      )}
    </ToolkitProvider>
  );
};

AccountList.propTypes = {
  searchProps: PropTypes.shape({
    onSearch: PropTypes.func,
    onClear: PropTypes.func,
    searchText: PropTypes.string,
  }),
  csvProps: PropTypes.object,
  baseProps: PropTypes.object,
};

export default AccountList;
