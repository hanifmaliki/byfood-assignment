import axios from 'axios';
import { Book, CreateBookRequest, UpdateBookRequest } from '@/types/book';
import { URLRequest, URLResponse } from '@/types/url';

const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Book API functions
export const bookApi = {
  getAll: async (): Promise<Book[]> => {
    const response = await api.get('/api/books');
    return response.data;
  },

  getById: async (id: string): Promise<Book> => {
    const response = await api.get(`/api/books/${id}`);
    return response.data;
  },

  create: async (book: CreateBookRequest): Promise<Book> => {
    const response = await api.post('/api/books', book);
    return response.data;
  },

  update: async (id: string, book: UpdateBookRequest): Promise<Book> => {
    const response = await api.put(`/api/books/${id}`, book);
    return response.data;
  },

  delete: async (id: string): Promise<void> => {
    await api.delete(`/api/books/${id}`);
  },
};

// URL API functions
export const urlApi = {
  process: async (request: URLRequest): Promise<URLResponse> => {
    const response = await api.post('/api/url/process', request);
    return response.data;
  },
}; 