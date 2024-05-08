"use client";

import { useAppSelector } from "@/redux/hooks";
import { getAuth } from "@/redux/slices/userSlice";
import SideBar from "@/sections/sideBar";
import WorkSpace from "@/sections/workSpace";
import { useRouter } from "next/navigation";
import { useEffect } from "react";

export default function Home() {
  const auth = useAppSelector(getAuth);
  const router = useRouter();

  useEffect(() => {
    if (!auth) {
      router.push("/login");
    }
  }, [auth]);

  let content = auth ? (
    <>
      <div className="text-center text-2xl py-5 bg-slate-800 font-bold flex">
        <h2 className="flex-1">Issuse Tracker App</h2>
      </div>
      <div
        style={{ height: "calc(100vh - 72px)" }}
        className="flex overflow-x-scroll overflow-y-hidden"
      >
        <SideBar />
        <WorkSpace />
      </div>
    </>
  ) : (
    ""
  );
  return <main>{content}</main>;
}
