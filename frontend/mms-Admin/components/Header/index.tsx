import React from "react";
import Logo from "@/public/images/Logo-Onley-3-01 1.png";
import { searchIcon, notificationIcon, chatIcon, avatarIcon } from "@/public";
import Image from "next/image";
import { signOut } from "next-auth/react";
import Link from "next/link";
import { redirect } from "next/dist/server/api-utils";

export const Header = () => {
  return (
    <nav className="bg-mmsPry3 h-[102px] w-full flex lg:px-[58px] p-3 py-[30px] justify-between items-center  fixed top-0 left-0 right-0 z-10">
      <div className="logo flex items-center space-x-3">
        <div className="logo__container">
          <Link href="/" passHref>
            <Image src={Logo} alt="logo" className="w-[69px]" />
          </Link>
        </div>
        <h4 className="font-semibold text-base text-green11">
          Mentor’s Managers System
        </h4>
      </div>

      <div className="left__container flex items-center space-x-[2.3rem]">
        <div className="search__container bg-white  lg:flex px-[2rem] h-[38px] rounded-[5px] items-center  w-[33rem] space-x-[1.4rem] hidden ">
          <div className="search__icon">
            <Image src={searchIcon} alt="search" width={24} height={24} />
          </div>
          <input
            type="text"
            placeholder="Search for anything"
            className="search__input bg-white p-1 rounded-[5px] w-full  outline-none placeholder:text-mmsBlack5 font-normal text-base text-mmsBlack5  "
          />
        </div>

        <div className="action__container flex items-center space-x-[2.3rem]">
          <Image
            src={chatIcon}
            alt="chat"
            width={24}
            height={24}
            className="cursor-pointer"
          />
          <Image
            src={notificationIcon}
            alt="notification"
            width={24}
            height={24}
            className="cursor-pointer"
          />

          <div className="avatar lg:flex hidden dropdown">
            <div
              className="w-[42px] h-[42px] rounded-full cursor-pointer "
              tabIndex={0}
            >
              <ul
                tabIndex={0}
                className="dropdown-content menu p-2 shadow bg-base-100 rounded-box absolute right-10 top-4 w-52"
              >
                <li onClick={() => signOut({

                  redirect : false
                })}>
                  <a>Sign out</a>
                </li>
              </ul>

              <Image src={avatarIcon} alt="avatar" />
            </div>
          </div>
        </div>
      </div>
    </nav>
  );
};
