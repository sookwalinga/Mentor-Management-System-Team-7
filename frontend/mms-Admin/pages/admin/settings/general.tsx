import { Button, InputField, SelectField, TextareaField } from "@/components";
import SettingsLayout from "@/components/Settings/SettingsLayout";
import { DashboardLayout } from "@/components/layouts/dashboard-layout";
import { avatarIcon } from "@/public";
import Image from "next/image";
import React, { ReactElement } from "react";

const Profile = () => {
  return (
    <SettingsLayout>
      <div className="border border-[#E6E6E6] rounded-md p-3">
        <div className="flex items-center">
          <div className="w-[42px] h-[42px] rounded-full cursor-pointer ">
            <Image src={avatarIcon} alt="avatar" />
          </div>
          <div className="mx-9">
            <h3 className="text-mmsBlack2 font-semibold text-xl">
              Set Profile Picture
            </h3>
            <Button variant="primary" className="py-1 px-[15px] text-xs">
              Upload Picture
            </Button>
          </div>
        </div>
        <div className="mt-5">
          <div className="form-group flex items-center justify-between">
            <span className="font-semibold text-base text-mmsBlack2">Full Name</span>
            <div className="ml-6 flex items-center">
            <InputField 
            label=""
            placeholder="First Name"
            id="firstName"
            type="text"
            className="w-[326px]"
            />
            <InputField 
            label=""
            placeholder="Last Name"
            id="lastName"
            type="text"
            className="md:w-[326px] ml-6"
            />
            </div>
          </div>
          <div className="form-group flex items-center justify-between">
            <span className="font-semibold text-base text-mmsBlack2">About</span>
            <div className="ml-6 flex items-center">
            <TextareaField 
            label=""
            placeholder="Your Bio"
            id="about"
            className="md:w-[679px]"
            type="text"
            />
            </div>
          </div>
          <div className="form-group flex items-center justify-between">
            <span className="font-semibold text-base text-mmsBlack2">Full Website</span>
            <div className="ml-6 flex items-center">
            <InputField 
            label=""
            placeholder="www.example.com"
            id="website"
            type="text"
            className="md:w-[679px]"
            />
            </div>
          </div>
          <div className="form-group flex items-center">
            <div className="flex items-center">
            <span className="font-semibold mr-20 text-base text-mmsBlack2">Country</span>
            <SelectField 
            label=""
            id="country"
            type="text"
            className="w-[274px]"
            >
              <option value="">Select Country</option>
            </SelectField>
            </div>
            <div className="ml-6 flex items-center">
            <span className="font-semibold text-base text-mmsBlack2 mx-9">City </span>
            <SelectField 
            label=""
            id="city"
            className="w-[274px]"
            >
              <option value="">Select City</option>
            </SelectField>
            </div>
          </div>
        </div>
      </div>
    </SettingsLayout>
  );
};

export default Profile;

Profile.getLayout = function getLayout(page: ReactElement) {
  return <DashboardLayout>{page}</DashboardLayout>;
};
