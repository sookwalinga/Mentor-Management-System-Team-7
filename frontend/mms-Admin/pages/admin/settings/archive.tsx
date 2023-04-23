import SettingsLayout from "@/components/Settings/SettingsLayout";
import { DashboardLayout } from "@/components/layouts/dashboard-layout";
import React, { ReactElement } from "react";

const Archive = () => {
  return <div>Archive</div>;
};

export default Archive;

Archive.getLayout = function getLayout(page: ReactElement) {
  return (
    <DashboardLayout>
      <SettingsLayout title="Archive">{page}</SettingsLayout>
    </DashboardLayout>
  );
};

Archive.requireAuth = true;