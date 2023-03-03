// import React from 'react';
import { fireEvent, render, screen, waitFor } from '@testing-library/react';
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';

import App from './App';

import { shorten, ShortenStatus } from "./api/shorten";

vi.mock("./api/shorten", () => ({ shorten: vi.fn() }));



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
    vi.mocked(shorten).mockImplementationOnce(async () => ({ status: 'created' as ShortenStatus }))

    render(<App />);
    const namespaceInput: HTMLInputElement = screen.getByLabelText('Namespace');
    fireEvent.change(namespaceInput, { target: { value: "test" } });
    fireEvent.click(screen.getByText('Next'));
    const urlInput: HTMLInputElement = screen.getByLabelText('Target URL');
    const segmentInput: HTMLInputElement = screen.getByLabelText('Alias');
    fireEvent.change(urlInput, { target: { value: "https://vitest.dev" } });
    fireEvent.change(segmentInput, { target: { value: "vitest" } });
    
    fireEvent.click(screen.getByText('Next'));
    
    await waitFor(() => expect(screen.getByText("Here's your new url")).toBeDefined())
    await waitFor(() => expect(screen.getByText("http://localhost:8000/test/vitest")).toBeDefined())
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

  it('should render generic error when server responds with error', async () => {
    vi.mocked(shorten).mockImplementationOnce(async () => Promise.resolve({ status: 'serverError' as ShortenStatus }))

    render(<App />);

    const namespaceInput: HTMLInputElement = screen.getByLabelText('Namespace');
    fireEvent.change(namespaceInput, { target: { value: "test" } });
    fireEvent.click(screen.getByText('Next'));
    const urlInput: HTMLInputElement = screen.getByLabelText('Target URL');
    const segmentInput: HTMLInputElement = screen.getByLabelText('Alias');
    fireEvent.change(urlInput, { target: { value: "https://vitest.dev" } });
    fireEvent.change(segmentInput, { target: { value: "vitest" } });
    
    fireEvent.click(screen.getByText('Next'));

    await waitFor(() => expect(screen.getByText("Internal error.", { exact: false })).toBeDefined())
  })

  it('should render generic error when server responds with already exists status', async () => {
    vi.mocked(shorten).mockImplementationOnce(async () => Promise.resolve({ status: 'alreadyExists' as ShortenStatus }))

    render(<App />);

    const namespaceInput: HTMLInputElement = screen.getByLabelText('Namespace');
    fireEvent.change(namespaceInput, { target: { value: "test" } });
    fireEvent.click(screen.getByText('Next'));
    const urlInput: HTMLInputElement = screen.getByLabelText('Target URL');
    const segmentInput: HTMLInputElement = screen.getByLabelText('Alias');
    fireEvent.change(urlInput, { target: { value: "https://vitest.dev" } });
    fireEvent.change(segmentInput, { target: { value: "vitest" } });
    
    fireEvent.click(screen.getByText('Next'));

    await waitFor(() => expect(screen.getByText("Url is already taken. Please enter another alias.")).toBeDefined())
  })
});