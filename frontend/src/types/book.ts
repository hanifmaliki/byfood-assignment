export interface Book {
  id: string;
  title: string;
  author: string;
  year: number;
  isbn: string;
  created_at: string;
  updated_at: string;
}

export interface CreateBookRequest {
  title: string;
  author: string;
  year: number;
  isbn: string;
}

export interface UpdateBookRequest {
  title?: string;
  author?: string;
  year?: number;
  isbn?: string;
} 