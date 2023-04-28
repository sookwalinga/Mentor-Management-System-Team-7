import SettingsLayout from "@/components/Settings/SettingsLayout";
import { DashboardLayout } from "@/components/layouts/dashboard-layout";
import React, { ReactElement } from "react";

const FAQ = () => {
  return <div>FAQ</div>;
};

export default FAQ;

FAQ.getLayout = function getLayout(page: ReactElement) {
  return (
    <DashboardLayout>
      <SettingsLayout title="FAQ">{page}</SettingsLayout>
    </DashboardLayout>
  );
};

FAQ.requireAuth = true;