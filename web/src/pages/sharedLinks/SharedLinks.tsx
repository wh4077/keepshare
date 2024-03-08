import TableFilter from "@/components/sharedLinks/TableFilter";
import SharedLinkTable from "../../components/sharedLinks/SharedLInkTable";
import { querySharedLinks } from "@/api/link";
import useStore from "@/store";
import { useEffect, useState } from "react";
import { shimSharedLinksTableData } from "@/util";
import { LinkFormatType } from "@/constant";
import { message } from "antd";

interface PaginationInfo {
  search: string;
  pageIndex: number;
  pageSize: number;
}
const Content = () => {
  const [
    setIsLoading,
    setTableData,
    setTotalSharedNum,
    totalSharedLinks,
    setTotalSharedLinks,
  ] = useStore((state) => [
    state.setIsLoading,
    state.setTableData,
    state.setTotalSharedNum,
    state.totalSharedLinks,
    state.setTotalSharedLinks,
  ]);

  const [currentHostName, setCurrentHostName] = useState<string>("RapidGator");
  const [pageInfo, setPageInfo] = useState<PaginationInfo>({
    search: "",
    // search: `host:"${currentHostName.toLowerCase()}"`,
    pageIndex: 1,
    pageSize: 10,
  });
  const { search, pageIndex, pageSize } = pageInfo;

  const { data, error, isLoading, mutate } = querySharedLinks({
    search,
    limit: pageSize,
    pageIndex: pageIndex,
  });

  error && message.error("query shared links error!");

  useEffect(() => {
    const tableData = data?.data?.list;
    const totalSharedNum = data?.data?.total;
    if (!isLoading && Array.isArray(tableData)) {
      const shimTableData = tableData.map((v) =>
        shimSharedLinksTableData({ ...v }),
      );
      setTableData(shimTableData);

      const localTotalSharedLinksKeys = totalSharedLinks.map((v) => v.auto_id);
      const newTotalSharedLinks = shimTableData.filter(
        (v) => !localTotalSharedLinksKeys.includes(v.auto_id),
      );
      setTotalSharedLinks([...totalSharedLinks, ...newTotalSharedLinks]);

      totalSharedNum && setTotalSharedNum(totalSharedNum);
    }
    setIsLoading(isLoading);
  }, [isLoading, pageInfo, data]);

  const handlePageChange = (pageIndex: number, pageSize: number) => {
    setPageInfo({ search, pageIndex, pageSize });
  };

  const [currentSearch, setCurrentSearch] = useState<string>("");
  const handleSearch = (search: string) => {
    setCurrentSearch(search)

    // search = search + " " + `host:"${currentHostName.toLowerCase()}"`;
    search = search + " " + `host:"${currentHostName}"`;

    setPageInfo({ search, pageIndex: 0, pageSize });
  };

  const [formatType, setFormatType] = useState<LinkFormatType>(
    LinkFormatType.TEXT,
  );
  const handleFormat = (formatType: LinkFormatType) =>
    setFormatType(formatType);

  const handleHost = (hostName: string) => {
    setCurrentHostName(hostName);

    // use hostName directly, currentHostName is not changed now;
    // let search = currentSearch + " " + `host:"${hostName.toLowerCase()}"`;
    let search = currentSearch + " " + `host:"${hostName}"`;

    setPageInfo({ search, pageIndex: 0, pageSize });
  };

  return (
    <>
      <TableFilter
        handleFormat={handleFormat}
        handleHost={handleHost}
        handleSearch={handleSearch}
        hostName={currentHostName}
      />
      <SharedLinkTable
        refresh={mutate}
        formatType={formatType}
        handlePageChange={handlePageChange}
      />
    </>
  );
};

export default Content;
