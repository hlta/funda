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
                
                <ProtectedRoute path="/dashboards/analytics" exact component={Analytics} />
                <ProtectedRoute path="/dashboards/projects" exact component={ProjectsDashboard} />
                <ProtectedRoute path="/dashboards/system" exact component={System} />
                <ProtectedRoute path="/dashboards/monitor" exact component={Monitor} />
                <ProtectedRoute path="/dashboards/financial" exact component={Financial} />
                <ProtectedRoute path="/dashboards/stock" exact component={Stock} />
                <ProtectedRoute path="/dashboards/reports" exact component={Reports} />

                <ProtectedRoute path='/widgets' exact component={Widgets} />
                
                { /*    Cards Routes     */ }
                <ProtectedRoute path='/cards/cards' exact component={Cards} />
                <ProtectedRoute path='/cards/cardsheaders' exact component={CardsHeaders} />
                
                { /*    Layouts     */ }
                <ProtectedRoute path='/layouts/navbar' component={NavbarOnly} />
                <ProtectedRoute path='/layouts/sidebar' component={SidebarDefault} />
                <ProtectedRoute path='/layouts/sidebar-a' component={SidebarA} />
                <ProtectedRoute path="/layouts/sidebar-with-navbar" component={SidebarWithNavbar} />
                <ProtectedRoute path='/layouts/dnd-layout' component={DragAndDropLayout} />

                { /*    Interface Routes   */ }
                <ProtectedRoute component={ Accordions } path="/interface/accordions" />
                <ProtectedRoute component={ Alerts } path="/interface/alerts" />
                <ProtectedRoute component={ Avatars } path="/interface/avatars" />
                <ProtectedRoute component={ BadgesLabels } path="/interface/badges-and-labels" />
                <ProtectedRoute component={ Breadcrumbs } path="/interface/breadcrumbs" />
                <ProtectedRoute component={ Buttons } path="/interface/buttons" />
                <ProtectedRoute component={ Colors } path="/interface/colors" />
                <ProtectedRoute component={ Dropdowns } path="/interface/dropdowns" />
                <ProtectedRoute component={ Images } path="/interface/images" />
                <ProtectedRoute component={ ListGroups } path="/interface/list-groups" />
                <ProtectedRoute component={ MediaObjects } path="/interface/media-objects" />
                <ProtectedRoute component={ Modals } path="/interface/modals" />
                <ProtectedRoute component={ Navbars } path="/interface/navbars" />
                <ProtectedRoute component={ Paginations } path="/interface/paginations" />
                <ProtectedRoute component={ ProgressBars } path="/interface/progress-bars" />
                <ProtectedRoute component={ TabsPills } path="/interface/tabs-pills" />
                <ProtectedRoute component={ TooltipPopovers } path="/interface/tooltips-and-popovers" />
                <ProtectedRoute component={ Typography } path="/interface/typography" />
                <ProtectedRoute component={ Notifications } path="/interface/notifications" />
                <ProtectedRoute component={ CropImage } path="/interface/crop-image" />
                <ProtectedRoute component={ DragAndDropElements } path="/interface/drag-and-drop-elements" />
                <ProtectedRoute component={ Calendar } path="/interface/calendar" />

                { /*    Forms Routes    */ }
                <ProtectedRoute component={ Forms } path="/forms/forms" />
                <ProtectedRoute component={ FormsLayouts } path="/forms/forms-layouts" />
                <ProtectedRoute component={ InputGroups } path="/forms/input-groups" />
                <ProtectedRoute component={ Wizard } path="/forms/wizard" />
                <ProtectedRoute component={ TextMask } path="/forms/text-mask" />
                <ProtectedRoute component={ Typeahead } path="/forms/typeahead" />
                <ProtectedRoute component={ Toggles } path="/forms/toggles" />
                <ProtectedRoute component={ Editor } path="/forms/editor" />
                <ProtectedRoute component={ DatePicker } path="/forms/date-picker" />
                <ProtectedRoute component={ Dropzone } path="/forms/dropzone" />
                <ProtectedRoute component={ Sliders } path="/forms/sliders" />
                
                { /*    Graphs Routes   */ }
                <ProtectedRoute component={ ReCharts } path="/graphs/re-charts" />

                { /*    Tables Routes   */ }
                <ProtectedRoute component={ Tables } path="/tables/tables" />
                <ProtectedRoute component={ ExtendedTable } path="/tables/extended-table" />
                <ProtectedRoute component={ AgGrid } path="/tables/ag-grid" />

                { /*    Apps Routes     */ }
                <ProtectedRoute component={ AccountEdit } path="/apps/account-edit" />
                <ProtectedRoute component={ BillingEdit } path="/apps/billing-edit" />
                <ProtectedRoute component={ Chat } path="/apps/chat" />
                <ProtectedRoute component={ Clients } path="/apps/clients" />
                <ProtectedRoute component={ EmailDetails } path="/apps/email-details" />
                <ProtectedRoute component={ Files } path="/apps/files/:type"/>
                <ProtectedRoute component={ GalleryGrid } path="/apps/gallery-grid" />
                <ProtectedRoute component={ GalleryTable } path="/apps/gallery-table" />
                <ProtectedRoute component={ ImagesResults } path="/apps/images-results" />
                <ProtectedRoute component={ Inbox } path="/apps/inbox" />
                <ProtectedRoute component={ NewEmail } path="/apps/new-email" />
                <ProtectedRoute component={ ProfileDetails } path="/apps/profile-details" />
                <ProtectedRoute component={ ProfileEdit } path="/apps/profile-edit" />
                <ProtectedRoute component={ Projects } path="/apps/projects/:type" />
                <ProtectedRoute component={ SearchResults } path="/apps/search-results" />
                <ProtectedRoute component={ SessionsEdit } path="/apps/sessions-edit" />
                <ProtectedRoute component={ SettingsEdit } path="/apps/settings-edit" />
                <ProtectedRoute component={ Tasks } path="/apps/tasks/:type" />
                <ProtectedRoute component={ TasksDetails } path="/apps/task-details" />
                <ProtectedRoute component={ TasksKanban } path="/apps/tasks-kanban" />
                <ProtectedRoute component={ Users } path="/apps/users/:type" />
                <ProtectedRoute component={ UsersResults } path="/apps/users-results" />
                <ProtectedRoute component={ VideosResults } path="/apps/videos-results" />

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

                <ProtectedRoute path='/icons' exact component={Icons} />

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
