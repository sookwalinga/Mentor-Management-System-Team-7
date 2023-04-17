import { useState } from "react";
import AuthLayout from "../components/layouts/auth-layout";
import Image from "next/image";
import { eyePasswordHideIcon, eyePasswordShowIcon } from "@/public";
const SetPassword = () => {
  const [showPassword, setShowPassword] = useState(false);
  return (
    <AuthLayout title="Set Password">
      <h2 className="text-black font-bold mb-4">Set new password</h2>
      <div className="w-full flex justify-center items-center relative  mb-4  border border-gray-400 rounded-lg px-4 h-12">
        <input
          type={showPassword ? "text" : "password"}
          placeholder="Password"
          className="w-full bg-transparent outline-none focus:outline-none"
        />
        <button onClick={() => setShowPassword(!showPassword)}>
          {showPassword ? (
            <Image
              src={eyePasswordShowIcon}
              alt="Google"
              width={20}
              height={20}
            />
          ) : (
            <Image
              src={eyePasswordHideIcon}
              alt="Google"
              width={20}
              height={20}
            />
          )}
        </button>
      </div>

      <p className="mb-4 mt-2">
        *Your new password must be different from previously used password.
      </p>
      <button className="btn bg-mmsPry3 hover:bg-teal-700 text-white font-bold py-2 px-4 rounded w-full normal-case">
        Reset Password
      </button>
    </AuthLayout>
  );
};

export default SetPassword;
