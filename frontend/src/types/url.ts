export interface URLRequest {
  url: string;
  operation: 'canonical' | 'redirection' | 'all';
}

export interface URLResponse {
  processed_url: string;
} 