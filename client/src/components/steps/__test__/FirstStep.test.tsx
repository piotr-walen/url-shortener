// import React from 'react';
import { render, screen } from '@testing-library/react';
import { FirstStep } from '../FirstStep';
import { expect, test, vi } from 'vitest';

const clickHandler = vi.fn()

test('render text', () => {
  render(<FirstStep namespace="" onNamespaceChange={clickHandler} namespaceError=""/>);
  expect(screen.getByText(`"Log in" to your namespace`)).toBeDefined();
});