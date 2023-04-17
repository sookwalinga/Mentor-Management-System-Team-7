/* eslint-disable require-jsdoc */
import React from "react";
import clsx from "clsx";

export interface InputScaffoldProps {
  prefix?: import("react").ReactNode;
  suffix?: import("react").ReactNode;
  label?: string;
  fullWidth?: boolean;
  size?: "medium" | "large" | "small";
  gutterBottom?: boolean;
  helperText?: string;
  error?: boolean;
  [x: string]: any;
}

function InputScaffold(
  props: InputScaffoldProps & React.ComponentPropsWithoutRef<"div">
) {
  const {
    prefix,
    suffix,
    label,
    fullWidth,
    className,
    children,
    size,
    gutterBottom,
    helperText,
    error,
    ...rest
  } = props;
  return (
    <div
      className={clsx("inline-block", className, {
        "w-full block": fullWidth,
        "mb-4": gutterBottom,
      })}
      {...rest}
    >
      {!!label && (
        <span className="text-sm text-gray-600 block mb-1 text-start">{label}</span>
      )}
      <div
        className={clsx("bg-white flex justify-center gap-2 text-gray-600 px-4 border border-gray-200 border-solid rounded-md", {
          "min-h-[40px]": size === "small",
          "min-h-[48px]": size === "medium",
          "min-h-[56px]": size === "large",
        })}
      >
        {!!prefix && <div className="self-center">{prefix}</div>}
        {children}
        {!!suffix && <div className="self-center">{suffix}</div>}
      </div>
      {!!helperText && (
        <p
          className={clsx(
            "text-sm text-action-disabled",
            !!error && "text-error-dark"
          )}
        >
          {helperText}
        </p>
      )}
    </div>
  );
}

InputScaffold.defaultProps = {
  size: "medium",
};

export default InputScaffold;
