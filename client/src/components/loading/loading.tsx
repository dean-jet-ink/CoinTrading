import React from "react";

const Loading = () => {
  return (
    <div className="flex justify-center my-20 space-x-2">
      <div className="animate-ping  h-4 w-1 bg-blue-600 rounded-full"></div>
      <div className="animate-ping  h-4 w-1 bg-blue-600 rounded-full animation-delay-100"></div>
      <div className="animate-ping  h-4 w-1 bg-blue-600 rounded-full animation-delay-200"></div>
      <div className="animate-ping  h-4 w-1 bg-blue-600 rounded-full animation-delay-300"></div>
      <div className="animate-ping  h-4 w-1 bg-blue-600 rounded-full animation-delay-400"></div>
    </div>
  );
};

export default Loading;
