import axios from 'axios';
import { useContext } from 'react';
import { AuthContext } from './AuthContext';
import { toast } from 'react-toastify';
import config from '../config';

const API_URL = config.apiUrl;

const apiClient = axios.create({
    baseURL: API_URL,
    withCredentials: true,
});

apiClient.interceptors.request.use(
    (config) => {
        const { token } = useContext(AuthContext);
        if (token) {
            config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
    },
    (error) => Promise.reject(error)
);

apiClient.interceptors.response.use(
    (response) => response,
    (error) => {
        if (error.response && error.response.status === 401) {
            toast.error('You do not have permission to access this resource.');
        }
        return Promise.reject(error);
    }
);

export default apiClient;
