import React from "react";

export default function Button({ title, onClick, className }: any) {
  return (
    <button
      onClick={onClick}
      className={`${
        className
          ? className
          : "bg-white font-semibold active:bg-black text-black rounded-md px-4 py-2 w-full"
      }`}
    >
      {title}
    </button>
  );
}
