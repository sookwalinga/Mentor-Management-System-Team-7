import React from "react";
import { Header, Sidebar } from "..";
import Head from "next/head";

interface dashboardLayout {
  children?: React.ReactNode;
  title?: string;
}

export const DashboardLayout = ({ children, title }: dashboardLayout) => {
  return (
    <div className="flex h-screen">
      {/* Sidebar */}
      <Sidebar />
      <Head>
        <title>{title}</title>
      </Head>

      {/* Content */}
      <div className="flex flex-col w-full">
        {/* Header */}
        <Header />
        {/* Main content */}
        <main className="flex-1 overflow-y-auto pl-72 pr-[61px] mt-32">
          {title && (
            <div className="flex items-center justify-between">
              <h1 className="text-2xl font-semibold mb-[0.75rem] text-mmsBlack1">
                {title}
              </h1>
            </div>
          )}

          {children}
        </main>
      </div>
    </div>
  );
};
