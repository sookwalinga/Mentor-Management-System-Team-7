import { Button, InputField, SelectField, TextareaField } from "@/components";
import Input from "@/components/InputFields";
import SettingsLayout from "@/components/Settings/SettingsLayout";
import { avatarIcon } from "@/public";
import Image from "next/image";
import React, { ReactElement } from "react";

const General = () => {
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
            <Button variant="primary" className="py-2 px-[18px] text-sm">
              Upload Picture
            </Button>
          </div>
        </div>
        <div className="mt-5">
          <div className="form-group flex items-center justify-between">
            <span className="font-semibold text-base text-mmsBlack2">
              Full Name
            </span>
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
            <span className="font-semibold text-base text-mmsBlack2">
              About
            </span>
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
            <span className="font-semibold text-base text-mmsBlack2">
              Full Website
            </span>
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
              <span className="font-semibold mr-[23%] text-base text-mmsBlack2">
                Country
              </span>
              <SelectField
                label=""
                id="country"
                type="text"
                className="w-[294px]"
              >
                <option value="">Select Country</option>
              </SelectField>
            </div>
            <div className="ml-[13%] flex items-center">
              <span className="font-semibold text-base text-mmsBlack2 mr-9">
                City{" "}
              </span>
              <SelectField label="" id="city" className="w-[294px]">
                <option value="">Select City</option>
              </SelectField>
            </div>
          </div>
          <div className="form-group flex items-start">
            <span className="font-semibold text-base text-mmsBlack2">
              Social
            </span>
            <div className="ml-[11%] grid grid-cols-2 gap-x-8 gap-y-4 items-center">
              <Input
                prefix={
                  <div className="flex items-center">
                    {" "}
                    <Image
                      src="/images/svgs/githubIcon.svg"
                      height={24}
                      width={24}
                      alt="github"
                    />{" "}
                    <span className="ml-2"> Github</span>
                  </div>
                }
                label=""
                placeholder="@githubuser"
                id="githubUser"
                type="text"
                className="w-[294px]"
              />
              <Input
                prefix={
                  <div className="flex items-center">
                    {" "}
                    <Image
                      src="/images/svgs/instagramIcon.svg"
                      height={24}
                      width={24}
                      alt="github"
                    />{" "}
                    <span className="ml-2"> Github</span>
                  </div>
                }
                label=""
                placeholder="@githubuser"
                id="githubUser"
                type="text"
                className="w-[294px] ml-8"
              />
              <Input
                prefix={
                  <div className="flex items-center">
                    {" "}
                    <Image
                      src="/images/svgs/linkedinIcon.svg"
                      height={24}
                      width={24}
                      alt="github"
                    />{" "}
                    <span className="ml-2"> Github</span>
                  </div>
                }
                label=""
                placeholder="@githubuser"
                id="githubUser"
                type="text"
                className="w-[294px]"
              />
              <Input
                prefix={
                  <div className="flex items-center">
                    <Image
                      src="/images/svgs/twitterIcon.svg"
                      height={24}
                      width={24}
                      alt="github"
                    />{" "}
                    <span className="ml-2"> Github</span>
                  </div>
                }
                label=""
                placeholder="@githubuser"
                id="githubUser"
                type="text"
                className="w-[294px] ml-8"
              />
            </div>
          </div>
          <div className="flex justify-end mt-5">
            <Button variant="primary" className="text-base px-4 py-2">
              Save Changes
            </Button>
          </div>
        </div>
      </div>
    </SettingsLayout>
  );
};

export default General;

General.getLayout = function getLayout(page: ReactElement) {
  return <SettingsLayout>{page}</SettingsLayout>;
};
