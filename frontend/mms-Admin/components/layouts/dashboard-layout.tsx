/* eslint-disable prettier/prettier */
import React from "react";
import { Header, Sidebar } from "..";
import Head from "next/head";
import { useSession } from "next-auth/react";

interface dashboardLayout {
  children?: React.ReactNode;
  title?: string;
}

export const DashboardLayout = ({ children, title }: dashboardLayout) => {


  return (
    <div className="flex h-screen">
      <Head>
        <title>{title}</title>
      </Head>
      {/* Sidebar */}
      <Sidebar />
      {/* Content */}
      <div className="flex flex-col w-full bg-white">
        {/* Header */}
        <Header />
        {/* Main content */}
        <main className="flex-1 overflow-y-auto lg:pl-72 p-6 lg:pr-[61px] mt-32">
          {title && (
            <div className="flex items-center justify-between mb-[12px] ">
              <h1 className="text-2xl font-semibold mb-[0.75rem] text-mmsBlack1">
                {title}
              </h1>

              {title === "Dashboard" && (
                <select className="bg-green11 border border-mmsPry10 rounded-[5px] w-[122px] h-[38px] text-mmsBlack2 outline-none text-lg font-normal pl-[10px]">
                  <option value="volvo">This Week </option>
                </select>
              )}
            </div>
          )}
          {children}
        </main>
      </div>
    </div>
  );
};


