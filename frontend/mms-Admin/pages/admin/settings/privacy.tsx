import SettingsLayout from "@/components/Settings/SettingsLayout";
import { Toggle } from "@/components/Toggle";
import { DashboardLayout } from "@/components/layouts/dashboard-layout";
import React, { ReactElement } from "react";

const Notifications = () => {
  return (
    <div className="border border-[#E6E6E6] rounded-md p-5 text-mmsBlack2">
      <div className="flex justify-end"></div>
      <div className="mt-3 lg:w-72">
        <div className="flex items-center justify-between my-3">
          <p className="text-base font-semibold  ">Show contact info</p>
          <div className="flex items-center gap-8 mx-3">
            <Toggle />
          </div>
        </div>
        <div className="flex items-center justify-between my-3">
          <p className="text-base font-semibold ">Show GitHub</p>
          <div className="flex items-center gap-8 mx-3">
            <Toggle />
          </div>
        </div>
        <div className="flex items-center justify-between my-3">
          <p className="text-base font-semibold ">Show Instagram</p>
          <div className="flex items-center gap-8 mx-3">
            <Toggle />
          </div>
        </div>

        <div className="flex items-center justify-between my-3">
          <p className="text-base font-semibold ">Show Linkdein </p>
          <div className="flex items-center gap-8 mx-3">
            <Toggle />
          </div>
        </div>

        <div className="flex items-center justify-between my-3">
          <p className="text-base font-semibold ">Show Twitter </p>
          <div className="flex items-center gap-8 mx-3">
            <Toggle />
          </div>
        </div>
      </div>
    </div>
  );
};

export default Notifications;

Notifications.getLayout = function getLayout(page: ReactElement) {
  return (
    <DashboardLayout>
      <SettingsLayout title="Notifications">{page}</SettingsLayout>
    </DashboardLayout>
  );
};
