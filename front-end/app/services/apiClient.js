import axios from 'axios';
import config from '../config';
import { toast } from 'react-toastify';

const API_URL = config.apiUrl;

export const createApiClient = (token) => {
    const apiClient = axios.create({
        baseURL: API_URL,
        withCredentials: true,
    });

    apiClient.interceptors.request.use(
        (config) => {
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

    return apiClient;
};
