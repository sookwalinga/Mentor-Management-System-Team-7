import Image from "next/image";
import Head from "next/head";
import { useRouter } from "next/router";
import { useEffect } from "react";

const SplashScreen = () => {
  const router = useRouter();
  useEffect(() => {
    const timeout = setTimeout(() => {
      router.push("/login");
    }, 3000);

    return () => clearTimeout(timeout);
  }, [router]);
  return (
    <div className="h-screen flex w-full bg-mmsPry3 items-center justify-center">
      <Head>
        <title>Splash Screen</title>
      </Head>

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
  );
};

export default SplashScreen;
