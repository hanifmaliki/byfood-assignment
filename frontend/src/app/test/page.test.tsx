import React from 'react'
import { render, screen } from '@testing-library/react'
import '@testing-library/jest-dom'

// Simple test to verify testing setup is working
describe('Testing Setup', () => {
  it('should have proper testing environment', () => {
    expect(1 + 1).toBe(2)
  })

  it('should have React Testing Library available', () => {
    expect(typeof render).toBe('function')
    expect(typeof screen).toBe('object')
  })

  it('should be able to render a simple component', () => {
    const TestComponent = () => <div>Test Component</div>
    render(<TestComponent />)
    expect(screen.getByText('Test Component')).toBeInTheDocument()
  })
}) 