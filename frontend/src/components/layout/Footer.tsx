import React from 'react';

const footerStyle: React.CSSProperties = {
  backgroundColor: 'var(--color-light-cream)',
  color: 'var(--color-gray-500)',
  padding: 'var(--spacing-md) var(--spacing-lg)',
  textAlign: 'center',
  fontSize: 'var(--font-size-sm)',
  marginTop: 'auto', // main要素の残りのスペースを埋める
};

export const Footer: React.FC = () => {
  const currentYear = new Date().getFullYear();

  return (
    <footer style={footerStyle}>
      <p>&copy; {currentYear} meal-compass. All Rights Reserved.</p>
    </footer>
  );
};