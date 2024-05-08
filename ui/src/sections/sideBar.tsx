"use client";

import Modal from "@/components/Modal";
import LogoutIcon from "@mui/icons-material/Logout";
import EditIcon from "@mui/icons-material/Edit";
import DeleteIcon from "@mui/icons-material/Delete";
import AddIcon from "@mui/icons-material/Add";
import StarOutlineIcon from "@mui/icons-material/StarOutline";
import StarIcon from "@mui/icons-material/Star";
import { useAppDispatch, useAppSelector } from "@/redux/hooks";
import { useState } from "react";
import {
  addNewWorkSpace,
  changeAuth,
  changeWorkSpaceName,
  deleteWorkSpace,
  getWorkSpaceId,
  selectWorkSpace,
  selectWorkSpaces,
  toggleFavorite,
} from "@/redux/slices/userSlice";

export default function SideBar() {
  const [nameWorkSpace, setNameWorkSpace] = useState({ title: "", id: "" });
  const [show, setShow] = useState(false);

  const workSpaces = useAppSelector(selectWorkSpaces);
  const workSpaceId = useAppSelector(getWorkSpaceId);
  const workSpacesSorted = [...workSpaces].sort((a, b) =>
    a.isFovrite === b.isFovrite ? 0 : a.isFovrite ? -1 : 1
  );
  console.log(workSpacesSorted, "work");

  const dispatch = useAppDispatch();

  function handlerCreateWorkSpace() {
    dispatch(addNewWorkSpace(nameWorkSpace.title));
    setNameWorkSpace({ title: "", id: "" });
    setShow(!show);
  }

  function handleToggleFovrite(id: string) {
    dispatch(toggleFavorite(id));
  }

  function handleDelete(id: string) {
    dispatch(deleteWorkSpace(id));
  }

  function handleChnageWorkSpaceTitle(e: any, id: string) {
    const { value } = e.target;
    setNameWorkSpace((prev) => {
      return { ...prev, id: id, title: value };
    });
  }

  function handleEditWorkSpace(idWorkSpace: string) {
    setNameWorkSpace((prevState) => ({
      ...prevState,
      id: prevState.id === idWorkSpace ? "" : idWorkSpace,
    }));
  }

  function updateWorkSpace(e: any) {
    if (e.key === "Enter") {
      dispatch(changeWorkSpaceName(nameWorkSpace));
      setNameWorkSpace({ title: "", id: "" });
    }
  }

  function onCloseModal() {
    setShow(!show);
  }

  function handlerChange(e: any) {
    const { value } = e.target;
    setNameWorkSpace({ title: value, id: "" });
  }

  function handleLogout() {
    dispatch(changeAuth(false));
  }

  function handleSelectWorkSpace(id: string) {
    dispatch(selectWorkSpace(id));
  }

  const createTitle = (
    <div>
      <h2 className="text-xl font-semibold">Create Work Spaces</h2>
    </div>
  );

  const createBody = (
    <div className="py-2">
      <input
        value={nameWorkSpace.title}
        onChange={handlerChange}
        placeholder="Work Space Name"
        className="block text-white/80 my-2 rounded-md p-3 bg-black/40 outline-none"
        type="text"
      />
    </div>
  );

  const createActions = (
    <div className="mx-auto w-full">
      <button
        onClick={handlerCreateWorkSpace}
        className="bg-white font-semibold active:bg-black text-black rounded-md px-4 py-2 w-full"
      >
        Submit
      </button>
    </div>
  );
  console.log(workSpaceId, "id");

  return (
    <>
      <nav className="p-5 bg-zinc-900 h-full">
        <div className="flex items-center p-5">
          <div className="bg-purple-700 h-8 w-8 rounded-lg"></div>
          <h2 className="px-3 text-lg min-w-28">user text</h2>
          <span className="rounded-full p-1 bg-transparent active:bg-white/20 cursor-pointer">
            <LogoutIcon onClick={handleLogout} />
          </span>
        </div>
        <h2 className="font-bold text-2xl py-5">Work Spaces</h2>
        <ul
          style={{ height: "calc(100vh - 260px)" }}
          className="text-lg overflow-y-scroll h-full"
        >
          {workSpacesSorted.map((i) => (
            <li
              key={i.id}
              onClick={() => handleSelectWorkSpace(i.id)}
              style={{
                backgroundColor: i.id === workSpaceId ? "#ffffff24" : "",
              }}
              className="hover:bg-black/40 p-2 rounded-md cursor-pointer"
            >
              <div className="flex">
                {i.id !== nameWorkSpace.id ? (
                  <h2 className="flex-1">{i.title}</h2>
                ) : (
                  <input
                    onChange={(e: any) => handleChnageWorkSpaceTitle(e, i.id)}
                    onKeyDown={updateWorkSpace}
                    type="text"
                    className="bg-white/10 max-w-[100px] text-white rounded-md py-2 px-3 mr-2"
                  />
                )}
                <span onClick={() => handleToggleFovrite(i.id)}>
                  {!i.isFovrite ? <StarOutlineIcon /> : <StarIcon />}
                </span>
                <span>
                  <EditIcon onClick={() => handleEditWorkSpace(i.id)} />
                </span>
                <span>
                  <DeleteIcon onClick={() => handleDelete(i.id)} />
                </span>
              </div>
            </li>
          ))}
          <button
            onClick={onCloseModal}
            className="text-center bg-black/50 rounded-full w-10 h-10 p-2 flex items-center justify-center active:bg-white my-5 mx-auto"
          >
            <AddIcon />
          </button>
        </ul>
      </nav>
      <Modal
        show={show}
        onClose={onCloseModal}
        createTitle={createTitle}
        createBody={createBody}
        createActions={createActions}
      />
    </>
  );
}
