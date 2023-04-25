import SettingsLayout from "@/components/Settings/SettingsLayout";
import { Toggle } from "@/components/Toggle";
import { DashboardLayout } from "@/components/layouts/dashboard-layout";
import React, { ReactElement } from "react";

const Notifications = () => {
  return (
    <div className="border border-[#E6E6E6] rounded-md p-5 text-mmsBlack2">
      <h3 className="font-semibold text-xl text-mmsBlack2">
        General Notifications
      </h3>
      <div className="flex justify-end">
        <div className="font-semibold text-lg text-mmsBlack2 flex items-center gap-3">
          <h2>E-mail</h2>
          <h2>In-app</h2>
        </div>
      </div>
      <div className="mt-3">
        <div className="flex items-center justify-between my-3">
          <p className="text-xl">All Notifications</p>
          <div className="flex items-center gap-8 mx-3">
            <Toggle />
            <Toggle />
          </div>
        </div>
        <div className="flex items-center justify-between my-3">
          <p className="text-xl">Programs</p>
          <div className="flex items-center gap-8 mx-3">
            <Toggle />
            <Toggle />
          </div>
        </div>
        <div className="flex items-center justify-between my-3">
          <p className="text-xl">Task</p>
          <div className="flex items-center gap-8 mx-3">
            <Toggle />
            <Toggle />
          </div>
        </div>
        <div className="flex items-center justify-between my-3">
          <p className="text-xl">Approval Request</p>
          <div className="flex items-center gap-8 mx-3">
            <Toggle />
            <Toggle />
          </div>
        </div>
        <div className="flex items-center justify-between my-3">
          <p className="text-xl">Reports</p>
          <div className="flex items-center gap-8 mx-3">
            <Toggle />
            <Toggle />
          </div>
        </div>
      </div>
      <h3 className="font-semibold text-xl text-mmsBlack2 mt-[40px]">
        Discussion Notifications
      </h3>
      <div className="flex justify-end">
        <div className="font-semibold text-lg text-mmsBlack2 flex items-center gap-3">
          <h2>E-mail</h2>
          <h2>In-app</h2>
        </div>
      </div>
      <div className="mt-3">
        <div className="flex items-center justify-between my-3">
          <p className="text-xl">Comments on my posts</p>
          <div className="flex items-center gap-8 mx-3">
            <Toggle />
            <Toggle />
          </div>
        </div>
        <div className="flex items-center justify-between my-3">
          <p className="text-xl">Posts</p>
          <div className="flex items-center gap-8 mx-3">
            <Toggle />
            <Toggle />
          </div>
        </div>
        <div className="flex items-center justify-between my-3">
          <p className="text-xl">Comments</p>
          <div className="flex items-center gap-8 mx-3">
            <Toggle />
            <Toggle />
          </div>
        </div>
        <div className="flex items-center justify-between my-3">
          <p className="text-xl">Mentions</p>
          <div className="flex items-center gap-8 mx-3">
            <Toggle />
            <Toggle />
          </div>
        </div>
        <div className="flex items-center justify-between my-3">
          <p className="text-xl">Directs Message</p>
          <div className="flex items-center gap-8 mx-3">
            <Toggle />
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

Notifications.requireAuth = true;