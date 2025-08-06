import axios from 'axios';
import { Book, CreateBookRequest, UpdateBookRequest } from '@/types/book';
import { URLRequest, URLResponse } from '@/types/url';

// Get API configuration from environment variables
const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';
const API_VERSION = process.env.NEXT_PUBLIC_API_VERSION || 'v1';
const API_PREFIX = process.env.NEXT_PUBLIC_API_PREFIX || '/api';

// Create axios instance with configuration
const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
  timeout: parseInt(process.env.NEXT_PUBLIC_API_TIMEOUT || '30000'),
});

// Add request interceptor for logging in debug mode
if (process.env.NEXT_PUBLIC_DEBUG === 'true') {
  api.interceptors.request.use(
    (config) => {
      console.log(`API Request: ${config.method?.toUpperCase()} ${config.url}`);
      return config;
    },
    (error) => {
      console.error('API Request Error:', error);
      return Promise.reject(error);
    }
  );

  api.interceptors.response.use(
    (response) => {
      console.log(`API Response: ${response.status} ${response.config.url}`);
      return response;
    },
    (error) => {
      console.error('API Response Error:', error.response?.status, error.response?.data);
      return Promise.reject(error);
    }
  );
}

// Book API functions
export const bookApi = {
  getAll: async (): Promise<Book[]> => {
    const response = await api.get(`${API_PREFIX}/books`);
    return response.data;
  },

  getById: async (id: string): Promise<Book> => {
    const response = await api.get(`${API_PREFIX}/books/${id}`);
    return response.data;
  },

  create: async (book: CreateBookRequest): Promise<Book> => {
    const response = await api.post(`${API_PREFIX}/books`, book);
    return response.data;
  },

  update: async (id: string, book: UpdateBookRequest): Promise<Book> => {
    const response = await api.put(`${API_PREFIX}/books/${id}`, book);
    return response.data;
  },

  delete: async (id: string): Promise<void> => {
    await api.delete(`${API_PREFIX}/books/${id}`);
  },
};

// URL API functions
export const urlApi = {
  process: async (request: URLRequest): Promise<URLResponse> => {
    const response = await api.post(`${API_PREFIX}/url/process`, request);
    return response.data;
  },
};

// Health check function
export const healthApi = {
  check: async (): Promise<{ status: string; service: string; version: string }> => {
    const healthEndpoint = process.env.NEXT_PUBLIC_HEALTH_CHECK_ENDPOINT || '/health';
    const response = await api.get(healthEndpoint);
    return response.data;
  },
};

// Export API configuration for debugging
export const apiConfig = {
  baseURL: API_BASE_URL,
  version: API_VERSION,
  prefix: API_PREFIX,
  timeout: parseInt(process.env.NEXT_PUBLIC_API_TIMEOUT || '30000'),
  debug: process.env.NEXT_PUBLIC_DEBUG === 'true',
}; 