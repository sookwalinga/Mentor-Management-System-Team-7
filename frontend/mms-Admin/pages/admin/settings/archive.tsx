import SettingsLayout from "@/components/Settings/SettingsLayout";
import { DashboardLayout } from "@/components/layouts/dashboard-layout";
import { googleIcon, searchIcon } from "@/public";
import Image from "next/image";
import React, { ReactElement } from "react";

const Archive = () => {
  const data = [
    {
      id: 1,
      profileImg: googleIcon,
      program: "Google Africa Scholarship Program",
      details:
        "Potenti ultricies habitant hendrerit semper feugiat pharetra leo etiam egestas nunc placerat proin posuere velit suscipit inceptos nostra consectetur quam pretium non aptent netus",
      postedAt: 1683098201
    },
    {
      id: 2,
      profileImg: googleIcon,
      program: "Google Africa Scholarship Program",
      details:
        "Quam nibh et ad habitant tortor class ridiculus auctor porttitor sagittis pretium luctus aliquet feugiat dolor porta suspendisse laoreet netus ipsum ex tristique lobortis",
      postedAt: 1683098201
    },
    {
      id: 3,
      profileImg: googleIcon,
      program: "Google Africa Scholarship Program",
      details:
        "Vulputate quam si feugiat conubia pede phasellus hendrerit eu posuere vehicula eleifend et suscipit molestie enim sodales duis viverra pellentesque aliquam elementum sagittis dictumst",
      postedAt: 1683098201
    },
    {
      id: 4,
      profileImg: googleIcon,
      program: "Google Africa Scholarship Program",
      details:
        "Nisi montes pharetra dictumst luctus sit congue parturient vestibulum porta cubilia lectus tristique lobortis lorem fames pellentesque quam ipsum suspendisse a ridiculus letius elit",
      postedAt: 1683098201
    }
  ];
  return (
    <div className="left__container flex items-center space-x-[2.3rem]">
      <div className="search__container bg-white  lg:flex px-[2rem] h-[38px] rounded-[5px] items-center  w-[33rem] space-x-[1.4rem] hidden border">
        <div className="search__icon">
          <Image src={searchIcon} alt="search" width={24} height={24} />
        </div>
        <input
          type="text"
          placeholder="Search archive"
          className="search__input bg-white p-1 rounded-[5px] w-full  outline-none placeholder:text-mmsBlack5 font-normal text-base text-mmsBlack5  "
        />
      </div>
      <div className="action__container flex items-center space-x-[2.3rem]">
        <div className="btn-group">
          <button className="btn bg-white">1</button>
          <button className="btn bg-white">2</button>
          <button className="btn bg-white">...</button>
          <button className="btn bg-white">99</button>
          <button className="btn bg-white">100</button>
        </div>
      </div>{" "}
    </div>
  );
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
