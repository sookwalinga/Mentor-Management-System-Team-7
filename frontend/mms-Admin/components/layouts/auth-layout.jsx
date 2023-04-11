import Head from "next/head";
import Image from "next/image";
import React from "react";

const AuthLayout = ({ children, title }) => (
  <div className="h-screen flex">
    <Head>
      <title>{title}</title>
    </Head>

    <div className="w-full bg-mmsPry3 items-center justify-center h-screen hidden md:flex">
      <div className="text-center">
        <figure>
          <Image
            src="/images/Logo-Onley-3-01 1.png"
            alt="Album"
            width={300}
            height={200}
          />
        </figure>
        <h1 className="text-white font-bold">Mentor Management System</h1>
      </div>
    </div>

    <div className="w-full bg-white flex flex-col justify-center h-screen px-10">
      {children}
    </div>
  </div>
);

export default AuthLayout;
