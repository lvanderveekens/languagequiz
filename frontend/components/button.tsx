import React, { ButtonHTMLAttributes } from 'react';

interface Props extends ButtonHTMLAttributes<HTMLButtonElement> {
  variant: 'primary-dark' | 'secondary-dark' | 'primary-light' | 'secondary-light';
}

const Button: React.FC<Props> = ({ variant, className, ...props }) => {
  let variantClasses = '';
  switch (variant) {
    case 'primary-dark':
      variantClasses = 'group text-white enabled:hover:text-[#003259] enabled:hover:bg-white bg-[#003259] font-bold px-4 py-2 border-2 border-[#003259] rounded-lg px-3 disabled:opacity-50';
      break;
    case 'secondary-dark':
      variantClasses = 'group text-[#003259] enabled:hover:text-white enabled:hover:bg-[#003259] font-bold px-4 py-2 border-2 border-[#003259] rounded-lg px-3 disabled:opacity-50';
      break;
    case 'primary-light':
      variantClasses = 'group text-[#003259] font-bold bg-white enabled:hover:bg-[#003259] enabled:hover:text-white border-2 border-white px-4 py-2 rounded-lg px-3 disabled:opacity-50';
      break;
    case 'secondary-light':
      variantClasses = 'group font-bold px-4 py-2 border-2 text-white border-white enabled:hover:bg-white enabled:hover:text-[#003259] rounded-lg px-3 disabled:opacity-50';
      break;
  }
  
  return (
    <button
      className={`${variantClasses} ${className}`}
      {...props}
    >
      {props.children}
    </button>
  );
};

export default Button;