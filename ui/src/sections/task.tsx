"use client";

import { useAppDispatch } from "@/redux/hooks";
import EditIcon from "@mui/icons-material/Edit";
import DeleteIcon from "@mui/icons-material/Delete";
import { deleteTask, editTask } from "@/redux/slices/userSlice";
import { GenerateRandomBg } from "@/utils/generateRandomBg";
import React, { useState } from "react";
import Modal from "@/components/Modal";

export default function Task({ task, workId }: any) {
  const [taskInfo, setTaskInfo] = useState({
    description: "",
    workSpaceId: "",
    columnId: "",
    assignment: "",
    id: "",
    status: "backLog",
  });
  const [iseditTask, setIsEditTask] = useState(false);

  const dispatch = useAppDispatch();

  function handlerOpenEditTask(id?: string) {
    setTaskInfo((prev) => {
      return { ...prev, id: id ? id : "" };
    });
    setIsEditTask(!iseditTask);
  }

  function handlerDeleteTask(id: string) {
    dispatch(deleteTask(id));
  }

  function handlerChange(e: any) {
    const { value, name } = e.target;
    setTaskInfo((prev) => {
      return { ...prev, [name]: value, workSpaceId: workId };
    });
  }

  function handlerEditTask() {
    const { description, id, assignment } = taskInfo;
    dispatch(editTask({ description, id, assignment }));
    handlerOpenEditTask();
  }

  const createTitleTask = (
    <div>
      <h2 className="text-xl text-center font-semibold">Edit Task</h2>
    </div>
  );

  const createBodyTask = (
    <div className="space-y-3">
      <div className="space-y-3">
        <label className="block">Description</label>
        <input
          name="description"
          className="bg-black/30 rounded-md text-white p-2 outline-none"
          type="text"
          onChange={handlerChange}
        />
      </div>
      <div className="space-y-3">
        <label className="block">assignment</label>
        <input
          name="assignment"
          className="bg-black/30 rounded-md text-white p-2 outline-none"
          type="text"
          onChange={handlerChange}
        />
      </div>
    </div>
  );

  const createActionsTask = (
    <button
      onClick={handlerEditTask}
      className="bg-white w-full py-2 rounded-md active:bg-black/40 text-black font-semibold my-4"
    >
      Submit
    </button>
  );

  return (
    <div
      key={task.id}
      className="bg-zinc-800 max-w-[200px] grid gap-2 h-fit rounded-lg p-3"
    >
      <h2 className="flex items-center justify-between break-words">
        {task?.description}
        <div className="flex">
          <span
            onClick={() => handlerOpenEditTask(task.id)}
            className="cursor-pointer active:bg-white/30 rounded-full p-1 flex items-center"
          >
            <EditIcon fontSize="small" />
          </span>
          <span
            onClick={() => handlerDeleteTask(task.id)}
            className="cursor-pointer active:bg-white/30 rounded-full p-1 flex items-center"
          >
            <DeleteIcon fontSize="small" />
          </span>
        </div>
      </h2>
      <div className="flex gap-1">
        {task?.assignment?.map((i: string) => (
          <div key={i} className="relative">
            <div
              style={{ backgroundColor: GenerateRandomBg() }}
              className="rounded-full w-9 h-9 font-bold flex items-center justify-center text-black"
            >
              {i?.slice(0, 2)}
            </div>
          </div>
        ))}
      </div>
      <Modal
        show={iseditTask}
        onClose={handlerOpenEditTask}
        createActions={createActionsTask}
        createBody={createBodyTask}
        createTitle={createTitleTask}
      />
    </div>
  );
}
