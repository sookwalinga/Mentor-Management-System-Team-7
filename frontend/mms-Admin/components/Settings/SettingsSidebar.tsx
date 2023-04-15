import Link from "next/link";
import React from "react";

const SettingsSidebar = () => {
  return (
    <div className="flex flex-col py-4 bg-white w-[225px] h-[295px]">
      {routes.map((route, index) => (
        <Link key={index} href={route.link} passHref>
          <div className="text-center px-8 py-[8px] text-mmsBlack5 bg-green11 border border-mmsPry10 text-base">{route.name}</div>
        </Link>
      ))}
    </div>
  );
};

export default SettingsSidebar;

const routes = [
  {
    name: "General",
    link: "admin/settings/profile"
  },
  {
    name: "Password",
    link: "admin/settings/password"
  },
  {
    name: "Notifications",
    link: "admin/settings/notifications"
  },
  {
    name: "Privacy",
    link: "admin/settings/privacy"
  },
  {
    name: "Archive",
    link: "admin/settings/archive"
  },
  {
    name: "Support",
    link: "admin/settings/support"
  },
  {
    name: "FAQ",
    link: "admin/settings/faq"
  }
];
