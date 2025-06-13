import React from 'react';
import { Link } from 'react-router-dom';

const headerStyle: React.CSSProperties = {
  backgroundColor: 'var(--color-dark-brown)',
  color: 'var(--color-white)',
  padding: 'var(--spacing-md) var(--spacing-lg)',
  boxShadow: 'var(--box-shadow-md)',
  textAlign: 'center',
};

const titleStyle: React.CSSProperties = {
  fontSize: 'var(--font-size-lg)',
  fontWeight: 'var(--font-weight-bold)',
  textDecoration: 'none',
  color: 'inherit', // 親要素の色を継承
};

export const Header: React.FC = () => {
  return (
    <header style={headerStyle}>
      <Link to="/" style={titleStyle}>
        meal-compass
      </Link>
    </header>
  );
};