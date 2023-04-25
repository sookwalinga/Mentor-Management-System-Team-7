import { useState, CSSProperties, useEffect } from "react";
import Image from "next/image";
import Link from "next/link";
import AuthLayout from "../components/layouts/auth-layout";
import { eyePasswordHideIcon, eyePasswordShowIcon } from "@/public";
import { useRouter } from "next/router";
import BeatLoader from "react-spinners/BeatLoader";

import { useSession, signIn, signOut } from "next-auth/react";

const Login = () => {
  const [showPassword, setShowPassword] = useState(false);
  const [isRedirecting, setIsRedirecting] = useState(false);
  const [loading, setLoading] = useState(false);
  const [inputvalues, setInputValues] = useState({
    email: "ikehfavourdeveloper@gmail.com",
    password: "secretsz"
  });

  const session = useSession();

  // const loading = session.status === "loading";
  console.log("session", session);

  const override: CSSProperties = {
    display: "block",
    margin: "0 auto",
    borderColor: "red"
  };

  const router = useRouter();

  const handleLogin = (e: any) => {
    // signIn('google', { callbackUrl: 'http://localhost:3000/admin/profile' })
    e.preventDefault();

    setLoading(true);
    signIn(
      "credentials",
      {
        redirect: false,
        email: inputvalues.email,
        password: inputvalues.password
      },
      {
        callbackUrl: "http://localhost:3000/admin/profile"
      }
    )
      .then(res => {
        console.log("res", res);
        if (res?.ok) {
          router.push("/admin/profile");
        }

        // alert("Invalid Credentials");

        if (res?.status === 401) {
          alert("Invalid Credentials");
        }
      })
      .catch(err => {
        console.log("err", err);
      })
      .finally(() => {
        setLoading(false);
      });
  };

  useEffect(() => {
    if (session.status == "authenticated" && !isRedirecting && router.isReady) {
      // display some message to the user that he is being redirected
      setIsRedirecting(true);
      // setTimeout(() => {
      // redirect to the return url or home page
      router.push((router.query.returnUrl as string) || "/admin/profile");
      // }, 2000);
    }
  }, [session, isRedirecting, router]);

  console.log(inputvalues);

  return (
    <AuthLayout title="Login">
      <div className="mb-4">
        <h2 className="text-black font-bold">Welcome!</h2>

        {/* {
         session && session?.data.user.name
        } */}

        <h3 className="text-grey ">Login to continue</h3>
      </div>
      <input
        type="email"
        placeholder="Email"
        className="input input-bordered input-success w-full mb-2 bg-transparent"
        value={inputvalues.email}
        onChange={e =>
          setInputValues({ ...inputvalues, email: e.target.value })
        }
      />
      <div className="w-full flex justify-center items-center relative  mb-4  border border-gray-400 rounded-lg px-4 h-12">
        <input
          type={showPassword ? "text" : "password"}
          placeholder="Password"
          value={inputvalues.password}
          className="w-full bg-transparent outline-none focus:outline-none"
          onChange={e =>
            setInputValues({ ...inputvalues, password: e.target.value })
          }
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
      <button
        className="btn bg-mmsPry3 hover:bg-teal-700 text-white font-bold py-2 px-4 rounded w-full normal-case"
        // onClick={() => router.push("/admin/profile")}

        onClick={handleLogin}
      >
        {loading ? (
          <BeatLoader
            color={"#fff"}
            loading={loading}
            cssOverride={override}
            size={15}
            aria-label="Loading Spinner"
            data-testid="loader"
          />
        ) : (
          "Login"
        )}
      </button>
      <Link href="/forgot-password">
        <div className="flex justify-end text-mmsBlack1 hover:text-blue-700 my-2 cursor-pointer">
          Forgot Password?
        </div>
      </Link>
      <button className="btn btn-outline btn-accent hover:bg-teal-700 text-white font-bold py-2 px-4 rounded w-full normal-case">
        <Image src="/images/image 1.png" alt="Google" width={30} height={20} />
        <span className="px-4 text-mmsBlack1"> Signin with Google</span>
      </button>
      <div className="flex justify-center mt-4  text-mmsBlack1">
        New User?{" "}
        <Link href="/">
          <span className="px-1 font-bold underline"> Signup</span>
        </Link>
      </div>
    </AuthLayout>
  );
};

export default Login;
