import Button from "@/components/Button";
import Column from "@/components/column";
import { useAppSelector } from "@/redux/hooks";
import {
  getColumns,
  getWorkSpaceId,
  getWorkSpaces,
} from "@/redux/slices/userSlice";
import CreateColumn from "./createColumn";

export default function WorkSpace() {
  const allColumns = useAppSelector(getColumns);
  const workspaceID = useAppSelector(getWorkSpaceId);
  const WorkSpaces = useAppSelector(getWorkSpaces);

  const columns = allColumns.filter((i) => i.workSpaceId === workspaceID);
  const isHaveWorkSoace = WorkSpaces.some((i) => i.id === workspaceID);

  console.log(columns, "colsss", workspaceID);

  return (
    <section className="p-10 w-full flex gap-10 relative">
      {columns?.length > 0 ? (
        columns?.map(
          (i) =>
            i.workSpaceId === workspaceID && (
              <Column
                title={i.title}
                key={i.id}
                id={i.id}
                worlId={i.workSpaceId}
              />
            )
        )
      ) : (
        <>
          {isHaveWorkSoace ? (
            <CreateColumn id={workspaceID} />
          ) : (
            <h2 className="text-center w-fit absolute -translate-x-1/2 left-1/2 top-1/2 -translate-y-1/2 text-4xl font-bold">
              Select a Workspace
            </h2>
          )}
        </>
      )}
    </section>
  );
}
