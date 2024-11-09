import type { AxiosRequestConfig, AxiosResponse } from 'axios';

import axios from 'axios';
import { get } from 'lodash';

import { CONFIG } from 'src/config-global';

// ----------------------------------------------------------------------

const axiosInstance = axios.create({
  baseURL: CONFIG.serverUrl,
  validateStatus: (status) => status < 400, // Resolve only if the status code is less than 400
});

axiosInstance.interceptors.request.use((config) => {
  const accessToken = sessionStorage.getItem(CONFIG.ACCESS_TOKEN);
  const newConfig = { ...config };

  if (accessToken) {
    newConfig.headers.Authorization = `Bearer ${accessToken}`;
  }

  return newConfig;
});

axiosInstance.interceptors.response.use(
  (response) => {
    refreshToken(response);
    return response;
  },
  (error) => {
    if (error.response.status === 401) {
      removeToken();
      window.location.href = '/auth/jwt/login';
    }
    return Promise.reject(error);
  }
);

export async function setSession(accessToken: string | null) {
  try {
    if (accessToken) {
      sessionStorage.setItem(CONFIG.ACCESS_TOKEN, accessToken);

      axios.defaults.headers.common.Authorization = `Bearer ${accessToken}`;
    } else {
      sessionStorage.removeItem(CONFIG.ACCESS_TOKEN);
      delete axios.defaults.headers.common.Authorization;
    }
  } catch (error) {
    console.error('Error during set session:', error);
    throw error;
  }
}

function refreshToken(response: AxiosResponse) {
  const token = get(response, 'data.header.token', '');
  if (token) {
    setSession(token);
  }
}

export function removeToken() {
  sessionStorage.removeItem(CONFIG.ACCESS_TOKEN);
  axiosInstance.defaults.headers.common.Authorization = '';
  delete axiosInstance.defaults.headers.common.Authorization;
}

export default axiosInstance;

// ----------------------------------------------------------------------

export const fetcher = async (args: string | [string, AxiosRequestConfig]) => {
  try {
    const [url, config] = Array.isArray(args) ? args : [args];

    const res = await axiosInstance.get(url, { ...config });

    return res.data;
  } catch (error) {
    console.error('Failed to fetch:', error);
    throw error;
  }
};

export const postFetcher = async (url: string, body: any) => {
  const res = await axiosInstance.post(url, {
    ...body,
  });
  return res.data;
};

// ----------------------------------------------------------------------

export const endpoints = {
  auth: {
    me: '/me',
    signIn: '/login',
    signUp: '/register',
  },

  models: {
    create: '/maas/create',
    update: '/maas/update',
    delete: (id: string | number) => `/model/delete/${id}`,
    detail: (id: string | number) => `/model/${id}`,
    list: '/models',
    invoke: '/maas/invoke',
  },

  product: {
    list: '/api/product/list',
    details: '/api/product/details',
    search: '/api/product/search',
  },
};
