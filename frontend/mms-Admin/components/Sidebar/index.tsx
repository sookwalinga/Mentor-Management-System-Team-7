import React from "react";
import {
  profileIcon,
  certificateIcon,
  programsIcon,
  settingsIcon,
  dashboardIcon,
  taskIcon,
  approvalIcon,
  chatIcon,
  discussionIcon,
  mentorsIcon,
  notificationIcon,
  messagesIcon,
  reportIcon,
  avatarIcon,
  mentorsManagerIcon
} from "@/public";
import Image from "next/image";
import Link from "next/link";
import { useRouter } from "next/router";

let links = [
  {
    name: "Profile",
    icon: profileIcon,
    link: "/admin/profile"
  },
  // {
  //   name: "Dashboard",
  //   icon: dashboardIcon,
  //   link: "/admin/dashboard"
  // },
  // {
  //   name: " Programs",
  //   icon: programsIcon,
  //   link: "/admin/programs"
  // },
  // {
  //   name: "Tasks",
  //   icon: taskIcon,
  //   link: "/admin/tasks"
  // },
  // {
  //   name: "Reports",
  //   icon: reportIcon,
  //   link: "/admin/reports"
  // },
  // {
  //   name: "Mentors",
  //   icon: mentorsIcon,
  //   link: "/admin/mentors"
  // },
  // {
  //   name: "Mentors Manager",
  //   icon: mentorsManagerIcon,
  //   link: "/admin/mentors-manager"
  // },
  // {
  //   name: "Messages",
  //   icon: messagesIcon,
  //   link: "/admin/messages"
  // },
  // {
  //   name: "Discussions",
  //   icon: discussionIcon,
  //   link: "/admin/discussions"
  // },
  // {
  //   name: "Approval",
  //   icon: approvalIcon,

  //   link: "/admin/approval"
  // },
  // {
  //   name: "Certificates",
  //   icon: certificateIcon,
  //   link: "/admin/certificates"
  // },
  {
    name: "Settings",
    icon: settingsIcon,
    link: "/admin/settings/general"
  }
];

const RenderSideBarLinks = () => {
  const router = useRouter();

  return links.map((link: any) => {
    return (
      <Link href={link.link} key={link.name} className="py-2 " legacyBehavior>
        <a
          className={`flex cursor-pointer ${
            router.asPath === link.link && " bg-white text-mmsBlack2 font-bold"
          }`}
        >
          <div className="flex items-center py-2   space-x-[1.3rem] font-normal text-base text-mmsBlack5  group is-published pl-[3.6rem] ">
            <div className="icon">
              <Image src={link.icon} alt={link.name} width={18} height={18} />
            </div>
            <div className="text">
              <p
                // className="text-base font-normal tracking-tighter"
                className={`text-base tracking-tighter ${
                  router.asPath === link.link && "  text-mmsBlack2 font-bold"
                }`}
              >
                {link.name}
              </p>
            </div>
          </div>
        </a>
      </Link>
    );
  });
};

export const Sidebar = () => {
  return (
    <aside className="bg-[#F7FEFF] lg:block  hidden h-screen w-[257px] mt-[30px]   py-[20px] flex-shrink-0 fixed top-0 left-0 bottom-0">
      <div className="user__greeting px-[3.6rem]">
        <h4 className="text-mmsBlack1 font-bold text-[20px] mt-20">
          Hi ,Kabiru
        </h4>
        <p className="font-normal text-[16px] text-mmsBlack5">Admin</p>
      </div>
      <div className="dashboard__links mt-[3rem]">{RenderSideBarLinks()}</div>
    </aside>
  );
};
