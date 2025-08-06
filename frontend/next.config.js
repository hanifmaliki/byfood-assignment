/** @type {import('next').NextConfig} */
const nextConfig = {
  experimental: {
    appDir: true,
  },
  async rewrites() {
    const apiUrl = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';
    const apiPrefix = process.env.NEXT_PUBLIC_API_PREFIX || '/api';
    
    return [
      {
        source: `${apiPrefix}/:path*`,
        destination: `${apiUrl}${apiPrefix}/:path*`,
      },
    ]
  },
  env: {
    CUSTOM_KEY: process.env.NEXT_PUBLIC_CUSTOM_KEY,
  },
}

module.exports = nextConfig 