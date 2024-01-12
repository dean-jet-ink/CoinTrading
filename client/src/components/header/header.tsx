import Link from "next/link";
import React, { ReactNode } from "react";

export type HeaderProps = {
  children: ReactNode;
};

const Header = ({ children }: HeaderProps) => {
  return (
    <header className="py-8 px-14 flex items-center justify-between fixed">
      <h1 className="text-2xl font-semibold">
        <Link href="/">Coin Trading</Link>
      </h1>
      {children}
    </header>
  );
};

export default Header;
