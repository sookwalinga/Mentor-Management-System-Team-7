import React, { ReactNode } from "react";
import SettingsSidebar from "./SettingsSidebar";

interface settingsLayout {
  children: ReactNode;
}
const SettingsLayout = ({ children }: settingsLayout) => {
  return (
    <div className="px-[28px]">
      <h1 className="text-2xl font-semibold mb-[0.75rem] text-mmsBlack1">
        Settings
      </h1>
      <div className="flex items-start">
        <SettingsSidebar />
        <div className="mx-7 w-[840px]">{children}</div>
      </div>
    </div>
  );
};

export default SettingsLayout;
