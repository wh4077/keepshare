import { Space, theme } from "antd";
import MultipleFilter from "./MultipleFilter";
import SelectionFilter from "./SelectionFilter";
import { LinkFormatType } from "@/constant";
import ColumnsFilter from "./ColumnsFilter";
import useStore from "@/store";

interface ComponentProps {
  handleSearch: (search: string) => void;
  handleFormat: (formatType: LinkFormatType) => void;
  handleHost: (hostName: string) => void;
  hostName: string;
}
const TableFilter = ({ handleSearch, handleFormat, handleHost, hostName }: ComponentProps) => {
  const { token } = theme.useToken();
  const isMobile = useStore((state) => state.isMobile);

  return (
    <Space.Compact
      block
      style={{
        justifyContent: "space-between",
        flexWrap: isMobile ? "wrap" : "nowrap",
      }}
    >
      <Space.Compact block direction="vertical">
        <SelectionFilter handleFormat={handleFormat} handleHost={handleHost} hostName={hostName}/>
        <MultipleFilter handleSearch={handleSearch} />
      </Space.Compact>
      <Space
        style={{
          alignSelf: "flex-start",
          marginLeft: isMobile ? "0" : "200px",
          marginTop: isMobile ? token.margin : "0",
        }}
      >
        <ColumnsFilter />
      </Space>
    </Space.Compact>
  );
};

export default TableFilter;
