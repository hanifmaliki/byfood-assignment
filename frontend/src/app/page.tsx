'use client';

import React, { useState } from 'react';
import { Plus, BookOpen, AlertCircle } from 'lucide-react';
import Link from 'next/link';
import { BookCard } from '@/components/BookCard';
import { BookForm } from '@/components/BookForm';
import { useBookContext } from '@/contexts/BookContext';
import { Book, CreateBookRequest, UpdateBookRequest } from '@/types/book';

export default function Home() {
  const { state, createBook, updateBook, deleteBook } = useBookContext();
  const [isFormOpen, setIsFormOpen] = useState(false);
  const [selectedBook, setSelectedBook] = useState<Book | undefined>(undefined);
  const [formMode, setFormMode] = useState<'create' | 'edit'>('create');
  const [showDeleteConfirm, setShowDeleteConfirm] = useState<string | null>(null);

  // Get app configuration from environment variables
  const appName = process.env.NEXT_PUBLIC_APP_NAME || 'Library Management System';
  const appVersion = process.env.NEXT_PUBLIC_APP_VERSION || '1.0.0';
  const isDebug = process.env.NEXT_PUBLIC_DEBUG === 'true';

  const handleAddBook = () => {
    setSelectedBook(undefined);
    setFormMode('create');
    setIsFormOpen(true);
  };

  const handleEditBook = (book: Book) => {
    setSelectedBook(book);
    setFormMode('edit');
    setIsFormOpen(true);
  };

  const handleViewBook = (book: Book) => {
    // Navigate to dynamic route for viewing details
    window.location.href = `/books/${book.id}`;
  };

  const handleDeleteBook = async (id: string) => {
    try {
      await deleteBook(id);
      setShowDeleteConfirm(null);
    } catch (error) {
      console.error('Failed to delete book:', error);
    }
  };

  const handleFormSubmit = async (bookData: CreateBookRequest | UpdateBookRequest) => {
    try {
      if (formMode === 'create') {
        await createBook(bookData as CreateBookRequest);
      } else {
        if (selectedBook) {
          await updateBook(selectedBook.id, bookData as UpdateBookRequest);
        }
      }
    } catch (error) {
      console.error('Form submission error:', error);
    }
  };

  return (
    <div className="min-h-screen bg-gray-50">
      {/* Header */}
      <header className="bg-white shadow-sm border-b">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center h-16">
            <div className="flex items-center space-x-3">
              <BookOpen className="h-8 w-8 text-blue-600" />
              <h1 className="text-2xl font-bold text-gray-900">
                {appName}
              </h1>
              {isDebug && (
                <span className="text-xs text-gray-500 bg-gray-100 px-2 py-1 rounded">
                  v{appVersion}
                </span>
              )}
            </div>
            <button
              onClick={handleAddBook}
              className="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 transition-colors"
            >
              <Plus className="h-4 w-4 mr-2" />
              Add Book
            </button>
          </div>
        </div>
      </header>

      {/* Main Content */}
      <main className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        {/* Error Message */}
        {state.error && (
          <div className="mb-6 bg-red-50 border border-red-200 rounded-md p-4">
            <div className="flex">
              <AlertCircle className="h-5 w-5 text-red-400" />
              <div className="ml-3">
                <h3 className="text-sm font-medium text-red-800">
                  Error
                </h3>
                <div className="mt-2 text-sm text-red-700">
                  {state.error}
                </div>
              </div>
            </div>
          </div>
        )}

        {/* Loading State */}
        {state.loading && (
          <div className="text-center py-12">
            <div className="inline-flex items-center px-4 py-2 font-semibold leading-6 text-sm shadow rounded-md text-white bg-blue-500 hover:bg-blue-400 transition ease-in-out duration-150 cursor-not-allowed">
              <svg className="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"></circle>
                <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              Loading books...
            </div>
          </div>
        )}

        {/* Books Grid */}
        {!state.loading && (
          <>
            {state.books.length === 0 ? (
              <div className="text-center py-12">
                <BookOpen className="mx-auto h-12 w-12 text-gray-400" />
                <h3 className="mt-2 text-sm font-medium text-gray-900">No books</h3>
                <p className="mt-1 text-sm text-gray-500">
                  Get started by adding your first book.
                </p>
                <div className="mt-6">
                  <button
                    onClick={handleAddBook}
                    className="inline-flex items-center px-4 py-2 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
                  >
                    <Plus className="h-4 w-4 mr-2" />
                    Add Book
                  </button>
                </div>
              </div>
            ) : (
              <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
                {state.books.map((book) => (
                  <BookCard
                    key={book.id}
                    book={book}
                    onEdit={handleEditBook}
                    onDelete={(id) => setShowDeleteConfirm(id)}
                    onView={handleViewBook}
                  />
                ))}
              </div>
            )}
          </>
        )}

        {/* Debug Information */}
        {isDebug && (
          <div className="mt-8 p-4 bg-gray-100 rounded-lg">
            <h3 className="text-sm font-medium text-gray-700 mb-2">Debug Information</h3>
            <div className="text-xs text-gray-600 space-y-1">
              <p>App Name: {appName}</p>
              <p>Version: {appVersion}</p>
              <p>API URL: {process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080'}</p>
              <p>Books Count: {state.books.length}</p>
              <p>Loading: {state.loading ? 'Yes' : 'No'}</p>
              <p>Error: {state.error || 'None'}</p>
            </div>
          </div>
        )}
      </main>

      {/* Book Form Modal */}
      <BookForm
        book={selectedBook}
        isOpen={isFormOpen}
        onClose={() => setIsFormOpen(false)}
        onSubmit={handleFormSubmit}
        mode={formMode}
      />

      {/* Delete Confirmation Modal */}
      {showDeleteConfirm && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
          <div className="bg-white rounded-lg p-6 w-full max-w-md mx-4">
            <h3 className="text-lg font-medium text-gray-900 mb-4">
              Confirm Delete
            </h3>
            <p className="text-sm text-gray-500 mb-6">
              Are you sure you want to delete this book? This action cannot be undone.
            </p>
            <div className="flex space-x-3">
              <button
                onClick={() => setShowDeleteConfirm(null)}
                className="flex-1 px-4 py-2 border border-gray-300 rounded-md text-gray-700 hover:bg-gray-50 transition-colors"
              >
                Cancel
              </button>
              <button
                onClick={() => handleDeleteBook(showDeleteConfirm)}
                className="flex-1 px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-700 transition-colors"
              >
                Delete
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
} 