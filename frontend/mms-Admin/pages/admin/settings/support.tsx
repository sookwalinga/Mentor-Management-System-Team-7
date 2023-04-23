import SettingsLayout from "@/components/Settings/SettingsLayout";
import { DashboardLayout } from "@/components/layouts/dashboard-layout";
import React, { ReactElement } from "react";

const Notifications = () => {
  return (
    <div className="border border-[#E6E6E6] rounded-[5px] p-5">
      <h1 className="mmsBlack2 font-bold text-2xl ">How can we help you ?</h1>

      <form>
        <div className="form-control w-full ">
          <input
            type="text"
            placeholder="Name"
            className="input rounded-[7px] border-[#E6E6E6] border-2 w-full my-3 "
          />
        </div>

        <div className="form-control w-full ">
          <input
            type="email"
            placeholder="Email"
            className="input rounded-[7px] border-[#E6E6E6] border-2 w-full my-5 "
          />
        </div>

        <div className="form-control w-full ">
          <input
            type="text"
            placeholder="title"
            className="input rounded-[7px] border-[#E6E6E6] border-2 w-full my-5 "
          />
        </div>

        <div className="form-control w-full ">
          <textarea
            className="textarea textarea-bordered"
            placeholder="body"
          ></textarea>
        </div>

        <div className="button__bottom flex justify-between my-6">
          <div className="file flex">
            <svg
              width="20"
              height="21"
              viewBox="0 0 20 21"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                d="M19 9.52626L10.8445 17.7888C8.5925 20.0705 4.94113 20.0705 2.68905 17.7888C0.436983 15.5072 0.436983 11.8079 2.68905 9.52626L9.4853 2.64082C10.9867 1.11973 13.4209 1.11973 14.9223 2.64082C16.4236 4.16191 16.4236 6.62808 14.9223 8.14917L8.126 15.0346C7.37535 15.7952 6.15824 15.7952 5.40754 15.0346C4.65685 14.2741 4.65685 13.041 5.40754 12.2805L12.2038 5.395"
                stroke="#808080"
                stroke-width="2"
                stroke-linecap="round"
                stroke-linejoin="round"
              />
            </svg>
            <input type="file" className="w-0" />
          </div>

          <button className="bg-mmsPry3 py-[10px] px-[40px] text-white rounded-[7px]">
            Send
          </button>
        </div>
      </form>
    </div>
  );
};

export default Notifications;

Notifications.getLayout = function getLayout(page: ReactElement) {
  return (
    <DashboardLayout>
      <SettingsLayout title="Notifications">{page}</SettingsLayout>
    </DashboardLayout>
  );
};
