import React, { useEffect, useState } from 'react';
import PropTypes from 'prop-types';
import BootstrapTable from 'react-bootstrap-table-next';
import paginationFactory from 'react-bootstrap-table2-paginator';
import filterFactory, { Comparator, dateFilter } from 'react-bootstrap-table2-filter';
import ToolkitProvider from 'react-bootstrap-table2-toolkit';
import { Badge, Button, CustomInput, ButtonGroup } from '../../../components';
import { ExportButton, SearchBar, PaginationPanel, PaginationTotal, SizePerPageButton } from './components/tables';
import { useAccounts } from '../../../hooks/useAccounts';
import {
  buildCustomTextFilter,
  buildCustomSelectFilter,
  buildCustomNumberFilter,
} from './components/filters';

export const AccountList = ({ initialAccounts }) => {
  const { accounts, loading, fetchAccounts, createAccount, updateAccount, deleteAccount } = useAccounts();
  const [accountData, setAccountData] = useState(initialAccounts);

  useEffect(() => {
    const fetchData = async () => {
      await fetchAccounts();
    };
    fetchData();
  }, [fetchAccounts]);

  useEffect(() => {
    if (accounts.length) {
      setAccountData(accounts);
    }
  }, [accounts]);

  const handleCreateAccount = async () => {
    await createAccount({ name: 'New Account', balance: 0 });
    fetchAccounts();
  };

  const handleUpdateAccount = async (id, data) => {
    await updateAccount(id, data);
    fetchAccounts();
  };

  const handleDeleteAccount = async (id) => {
    await deleteAccount(id);
    fetchAccounts();
  };

  const handleResetFilters = () => {
    this.idFilter("");
    this.codeFilter("");
    this.nameFilter("");
    this.typeFilter("");
    this.taxRateFilter("");
    this.balanceFilter("");
    this.orgNameFilter("");
    this.ytdFilter("");
  };

  const createColumnDefinitions = () => [
    {
      dataField: 'id',
      text: 'Account ID',
      sort: true,
      sortCaret,
      ...buildCustomTextFilter({
        placeholder: 'Enter account ID...',
        getFilter: (filter) => {
          this.idFilter = filter;
        },
      }),
    },
    {
      dataField: 'code',
      text: 'Code',
      sort: true,
      sortCaret,
      ...buildCustomTextFilter({
        placeholder: 'Enter code...',
        getFilter: (filter) => {
          this.codeFilter = filter;
        },
      }),
    },
    {
      dataField: 'name',
      text: 'Name',
      sort: true,
      sortCaret,
      ...buildCustomTextFilter({
        placeholder: 'Enter name...',
        getFilter: (filter) => {
          this.nameFilter = filter;
        },
      }),
    },
    {
      dataField: 'type',
      text: 'Type',
      sort: true,
      sortCaret,
      ...buildCustomTextFilter({
        placeholder: 'Enter type...',
        getFilter: (filter) => {
          this.typeFilter = filter;
        },
      }),
    },
    {
      dataField: 'taxRate',
      text: 'Tax Rate',
      sort: true,
      sortCaret,
      ...buildCustomTextFilter({
        placeholder: 'Enter tax rate...',
        getFilter: (filter) => {
          this.taxRateFilter = filter;
        },
      }),
    },
    {
      dataField: 'balance',
      text: 'Balance',
      sort: true,
      sortCaret,
      ...buildCustomNumberFilter({
        comparators: [Comparator.EQ, Comparator.GT, Comparator.LT],
        getFilter: (filter) => {
          this.balanceFilter = filter;
        },
      }),
    },
    {
      dataField: 'orgName',
      text: 'Organization',
      sort: true,
      sortCaret,
      ...buildCustomTextFilter({
        placeholder: 'Enter organization name...',
        getFilter: (filter) => {
          this.orgNameFilter = filter;
        },
      }),
    },
    {
      dataField: 'ytd',
      text: 'Year to Date Balance',
      sort: true,
      sortCaret,
      ...buildCustomNumberFilter({
        comparators: [Comparator.EQ, Comparator.GT, Comparator.LT],
        getFilter: (filter) => {
          this.ytdFilter = filter;
        },
      }),
    },
  ];

  const columnDefs = createColumnDefinitions();
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
    sizePerPageRenderer: (props) => <SizePerPageButton {...props} />,
    paginationTotalRenderer: (from, to, size) => (
      <PaginationTotal {...{ from, to, size }} />
    ),
  });

  return (
    <ToolkitProvider
      keyField="id"
      data={accountData.length ? accountData : initialAccounts}
      columns={columnDefs}
      search
      exportCSV
    >
      {(props) => (
        <React.Fragment>
          <div className="d-flex justify-content-end align-items-center mb-2">
            <h6 className="my-0">Accounts</h6>
            <div className="d-flex ml-auto">
              <SearchBar className="mr-2" {...props.searchProps} />
              <ButtonGroup>
                <ExportButton {...props.csvProps}>Export</ExportButton>
                <Button
                  size="sm"
                  outline
                  onClick={handleCreateAccount}
                >
                  <i className="fa fa-fw fa-plus"></i>
                </Button>
                <Button
                  size="sm"
                  outline
                  onClick={handleResetFilters}
                >
                  Reset Filters
                </Button>
              </ButtonGroup>
            </div>
          </div>
          <BootstrapTable
            classes="table-responsive"
            pagination={paginationDef}
            filter={filterFactory()}
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
  initialAccounts: PropTypes.arrayOf(
    PropTypes.shape({
      id: PropTypes.number.isRequired,
      code: PropTypes.number.isRequired,
      name: PropTypes.string.isRequired,
      type: PropTypes.string.isRequired,
      taxRate: PropTypes.string.isRequired,
      balance: PropTypes.number.isRequired,
      orgID: PropTypes.number.isRequired,
      orgName: PropTypes.string.isRequired,
      ytd: PropTypes.number.isRequired,
    })
  ),
  searchProps: PropTypes.shape({
    placeholder: PropTypes.string,
    delay: PropTypes.number,
    searchText: PropTypes.string,
    onSearch: PropTypes.func,
  }),
  csvProps: PropTypes.object,
  baseProps: PropTypes.object,
};

AccountList.defaultProps = {
  initialAccounts: [],
  searchProps: {},
  csvProps: {},
  baseProps: {},
};

const sortCaret = (order) => {
  if (!order) return <i className="fa fa-fw fa-sort text-muted"></i>;
  if (order === 'asc') return <i className="fa fa-fw fa-sort-asc text-muted"></i>;
  if (order === 'desc') return <i className="fa fa-fw fa-sort-desc text-muted"></i>;
  return null;
};
