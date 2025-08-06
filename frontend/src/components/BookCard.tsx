'use client';

import React from 'react';
import { Book } from '@/types/book';
import { Edit, Trash2, Eye } from 'lucide-react';

interface BookCardProps {
  book: Book;
  onEdit: (book: Book) => void;
  onDelete: (id: string) => void;
  onView: (book: Book) => void;
}

export const BookCard: React.FC<BookCardProps> = ({
  book,
  onEdit,
  onDelete,
  onView,
}) => {
  return (
    <div className="bg-white rounded-lg shadow-md p-6 hover:shadow-lg transition-shadow">
      <div className="flex justify-between items-start mb-4">
        <h3 className="text-xl font-semibold text-gray-900 truncate">
          {book.title}
        </h3>
        <div className="flex space-x-2">
          <button
            onClick={() => onView(book)}
            className="p-2 text-blue-600 hover:bg-blue-50 rounded-full transition-colors"
            title="View details"
          >
            <Eye size={16} />
          </button>
          <button
            onClick={() => onEdit(book)}
            className="p-2 text-green-600 hover:bg-green-50 rounded-full transition-colors"
            title="Edit book"
          >
            <Edit size={16} />
          </button>
          <button
            onClick={() => onDelete(book.id)}
            className="p-2 text-red-600 hover:bg-red-50 rounded-full transition-colors"
            title="Delete book"
          >
            <Trash2 size={16} />
          </button>
        </div>
      </div>
      
      <div className="space-y-2">
        <p className="text-gray-600">
          <span className="font-medium">Author:</span> {book.author}
        </p>
        <p className="text-gray-600">
          <span className="font-medium">Year:</span> {book.year}
        </p>
        <p className="text-gray-600">
          <span className="font-medium">ISBN:</span> {book.isbn}
        </p>
      </div>
      
      <div className="mt-4 pt-4 border-t border-gray-200">
        <p className="text-xs text-gray-500">
          Added: {new Date(book.created_at).toLocaleDateString()}
        </p>
      </div>
    </div>
  );
}; 