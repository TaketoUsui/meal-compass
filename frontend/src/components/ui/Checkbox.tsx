import React, { useId } from 'react';
import styles from './Checkbox.module.css';

type CheckboxProps = Omit<React.ComponentPropsWithRef<'input'>, 'type'> & {
  label: React.ReactNode;
};

export const Checkbox = React.forwardRef<HTMLInputElement, CheckboxProps>(
  ({ label, className, ...props }, ref) => {
    // コンポーネント内で一意なIDを生成し、labelとinputを関連付ける
    const id = useId();
    
    const containerClasses = [styles.container, className].filter(Boolean).join(' ');

    return (
      <div className={containerClasses}>
        <input
          ref={ref}
          type="checkbox"
          id={id}
          className={styles.input}
          {...props}
        />
        <label htmlFor={id} className={styles.label}>
          {label}
        </label>
      </div>
    );
  }
);

Checkbox.displayName = 'Checkbox';