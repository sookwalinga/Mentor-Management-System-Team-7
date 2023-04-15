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
      className={clsx("InputScaffold", className, {
        "InputScaffold--fullWidth": fullWidth,
        "InputScaffold--gutter-bottom": gutterBottom,
      })}
      {...rest}
    >
      {!!label && (
        <span className="InputScaffold__label text-start">{label}</span>
      )}
      <div
        className={clsx("InputScaffold__input", {
          "InputScaffold__input--small": size === "small",
          "InputScaffold__input--medium": size === "medium",
          "InputScaffold__input--large": size === "large",
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
