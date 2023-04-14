import { useState } from "react";
import AuthLayout from "../components/layouts/auth-layout";
const SetPassword = () => {
  const [showPassword, setShowPassword] = useState(false);
  return (
    <AuthLayout title="Set Password">
      <h2 className="text-black font-bold mb-4">Set new password</h2>
      <div className="w-full flex justify-center items-center relative  border border-gray-400 rounded-lg p-4">
        <input
          type={showPassword ? "text" : "password"}
          placeholder="Password"
          className="w-full bg-transparent"
        />
        <svg
          width="22"
          height="16"
          viewBox="0 0 22 16"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
          onClick={() => setShowPassword(!showPassword)}
        >
          {showPassword ? (
            <>
              <path
                d="M1 12C1 12 5 4 12 4C19 4 23 12 23 12"
                stroke="black"
                stroke-width="2"
                stroke-linecap="round"
                strokeLinejoin="round"
              />
              <path
                d="M1 12C1 12 5 20 12 20C19 20 23 12 23 12"
                stroke="black"
                stroke-width="2"
                stroke-linecap="round"
                strokeLinejoin="round"
              />
              <path
                d="M12 15C13.6569 15 15 13.6569 15 12C15 10.3431 13.6569 9 12 9C10.3431 9 9 10.3431 9 12C9 13.6569 10.3431 15 12 15Z"
                stroke="black"
                stroke-width="2"
                stroke-linecap="round"
                strokeLinejoin="round"
              />
            </>
          ) : (
            <>
              <path
                d="M2 2L22 22"
                stroke="black"
                stroke-width="2"
                stroke-linecap="round"
                strokeLinejoin="round"
              />
              <path
                d="M6.71277 6.7226C3.66479 8.79527 2 12 2 12C2 12 5.63636 19 12 19C14.0503 19 15.8174 18.2734 17.2711 17.2884M11 5.05822C11.3254 5.02013 11.6588 5 12 5C18.3636 5 22 12 22 12C22 12 21.3082 13.3317 20 14.8335"
                stroke="black"
                stroke-width="2"
                stroke-linecap="round"
                strokeLinejoin="round"
              />
              <path
                d="M14 14.2362C13.4692 14.7112 12.7684 15.0001 12 15.0001C10.3431 15.0001 9 13.657 9 12.0001C9 11.1764 9.33193 10.4303 9.86932 9.88818"
                stroke="black"
                stroke-width="2"
                stroke-linecap="round"
                strokeLinejoin="round"
              />
            </>
          )}
        </svg>
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
