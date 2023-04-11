import { useRouter } from "next/router";
import AuthLayout from "../components/layouts/auth-layout";

const ForgotPassword = () => {
  const router = useRouter();

  return (
    <AuthLayout title="Forgot Password">
      <h1 className="text-black font-bold">Forgot Password?</h1>
      <p>An email has been sent to your registered email.</p>
      <p>Follow the link to reset your password.</p>
      <button
        className="btn bg-mmsPry3 hover:bg-teal-700 text-white font-bold py-2 px-4 rounded w-full normal-case mt-4"
        onClick={() => router.push("/")}
      >
        Done
      </button>
    </AuthLayout>
  );
};

export default ForgotPassword;
