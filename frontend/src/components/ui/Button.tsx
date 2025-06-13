import React from 'react';
import styles from './Button.module.css';

// React.ComponentPropsWithRef<'button'> を使うことで、
// <button>が受け取る全ての属性（onClick, disabled, classNameなど）とrefを型安全に扱える
type ButtonProps = React.ComponentPropsWithRef<'button'> & {
  variant?: 'primary' | 'secondary';
  isLoading?: boolean;
};

export const Button = React.forwardRef<HTMLButtonElement, ButtonProps>(
  ({ children, className, variant = 'primary', isLoading = false, ...props }, ref) => {
    
    // 複数のCSSクラスを結合するためのヘルパー
    const buttonClasses = [
      styles.button,
      styles[variant],
      className
    ].filter(Boolean).join(' ');

    return (
      <button
        ref={ref}
        className={buttonClasses}
        disabled={isLoading || props.disabled}
        {...props}
      >
        {isLoading ? '処理中...' : children}
      </button>
    );
  }
);

Button.displayName = 'Button'; // デバッグ時の表示名を定義