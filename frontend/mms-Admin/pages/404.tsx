import router from "next/router";
import { DashboardLayout } from "@/components/layouts/dashboard-layout";
import React, { ReactElement } from "react";
import Image from "next/image";
import { errorPageIcon } from "@/public";

export default function ErrorPage() {
  return (
    <>
      <Image src={errorPageIcon} alt="logo" objectFit="cover" />
      <div className="flex justify-center">
        <button
          className="bg-mmsPry3 hover:bg-teal-700 text-white font-bold py-2 px-4 rounded mt-4"
          onClick={() => router.push("/admin/dashboard")}
        >
          Take me home
        </button>
      </div>
    </>
  );
}

ErrorPage.getLayout = function getLayout(page: ReactElement) {
  return <DashboardLayout title="Not Found">{page}</DashboardLayout>;
};
