import "@/styles/globals.css";
import type { ReactElement, ReactNode } from "react";
import { NextPage } from "next";
import type { AppProps } from "next/app";
import { SWRConfig } from "swr";
import { useRouter } from "next/router";

type NextPageWithLayout = NextPage & {
  getLayout?: (page: ReactElement) => ReactNode;
};

type AppPropsWithLayout = AppProps & {
  Component: NextPageWithLayout;
};
export default function App({ Component, pageProps }: AppPropsWithLayout) {
  const getLayout = Component.getLayout ?? (page => page);

  const router = useRouter();

  return getLayout(
    <SWRConfig
      value={{
        onError: () => router.push("/admin/profile")
      }}
    >
      <Component {...pageProps} />
    </SWRConfig>
  );
}
