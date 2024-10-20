export default function AddButton(cur: number, total: number) {
  return (
    <div className="rounded-lg items-center py-2 px-3 border border-dashed  border-gray-300 w-[180px] h-[56px] flex-shrink-0">
      <div className="bg-white  rounded-lg flex justify-between items-center px-4 py-2 gap-2 cursor-pointer w-full">
        <div className="flex justify-between gap-1.5">
          <div className="flex justify-center items-center " />
          <div className="font-bold text-[#2A55EC]">添加模型</div>
        </div>
        <div>
          <div className=" text-gray-400">
            {cur}/{total}
          </div>
        </div>
      </div>
    </div>
  );
}
