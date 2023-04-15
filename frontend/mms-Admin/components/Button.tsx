import clsx from "clsx";
import React from "react";

export interface ButtonProps {
    variant?: keyof typeof ButtonVariantClass;
    prefix?: import("react").ReactNode;
    suffix?: import("react").ReactNode;
    [x: string]: any;
  }
  
  export const Button = React.forwardRef(function Button4(
    props: ButtonProps,
    ref: any
  ) {
    const { className, variant, prefix, suffix, children, ...rest } = props;
  
    return (
      <button
        ref={ref}
        className={clsx("button", ButtonVariantClass[variant!], className)}
        {...rest}
      >
        {prefix}
        {children}
        {suffix}
      </button>
    );
  });
  
  export default Button;
  
  const ButtonVariantClass: any = {
    primary: "button-primary",
    secondary: "button-secondary",
  };