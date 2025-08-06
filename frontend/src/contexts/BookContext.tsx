'use client';

import React, { createContext, useContext, useReducer, useEffect } from 'react';
import { Book, CreateBookRequest, UpdateBookRequest } from '@/types/book';
import { bookApi } from '@/utils/api';

interface BookState {
  books: Book[];
  loading: boolean;
  error: string | null;
}

type BookAction =
  | { type: 'SET_LOADING'; payload: boolean }
  | { type: 'SET_ERROR'; payload: string | null }
  | { type: 'SET_BOOKS'; payload: Book[] }
  | { type: 'ADD_BOOK'; payload: Book }
  | { type: 'UPDATE_BOOK'; payload: Book }
  | { type: 'DELETE_BOOK'; payload: string };

const initialState: BookState = {
  books: [],
  loading: false,
  error: null,
};

const bookReducer = (state: BookState, action: BookAction): BookState => {
  switch (action.type) {
    case 'SET_LOADING':
      return { ...state, loading: action.payload };
    case 'SET_ERROR':
      return { ...state, error: action.payload };
    case 'SET_BOOKS':
      return { ...state, books: action.payload };
    case 'ADD_BOOK':
      return { ...state, books: [action.payload, ...state.books] };
    case 'UPDATE_BOOK':
      return {
        ...state,
        books: state.books.map((book) =>
          book.id === action.payload.id ? action.payload : book
        ),
      };
    case 'DELETE_BOOK':
      return {
        ...state,
        books: state.books.filter((book) => book.id !== action.payload),
      };
    default:
      return state;
  }
};

interface BookContextType {
  state: BookState;
  fetchBooks: () => Promise<void>;
  createBook: (book: CreateBookRequest) => Promise<void>;
  updateBook: (id: string, book: UpdateBookRequest) => Promise<void>;
  deleteBook: (id: string) => Promise<void>;
}

const BookContext = createContext<BookContextType | undefined>(undefined);

export const useBookContext = () => {
  const context = useContext(BookContext);
  if (!context) {
    throw new Error('useBookContext must be used within a BookProvider');
  }
  return context;
};

interface BookProviderProps {
  children: React.ReactNode;
}

export const BookProvider: React.FC<BookProviderProps> = ({ children }) => {
  const [state, dispatch] = useReducer(bookReducer, initialState);
  
  // Get debug configuration from environment
  const isDebug = process.env.NEXT_PUBLIC_DEBUG === 'true';

  const fetchBooks = async () => {
    try {
      dispatch({ type: 'SET_LOADING', payload: true });
      dispatch({ type: 'SET_ERROR', payload: null });
      
      if (isDebug) {
        console.log('Fetching books from API...');
      }
      
      const books = await bookApi.getAll();
      dispatch({ type: 'SET_BOOKS', payload: books });
      
      if (isDebug) {
        console.log(`Fetched ${books.length} books`);
      }
    } catch (error) {
      const errorMessage = error instanceof Error ? error.message : 'Failed to fetch books';
      dispatch({ type: 'SET_ERROR', payload: errorMessage });
      
      if (isDebug) {
        console.error('Error fetching books:', error);
      }
    } finally {
      dispatch({ type: 'SET_LOADING', payload: false });
    }
  };

  const createBook = async (book: CreateBookRequest) => {
    try {
      dispatch({ type: 'SET_ERROR', payload: null });
      
      if (isDebug) {
        console.log('Creating book:', book);
      }
      
      const newBook = await bookApi.create(book);
      dispatch({ type: 'ADD_BOOK', payload: newBook });
      
      if (isDebug) {
        console.log('Book created successfully:', newBook);
      }
    } catch (error) {
      const errorMessage = error instanceof Error ? error.message : 'Failed to create book';
      dispatch({ type: 'SET_ERROR', payload: errorMessage });
      
      if (isDebug) {
        console.error('Error creating book:', error);
      }
      throw error;
    }
  };

  const updateBook = async (id: string, book: UpdateBookRequest) => {
    try {
      dispatch({ type: 'SET_ERROR', payload: null });
      
      if (isDebug) {
        console.log('Updating book:', { id, book });
      }
      
      const updatedBook = await bookApi.update(id, book);
      dispatch({ type: 'UPDATE_BOOK', payload: updatedBook });
      
      if (isDebug) {
        console.log('Book updated successfully:', updatedBook);
      }
    } catch (error) {
      const errorMessage = error instanceof Error ? error.message : 'Failed to update book';
      dispatch({ type: 'SET_ERROR', payload: errorMessage });
      
      if (isDebug) {
        console.error('Error updating book:', error);
      }
      throw error;
    }
  };

  const deleteBook = async (id: string) => {
    try {
      dispatch({ type: 'SET_ERROR', payload: null });
      
      if (isDebug) {
        console.log('Deleting book:', id);
      }
      
      await bookApi.delete(id);
      dispatch({ type: 'DELETE_BOOK', payload: id });
      
      if (isDebug) {
        console.log('Book deleted successfully:', id);
      }
    } catch (error) {
      const errorMessage = error instanceof Error ? error.message : 'Failed to delete book';
      dispatch({ type: 'SET_ERROR', payload: errorMessage });
      
      if (isDebug) {
        console.error('Error deleting book:', error);
      }
      throw error;
    }
  };

  useEffect(() => {
    fetchBooks();
  }, []);

  const value: BookContextType = {
    state,
    fetchBooks,
    createBook,
    updateBook,
    deleteBook,
  };

  return <BookContext.Provider value={value}>{children}</BookContext.Provider>;
}; 