type LoadingProps = {
  variant: "sm" | "md" | "lg";
};

const variants = {
  sm: {
    size: "h-2 w-1",
    gap: "gap-2",
  },
  md: {
    size: "h-4 w-2",
    gap: "gap-4",
  },
  lg: {
    size: "h-6 w-3",
    gap: "gap-6",
  },
};

export const Loading = ({ variant }: LoadingProps) => {
  const { size, gap } = variants[variant];

  return (
    <div className={`flex justify-center ${gap}`}>
      <div className={`animate-ping  ${size} bg-zinc-400 rounded-full`}></div>
      <div
        className={`animate-ping  ${size} bg-zinc-400 rounded-full animation-delay-100`}
      ></div>
      <div
        className={`animate-ping  ${size} bg-zinc-400 rounded-full animation-delay-200`}
      ></div>
    </div>
  );
};
