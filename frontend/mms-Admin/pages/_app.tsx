import "@/styles/globals.css";
import type { ReactElement, ReactNode } from "react";
import { NextPage } from "next";
import type { AppProps } from "next/app";
import { SWRConfig } from "swr";
import { useRouter } from "next/router";
import { SessionProvider } from "next-auth/react";
import { ProtectedLayout } from "@/components/layouts/protected";

type NextPageWithLayout = NextPage & {
  getLayout?: (page: ReactElement) => ReactNode;
};

type AppPropsWithLayout = AppProps & {
  Component: {
    getLayout?: (page: ReactElement) => ReactNode;
    requireAuth: boolean;
  };
  session: any;
};
export default function App({
  Component,
  pageProps,
  session
}: AppPropsWithLayout) {
  const getLayout = Component.getLayout ?? (page => page);

  const router = useRouter();

  return getLayout(
    <SessionProvider session={session}>

      {Component.requireAuth ? (

        <ProtectedLayout>
          <SWRConfig
            value={{
              onError: () => router.push("/admin/profile")
            }}
          >
            {/* <ProtectedLayout> */}
            <Component {...pageProps} />

            {/* </ProtectedLayout> */}
          </SWRConfig>
        </ProtectedLayout>
   
      ) : (
        
        <SWRConfig
          value={{
            onError: () => router.push("/admin/profile")
          }}
        >
          {/* <ProtectedLayout> */}
          <Component {...pageProps} />

          {/* </ProtectedLayout> */}
        </SWRConfig>
      )}
    </SessionProvider>
  );
}
