import { DashboardLayout } from "@/components/layouts/dashboard-layout";

import {
  profileIcon,
  mentorsIconLarge,
  mentorsManagerIconLarge,
  taskIconLarge,
  reportIconLarge,
  dashboardIcon,
  GADSIcon
} from "@/public";
import Image from "next/image";
import React, { ReactElement } from "react";

const data = [
  {
    name: "Mentors",
    value: 20,
    icon: mentorsIconLarge,
    percent: 20
  },

  {
    name: "Mentors Manager",
    value: 20,
    icon: mentorsManagerIconLarge,
    percent: 20
  },
  {
    name: "Tasks",
    value: 20,
    icon: taskIconLarge,
    percent: 20
  },
  {
    name: "Reports",
    value: 20,
    icon: reportIconLarge,
    percent: 20
  }
];

let programOverview = [
  {
    name: "GADS Program 2022",
    value: 50,
    percent: 20,
    icon: GADSIcon
  },
  {
    name: "GADS Program 2022",
    value: 50,
    percent: 20,
    icon: GADSIcon
  },

  {
    name: "GADS Program 2022",
    value: 20,
    percent: 20,
    icon: GADSIcon
  }
];

const reports = [
  {
    name: "Google Africa Scholarship",
    author: " Ibrahim Kabir ",
    value: 20
  },

  {
    name: "Google Africa Scholarship",
    author: " Ibrahim Kabir ",

    value: 20
  },

  {
    name: "Google Africa Scholarship",
    author: " Ibrahim Kabir ",

    value: 20
  }
];

const tasks = [
  {
    name: " Room liberary article write........",

    days: "2 days ago",
    icon: taskIconLarge
  },
  {
    name: " Room liberary article write........",

    days: "4 days ago",
    icon: taskIconLarge
  },
  {
    name: " Room liberary article write........",

    days: "45 days ago",
    icon: taskIconLarge
  }
];

const DashboardHome = () => {
  return (
  <>
      <div className="dashboard__details flex lg:items-center lg:flex-row flex-col  justify-between ">
        <div className="bg-mmsPry3 lg:w-[194px] h-[92px]  rounded-[7px] pl-[20px] pr-4 py-2">
          <div className="flex justify-end  left-[200px] ">
            <div className=" w-[55px]  h-[24px] bg-green11 rounded-[5px] flex items-center justify-center">
              <p className="text-xs font-normal text-mmsPry3 text-center ">
                View
              </p>
            </div>
          </div>

          <div className="flex items-center  h-[30px] w-36 lg:w-full ">
            <div className="text-[64px] text-green11 font-bold">6</div>
            <div className="text-xl ml-2 text-green11 font-bold tracking-tight leading-4">
              Active programs
            </div>
          </div>
        </div>

        <div className="bg-green11 border px-[12px] py-[11.5px]  border-[#E6FDFE] lg:h-[92px]  rounded-[10px] flex lg:flex-row flex-col  space-y-4 lg:space-y-0  lg:space-x-[28px]">
          {data.map((item, index) => (
            <div
              key={index}
              className="dashboard__info rounded-[7px]   h-[69px] bg-mmsPry10 px-[19px] py-[7px] flex items-center lg:space-x-[29px] justify-between "
            >
              <div className="dashboard__details ">
                <div className="dashboard__top">
                  <h1 className="text-xl font-normal text-mmsBlack2 tracking-tighter">
                    {item.name}
                  </h1>
                </div>

                <div className="dashboard__bottom flex items-center">
                  <h3 className="font-semibold text-mmsBlack2 text-xl">
                    {item.value}
                  </h3>
                  <h4 className="text-mmsBlack5 text-base font-normal ">
                    +{item.percent}%
                  </h4>
                </div>
              </div>

              <div className="dashboard__icon ">
                <Image src={item.icon} alt="" className="" />
              </div>
            </div>
          ))}
        </div>
      </div>

      <div className="programoverview__container w-full bg-green11 border border-mmsPry10 lg:h-[188px] my-[28px] rounded-[10px] px-[19px] py-[9px]">
        <div className="row flex justify-between">
          <h3 className="text-mmsBlack1 font-semibold text-lg mb-[10px]">
            Program overview
          </h3>

          <div className="reports__summited__container  h-[30px] text-lg tracking-tighter border border-mmsPry10 bg-white rounded-[5px] px-[8px] text-mmsBlack2 font-semibold ">
            6 Active
          </div>
        </div>

        <div className="programoverview__cards__container flex flex-col lg:flex-row space-y-4 lg:space-y-0 justify-between">
          {programOverview.map((item, index) => (
            <>
              <div className="bg-mmsPry10 rounded-[10px] py-[13px] px-[10px] h-[92px] w-[332px]">
                <div className="programoverview__card__top flex justify-between items-center">
                  <div className="programoverview__card__top__left flex items-center space-x-[30px]">
                    <div className="programoverview__card__icon">
                      <Image src={item.icon} alt="" className="" />
                    </div>

                    <div className="programoverview__card__details">
                      <h3 className="text-mmsBlack1 font-normal text-xl tracking-tighter">
                        {item.name}
                      </h3>

                      <div className="progress__container flex items-center space-x-[29px]">
                        <h4 className="text-mmsBlack3 text-xs font-normal">
                          {item.value}%
                        </h4>
                        <ProgressBar progress={item.value} />
                      </div>
                    </div>
                  </div>
                </div>

                <span className="font-normal text-mmsBlack3 text-xs mt-[7px]">
                  Jun 13, 2022 - `{">"}` Feb 10, 2023
                </span>
              </div>
            </>
          ))}
        </div>

        <div className="viewall__button__container flex justify-end">
          <button className="bg-mmsPry3 viweall__btn  text-white font-semibold text-sm  rounded-[5px] w-[70px] h-[24px] mt-[10px]">
            View all
          </button>
        </div>
      </div>

      <div className="reportoverview__container w-full bg-green11 border border-mmsPry10 lg:h-[161px] my-[28px] rounded-[10px] px-[19px] py-[9px]">
        <div className="row flex justify-between">
          <h3 className="text-mmsBlack1 font-semibold text-lg mb-[10px] items-center">
            Reports overview
          </h3>

          <div className="reports__summited__container w-[168px] h-[30px] text-lg tracking-tighter border border-mmsPry10 bg-white rounded-[5px] px-[8px] text-mmsBlack2 font-semibold ">
            10 report submitted
          </div>
        </div>

        <div className="reportoverview__container flex flex-col space-y-5 lg:space-y-0 lg:flex-row justify-between">
          {reports.map((item, index) => (
            <>
              <div className="bg-mmsPry10 rounded-[10px] py-[13px] px-[10px] h-[64px] w-[332px]">
                <div className="reportoverview__card flex justify-between items-center">
                  <div className="reportoverview__card__top__left flex items-center space-x-[30px]">
                    <div className="reportoverview__card__icon">
                      <svg
                        width="39"
                        height="40"
                        viewBox="0 0 39 40"
                        fill="none"
                        xmlns="http://www.w3.org/2000/svg"
                      >
                        <path
                          d="M26.1112 9.99967H13.2223M26.1112 16.6663H13.2223M26.1112 23.333H21.2779M29.3334 36.6663H10.0001C8.22048 36.6663 6.77789 35.174 6.77789 33.333V6.66634C6.77789 4.82539 8.22048 3.33301 10.0001 3.33301H29.3334C31.113 3.33301 32.5557 4.82539 32.5557 6.66634V33.333C32.5557 35.174 31.113 36.6663 29.3334 36.6663Z"
                          stroke="#058B94"
                          stroke-width="2"
                          stroke-linecap="round"
                          stroke-linejoin="round"
                        />
                      </svg>
                    </div>

                    <div className="reportoverview__card__details">
                      <h3 className="text-mmsBlack1 font-normal text-xl ">
                        {item.name}
                      </h3>

                      <div className="author__container flex items-center ">
                        <h4 className="text-mmsBlack3 text-xs font-semibold">
                          {item.author} -
                        </h4>
                        <span className="text-mmsBlack5 text-xs font-normal">
                          19th - 25th Oct 22
                        </span>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </>
          ))}
        </div>

        <div className="viewall__button__container flex justify-end">
          <button className="bg-mmsPry3 viweall__btn  text-white font-semibold text-sm  rounded-[5px] w-[70px] h-[24px] mt-[10px]">
            View all
          </button>
        </div>
      </div>

      <div className="taskoverview__container w-full bg-green11 border border-mmsPry10 lg:h-[341px] my-[28px] rounded-[10px] px-[19px] py-[9px]">
        <h3 className="text-mmsBlack1 font-semibold text-lg mb-[10px]">
          Tasks overview
        </h3>

        <div className="task__container mb-[28px] ">
          <div className="tasks__tasks flex lg:flex-row flex-col  lg:items-center  lg:space-x-[28px] w-full">
            <div className="tasks__indicatorlg:mb-0 mb-5 h-[92px] w-[120px] rounded-[10px] bg-mmsPry3 flex items-center justify-center">
              <p className="text-green11 font-semibold text-xl tracking-tighter">
                In Progress
              </p>
            </div>

            <div className="tasks__row w-full flex flex-col lg:flex-row justify-between lg:space-y-0 space-y-4 ">
              {tasks.map((item, index) => (
                <div
                  className="task__container flex bg-mmsPry10 rounded-[10px] h-[92px] lg:w-[283px] space-x-[24px] px-[9px] items-center py-[17px]"
                  key={index}
                >
                  <div className="icon">
                    <Image src={item.icon} alt="" className="" />
                  </div>

                  <div className="tasks__detials">
                    <h2 className="text-mmsBlack2 font-normal text-xl tracking-tighter">
                      {item.name}
                    </h2>

                    <div className="task__date flex space-x-[3px] items-center mt-[3px] ">
                      <svg
                        width="18"
                        height="18"
                        viewBox="0 0 18 18"
                        fill="none"
                        xmlns="http://www.w3.org/2000/svg"
                      >
                        <path
                          d="M8.99984 10.667C9.16465 10.667 9.32577 10.6181 9.46281 10.5266C9.59985 10.435 9.70666 10.3048 9.76974 10.1526C9.83281 10.0003 9.84931 9.83273 9.81716 9.67108C9.785 9.50943 9.70564 9.36095 9.58909 9.2444C9.47255 9.12786 9.32406 9.04849 9.16241 9.01634C9.00076 8.98418 8.83321 9.00069 8.68093 9.06376C8.52866 9.12683 8.39851 9.23364 8.30695 9.37068C8.21538 9.50772 8.1665 9.66884 8.1665 9.83366C8.1665 10.0547 8.2543 10.2666 8.41058 10.4229C8.56686 10.5792 8.77882 10.667 8.99984 10.667ZM13.1665 10.667C13.3313 10.667 13.4924 10.6181 13.6295 10.5266C13.7665 10.435 13.8733 10.3048 13.9364 10.1526C13.9995 10.0003 14.016 9.83273 13.9838 9.67108C13.9517 9.50943 13.8723 9.36095 13.7558 9.2444C13.6392 9.12786 13.4907 9.04849 13.3291 9.01634C13.1674 8.98418 12.9999 9.00069 12.8476 9.06376C12.6953 9.12683 12.5652 9.23364 12.4736 9.37068C12.382 9.50772 12.3332 9.66884 12.3332 9.83366C12.3332 10.0547 12.421 10.2666 12.5772 10.4229C12.7335 10.5792 12.9455 10.667 13.1665 10.667ZM8.99984 14.0003C9.16465 14.0003 9.32577 13.9515 9.46281 13.8599C9.59985 13.7683 9.70666 13.6382 9.76974 13.4859C9.83281 13.3336 9.84931 13.1661 9.81716 13.0044C9.785 12.8428 9.70564 12.6943 9.58909 12.5777C9.47255 12.4612 9.32406 12.3818 9.16241 12.3497C9.00076 12.3175 8.83321 12.334 8.68093 12.3971C8.52866 12.4602 8.39851 12.567 8.30695 12.704C8.21538 12.8411 8.1665 13.0022 8.1665 13.167C8.1665 13.388 8.2543 13.6 8.41058 13.7562C8.56686 13.9125 8.77882 14.0003 8.99984 14.0003ZM13.1665 14.0003C13.3313 14.0003 13.4924 13.9515 13.6295 13.8599C13.7665 13.7683 13.8733 13.6382 13.9364 13.4859C13.9995 13.3336 14.016 13.1661 13.9838 13.0044C13.9517 12.8428 13.8723 12.6943 13.7558 12.5777C13.6392 12.4612 13.4907 12.3818 13.3291 12.3497C13.1674 12.3175 12.9999 12.334 12.8476 12.3971C12.6953 12.4602 12.5652 12.567 12.4736 12.704C12.382 12.8411 12.3332 13.0022 12.3332 13.167C12.3332 13.388 12.421 13.6 12.5772 13.7562C12.7335 13.9125 12.9455 14.0003 13.1665 14.0003ZM4.83317 10.667C4.99799 10.667 5.1591 10.6181 5.29615 10.5266C5.43319 10.435 5.54 10.3048 5.60307 10.1526C5.66614 10.0003 5.68265 9.83273 5.65049 9.67108C5.61834 9.50943 5.53897 9.36095 5.42243 9.2444C5.30588 9.12786 5.1574 9.04849 4.99575 9.01634C4.8341 8.98418 4.66654 9.00069 4.51427 9.06376C4.362 9.12683 4.23185 9.23364 4.14028 9.37068C4.04871 9.50772 3.99984 9.66884 3.99984 9.83366C3.99984 10.0547 4.08763 10.2666 4.24392 10.4229C4.4002 10.5792 4.61216 10.667 4.83317 10.667ZM14.8332 2.33366H13.9998V1.50033C13.9998 1.27931 13.912 1.06735 13.7558 0.91107C13.5995 0.75479 13.3875 0.666992 13.1665 0.666992C12.9455 0.666992 12.7335 0.75479 12.5772 0.91107C12.421 1.06735 12.3332 1.27931 12.3332 1.50033V2.33366H5.6665V1.50033C5.6665 1.27931 5.57871 1.06735 5.42243 0.91107C5.26615 0.75479 5.05418 0.666992 4.83317 0.666992C4.61216 0.666992 4.4002 0.75479 4.24392 0.91107C4.08763 1.06735 3.99984 1.27931 3.99984 1.50033V2.33366H3.1665C2.50346 2.33366 1.86758 2.59705 1.39874 3.06589C0.929896 3.53473 0.666504 4.17062 0.666504 4.83366V14.8337C0.666504 15.4967 0.929896 16.1326 1.39874 16.6014C1.86758 17.0703 2.50346 17.3337 3.1665 17.3337H14.8332C15.4962 17.3337 16.1321 17.0703 16.6009 16.6014C17.0698 16.1326 17.3332 15.4967 17.3332 14.8337V4.83366C17.3332 4.17062 17.0698 3.53473 16.6009 3.06589C16.1321 2.59705 15.4962 2.33366 14.8332 2.33366ZM15.6665 14.8337C15.6665 15.0547 15.5787 15.2666 15.4224 15.4229C15.2661 15.5792 15.0542 15.667 14.8332 15.667H3.1665C2.94549 15.667 2.73353 15.5792 2.57725 15.4229C2.42097 15.2666 2.33317 15.0547 2.33317 14.8337V7.33366H15.6665V14.8337ZM15.6665 5.66699H2.33317V4.83366C2.33317 4.61264 2.42097 4.40068 2.57725 4.2444C2.73353 4.08812 2.94549 4.00033 3.1665 4.00033H14.8332C15.0542 4.00033 15.2661 4.08812 15.4224 4.2444C15.5787 4.40068 15.6665 4.61264 15.6665 4.83366V5.66699ZM4.83317 14.0003C4.99799 14.0003 5.1591 13.9515 5.29615 13.8599C5.43319 13.7683 5.54 13.6382 5.60307 13.4859C5.66614 13.3336 5.68265 13.1661 5.65049 13.0044C5.61834 12.8428 5.53897 12.6943 5.42243 12.5777C5.30588 12.4612 5.1574 12.3818 4.99575 12.3497C4.8341 12.3175 4.66654 12.334 4.51427 12.3971C4.362 12.4602 4.23185 12.567 4.14028 12.704C4.04871 12.8411 3.99984 13.0022 3.99984 13.167C3.99984 13.388 4.08763 13.6 4.24392 13.7562C4.4002 13.9125 4.61216 14.0003 4.83317 14.0003Z"
                          fill="#058B94"
                        />
                      </svg>

                      <span className="text-sm text-mmsBlack5 font-normal tracking-tighter">
                        {item.days}
                      </span>
                    </div>
                  </div>
                </div>
              ))}
            </div>
          </div>
          <div className="viewall__button__container flex justify-end">
            <button className="bg-mmsPry3 viweall__btn  text-white font-semibold text-sm  rounded-[5px] w-[70px] h-[24px] mt-[10px]">
              View all
            </button>
          </div>
        </div>

        <div className="task__container mb-[28px] ">
          <div className="tasks__tasks flex lg:flex-row flex-col  lg:items-center  lg:space-x-[28px] w-full">
            <div className="tasks__indicatorlg:mb-0 mb-5 h-[92px] w-[120px] rounded-[10px] bg-mmsPry3 flex items-center justify-center">
              <p className="text-green11 font-semibold text-xl tracking-tighter">
                Completed
              </p>
            </div>

            <div className="tasks__row w-full flex flex-col lg:flex-row justify-between lg:space-y-0 space-y-4 ">
              {tasks.map((item, index) => (
                <div
                  className="task__container flex bg-mmsPry10 rounded-[10px] h-[92px] lg:w-[283px] space-x-[24px] px-[9px] items-center py-[17px]"
                  key={index}
                >
                  <div className="icon">
                    <Image src={item.icon} alt="" className="" />
                  </div>

                  <div className="tasks__detials">
                    <h2 className="text-mmsBlack2 font-normal text-xl tracking-tighter">
                      {item.name}
                    </h2>

                    <div className="task__date flex space-x-[3px] items-center mt-[3px] ">
                      <svg
                        width="18"
                        height="18"
                        viewBox="0 0 18 18"
                        fill="none"
                        xmlns="http://www.w3.org/2000/svg"
                      >
                        <path
                          d="M8.99984 10.667C9.16465 10.667 9.32577 10.6181 9.46281 10.5266C9.59985 10.435 9.70666 10.3048 9.76974 10.1526C9.83281 10.0003 9.84931 9.83273 9.81716 9.67108C9.785 9.50943 9.70564 9.36095 9.58909 9.2444C9.47255 9.12786 9.32406 9.04849 9.16241 9.01634C9.00076 8.98418 8.83321 9.00069 8.68093 9.06376C8.52866 9.12683 8.39851 9.23364 8.30695 9.37068C8.21538 9.50772 8.1665 9.66884 8.1665 9.83366C8.1665 10.0547 8.2543 10.2666 8.41058 10.4229C8.56686 10.5792 8.77882 10.667 8.99984 10.667ZM13.1665 10.667C13.3313 10.667 13.4924 10.6181 13.6295 10.5266C13.7665 10.435 13.8733 10.3048 13.9364 10.1526C13.9995 10.0003 14.016 9.83273 13.9838 9.67108C13.9517 9.50943 13.8723 9.36095 13.7558 9.2444C13.6392 9.12786 13.4907 9.04849 13.3291 9.01634C13.1674 8.98418 12.9999 9.00069 12.8476 9.06376C12.6953 9.12683 12.5652 9.23364 12.4736 9.37068C12.382 9.50772 12.3332 9.66884 12.3332 9.83366C12.3332 10.0547 12.421 10.2666 12.5772 10.4229C12.7335 10.5792 12.9455 10.667 13.1665 10.667ZM8.99984 14.0003C9.16465 14.0003 9.32577 13.9515 9.46281 13.8599C9.59985 13.7683 9.70666 13.6382 9.76974 13.4859C9.83281 13.3336 9.84931 13.1661 9.81716 13.0044C9.785 12.8428 9.70564 12.6943 9.58909 12.5777C9.47255 12.4612 9.32406 12.3818 9.16241 12.3497C9.00076 12.3175 8.83321 12.334 8.68093 12.3971C8.52866 12.4602 8.39851 12.567 8.30695 12.704C8.21538 12.8411 8.1665 13.0022 8.1665 13.167C8.1665 13.388 8.2543 13.6 8.41058 13.7562C8.56686 13.9125 8.77882 14.0003 8.99984 14.0003ZM13.1665 14.0003C13.3313 14.0003 13.4924 13.9515 13.6295 13.8599C13.7665 13.7683 13.8733 13.6382 13.9364 13.4859C13.9995 13.3336 14.016 13.1661 13.9838 13.0044C13.9517 12.8428 13.8723 12.6943 13.7558 12.5777C13.6392 12.4612 13.4907 12.3818 13.3291 12.3497C13.1674 12.3175 12.9999 12.334 12.8476 12.3971C12.6953 12.4602 12.5652 12.567 12.4736 12.704C12.382 12.8411 12.3332 13.0022 12.3332 13.167C12.3332 13.388 12.421 13.6 12.5772 13.7562C12.7335 13.9125 12.9455 14.0003 13.1665 14.0003ZM4.83317 10.667C4.99799 10.667 5.1591 10.6181 5.29615 10.5266C5.43319 10.435 5.54 10.3048 5.60307 10.1526C5.66614 10.0003 5.68265 9.83273 5.65049 9.67108C5.61834 9.50943 5.53897 9.36095 5.42243 9.2444C5.30588 9.12786 5.1574 9.04849 4.99575 9.01634C4.8341 8.98418 4.66654 9.00069 4.51427 9.06376C4.362 9.12683 4.23185 9.23364 4.14028 9.37068C4.04871 9.50772 3.99984 9.66884 3.99984 9.83366C3.99984 10.0547 4.08763 10.2666 4.24392 10.4229C4.4002 10.5792 4.61216 10.667 4.83317 10.667ZM14.8332 2.33366H13.9998V1.50033C13.9998 1.27931 13.912 1.06735 13.7558 0.91107C13.5995 0.75479 13.3875 0.666992 13.1665 0.666992C12.9455 0.666992 12.7335 0.75479 12.5772 0.91107C12.421 1.06735 12.3332 1.27931 12.3332 1.50033V2.33366H5.6665V1.50033C5.6665 1.27931 5.57871 1.06735 5.42243 0.91107C5.26615 0.75479 5.05418 0.666992 4.83317 0.666992C4.61216 0.666992 4.4002 0.75479 4.24392 0.91107C4.08763 1.06735 3.99984 1.27931 3.99984 1.50033V2.33366H3.1665C2.50346 2.33366 1.86758 2.59705 1.39874 3.06589C0.929896 3.53473 0.666504 4.17062 0.666504 4.83366V14.8337C0.666504 15.4967 0.929896 16.1326 1.39874 16.6014C1.86758 17.0703 2.50346 17.3337 3.1665 17.3337H14.8332C15.4962 17.3337 16.1321 17.0703 16.6009 16.6014C17.0698 16.1326 17.3332 15.4967 17.3332 14.8337V4.83366C17.3332 4.17062 17.0698 3.53473 16.6009 3.06589C16.1321 2.59705 15.4962 2.33366 14.8332 2.33366ZM15.6665 14.8337C15.6665 15.0547 15.5787 15.2666 15.4224 15.4229C15.2661 15.5792 15.0542 15.667 14.8332 15.667H3.1665C2.94549 15.667 2.73353 15.5792 2.57725 15.4229C2.42097 15.2666 2.33317 15.0547 2.33317 14.8337V7.33366H15.6665V14.8337ZM15.6665 5.66699H2.33317V4.83366C2.33317 4.61264 2.42097 4.40068 2.57725 4.2444C2.73353 4.08812 2.94549 4.00033 3.1665 4.00033H14.8332C15.0542 4.00033 15.2661 4.08812 15.4224 4.2444C15.5787 4.40068 15.6665 4.61264 15.6665 4.83366V5.66699ZM4.83317 14.0003C4.99799 14.0003 5.1591 13.9515 5.29615 13.8599C5.43319 13.7683 5.54 13.6382 5.60307 13.4859C5.66614 13.3336 5.68265 13.1661 5.65049 13.0044C5.61834 12.8428 5.53897 12.6943 5.42243 12.5777C5.30588 12.4612 5.1574 12.3818 4.99575 12.3497C4.8341 12.3175 4.66654 12.334 4.51427 12.3971C4.362 12.4602 4.23185 12.567 4.14028 12.704C4.04871 12.8411 3.99984 13.0022 3.99984 13.167C3.99984 13.388 4.08763 13.6 4.24392 13.7562C4.4002 13.9125 4.61216 14.0003 4.83317 14.0003Z"
                          fill="#058B94"
                        />
                      </svg>

                      <span className="text-sm text-mmsBlack5 font-normal tracking-tighter">
                        {item.days}
                      </span>
                    </div>
                  </div>
                </div>
              ))}
            </div>
          </div>
          <div className="viewall__button__container flex justify-end">
            <button className="bg-mmsPry3 viweall__btn  text-white font-semibold text-sm  rounded-[5px] w-[70px] h-[24px] mt-[10px]">
              View all
            </button>
          </div>
        </div>
      </div>
      </>

  );
};

const ProgressBar = ({ progress }: any) => {
  return (
    <div className="relative h-[7px] bg-gray-200 rounded w-[106px]">
      <div
        className="absolute top-0 left-0 h-full bg-mmsPry3 rounded"
        style={{ width: `${progress}%` }}
      />
      <div className="absolute top-0 left-0 h-full w-full flex justify-center items-center" />
    </div>
  );
};

export default DashboardHome;


DashboardHome.requireAuth = true;


DashboardHome.getLayout = function getLayout(page: ReactElement) {
  return <DashboardLayout title="DashbOARD">{page}</DashboardLayout>;
};

