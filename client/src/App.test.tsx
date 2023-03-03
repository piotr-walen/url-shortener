// import React from 'react';
import { fireEvent, render, screen, waitFor } from '@testing-library/react';
import { describe, expect, it, vi } from 'vitest';

import App from './App';

vi.mock("./api/shorten", async () => {
  return {
    shorten: vi.fn().mockImplementation(async () => ({ status: 201 })),
  };
})

describe("App", () => {
  it('should render "Log in" text on initial view', () => {
    render(<App />);
    
    expect(screen.getByText(`"Log in" to your namespace`)).toBeDefined();
    expect(screen.getByText('Namespace is required')).toBeDefined();
  });

  it('should allow to enter namespace', () => {
    render(<App />);
    const namespaceInput: HTMLInputElement = screen.getByLabelText('Namespace');
    
    fireEvent.change(namespaceInput, { target: { value: "test" } });
   
    expect(namespaceInput.value).toBe("test");
  });

  it('should allow to proceed to next step when valid namespace is entered', () => {
    render(<App />);
    const namespaceInput: HTMLInputElement = screen.getByLabelText('Namespace');
    fireEvent.change(namespaceInput, { target: { value: "test" } });
    
    fireEvent.click(screen.getByText('Next'));
    
    expect(screen.getByText('Enter URL that you would like to create alias for')).toBeDefined();
    expect(screen.getByText('Target URL is required')).toBeDefined();
    expect(screen.getByText('Alias is required')).toBeDefined();
  });

  it('should allow to submit when all data is entered correctly', async () => {
    render(<App />);
    const namespaceInput: HTMLInputElement = screen.getByLabelText('Namespace');
    fireEvent.change(namespaceInput, { target: { value: "test" } });
    fireEvent.click(screen.getByText('Next'));
    const urlInput: HTMLInputElement = screen.getByLabelText('Target URL');
    const segmentInput: HTMLInputElement = screen.getByLabelText('Alias');
    fireEvent.change(urlInput, { target: { value: "https://vitest.dev" } });
    fireEvent.change(segmentInput, { target: { value: "vitest" } });
    
    fireEvent.click(screen.getByText('Next'));
    
    await waitFor(() => {
      expect(screen.getByText("Here's your new url")).toBeDefined()
      expect(screen.getByText("http://localhost:8000/test/vitest")).toBeDefined()
    });
  });

  it('should render error when namespace invalid', () => {
    render(<App />);
    const namespaceInput: HTMLInputElement = screen.getByLabelText('Namespace');
    
    fireEvent.change(namespaceInput, { target: { value: "//" } });
    
    expect(namespaceInput.value).toBe("//");
    expect(screen.getByText('Namespace field can only contain alphanumeric or "-", "_" characters')).toBeDefined();
  });

  it('should render error when invalid url or alias is entered', () => {
    render(<App />);
    const namespaceInput: HTMLInputElement = screen.getByLabelText('Namespace');
    fireEvent.change(namespaceInput, { target: { value: "test" } });
    fireEvent.click(screen.getByText('Next'));
    const urlInput: HTMLInputElement = screen.getByLabelText('Target URL');
    const aliasInput: HTMLInputElement = screen.getByLabelText('Alias');
    
    fireEvent.change(urlInput, { target: { value: "//" } });
    fireEvent.change(aliasInput, { target: { value: "$#@" } });
    
    expect(screen.getByText('Target URL should contain valid http or https URL')).toBeDefined();
    expect(screen.getByText('Alias field can only contain alphanumeric or "-", "_" characters')).toBeDefined();
  });

  it('should not allow to proceed to second step when invalid namespace is entered', () => {
    render(<App />);
    const namespaceInput: HTMLInputElement = screen.getByLabelText('Namespace');
    fireEvent.change(namespaceInput, { target: { value: "//" } });
   
    fireEvent.click(screen.getByText('Next'));
    
    expect(screen.getByText(`"Log in" to your namespace`)).toBeDefined();
  });

  it('should not allow to proceed to third step when invalid alias or target url is entered', () => {
    render(<App />);
    const namespaceInput: HTMLInputElement = screen.getByLabelText('Namespace');
    fireEvent.change(namespaceInput, { target: { value: "test" } });
    fireEvent.click(screen.getByText('Next'));
    const urlInput: HTMLInputElement = screen.getByLabelText('Target URL');
    fireEvent.change(urlInput, { target: { value: "" } });

    fireEvent.click(screen.getByText('Next'));

    expect(screen.getByText('Target URL is required')).toBeDefined();
  });
});