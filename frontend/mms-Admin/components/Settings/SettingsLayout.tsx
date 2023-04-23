import React, { ReactNode } from "react";
import SettingsSidebar from "./SettingsSidebar";
import Head from "next/head";

interface settingsLayout {
  children: ReactNode;
  title?: string;
}
const SettingsLayout = ({ children, title }: settingsLayout) => {
  return (
    <div className="px-[28px]">
      <Head>
        <title>Settings | {title}</title>
      </Head>
      <h1 className="text-2xl font-semibold mb-[0.75rem] text-mmsBlack1">
        Settings
      </h1>
      <div className="flex lg:flex-row flex-col items-start">
        <SettingsSidebar />
        <div className="lg:mx-7 lg:w-[840px] w-full my-10 lg:m-0">{children}</div>
      </div>
    </div>
  );
};

export default SettingsLayout;


SettingsLayout.requireAuth = true;