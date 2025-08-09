'use client';

import React, { useEffect, useState } from 'react';
import { useParams, useRouter } from 'next/navigation';
import { bookApi } from '@/utils/api';
import type { Book } from '@/types/book';
import { ArrowLeft } from 'lucide-react';

export default function BookDetailPage() {
  const params = useParams();
  const router = useRouter();
  const { id } = (params || {}) as { id: string };

  const [book, setBook] = useState<Book | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    let isMounted = true;
    async function fetchBook() {
      if (!id) return;
      try {
        setLoading(true);
        setError(null);
        const data = await bookApi.getById(id);
        if (isMounted) setBook(data);
      } catch (e) {
        const message = e instanceof Error ? e.message : 'Failed to load book';
        if (isMounted) setError(message);
      } finally {
        if (isMounted) setLoading(false);
      }
    }
    fetchBook();
    return () => {
      isMounted = false;
    };
  }, [id]);

  return (
    <div className="min-h-screen bg-gray-50">
      <header className="bg-white shadow-sm border-b">
        <div className="max-w-5xl mx-auto px-4 sm:px-6 lg:px-8 h-16 flex items-center">
          <button
            onClick={() => router.back()}
            className="inline-flex items-center text-sm text-gray-700 hover:text-gray-900"
          >
            <ArrowLeft className="h-4 w-4 mr-2" /> Back
          </button>
        </div>
      </header>

      <main className="max-w-5xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        {loading && (
          <div className="text-center text-gray-600">Loading book details...</div>
        )}
        {error && (
          <div className="mb-6 bg-red-50 border border-red-200 rounded-md p-4 text-sm text-red-700">
            {error}
          </div>
        )}
        {!loading && !error && book && (
          <div className="bg-white rounded-lg shadow p-6">
            <h1 className="text-2xl font-bold text-gray-900 mb-4">{book.title}</h1>
            <div className="space-y-3">
              <div>
                <p className="text-sm text-gray-500">Author</p>
                <p className="text-gray-900 font-medium">{book.author}</p>
              </div>
              <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
                <div>
                  <p className="text-sm text-gray-500">Year</p>
                  <p className="text-gray-900 font-medium">{book.year}</p>
                </div>
                <div>
                  <p className="text-sm text-gray-500">ISBN</p>
                  <p className="text-gray-900 font-medium">{book.isbn}</p>
                </div>
              </div>
              <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
                <div>
                  <p className="text-sm text-gray-500">Created</p>
                  <p className="text-gray-900 font-medium">{new Date(book.created_at).toLocaleString()}</p>
                </div>
                <div>
                  <p className="text-sm text-gray-500">Updated</p>
                  <p className="text-gray-900 font-medium">{new Date(book.updated_at).toLocaleString()}</p>
                </div>
              </div>
            </div>
          </div>
        )}
      </main>
    </div>
  );
} 