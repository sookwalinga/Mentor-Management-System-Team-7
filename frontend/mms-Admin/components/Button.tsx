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
        className={clsx("inline-flex cursor-pointer justify-center items-center font-semibold gap-1  rounded-lg", ButtonVariantClass[variant!], className)}
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
    primary: "bg-mmsPry3 text-white",
    secondary: "bg-green11 text-mmsPry3 border border-mmsPry3",
  };