import { ReactElement, ReactNode, ComponentProps } from "react";
import { Loading } from "@/components/loading";

const variants = {
  primary: {
    base: "bg-sub",
    loading: "bg-indigo-600",
  },
  secondary: {
    base: "bg-gray-400",
    loading: "bg-gray-500",
  },
};

const sizes = {
  sm: "w-24 text-xs h-8",
  md: "w-32",
  lg: "w-72",
};

export type ButtonProps = {
  variant?: keyof typeof variants;
  size?: "lg" | "md" | "sm";
  children: ReactNode;
  isLoading?: boolean;
  icon?: ReactElement;
} & Omit<ComponentProps<"button">, "className">;

export const Button = ({
  variant = "primary",
  size = "md",
  type = "button",
  children,
  isLoading,
  icon,
  disabled,
  ...props
}: ButtonProps) => {
  const { base, loading } = variants[variant];
  const width = sizes[size];

  return (
    <button
      className={`text-white text-sm font-bold tracking-wider h-10 ${width} rounded-sm transition-all duration-200 text-center shadow-md hover:opacity-85 ${
        isLoading || disabled ? `${loading} pointer-events-none` : `${base}`
      }`}
      type={type}
      disabled={disabled}
      {...props}
    >
      {isLoading ? (
        <Loading />
      ) : icon ? (
        <div className="flex items-center justify-center">
          <div className="mr-2">{icon}</div>
          {children}
        </div>
      ) : (
        <>{children}</>
      )}
    </button>
  );
};
