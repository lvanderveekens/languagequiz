import Link from 'next/link';
import React, { ButtonHTMLAttributes } from 'react';

interface Props {
  className?: string
}

const Navbar: React.FC<Props> = ({ className }) => {
  return (
    <nav className={`py-4 b-4 bg-[#003259] ${className}`}>
      <div className="container text-xl text-white flex justify-between align-center">
        <div className="text-2xl font-bold self-center">
          <Link href="/">LanguageQuiz</Link>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;