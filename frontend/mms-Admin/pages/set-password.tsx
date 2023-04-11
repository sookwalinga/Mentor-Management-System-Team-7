import AuthLayout from "../components/layouts/auth-layout";
const ForgotPassword = () => {
  return (
    <AuthLayout title="Login">
      <h1 className="text-black font-bold">Forgot Password?</h1>
      <p>An email has been sent to your registered email.</p>
      <p>Follow the link to reset your password.</p>
      <button className="btn bg-btnPrimary hover:bg-teal-700 text-white font-bold py-2 px-4 rounded w-full normal-case mt-4">
        ForgotPassword
      </button>
    </AuthLayout>
  );
};

export default ForgotPassword;
