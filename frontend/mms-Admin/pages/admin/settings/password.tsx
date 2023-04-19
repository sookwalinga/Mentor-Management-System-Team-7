import { Button, InputField } from "@/components";
import SettingsLayout from "@/components/Settings/SettingsLayout";
import { DashboardLayout } from "@/components/layouts/dashboard-layout";
import Link from "next/link";
import React, { ReactElement } from "react";

const Password = () => {
  return (
    <div className="border border-[#E6E6E6] rounded-md p-3">
      <div className="mt-5">
        <div className="form-group flex items-center justify-between">
          <span className="font-semibold text-base text-mmsBlack2">
            Current Password
          </span>
          <div className="ml-6 flex items-center">
            <InputField
              label=""
              placeholder="your current password"
              id="currentPassword"
              type="password"
              className="w-full md:w-[593px] bg-white"
            />
          </div>
        </div>
        <div className="form-group flex items-center justify-between">
          <span className="font-semibold text-base text-mmsBlack2">
            New Password
          </span>
          <div className="ml-6 flex items-center">
            <InputField
              label=""
              placeholder="Must be at least 8 characters"
              id="newPassword"
              type="password"
              className="w-full md:w-[593px] bg-white"
            />
          </div>
        </div>
        <div className="form-group flex items-center justify-between">
          <span className="font-semibold text-base text-mmsBlack2">
            Confirm New Password
          </span>
          <div className="ml-6 flex items-center">
            <InputField
              label=""
              placeholder="Must match your new password"
              id="confirmNewPassword"
              type="password"
              className="w-full md:w-[593px] bg-white"
            />
          </div>
        </div>
        <div className="flex justify-end mt-5">
          <Button variant="primary" className="text-base px-4 py-2">
            Save New Password
          </Button>
        </div>
        <div className="text-center mt-8">
          <Link href="/#" passHref>
            <span className="font-semibold text-[#023C40]">
              Forgot Password ?
            </span>
          </Link>
        </div>
      </div>
    </div>
  );
};

export default Password;

Password.getLayout = function getLayout(page: ReactElement) {
  return (
    <DashboardLayout>
      <SettingsLayout title="Password">{page}</SettingsLayout>
    </DashboardLayout>
  );
};
