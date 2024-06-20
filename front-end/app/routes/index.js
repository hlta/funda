import React from 'react';
import {
    Route,
    Switch,
    Redirect
} from 'react-router';

// ----------- Pages Imports ---------------
import Analytics from './Dashboards/Analytics';
import ProjectsDashboard from './Dashboards/Projects';
import System from './Dashboards/System';
import Monitor from './Dashboards/Monitor'; 
import Financial from './Dashboards/Financial';
import Stock from './Dashboards/Stock';
import Reports from './Dashboards/Reports';

import Widgets from './Widgets';

import Cards from './Cards/Cards';
import CardsHeaders from './Cards/CardsHeaders';

import NavbarOnly from './Layouts/NavbarOnly';
import SidebarDefault from './Layouts/SidebarDefault';
import SidebarA from './Layouts/SidebarA';
import DragAndDropLayout from './Layouts/DragAndDropLayout';
import SidebarWithNavbar from './Layouts/SidebarWithNavbar';

import Accordions from './Interface/Accordions';
import Alerts from './Interface/Alerts';
import Avatars from './Interface/Avatars';
import BadgesLabels from './Interface/BadgesLabels';
import Breadcrumbs from './Interface/Breadcrumbs';
import Buttons from './Interface/Buttons';
import Colors from './Interface/Colors';
import Dropdowns from './Interface/Dropdowns';
import Images from './Interface/Images';
import ListGroups from './Interface/ListGroups';
import MediaObjects from './Interface/MediaObjects';
import Modals from './Interface/Modals';
import Navbars from './Interface/Navbars';
import Paginations from './Interface/Paginations';
import ProgressBars from './Interface/ProgressBars';
import TabsPills from './Interface/TabsPills';
import TooltipPopovers from './Interface/TooltipsPopovers';
import Typography from './Interface/Typography';
import Notifications from './Interface/Notifications';
import CropImage from './Interface/CropImage';
import DragAndDropElements from './Interface/DragAndDropElements';
import Calendar from './Interface/Calendar';
import ReCharts from './Graphs/ReCharts';

import Forms from './Forms/Forms';
import FormsLayouts from './Forms/FormsLayouts';
import InputGroups from './Forms/InputGroups';
import Wizard from './Forms/Wizard';
import TextMask from './Forms/TextMask';
import Typeahead from './Forms/Typeahead';
import Toggles from './Forms/Toggles';
import Editor from './Forms/Editor';
import DatePicker from './Forms/DatePicker';
import Dropzone from './Forms/Dropzone';
import Sliders from './Forms/Sliders';

import Tables from './Tables/Tables';
import ExtendedTable from './Tables/ExtendedTable';
import AgGrid from './Tables/AgGrid';

import AccountEdit from './Apps/AccountEdit';
import BillingEdit from './Apps/BillingEdit';
import Chat from './Apps/Chat';
import Clients from './Apps/Clients';
import EmailDetails from './Apps/EmailDetails';
import Files from './Apps/Files';
import GalleryGrid from './Apps/GalleryGrid';
import GalleryTable from './Apps/GalleryTable';
import ImagesResults from './Apps/ImagesResults';
import Inbox from './Apps/Inbox';
import NewEmail from './Apps/NewEmail';
import ProfileDetails from './Apps/ProfileDetails';
import ProfileEdit from './Apps/ProfileEdit';
import Projects from './Apps/Projects';
import SearchResults from './Apps/SearchResults';
import SessionsEdit from './Apps/SessionsEdit';
import SettingsEdit from './Apps/SettingsEdit';
import Tasks from './Apps/Tasks';
import TasksDetails from './Apps/TasksDetails';
import TasksKanban from './Apps/TasksKanban';
import Users from './Apps/Users';
import UsersResults from './Apps/UsersResults';
import VideosResults from './Apps/VideosResults';

import ComingSoon from './Pages/ComingSoon';
import Confirmation from './Pages/Confirmation';
import Danger from './Pages/Danger';
import Error404 from './Pages/Error404';
import ForgotPassword from './Pages/ForgotPassword';
import LockScreen from './Pages/LockScreen';
import Login from './Pages/Login';
import Register from './Pages/Register';
import Success from './Pages/Success';
import Timeline from './Pages/Timeline';

import Icons from './Icons';
import ProtectedRoute from './ProtectedRoute';

// ----------- Layout Imports ---------------
import { DefaultNavbar } from './../layout/components/DefaultNavbar';
import { DefaultSidebar } from './../layout/components/DefaultSidebar';

import { SidebarANavbar } from './../layout/components/SidebarANavbar';
import { SidebarASidebar } from './../layout/components/SidebarASidebar';



//------ Route Definitions --------
// eslint-disable-next-line no-unused-vars
export const RoutedContent = () => {
    return (
            <Switch>
                <Redirect from="/" to="/dashboards/projects" exact />
                
                <ProtectedRoute requiredPermissions={['View Reports']} path="/dashboards/analytics" exact component={Analytics} />
                <ProtectedRoute requiredPermissions={['View Reports']} path="/dashboards/projects" exact component={ProjectsDashboard} />
                <ProtectedRoute requiredPermissions={['View Reports']} path="/dashboards/system" exact component={System} />
                <ProtectedRoute requiredPermissions={['View Reports']} path="/dashboards/monitor" exact component={Monitor} />
                <ProtectedRoute requiredPermissions={['View Reports']} path="/dashboards/financial" exact component={Financial} />
                <ProtectedRoute requiredPermissions={['View Reports']} path="/dashboards/stock" exact component={Stock} />
                <ProtectedRoute requiredPermissions={['View Reports']} path="/dashboards/reports" exact component={Reports} />

                <ProtectedRoute requiredPermissions={['View Reports']} path='/widgets' exact component={Widgets} />
                
                { /*    Cards Routes     */ }
                <ProtectedRoute requiredPermissions={['View Reports']} path='/cards/cards' exact component={Cards} />
                <ProtectedRoute requiredPermissions={['View Reports']} path='/cards/cardsheaders' exact component={CardsHeaders} />
                
                { /*    Layouts     */ }
                <ProtectedRoute requiredPermissions={['View Reports']} path='/layouts/navbar' component={NavbarOnly} />
                <ProtectedRoute requiredPermissions={['View Reports']} path='/layouts/sidebar' component={SidebarDefault} />
                <ProtectedRoute requiredPermissions={['View Reports']} path='/layouts/sidebar-a' component={SidebarA} />
                <ProtectedRoute requiredPermissions={['View Reports']} path="/layouts/sidebar-with-navbar" component={SidebarWithNavbar} />
                <ProtectedRoute requiredPermissions={['View Reports']} path='/layouts/dnd-layout' component={DragAndDropLayout} />

                { /*    Interface Routes   */ }
                <ProtectedRoute requiredPermissions={['View Reports']} component={ Accordions } path="/interface/accordions" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ Alerts } path="/interface/alerts" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ Avatars } path="/interface/avatars" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ BadgesLabels } path="/interface/badges-and-labels" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ Breadcrumbs } path="/interface/breadcrumbs" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ Buttons } path="/interface/buttons" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ Colors } path="/interface/colors" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ Dropdowns } path="/interface/dropdowns" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ Images } path="/interface/images" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ ListGroups } path="/interface/list-groups" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ MediaObjects } path="/interface/media-objects" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ Modals } path="/interface/modals" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ Navbars } path="/interface/navbars" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ Paginations } path="/interface/paginations" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ ProgressBars } path="/interface/progress-bars" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ TabsPills } path="/interface/tabs-pills" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ TooltipPopovers } path="/interface/tooltips-and-popovers" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ Typography } path="/interface/typography" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ Notifications } path="/interface/notifications" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ CropImage } path="/interface/crop-image" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ DragAndDropElements } path="/interface/drag-and-drop-elements" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ Calendar } path="/interface/calendar" />

                { /*    Forms Routes    */ }
                <ProtectedRoute requiredPermissions={['View Reports']} component={ Forms } path="/forms/forms" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ FormsLayouts } path="/forms/forms-layouts" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ InputGroups } path="/forms/input-groups" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ Wizard } path="/forms/wizard" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ TextMask } path="/forms/text-mask" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ Typeahead } path="/forms/typeahead" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ Toggles } path="/forms/toggles" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ Editor } path="/forms/editor" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ DatePicker } path="/forms/date-picker" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ Dropzone } path="/forms/dropzone" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ Sliders } path="/forms/sliders" />
                
                { /*    Graphs Routes   */ }
                <ProtectedRoute requiredPermissions={['View Reports']} component={ ReCharts } path="/graphs/re-charts" />

                { /*    Tables Routes   */ }
                <ProtectedRoute requiredPermissions={['View Reports']} component={ Tables } path="/tables/tables" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ ExtendedTable } path="/tables/extended-table" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ AgGrid } path="/tables/ag-grid" />

                { /*    Apps Routes     */ }
                <ProtectedRoute requiredPermissions={['View Reports']} component={ AccountEdit } path="/apps/account-edit" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ BillingEdit } path="/apps/billing-edit" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ Chat } path="/apps/chat" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ Clients } path="/apps/clients" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ EmailDetails } path="/apps/email-details" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ Files } path="/apps/files/:type"/>
                <ProtectedRoute requiredPermissions={['View Reports']} component={ GalleryGrid } path="/apps/gallery-grid" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ GalleryTable } path="/apps/gallery-table" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ ImagesResults } path="/apps/images-results" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ Inbox } path="/apps/inbox" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ NewEmail } path="/apps/new-email" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ ProfileDetails } path="/apps/profile-details" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ ProfileEdit } path="/apps/profile-edit" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ Projects } path="/apps/projects/:type" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ SearchResults } path="/apps/search-results" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ SessionsEdit } path="/apps/sessions-edit" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ SettingsEdit } path="/apps/settings-edit" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ Tasks } path="/apps/tasks/:type" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ TasksDetails } path="/apps/task-details" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ TasksKanban } path="/apps/tasks-kanban" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ Users } path="/apps/users/:type" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ UsersResults } path="/apps/users-results" />
                <ProtectedRoute requiredPermissions={['View Reports']} component={ VideosResults } path="/apps/videos-results" />

                { /*    Pages Routes    */ }
                <Route component={ ComingSoon } path="/pages/coming-soon" />
                <Route component={ Confirmation } path="/pages/confirmation" />
                <Route component={ Danger } path="/pages/danger" />
                <Route component={ Error404 } path="/error-404" />
                <Route component={ ForgotPassword } path="/forgot-password" />
                <Route component={ LockScreen } path="/pages/lock-screen" />
                <Route component={ Login } path="/login" />
                <Route component={ Register } path="/register" />
                <Route component={ Success } path="/pages/success" />
                <Route component={ Timeline } path="/pages/timeline" />

                <ProtectedRoute requiredPermissions={['View Reports']} path='/icons' exact component={Icons} />

                { /*    404    */ }
                <Redirect to="/error-404" />
            </Switch>
    );
};

//------ Custom Layout Parts --------
export const RoutedNavbars  = () => (
    <Switch>
        { /* Other Navbars: */}
        <Route
            component={ SidebarANavbar }
            path="/layouts/sidebar-a"
        />
        <Route
            component={ NavbarOnly.Navbar }
            path="/layouts/navbar"
        />
        <Route
            component={ SidebarWithNavbar.Navbar }
            path="/layouts/sidebar-with-navbar"
        />
        { /* Default Navbar: */}
        <Route
            component={ DefaultNavbar }
        />
    </Switch>  
);

export const RoutedSidebars = () => (
    <Switch>
        { /* Other Sidebars: */}
        <Route
            component={ SidebarASidebar }
            path="/layouts/sidebar-a"
        />
        <Route
            component={ SidebarWithNavbar.Sidebar }
            path="/layouts/sidebar-with-navbar"
        />
        { /* Default Sidebar: */}
        <Route
            component={ DefaultSidebar }
        />
    </Switch>
);
