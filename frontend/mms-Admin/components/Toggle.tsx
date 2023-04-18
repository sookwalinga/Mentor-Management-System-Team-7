import { Switch } from "@headlessui/react";
import clsx from "clsx";
import React from "react";

function classNames(...classes: string[]) {
  return classes.filter(Boolean).join(" ");
}

export const Toggle = (props: any) => {
  // const { onChange, enabled, handleAction } = props;
  const [enable, setEnabled] = React.useState(true)

  return (
    <Switch
      checked={enable}
      onChange={() => {
        enable
        // onChange();
        // handleAction();
      }}
      className={classNames(
        enable ? "bg-mmsPry3" : "bg-gray-200",
        clsx(
          "relative inline-flex h-5 w-8 flex-shrink-0 cursor-pointer rounded-full border-2 border-transparent transition-colors duration-200 ease-in-out focus:outline-none focus:ring-2 focus:ring-brand-300 focus:ring-offset-2",
          props.className
        )
      )}
    >
      <span className="sr-only">Use setting</span>
      <span
        aria-hidden="true"
        className={classNames(
          enable ? "translate-x-3" : "translate-x-0",
          clsx(
            "pointer-events-none inline-block h-4 w-4 transform rounded-full bg-white shadow ring-0 transition duration-200 ease-in-out",
            props.togglerStyle
          )
        )}
      />
    </Switch>
  );
};