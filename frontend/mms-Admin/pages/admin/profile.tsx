import { Button } from "@/components";
import { DashboardLayout } from "@/components/layouts/dashboard-layout";
import Image from "next/image";
import { useRouter } from "next/router";
import React, { ReactElement } from "react";
import { Calendar, Globe, Mail, MapPin } from "react-feather";

const Profile = () => {
  const router = useRouter();
  return (
    <div>
      <div className="mb-[60px] flex items-center justify-between">
        <div className="flex items-center">
          <div className="rounded-full cursor-pointer ">
            <Image
              src="/images/personImg.png"
              alt="avatar"
              width={90}
              height={90}
            />
          </div>
          <div className="mx-5">
            <h3 className="text-mmsBlack2 font-semibold text-2xl">
              Peculiar Umeh
            </h3>
            <p>Admin</p>
          </div>
        </div>
        <div>
          <Button
            variant="primary"
            className="text-base px-[40px] py-[10px]"
            onClick={() => router.push("/admin/settings/general")}
          >
            Edit Profile
          </Button>
        </div>
      </div>
      <div className="border border-gray-200 rounded-md p-5">
        <div>
          <h2 className="text-mmsBlack2 font-semibold mb-3 text-2xl">About</h2>
          <p className="p-5 bg-green11 text-mmsBlack3 text-base">
            Lorem ipsum dolor sit amet, consectetur adipiscing elit. Praesent
            dignissim ut cursus purus efficitur et. Duis ac enim tellus.
            Phasellus pharetra metus, ut cursus purus efficitur et. Duis ac enim
            tellus. Phasellus eget tortor dapibus, laoreet mauris sed, dignissim
            lectus. Duis ac enim tellus. Phasellus pharetra metus, ut cursus
            purus efficitur et. Duis ac enim tellus. Phasellus eget tortor
            dapibus, laoreet mauris sed, dignissim lectus.
          </p>
        </div>
        <div className="mt-6">
          <div className="flex items-center">
            <div className="flex items-center">
              <Mail size={20} className="text-mmsPry3 mr-3" />
              <span>peculiah@andela.com</span>
            </div>
            <div className="flex items-center ml-[15%]">
              <MapPin size={20} className="text-mmsPry3 mr-3" />
              <span>Lagos, Nigeria</span>
            </div>
          </div>
          <div className="flex items-center mt-[33px]">
            <div className="flex items-center">
              <Globe size={20} className="text-mmsPry3 mr-3" />
              <span>www.peculiah.com</span>
            </div>
            <div className="flex items-center ml-[17%]">
              <Calendar size={20} className="text-mmsPry3 mr-3" />
              <span>Member since June 22, 2021</span>
            </div>
          </div>
        </div>
        <div className="mt-[50px]">
          <h2 className="text-mmsBlack2 font-semibold mb-3 text-2xl">Social</h2>
          <div className="flex items-center">
            <div className="flex items-center bg-mmsPry10 py-[5px] px-[11px]">
              <Image
                src="/images/svgs/githubIcon.svg"
                width={24}
                height={23}
                alt="github"
              />
              <span className="font-semibold text-xl text-mmsBlack3 ml-3">
                @peculiah.umeh
              </span>
            </div>
            <div className="flex items-center bg-mmsPry10 py-[5px] px-[11px] ml-[15%]">
              <Image
                src="/images/svgs/linkedinIcon.svg"
                width={24}
                height={23}
                alt="github"
              />
              <span className="font-semibold text-xl text-mmsBlack3 ml-3">
                @peculiah.umeh
              </span>
            </div>
          </div>
        </div>
        <div className="flex items-center mt-[33px]">
          <div className="flex items-center bg-mmsPry10 py-[5px] px-[11px]">
            <Image
              src="/images/svgs/twitterIcon.svg"
              width={24}
              height={23}
              alt="github"
            />
            <span className="font-semibold text-xl text-mmsBlack3 ml-3">
              @peculiah.umeh
            </span>
          </div>
          <div className="flex items-center bg-mmsPry10 py-[5px] px-[11px] ml-[15%]">
            <Image
              src="/images/svgs/instagramIcon.svg"
              width={24}
              height={23}
              alt="github"
            />
            <span className="font-semibold text-xl text-mmsBlack3 ml-3">
              @peculiah.umeh
            </span>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Profile;

Profile.requireAuth = true;

Profile.getLayout = function getLayout(page: ReactElement) {
  return <DashboardLayout title="Profile">{page}</DashboardLayout>;
};
