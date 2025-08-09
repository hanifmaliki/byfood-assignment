import type { Metadata } from 'next'
import './globals.css'
import { BookProvider } from '@/contexts/BookContext'

export const metadata: Metadata = {
  title: 'Library Management System',
  description: 'Manage books with Next.js (TypeScript) and Golang backend',
}

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en">
      <body>
        <BookProvider>{children}</BookProvider>
      </body>
    </html>
  )
} 