import { useEffect, useState } from "react";
import { MailOutlined } from "@ant-design/icons";
import { Button, Divider, Space, Tabs, Typography, theme, message } from "antd";
import useStore from "@/store";
import { getRapidGatorHostInfo } from "@/api/rapidgator";
import SetAccountModal from './SetAccountModal';

const { Text } = Typography;

const RapidGator = () => {
  const [openSetAccount, setOpenSetAccount] = useState(false);
  const { token } = theme.useToken();
  const [rapidGatorInfo, setRapidGatorInfo] = useStore((state) => [
    state.rapidGatorInfo,
    state.setRapidGatorInfo,
  ]);
  useEffect(() => {
    getRapidGatorHostInfo().then(({ data, error }) => {
      if (error) {
        message.error(error.message);
        return;
      }
      data && setRapidGatorInfo(data);
    });
  }, []);

  const handleOpenSetAccount = () => {
    setOpenSetAccount(true);
  };

  const handleSetAccountClose = () => {
    setOpenSetAccount(false);
  };

  return (
    <>
      <Tabs items={[{ key: "account", label: "Account" }]} />
      <Space align="start" wrap>
        <Space style={{ width: "200px" }}>
          <Text style={{ color: token.colorTextSecondary }}>Account</Text>
        </Space>
        <Space direction="vertical">
          <Text copyable strong>
            {rapidGatorInfo.account || "-"}
          </Text>
          <Button icon={<MailOutlined style={{ color: token.colorPrimary }} />} onClick={handleOpenSetAccount}>
            Set Account
          </Button>
          <Space direction="vertical" align="start" >
            <Text>
              You can set your RapidGator account.
            </Text>
          </Space>
        </Space>
      </Space>
      <Divider />
      <SetAccountModal
        open={openSetAccount}
        onClose={handleSetAccountClose}
      />
    </>
  );
};

export default RapidGator;
