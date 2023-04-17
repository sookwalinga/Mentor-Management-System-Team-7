import SettingsLayout from "@/components/Settings/SettingsLayout";
import { DashboardLayout } from "@/components/layouts/dashboard-layout";
import React, { ReactElement } from "react";

const Notifications = () => {
  return <div>Notifications</div>;
};

export default Notifications;

Notifications.getLayout = function getLayout(page: ReactElement) {
  return (
    <DashboardLayout>
      <SettingsLayout title="Notifications">{page}</SettingsLayout>
    </DashboardLayout>
  );
};
