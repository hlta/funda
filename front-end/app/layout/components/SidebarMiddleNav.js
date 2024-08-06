import React from 'react';

import { SidebarMenu } from './../../components';

export const SidebarMiddleNav = () => (
    <SidebarMenu>
        <SidebarMenu.Item
            icon={<i className="fa fa-fw fa-home"></i>}
            title="Dashboards"
            to="/dashboards/financial" exact
        />
           
        { /* -------- Business ---------*/ }
        <SidebarMenu.Item
            icon={<i className="fa fa-fw fa-money"></i>}
            title="Business"
        >
            <SidebarMenu.Item title="Invoices" to='/cards/cards' exact />
            <SidebarMenu.Item title="Quotes" to='/cards/cardsheaders' exact />
            <SidebarMenu.Item title="Sales overview" to='/cards/cardsheaders' exact />
            <SidebarMenu.Item title="Bills to pay" to='/cards/cardsheaders' exact />
            <SidebarMenu.Item title="Purchase orders" to='/cards/cardsheaders' exact />
            <SidebarMenu.Item title="Purchase overview" to='/cards/cardsheaders' exact />
            <SidebarMenu.Item title="Expense claims" to='/cards/cardsheaders' exact />
            <SidebarMenu.Item title="Products and services" to='/cards/cardsheaders' exact />


        </SidebarMenu.Item>
        { /* -------- Accounting ---------*/ }
        <SidebarMenu.Item
            icon={<i className="fa fa-fw fa-calculator"></i>}
            title="Accounting"
        >
            <SidebarMenu.Item title="Bank accounts" to='/layouts/navbar' exact />
            <SidebarMenu.Item title="Reports" to='/layouts/sidebar' exact />
            <SidebarMenu.Item title="Chart of accounts" to='/accounting/chart-of-accounts' exact />
            <SidebarMenu.Item title="Fixed assets" to='/layouts/dnd-layout' exact />

        </SidebarMenu.Item>
        { /* -------- Reports ---------*/ }
        <SidebarMenu.Item
            icon={<i className="fa fa-fw fa-line-chart"></i>}
            title="Reports"
        >
            <SidebarMenu.Item title="Account Transactions" to='/layouts/sidebar' exact />
            <SidebarMenu.Item title="Activity Statement" to='/layouts/sidebar' exact />
            <SidebarMenu.Item title="Balance Sheet" to='/layouts/sidebar-a' exact />
            <SidebarMenu.Item title="Profit Loss" to="/layouts/sidebar-with-navbar" exact />
        </SidebarMenu.Item>
        { /* -------- Contacts ---------*/ }
        <SidebarMenu.Item
            icon={<i className="fa fa-fw fa-address-book"></i>}
            title="Contacts"
        >
            <SidebarMenu.Item title="All contacts" to='/interface/colors' />
            <SidebarMenu.Item title="Customers" to='/interface/typography' />
            <SidebarMenu.Item title="Suppliers" to='/interface/buttons' />
            
        </SidebarMenu.Item>
    </SidebarMenu >
);
