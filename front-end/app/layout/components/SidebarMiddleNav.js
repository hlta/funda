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
            <SidebarMenu.Item title="Cards" to='/cards/cards' exact />
            <SidebarMenu.Item title="Cards Headers" to='/cards/cardsheaders' exact />
        </SidebarMenu.Item>
        { /* -------- Accounting ---------*/ }
        <SidebarMenu.Item
            icon={<i className="fa fa-fw fa-calculator"></i>}
            title="Accounting"
        >
            <SidebarMenu.Item title="Navbar" to='/layouts/navbar' exact />
            <SidebarMenu.Item title="Sidebar" to='/layouts/sidebar' exact />
            <SidebarMenu.Item title="Sidebar A" to='/layouts/sidebar-a' exact />
            <SidebarMenu.Item title="Sidebar With Navbar" to="/layouts/sidebar-with-navbar" exact />
            <SidebarMenu.Item title="Drag &amp; Drop" to='/layouts/dnd-layout' exact />
        </SidebarMenu.Item>
        { /* -------- Reports ---------*/ }
        <SidebarMenu.Item
            icon={<i className="fa fa-fw fa-line-chart"></i>}
            title="Reports"
        >
            <SidebarMenu.Item title="Navbar" to='/layouts/navbar' exact />
            <SidebarMenu.Item title="Sidebar" to='/layouts/sidebar' exact />
            <SidebarMenu.Item title="Sidebar A" to='/layouts/sidebar-a' exact />
            <SidebarMenu.Item title="Sidebar With Navbar" to="/layouts/sidebar-with-navbar" exact />
            <SidebarMenu.Item title="Drag &amp; Drop" to='/layouts/dnd-layout' exact />
        </SidebarMenu.Item>
        { /* -------- Contacts ---------*/ }
        <SidebarMenu.Item
            icon={<i className="fa fa-fw fa-address-book"></i>}
            title="Contacts"
        >
            <SidebarMenu.Item title="Colors" to='/interface/colors' />
            <SidebarMenu.Item title="Typography" to='/interface/typography' />
            <SidebarMenu.Item title="Buttons" to='/interface/buttons' />
            <SidebarMenu.Item title="Paginations" to='/interface/paginations' />
            <SidebarMenu.Item title="Images" to='/interface/images' />
            <SidebarMenu.Item title="Avatars" to='/interface/avatars' />
            <SidebarMenu.Item title="Progress Bars" to='/interface/progress-bars' />
            <SidebarMenu.Item title="Badges &amp; Labels" to='/interface/badges-and-labels' />
            <SidebarMenu.Item title="Media Objects" to='/interface/media-objects' />
            <SidebarMenu.Item title="List Groups" to='/interface/list-groups' />
            <SidebarMenu.Item title="Alerts" to='/interface/alerts' />
            <SidebarMenu.Item title="Accordions" to='/interface/accordions' />
            <SidebarMenu.Item title="Tabs Pills" to='/interface/tabs-pills' />
            <SidebarMenu.Item title="Tooltips &amp; Popovers" to='/interface/tooltips-and-popovers' />
            <SidebarMenu.Item title="Dropdowns" to='/interface/dropdowns' />
            <SidebarMenu.Item title="Modals" to='/interface/modals' />
            <SidebarMenu.Item title="Breadcrumbs" to='/interface/breadcrumbs' />
            <SidebarMenu.Item title="Navbars" to='/interface/navbars' />
            <SidebarMenu.Item title="Notifications" to='/interface/notifications' />
            <SidebarMenu.Item title="Crop Image" to='/interface/crop-image' />
            <SidebarMenu.Item title="Drag &amp; Drop Elements" to='/interface/drag-and-drop-elements' />
            <SidebarMenu.Item title="Calendar" to='/interface/calendar' />
        </SidebarMenu.Item>
    </SidebarMenu >
);
