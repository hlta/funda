import { textFilter, selectFilter, numberFilter } from 'react-bootstrap-table2-filter';

export const buildCustomTextFilter = ({ placeholder, getFilter }) => ({
  filter: textFilter({
    placeholder,
    getFilter,
  }),
});

export const buildCustomSelectFilter = ({ placeholder, options, getFilter }) => ({
  filter: selectFilter({
    placeholder,
    options,
    getFilter,
  }),
});

export const buildCustomNumberFilter = ({ comparators, getFilter }) => ({
  filter: numberFilter({
    comparators,
    getFilter,
  }),
});
